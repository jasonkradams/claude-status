package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/jasonkradams/claude/pkg/status"
)

func main() {
	input, err := status.ReadInput()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	pct := int(input.ContextWindow.UsedPercentage)
	barColor := status.ColorForPct(pct)
	bar := status.Bar(pct, "█", "░")

	// Line 1: model, dir, branch, repo link
	branch := status.Branch()
	branchStr := ""
	if branch != "" {
		branchStr = " | 🌿 " + branch
	}

	linkStr := ""
	if remote, err := exec.Command("git", "remote", "get-url", "origin").Output(); err == nil {
		url := strings.TrimSpace(string(remote))
		url = strings.Replace(url, "git@github.com:", "https://github.com/", 1)
		url = strings.TrimSuffix(url, ".git")
		repoName := filepath.Base(url)
		linkStr = fmt.Sprintf(" | 🔗 \033]8;;%s\007%s\033]8;;\007", url, repoName)
	}

	line1 := fmt.Sprintf("%s[%s]%s 📁 %s%s%s",
		status.Cyan, input.Model.DisplayName, status.Reset,
		input.DirName(), branchStr, linkStr)

	// Line 2: context bar, cost, duration, rate limits
	var rateParts []string
	if input.RateLimits.FiveHour != nil {
		rateParts = append(rateParts, fmt.Sprintf("5h: %.0f%%", input.RateLimits.FiveHour.UsedPercentage))
	}
	if input.RateLimits.SevenDay != nil {
		rateParts = append(rateParts, fmt.Sprintf("7d: %.0f%%", input.RateLimits.SevenDay.UsedPercentage))
	}

	line2 := fmt.Sprintf("%s%s%s %d%% | %s%s%s | ⏱️ %s",
		barColor, bar, status.Reset,
		pct,
		status.Yellow, status.Cost(input.Cost.TotalCostUSD), status.Reset,
		status.Duration(input.Cost.TotalDurationMS))

	if len(rateParts) > 0 {
		line2 += fmt.Sprintf(" | %s", strings.Join(rateParts, " "))
	}

	fmt.Println(line1)
	fmt.Println(line2)
}
