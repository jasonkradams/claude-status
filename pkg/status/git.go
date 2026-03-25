package status

import (
	"os/exec"
	"strings"
)

// GitInfo holds current branch and file change counts.
type GitInfo struct {
	Branch   string `json:"branch"`
	Staged   int    `json:"staged"`
	Modified int    `json:"modified"`
}

// InRepo reports whether the current directory is inside a git repository.
func InRepo() bool {
	return exec.Command("git", "rev-parse", "--git-dir").Run() == nil
}

// Branch returns the current git branch name, or empty string if not in a repo.
func Branch() string {
	out, _ := exec.Command("git", "branch", "--show-current").Output()
	return strings.TrimSpace(string(out))
}

// Info returns branch name and staged/modified file counts.
// Returns a zero-value GitInfo when not inside a git repository.
func Info() GitInfo {
	if !InRepo() {
		return GitInfo{}
	}
	stagedOut, _ := exec.Command("git", "diff", "--cached", "--numstat").Output()
	modifiedOut, _ := exec.Command("git", "diff", "--numstat").Output()
	return GitInfo{
		Branch:   Branch(),
		Staged:   countLines(stagedOut),
		Modified: countLines(modifiedOut),
	}
}

func countLines(b []byte) int {
	s := strings.TrimSpace(string(b))
	if s == "" {
		return 0
	}
	return len(strings.Split(s, "\n"))
}
