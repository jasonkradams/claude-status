package status

import "fmt"

// Duration formats a millisecond duration as "Xm Ys".
func Duration(ms int64) string {
	mins := ms / 60000
	secs := (ms % 60000) / 1000
	return fmt.Sprintf("%dm %ds", mins, secs)
}

// Cost formats a USD amount as "$X.XX".
func Cost(usd float64) string {
	return fmt.Sprintf("$%.2f", usd)
}
