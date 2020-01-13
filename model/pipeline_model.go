package model

type Pipeline struct {
	CreatedAt string `json:"created_at"`
	ID        int    `json:"id"`
	Ref       string `json:"ref"`
	Sha       string `json:"sha"`
	Status    string `json:"status"`
	UpdatedAt string `json:"updated_at"`
	WebURL    string `json:"web_url"`
}
