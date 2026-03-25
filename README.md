# claude-status

Go-based status line widgets for [Claude Code](https://claude.ai/code).

## Widgets

### multi-line

Two-line status bar showing:
- Line 1: model, current directory, git branch, repo link
- Line 2: context window usage bar, cost, duration, rate limits

## Build

```sh
make build        # build all widgets to bin/
make build-multi-line  # build a specific widget
```

## Install

```sh
make install      # copies to ~/bin/claude-status
```

Then point Claude Code at the binary in your settings:

```json
{
  "statusCommand": "~/bin/claude-status"
}
```
