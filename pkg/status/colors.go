package status

// ANSI terminal color codes.
const (
	Cyan   = "\033[36m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Red    = "\033[31m"
	Bold   = "\033[1m"
	Reset  = "\033[0m"
)

// ColorForPct returns a terminal color based on usage thresholds:
// green below 70%, yellow below 90%, red at 90%+.
func ColorForPct(pct int) string {
	switch {
	case pct >= 90:
		return Red
	case pct >= 70:
		return Yellow
	default:
		return Green
	}
}
