#!/usr/bin/env bash
set -euo pipefail

INSTALL_DIR="${INSTALL_DIR:-$HOME/bin}"
INSTALL_NAME="claude-status"

mapfile -t binaries < <(find bin/ -maxdepth 1 -type f -perm +111 | sort)

if [ ${#binaries[@]} -eq 0 ]; then
    echo "No binaries found in bin/. Run 'make build' first." >&2
    exit 1
fi

echo "Available widgets:"
for i in "${!binaries[@]}"; do
    printf "  %d) %s\n" "$((i+1))" "${binaries[$i]##*/}"
done

read -rp "Select a widget to activate [1-${#binaries[@]}]: " choice

if ! [[ "$choice" =~ ^[0-9]+$ ]] || [ "$choice" -lt 1 ] || [ "$choice" -gt ${#binaries[@]} ]; then
    echo "Invalid selection." >&2
    exit 1
fi

selected="${binaries[$((choice-1))]}"

echo "==> Installing ${selected##*/} as $INSTALL_DIR/$INSTALL_NAME..."
mkdir -p "$INSTALL_DIR"
cp "$selected" "$INSTALL_DIR/$INSTALL_NAME"
