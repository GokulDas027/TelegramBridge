
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/yanzay/tbot/v2"
)

func main() {
	// Just to do it a bit fancy
	icons := map[string]string{
		"failure":   "❗️❗️❗️",
		"cancelled": "❕❕❕",
		"success":   "✅✅✅",
	}
	texts:=map[string]string{
		"failure":   "Eda Mone Nee pettu",
		"cancelled": "Enthpattiyeda Uvve?",
		"success":   "Adipoli Mounse ath Work aayi",
	}
		
	var (
		// inputs provided by Github Actions runtime
		// we should define them in action.yml
		token  = os.Getenv("INPUT_TOKEN")
		chat   = os.Getenv("INPUT_CHAT")
		status = os.Getenv("INPUT_STATUS")

		// github environment context
		workflow = os.Getenv("GITHUB_WORKFLOW")
		repo     = os.Getenv("GITHUB_REPOSITORY")
		commit   = os.Getenv("GITHUB_SHA")
	)

	// Create Telegram client using token
	c := tbot.NewClient(token, http.DefaultClient, "https://api.telegram.org")

	icon := icons[strings.ToLower(status)]
	text:=texts[strings.ToLower(status)]// which icon to use?
	link := fmt.Sprintf("https://github.com/%s/commit/%s/checks", repo, commit)
	// Prepare message to send
	msg := fmt.Sprintf(`%s
	%s  *%s*: %s ([%s](%s))`, icon, text, status, repo, workflow, link)

	// Send to chat using Markdown format
	_, err := c.SendMessage(chat, msg, tbot.OptParseModeMarkdown)
	if err != nil {
		log.Fatalf("unable to send message: %v", err)
	}
}
