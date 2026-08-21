package mcp_test

import (
	"encoding/json"
	"testing"

	"github.com/twoBoots/bender/pkg/mcp"
)

func TestNewResponse(t *testing.T) {
	idRaw := json.RawMessage(`1`)
	res := mcp.NewResponse(&idRaw, map[string]string{"status": "ok"})

	if res.JSONRPC != "2.0" {
		t.Errorf("got JSONRPC %q; want 2.0", res.JSONRPC)
	}
	if res.ID == nil || string(*res.ID) != "1" {
		t.Errorf("unexpected ID in response: %v", res.ID)
	}

	bytes, err := json.Marshal(res)
	if err != nil {
		t.Fatalf("failed to marshal response: %v", err)
	}

	var parsed mcp.Response
	if err := json.Unmarshal(bytes, &parsed); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}
}

func TestNewErrorResponse(t *testing.T) {
	idRaw := json.RawMessage(`"req-123"`)
	errRes := mcp.NewErrorResponse(&idRaw, mcp.MethodNotFoundCode, "unknown method", nil)

	if errRes.Error == nil {
		t.Fatalf("expected error object, got nil")
	}
	if errRes.Error.Code != mcp.MethodNotFoundCode {
		t.Errorf("got error code %d; want %d", errRes.Error.Code, mcp.MethodNotFoundCode)
	}
	if errRes.Error.Message != "unknown method" {
		t.Errorf("got error message %q; want 'unknown method'", errRes.Error.Message)
	}
}

func TestNewTextResult(t *testing.T) {
	res := mcp.NewTextResult("hello world", false)
	if len(res.Content) != 1 {
		t.Fatalf("expected 1 content item, got %d", len(res.Content))
	}
	if res.Content[0].Text != "hello world" {
		t.Errorf("got text %q; want 'hello world'", res.Content[0].Text)
	}
	if res.IsError {
		t.Errorf("expected isError=false")
	}

	errRes := mcp.NewErrorResult("failed")
	if !errRes.IsError {
		t.Errorf("expected isError=true for NewErrorResult")
	}
}
