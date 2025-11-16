package types

type Team struct {
	ID       string
	TeamName string       `json:"team_name"`
	Members  []TeamMember `json:"members"`
}

type TeamMember struct {
	ID       string
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	IsActive bool   `json:"is_active"`
}

type User struct {
	UserID   string  `json:"user_id"`
	Username string  `json:"username"`
	TeamName *string `json:"team_name"`
	IsActive bool    `json:"is_active"`
}

type PullRequest struct {
	ID                string
	PullRequestID     string   `json:"pull_request_id"`
	PullRequestName   string   `json:"pull_request_name"`
	AuthorID          string   `json:"author_id"`
	Status            string   `json:"status"`
	AssignedReviewers []string `json:"assigned_reviewers"`
}

type PullRequestShort struct {
	ID              string
	PullRequestID   string `json:"pull_request_id"`
	PullRequestName string `json:"pull_request_name"`
	AuthorID        string `json:"author_id"`
	Status          string `json:"status"`
}
