package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

const geminiUrl = "https://generativelanguage.googleapis.com/v1beta/models/gemini-2.0-flash:generateContent"
const geminiPromt = "Generate %d professional git commit messages (type '%s') in %s, separated by '|'. ach option should be a concise, full sentence in the present tense, adhering to the conventional commits specification (' %s: <subject>') : \n\n%s\n\n"

type GeminiResponse struct {
	Candidates []struct {
		Content struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		} `json:"content"`
	} `json:"candidates"`
}

type Gemini struct {
	url    string
	apiKey string
}

func NewGemini() Gemini {
	return Gemini{
		url:    geminiUrl,
		apiKey: os.Getenv("GEMINI_API_KEY"),
	}
}

func (g *Gemini) IsInstalled() bool {
	return g.apiKey != ""
}

func (g *Gemini) SendPrompt(options int, commitType, lang, changes string) ([]string, error) {

	prompt := fmt.Sprintf(geminiPromt, options, commitType, lang, commitType, changes)
	payload := map[string]interface{}{
		"contents": []map[string]interface{}{
			{
				"parts": []map[string]interface{}{
					{"text": prompt},
				},
			},
		},
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("error marshalling payload: %v", err)
	}

	req, err := http.NewRequest(http.MethodPost, g.url+fmt.Sprintf("?key=%s", g.apiKey), bytes.NewBuffer(payloadBytes))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	// You'll likely want to decode the response body into a Go struct
	// to access the generated content. Here's a basic example:
	var responseData GeminiResponse
	err = json.NewDecoder(resp.Body).Decode(&responseData)
	if err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	return strings.Split(responseData.Candidates[0].Content.Parts[0].Text, "|"), nil
}

func (g *Gemini) GetMaxToken() int {
	return 4096
}
