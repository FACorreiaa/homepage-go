package service

import (
	"bytes"
	"encoding/json"
	"net/http"

	"myapp/internal/db"
)

func SendLeadNotification(webhookURL string, lead db.ProposalLead) {
	if webhookURL == "" {
		return
	}
	go func() {
		msg := "**New Proposal Lead**\n" +
			"**Name:** " + lead.Name + "\n" +
			"**Email:** " + lead.Email + "\n" +
			"**Type:** " + lead.ProjectType + "\n" +
			"**Details:** " + lead.Details
		payload, _ := json.Marshal(map[string]string{"content": msg})
		http.Post(webhookURL, "application/json", bytes.NewReader(payload)) //nolint:errcheck
	}()
}
