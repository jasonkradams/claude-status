package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

type Input struct {
	Model struct {
		DisplayName string `json:"display_name"`
	} `json:"model"`
	Workspace struct {
		CurrentDir string `json:"current_dir"`
	} `json:"workspace"`
	Cost struct {
		TotalCostUSD    float64 `json:"total_cost_usd"`
		TotalDurationMS int64   `json:"total_duration_ms"`
	} `json:"cost"`
	ContextWindow struct {
		UsedPercentage float64 `json:"used_percentage"`
	} `json:"context_window"`
}

const (
	cyan   = "\033[36m"
	green  = "\033[32m"
	yellow = "\033[33m"
	red    = "\033[31m"
	reset  = "\033[0m"
)

func main() {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error reading stdin:", err)
		os.Exit(1)
	}

	var input Input
	if err := json.Unmarshal(data, &input); err != nil {
		fmt.Fprintln(os.Stderr, "error parsing JSON:", err)
		os.Exit(1)
	}

	pct := int(input.ContextWindow.UsedPercentage)

	barColor := green
	switch {
	case pct >= 90:
		barColor = red
	case pct >= 70:
		barColor = yellow
	}

	filled := min(pct/10, 10)
	bar := strings.Repeat("█", filled) + strings.Repeat("░", 10-filled)

	mins := input.Cost.TotalDurationMS / 60000
	secs := (input.Cost.TotalDurationMS % 60000) / 1000

	dir := input.Workspace.CurrentDir
	if idx := strings.LastIndex(dir, "/"); idx >= 0 {
		dir = dir[idx+1:]
	}

	branch := gitBranch()

	line1 := fmt.Sprintf("%s[%s]%s 📁 %s%s", cyan, input.Model.DisplayName, reset, dir, branch)
	line2 := fmt.Sprintf("%s%s%s %d%% | %s$%.2f%s | ⏱️ %dm %ds",
		barColor, bar, reset,
		pct,
		yellow, input.Cost.TotalCostUSD, reset,
		mins, secs,
	)

	fmt.Println(line1)
	fmt.Println(line2)
}

func gitBranch() string {
	cmd := exec.Command("git", "rev-parse", "--git-dir")
	cmd.Stderr = nil
	if err := cmd.Run(); err != nil {
		return ""
	}
	out, err := exec.Command("git", "branch", "--show-current").Output()
	if err != nil {
		return ""
	}
	b := strings.TrimSpace(string(out))
	if b == "" {
		return ""
	}
	return " | 🌿 " + b
}
