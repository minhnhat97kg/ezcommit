package main

import (
	"flag"
	"fmt"
	"strings"
)

type AIAgent interface {
	SendPrompt(options int, commitType, lang, changes string) ([]string, error)
	GetMaxTokens() int
	IsInstalled() bool
}

var (
	optionsNum  = flag.Int("options", 3, "Number of options to generate")
	lang        = flag.String("lang", "English", "Language of the commit message")
	commitTypes = []string{"feat", "fix", "chore", "docs", "style", "refactor", "perf", "test", "build", "ci", "revert"}
)

func main() {

	// Create an instance of GitTools
	git := &GitTools{}

	// Check if Git is installed
	if !git.IsInstalled() {
		fmt.Println("⚠️ Git is not installed. Please install Git to use this tool.")
		return
	}
	// Check if in a Git repository
	if !git.IsRepository() {
		fmt.Println("⚠️ Not in a Git repository. Please run this tool inside a Git repository.")
		return
	}

	// Get changes
	changes := git.GetChanges()
	if changes == "" {
		fmt.Println("⚠️ No changes detected. Please make some changes to your files.")
		return
	}

	// Create an instance of Gemini
	agent := NewGemini()

	// Check if agent is installed
	if !agent.IsInstalled() {
		fmt.Println("❌ Gemini is not ready. Please export the environment variable GEMINI_API_KEY.")
		return
	}

	menuType := NewMenu("💡 Select a commit type")
	for i, option := range commitTypes {
		menuType.AddItem(strings.TrimSpace(option), i)
	}

	selectedCommitType, err := menuType.Display()
	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
		return
	}

	fmt.Println("💬 Generating commit message...")
	options, err := agent.SendPrompt(*optionsNum, commitTypes[selectedCommitType.(int)], *lang, changes)
	if err != nil {
		fmt.Println("❌ Error sending prompt:", err)
		return
	}

	menu := NewMenu("🚨 Select a commit message")
	for i, option := range options {
		menu.AddItem(strings.TrimSpace(option), i)
	}
	selectedMsg, err := menu.Display()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	if err := git.CommitChanges(options[selectedMsg.(int)]); err != nil {
		fmt.Println("❌ Error committing changes:", err)
		return
	}

	fmt.Println("🎁 Changes committed successfully!")
}
