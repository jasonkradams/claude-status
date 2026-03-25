package status

import "strings"

// Bar renders a 10-character progress bar for the given percentage (0–100).
// fillChar fills completed segments; emptyChar fills the remainder.
func Bar(pct int, fillChar, emptyChar string) string {
	filled := min(pct*10/100, 10)
	return strings.Repeat(fillChar, filled) + strings.Repeat(emptyChar, 10-filled)
}
