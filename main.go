package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/yanzay/tbot/v2"
)

func composer(status, event, actor, repo, workflow, link string) string {
	var text string

	// choose icon based on the build status
	icons := map[string]string{
		"failure":   "❗️❗️❗️",
		"cancelled": "❕❕❕",
		"success":   "✅✅✅",
	}

	// Message text composing
	text = icons[strings.ToLower(status)] + "\n"
	text += event + " was made at " + repo + " by " + actor + "\n"
	text += "Check here " + "[" + workflow + "](" + link + ")"

	return text
}

func main() {

	var (
		// inputs provided by Github Actions runtime
		// should be defined in the action.yml
		token  = os.Getenv("INPUT_TOKEN")
		chat   = os.Getenv("INPUT_CHAT")
		status = os.Getenv("INPUT_STATUS")
		event  = os.Getenv("INPUT_EVENT")
		actor  = os.Getenv("INPUT_ACTOR")

		// github environment context
		workflow = os.Getenv("GITHUB_WORKFLOW")
		repo     = os.Getenv("GITHUB_REPOSITORY")
		commit   = os.Getenv("GITHUB_SHA")
	)

	// Create Telegram client using token
	c := tbot.NewClient(token, http.DefaultClient, "https://api.telegram.org")

	// link to the commit
	link := fmt.Sprintf("https://github.com/%s/commit/%s/", repo, commit)

	// Prepare message to send
	msg := composer(status, event, actor, repo, workflow, link)

	// Send to chat using Markdown format
	_, err := c.SendMessage(chat, msg, tbot.OptParseModeMarkdown)
	if err != nil {
		log.Fatalf("unable to send message: %v", err)
	}
}
