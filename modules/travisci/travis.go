package travisci

type Builds struct {
	Builds []Build `json:"builds"`
}

type Build struct {
	ID         int        `json:"id"`
	CreatedBy  Owner      `json:"created_by"`
	Branch     Branch     `json:"branch"`
	Number     string     `json:"number"`
	Repository Repository `json:"repository"`
	Commit     Commit     `json:"commit"`
	State      string     `json:"state"`
}

type Owner struct {
	Login string `json:"login"`
}

type Branch struct {
	Name string `json:"name"`
}

type Repository struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type Commit struct {
	Message string `json:"message"`
}
