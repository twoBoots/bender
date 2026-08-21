package mcp_test

import (
	"bytes"
	"context"
	"encoding/json"
	"strings"
	"testing"

	"github.com/twoBoots/bender/pkg/mcp"
)

func TestServer_Initialize(t *testing.T) {
	srv := mcp.NewServer("bender-mcp", "1.0.0", "/test/cwd")
	id := json.RawMessage(`1`)
	req := mcp.Request{
		JSONRPC: "2.0",
		ID:      &id,
		Method:  "initialize",
	}

	resp := srv.HandleRequest(context.Background(), req)
	if resp.Error != nil {
		t.Fatalf("unexpected error: %v", resp.Error)
	}

	initRes, ok := resp.Result.(mcp.InitializeResult)
	if !ok {
		t.Fatalf("expected InitializeResult, got %T", resp.Result)
	}

	if initRes.ServerInfo.Name != "bender-mcp" {
		t.Errorf("got server name %q; want bender-mcp", initRes.ServerInfo.Name)
	}
	if initRes.ServerInfo.Version != "v1.0.0" {
		t.Errorf("got server version %q; want v1.0.0", initRes.ServerInfo.Version)
	}
	if srv.Name() != "bender-mcp" {
		t.Errorf("got srv.Name() %q; want bender-mcp", srv.Name())
	}
	if srv.Version() != "v1.0.0" {
		t.Errorf("got srv.Version() %q; want v1.0.0", srv.Version())
	}
	if srv.Cwd() != "/test/cwd" {
		t.Errorf("got srv.Cwd() %q; want /test/cwd", srv.Cwd())
	}
}

func TestServer_ToolsLifecycle(t *testing.T) {
	srv := mcp.NewServer("test-mcp", "0.1.0", "/cwd")

	tool := mcp.Tool{
		Name:        "echo",
		Description: "echoes back message",
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"msg": map[string]string{"type": "string"},
			},
		},
	}

	srv.RegisterTool(tool, func(ctx context.Context, args map[string]interface{}) (mcp.CallToolResult, error) {
		msg, _ := args["msg"].(string)
		if msg == "error" {
			return mcp.CallToolResult{}, fmtError("tool execution error")
		}
		return mcp.NewTextResult("Echo: "+msg, false), nil
	})

	idList := json.RawMessage(`2`)
	listResp := srv.HandleRequest(context.Background(), mcp.Request{
		JSONRPC: "2.0",
		ID:      &idList,
		Method:  "tools/list",
	})
	listRes := listResp.Result.(mcp.ListToolsResult)
	if len(listRes.Tools) != 1 || listRes.Tools[0].Name != "echo" {
		t.Fatalf("unexpected tools list: %v", listRes.Tools)
	}

	callParams, _ := json.Marshal(mcp.CallToolParams{
		Name: "echo",
		Arguments: map[string]interface{}{
			"msg": "bender",
		},
	})
	idCall := json.RawMessage(`3`)
	callResp := srv.HandleRequest(context.Background(), mcp.Request{
		JSONRPC: "2.0",
		ID:      &idCall,
		Method:  "tools/call",
		Params:  callParams,
	})
	callRes := callResp.Result.(mcp.CallToolResult)
	if len(callRes.Content) == 0 || callRes.Content[0].Text != "Echo: bender" {
		t.Fatalf("unexpected tool call result: %v", callRes)
	}

	// Tool execution error handler returns error result
	errCallParams, _ := json.Marshal(mcp.CallToolParams{
		Name: "echo",
		Arguments: map[string]interface{}{
			"msg": "error",
		},
	})
	errResp := srv.HandleRequest(context.Background(), mcp.Request{
		JSONRPC: "2.0",
		ID:      &idCall,
		Method:  "tools/call",
		Params:  errCallParams,
	})
	errRes := errResp.Result.(mcp.CallToolResult)
	if !errRes.IsError {
		t.Errorf("expected isError=true for tool error")
	}

	// Unknown tool
	unknownToolParams, _ := json.Marshal(mcp.CallToolParams{Name: "nonexistent"})
	unknownResp := srv.HandleRequest(context.Background(), mcp.Request{
		JSONRPC: "2.0",
		ID:      &idCall,
		Method:  "tools/call",
		Params:  unknownToolParams,
	})
	if unknownResp.Error == nil {
		t.Errorf("expected error for nonexistent tool")
	}
}

