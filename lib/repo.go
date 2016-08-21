package repo

type Repo struct {
	Name          string `json:"name"`
	Description   string `json:"description"`
	Homepage      string `json:"homepage"`
	Private       bool   `json:"private"`
	Has_issues    bool   `json:"has_issues"`
	Has_wiki      bool   `json:"has_wiki"`
	Has_downloads bool   `json:"has_downloads"`
}

// gotta create a constructor

// repo := r.Repo{Username: username} as an example

func CreateDefaultRepo(name string) Repo {

	r := Repo{}
	r.Name = name
	r.Description = "Repo for " + name
	r.Homepage = "https://github.com"
	r.Private = true
	r.Has_issues = true
	r.Has_wiki = false
	r.Has_downloads = true

	return r
}
