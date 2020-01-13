package model

type Namespace struct {
	AvatarURL string      `json:"avatar_url"`
	FullPath  string      `json:"full_path"`
	ID        int         `json:"id"`
	Kind      string      `json:"kind"`
	Name      string      `json:"name"`
	ParentID  interface{} `json:"parent_id"`
	Path      string      `json:"path"`
	WebURL    string      `json:"web_url"`
}

type Project struct {
	AvatarURL         interface{}   `json:"avatar_url"`
	CreatedAt         string        `json:"created_at"`
	DefaultBranch     string        `json:"default_branch"`
	Description       string        `json:"description"`
	ForksCount        int           `json:"forks_count"`
	HTTPURLToRepo     string        `json:"http_url_to_repo"`
	ID                int           `json:"id"`
	LastActivityAt    string        `json:"last_activity_at"`
	Name              string        `json:"name"`
	NameWithNamespace string        `json:"name_with_namespace"`
	Namespace         *Namespace    `json:"namespace"`
	Path              string        `json:"path"`
	PathWithNamespace string        `json:"path_with_namespace"`
	ReadmeURL         interface{}   `json:"readme_url"`
	SSHURLToRepo      string        `json:"ssh_url_to_repo"`
	StarCount         int           `json:"star_count"`
	TagList           []interface{} `json:"tag_list"`
	WebURL            string        `json:"web_url"`
}
