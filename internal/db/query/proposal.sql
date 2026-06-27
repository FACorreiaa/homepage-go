-- name: CreateProposalLead :exec
INSERT INTO proposal_leads (id, name, email, company, project_type, budget, timeline, details, created_at)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, datetime('now'));

-- name: ListProposalLeads :many
SELECT * FROM proposal_leads ORDER BY created_at DESC;
