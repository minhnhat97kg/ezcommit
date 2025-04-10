package main

import (
	"os/exec"
	"regexp"
	"strings"
)

type GitTools struct {
}

func (g *GitTools) IsInstalled() bool {
	_, err := exec.LookPath("git")
	if err != nil {
		return false
	}
	return true
}

func (g *GitTools) GetChanges() string {
	cmd := exec.Command("git", "diff", "--staged")
	output, err := cmd.Output()
	if err != nil {
		return ""
	}
	lines := filterLines(string(output))
	return lines
}

func (g *GitTools) CommitChanges(message string) error {
	cmd := exec.Command("git", "commit", "-m", message)
	return cmd.Run()
}

func (g *GitTools) LatestCommit() (string, error) {
	cmd := exec.Command("git", "log", "-1", "--pretty=format:%H")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

func (g *GitTools) IsRepository() bool {
	cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")
	output, err := cmd.Output()
	if err != nil {
		return false
	}
	return strings.TrimSpace(string(output)) == "true"
}

func filterLines(diff string) string {
	lines := strings.Split(diff, "\n")
	lockFileRegex := regexp.MustCompile(`^diff --git a/(.*)?`)
	filteredLines := make([]string, 0, len(lines))
	for _, line := range lines {
		isLockFile := false
		if lockFileRegex.MatchString(line) {
			isLockFile = true
			continue
		}
		if !isLockFile {
			filteredLines = append(filteredLines, line)
		}
	}
	return strings.Join(filteredLines, "\n")
}
