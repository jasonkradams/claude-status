package status

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// Input is the full JSON payload Claude Code sends to status line scripts.
type Input struct {
	Model struct {
		ID          string `json:"id"`
		DisplayName string `json:"display_name"`
	} `json:"model"`
	Workspace struct {
		CurrentDir string `json:"current_dir"`
		ProjectDir string `json:"project_dir"`
	} `json:"workspace"`
	Cost struct {
		TotalCostUSD       float64 `json:"total_cost_usd"`
		TotalDurationMS    int64   `json:"total_duration_ms"`
		TotalAPIDurationMS int64   `json:"total_api_duration_ms"`
		TotalLinesAdded    int     `json:"total_lines_added"`
		TotalLinesRemoved  int     `json:"total_lines_removed"`
	} `json:"cost"`
	ContextWindow struct {
		TotalInputTokens  int     `json:"total_input_tokens"`
		TotalOutputTokens int     `json:"total_output_tokens"`
		ContextWindowSize int     `json:"context_window_size"`
		UsedPercentage    float64 `json:"used_percentage"`
		RemainingPct      float64 `json:"remaining_percentage"`
	} `json:"context_window"`
	RateLimits struct {
		FiveHour *RateWindow `json:"five_hour"`
		SevenDay *RateWindow `json:"seven_day"`
	} `json:"rate_limits"`
}

// RateWindow holds usage data for a rate limit window.
type RateWindow struct {
	UsedPercentage float64 `json:"used_percentage"`
	ResetsAt       int64   `json:"resets_at"`
}

// ReadInput reads and parses the JSON payload from stdin.
func ReadInput() (*Input, error) {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		return nil, fmt.Errorf("reading stdin: %w", err)
	}
	var input Input
	if err := json.Unmarshal(data, &input); err != nil {
		return nil, fmt.Errorf("parsing JSON: %w", err)
	}
	return &input, nil
}

// DirName returns the last path component of the current working directory.
func (i *Input) DirName() string {
	dir := i.Workspace.CurrentDir
	for j := len(dir) - 1; j >= 0; j-- {
		if dir[j] == '/' {
			return dir[j+1:]
		}
	}
	return dir
}
