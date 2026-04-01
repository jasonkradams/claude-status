#!/bin/sh
set -eu

INSTALL_DIR="${INSTALL_DIR:-$HOME/bin}"
INSTALL_NAME="claude-status"

count=0
for f in bin/*; do
    [ -f "$f" ] && [ -x "$f" ] && count=$((count + 1))
done

if [ "$count" -eq 0 ]; then
    echo "No binaries found in bin/. Run 'make build' first." >&2
    exit 1
fi

echo "Available widgets:"
i=0
for f in bin/*; do
    if [ -f "$f" ] && [ -x "$f" ]; then
        i=$((i + 1))
        printf "  %d) %s\n" "$i" "${f##*/}"
    fi
done

printf "Select a widget to activate [1-%d]: " "$count"
read -r choice

case "$choice" in
    ''|*[!0-9]*) echo "Invalid selection." >&2; exit 1 ;;
esac

if [ "$choice" -lt 1 ] || [ "$choice" -gt "$count" ]; then
    echo "Invalid selection." >&2
    exit 1
fi

i=0
selected=""
for f in bin/*; do
    if [ -f "$f" ] && [ -x "$f" ]; then
        i=$((i + 1))
        if [ "$i" -eq "$choice" ]; then
            selected="$f"
            break
        fi
    fi
done

echo "==> Installing ${selected##*/} as $INSTALL_DIR/$INSTALL_NAME..."
mkdir -p "$INSTALL_DIR"
cp "$selected" "$INSTALL_DIR/$INSTALL_NAME"
