# ezcommit

Effortlessly generate professional git commit messages based on your staged changes using the power of Google Gemini. Say goodbye to commit message writer's block and ensure clear, consistent commit history!

## Features

- **Automated Commit Message Generation:** Analyzes your staged git diff and uses Google Gemini to suggest relevant and informative commit messages.
- **Conventional Commits Support:** Generates messages that adhere to the [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) specification (e.g., `feat:`, `fix:`, `chore:`).
- **Customizable Prompts (Future):** Allows you to fine-tune the prompts sent to Gemini for more specific commit message generation.
- **Easy Integration:** Simple to use with your existing git workflow.

## How It Works

`ezcommit` leverages the Google Gemini API to understand the context of your code changes and generate meaningful commit messages. It analyzes the output of `git diff --staged` and sends a carefully crafted prompt to Gemini, requesting a commit message that summarizes the changes.

## Prerequisites

- **Go Installation:** `ezcommit` is written in Go. Make sure you have Go installed on your system. You can download it from [https://go.dev/dl/](https://go.dev/dl/).
- **Git:** You need Git installed to stage your changes.
- **Google Gemini API Key:** You will need an API key from Google AI Studio to use the Gemini models.

## Getting Your Gemini API Key

1.  Go to the Google AI Studio website: [https://aistudio.google.com/app/apikey](https://aistudio.google.com/app/apikey)
2.  If you haven't already, you'll need to create a project.
3.  Once your project is set up, navigate to the "API keys" section (usually in the sidebar).
4.  Click on "Create new API key".
5.  Copy the generated API key. **Keep this key secure and do not share it publicly.**

## Installation

Currently, `ezcommit` is under development. Once a stable version is released, installation instructions will be provided here.

**(Example - Future Installation via Go):**

```bash
go install [github.com/minhnhat97kg/ezcommit@latest](https://github.com/minhnhat97kg/ezcommit@latest)
```