func TestServer_ResourcesAndPrompts(t *testing.T) {
	srv := mcp.NewServer("test-mcp", "0.1.0", "/cwd")

	// Resource
	res := mcp.Resource{
		URI:  "bender://docs/{topic}",
		Name: "docs",
	}
	srv.RegisterResource(res, func(ctx context.Context, uri string) (mcp.ReadResourceResult, error) {
		if uri == "bender://docs/error" {
			return mcp.ReadResourceResult{}, fmtError("resource read error")
		}
		return mcp.ReadResourceResult{
			Contents: []mcp.ResourceContent{
				{URI: uri, MIMEType: "text/plain", Text: "Doc content for " + uri},
			},
		}, nil
	})

	idResList := json.RawMessage(`4`)
	resListResp := srv.HandleRequest(context.Background(), mcp.Request{
		JSONRPC: "2.0",
		ID:      &idResList,
		Method:  "resources/list",
	})
	if len(resListResp.Result.(mcp.ListResourcesResult).Resources) != 1 {
		t.Errorf("expected 1 resource in list")
	}

	readParams, _ := json.Marshal(mcp.ReadResourceParams{URI: "bender://docs/quickstart"})
	idRead := json.RawMessage(`5`)
	readResp := srv.HandleRequest(context.Background(), mcp.Request{
		JSONRPC: "2.0",
		ID:      &idRead,
		Method:  "resources/read",
		Params:  readParams,
	})
	readRes := readResp.Result.(mcp.ReadResourceResult)
	if len(readRes.Contents) == 0 || !strings.Contains(readRes.Contents[0].Text, "quickstart") {
		t.Errorf("unexpected resource read result: %v", readRes)
	}

	// Unknown resource
	unknownResParams, _ := json.Marshal(mcp.ReadResourceParams{URI: "bender://unknown"})
	unknownResResp := srv.HandleRequest(context.Background(), mcp.Request{
		JSONRPC: "2.0",
		ID:      &idRead,
		Method:  "resources/read",
		Params:  unknownResParams,
	})
	if unknownResResp.Error == nil {
		t.Errorf("expected error for unknown resource")
	}

	// Prompt
	prompt := mcp.Prompt{
		Name:        "review",
		Description: "code review prompt",
	}
	srv.RegisterPrompt(prompt, func(ctx context.Context, args map[string]string) (mcp.GetPromptResult, error) {
		if args["fail"] == "true" {
			return mcp.GetPromptResult{}, fmtError("prompt failure")
		}
		return mcp.GetPromptResult{
			Description: "Code review instructions",
			Messages: []mcp.PromptMessage{
				{Role: "user", Content: mcp.ContentItem{Type: "text", Text: "Review this diff"}},
			},
		}, nil
	})

	idPromptList := json.RawMessage(`6`)
	promptListResp := srv.HandleRequest(context.Background(), mcp.Request{
		JSONRPC: "2.0",
		ID:      &idPromptList,
		Method:  "prompts/list",
	})
	if len(promptListResp.Result.(mcp.ListPromptsResult).Prompts) != 1 {
		t.Errorf("expected 1 prompt in list")
	}

	getPromptParams, _ := json.Marshal(mcp.GetPromptParams{Name: "review"})
	idGetPrompt := json.RawMessage(`7`)
	getPromptResp := srv.HandleRequest(context.Background(), mcp.Request{
		JSONRPC: "2.0",
		ID:      &idGetPrompt,
		Method:  "prompts/get",
		Params:  getPromptParams,
	})
	getPromptRes := getPromptResp.Result.(mcp.GetPromptResult)
	if len(getPromptRes.Messages) != 1 {
		t.Errorf("expected 1 prompt message")
	}

	// Unknown prompt
	unknownPromptParams, _ := json.Marshal(mcp.GetPromptParams{Name: "unknown"})
	unknownPromptResp := srv.HandleRequest(context.Background(), mcp.Request{
		JSONRPC: "2.0",
		ID:      &idGetPrompt,
		Method:  "prompts/get",
		Params:  unknownPromptParams,
	})
	if unknownPromptResp.Error == nil {
		t.Errorf("expected error for unknown prompt")
	}
}

func TestServer_ServeStdio(t *testing.T) {
	srv := mcp.NewServer("test-mcp", "0.1.0", "/cwd")
	in := bytes.NewBufferString("{\"jsonrpc\":\"2.0\",\"id\":1,\"method\":\"ping\"}\n{\"jsonrpc\":\"2.0\",\"method\":\"notifications/initialized\"}\n")
	var out bytes.Buffer

	err := srv.Serve(context.Background(), in, &out)
	if err != nil {
		t.Fatalf("Serve returned error: %v", err)
	}

	var resp mcp.Response
	if err := json.Unmarshal(out.Bytes(), &resp); err != nil {
		t.Fatalf("failed to decode response: %v, raw: %s", err, out.String())
	}
}

type testErr string

func (e testErr) Error() string { return string(e) }
func fmtError(msg string) error { return testErr(msg) }
