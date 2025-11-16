package types

type User struct {
	ID       string
	IsActive bool
}

type Team struct {
	ID    string
	Name  string
	Users []User
}

type PullRequest struct {
	ID         string
	AuthorID   string
	ReviewerID string
	TeamID     string
	Status     string
}
