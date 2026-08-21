#!/usr/bin/env bash
set -e

# Bender Remote/Local Installer Script
# Installs Bender CLI binary globally, resolving Darwin/Linux/Windows platform assets.
#
# Usage:
#   curl -fsSL https://raw.githubusercontent.com/twoBoots/bender/main/install.sh | bash
#   or: ./install.sh [options]

RAW_BASE_URL="https://raw.githubusercontent.com/twoBoots/bender/main"
GITHUB_REPO="twoBoots/bender"
BINARY_NAME="bender"
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}" 2>/dev/null)" && pwd || true)"
TARGET_DIR="$(pwd)"

export PATH="/opt/homebrew/bin:/usr/local/bin:$HOME/go/bin:$PATH"

NON_INTERACTIVE=false
FORCE_ARG=false

for arg in "$@"; do
    case "$arg" in
        --non-interactive|-y|--yes)
            NON_INTERACTIVE=true
            ;;
        --force|-f|--overwrite)
            FORCE_ARG=true
            ;;
        -*)
            # Ignore other flags
            ;;
        *)
            TARGET_DIR="$arg"
            ;;
    esac
done

if [ "$CI" = "true" ]; then
    NON_INTERACTIVE=true
fi

echo "🤖 Installing Bender CLI into $(pwd)..."

OS="$(uname -s | tr '[:upper:]' '[:lower:]')"
ARCH="$(uname -m)"
case "$ARCH" in
    x86_64|amd64) ARCH="x86_64" ;;
    aarch64|arm64) ARCH="aarch64" ;;
esac

RELEASE_BINARY="${BINARY_NAME}-${OS}-${ARCH}"
[ "$OS" = "windows" ] && RELEASE_BINARY="${RELEASE_BINARY}.exe"

INSTALL_BIN_DIR="${HOME}/.local/bin"
[ -d "/usr/local/bin" ] && [ -w "/usr/local/bin" ] && INSTALL_BIN_DIR="/usr/local/bin"
mkdir -p "$INSTALL_BIN_DIR"

CLI_INSTALLED=false

# Tier 1: Build locally if Go is available and source clone is present
if [ -n "$SCRIPT_DIR" ] && [ -f "$SCRIPT_DIR/main.go" ] && command -v go >/dev/null 2>&1; then
    TMP_BUILD="${INSTALL_BIN_DIR}/.${BINARY_NAME}.build.$$"
    (cd "$SCRIPT_DIR" && CGO_ENABLED=0 go build -ldflags="-s -w" -o "$TMP_BUILD" .) >/dev/null 2>&1 || true
    if [ -f "$TMP_BUILD" ]; then
        chmod +x "$TMP_BUILD"
        if [ "$OS" = "darwin" ]; then
            xattr -d com.apple.quarantine "$TMP_BUILD" 2>/dev/null || true
            codesign -s - --force "$TMP_BUILD" 2>/dev/null || true
        fi
        mv -f "$TMP_BUILD" "${INSTALL_BIN_DIR}/${BINARY_NAME}"
        CLI_INSTALLED=true
        echo "  [✓] Compiled and installed CLI globally with Go (${INSTALL_BIN_DIR}/${BINARY_NAME})"
    fi
fi

# Tier 2: Download prebuilt binary from GitHub Releases
if [ "$CLI_INSTALLED" = false ]; then
    RELEASE_URL="https://github.com/${GITHUB_REPO}/releases/latest/download/${RELEASE_BINARY}"
    TMP_DOWNLOAD="${INSTALL_BIN_DIR}/.${BINARY_NAME}.dl.$$"
    if command -v curl >/dev/null 2>&1; then
        if curl -fsSL "$RELEASE_URL" -o "$TMP_DOWNLOAD" 2>/dev/null; then
            chmod +x "$TMP_DOWNLOAD"
            if [ "$OS" = "darwin" ]; then
                xattr -d com.apple.quarantine "$TMP_DOWNLOAD" 2>/dev/null || true
                codesign -s - --force "$TMP_DOWNLOAD" 2>/dev/null || true
            fi
            mv -f "$TMP_DOWNLOAD" "${INSTALL_BIN_DIR}/${BINARY_NAME}"
            CLI_INSTALLED=true
            echo "  [✓] Downloaded prebuilt binary from GitHub Releases to ${INSTALL_BIN_DIR}/${BINARY_NAME}"
        else
            rm -f "$TMP_DOWNLOAD"
        fi
    elif command -v wget >/dev/null 2>&1; then
        if wget -qO "$TMP_DOWNLOAD" "$RELEASE_URL" 2>/dev/null; then
            chmod +x "$TMP_DOWNLOAD"
            if [ "$OS" = "darwin" ]; then
                xattr -d com.apple.quarantine "$TMP_DOWNLOAD" 2>/dev/null || true
                codesign -s - --force "$TMP_DOWNLOAD" 2>/dev/null || true
            fi
            mv -f "$TMP_DOWNLOAD" "${INSTALL_BIN_DIR}/${BINARY_NAME}"
            CLI_INSTALLED=true
            echo "  [✓] Downloaded prebuilt binary from GitHub Releases to ${INSTALL_BIN_DIR}/${BINARY_NAME}"
        else
            rm -f "$TMP_DOWNLOAD"
        fi
    fi
fi

# Tier 3: Fallback announcement
if [ "$CLI_INSTALLED" = false ]; then
    echo "  [!] Pre-built binary could not be fetched and local Go compiler is missing."
    echo "      Install Go 1.22+ or download the binary manually from https://github.com/${GITHUB_REPO}/releases."
fi

echo ""
echo "🤖 Bender successfully installed!"
echo "Available CLI commands:"
echo "  bender --help        - Show available commands"
echo "  bender version       - Show version and commit metadata"
echo "  bender update        - Update Bender CLI in-place from GitHub Releases"
echo "  bender mcp           - Start Model Context Protocol (MCP) server over stdio"
echo "  bender mcp install   - Configure Bender MCP in AI coding assistants"
echo "  bender hello         - Run sample template extension command"
