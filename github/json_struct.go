package github

import (
	"gopkg.in/guregu/null.v3"
)

type SearchIssues struct {
	TotalCount        int         `json:"total_count"`
	IncompleteResults bool        `json:"incomplete_results"`
	Items             []IssueItem `json:"items"`
}

type IssueItem struct {
	Url         string          `json:"url"`
	LabelsUrl   string          `json:"labels_url"`
	CommentsUrl string          `json:"comments_url"`
	EventsUrl   string          `json:"events_url"`
	HtmlUrl     string          `json:"html_url"`
	Id          int             `json:"id"`
	Number      int             `json:"number"`
	Title       string          `json:"title"`
	User        UserItem        `json:"user"`
	Labels      []string        `json:"labels"`
	State       string          `json:"state"`
	Locked      bool            `json:"locked"`
	Assignee    null.String     `json:"assignee"`
	Milestone   null.String     `json:"milestone"`
	Comments    int             `json:"comments"`
	CreatedAt   string          `json:"created_at"`
	UpdatedAt   null.String     `json:"updated_at"`
	ClosedAt    null.String     `json:"closed_at"`
	PullRequest PullRequestItem `json:"pull_request"`
	Body        string          `json:"body"`
	Score       float32         `json:"score"`
}

type UserItem struct {
	Login             string `json:"login"`
	Id                int    `json:"id"`
	AvatarUrl         string `json:"avatar_url"`
	GravatarId        string `json:"gravatar_id"`
	Url               string `json:"url"`
	HtmlUrl           string `json:"html_url"`
	FollowersUrl      string `json:"followers_url"`
	FollowingUrl      string `json:"following_url"`
	GistsUrl          string `json:"gists_url"`
	StarredUrl        string `json:"starred_url"`
	SubscriptionsUrl  string `json:"subscriptions_url"`
	OrganizationsUrl  string `json:"organizations_url"`
	ReposUrl          string `json:"repos_url"`
	EventsUrl         string `json:"events_url"`
	ReceivedEventsUrl string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}

type PullRequestItem struct {
	Url      string `json:"url"`
	HtmlUrl  string `json:"html_url"`
	DiffUrl  string `json:"diff_url"`
	PatchUrl string `json:"patch_url"`
}

type Authorizations struct {
	Id             int         `json:"id"`
	Url            string      `json:"url"`
	Scopes         []string    `json:"scopes"`
	Token          string      `json:"token"`
	TokenLastEight string      `json:"token_last_eight"`
	HashedToken    string      `json:"hashed_token"`
	App            App         `json:"app"`
	Note           string      `json:"note"`
	NoteUrl        string      `json:"note_url"`
	UpdatedAt      null.String `json:"updated_at"`
	CreatedAt      string      `json:"created_at"`
	Fingerprint    string      `json:"fingerprint"`
}

type App struct {
	Url      string `json:"url"`
	Name     string `json:"name"`
	ClientId string `json:"client_id"`
}

type SiteConfig struct {
	Host        string `json:"host"`
	AccessToken string `json:"access_token"`
}
