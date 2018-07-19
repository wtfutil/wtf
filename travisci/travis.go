package travisci

type Builds struct {
	Builds []Build `json:"builds"`
}

type Build struct {
	CreatedBy  Owner      `json:"created_by"`
	Branch     Branch     `json:"branch"`
	Number     string     `json:"number"`
	Repository Repository `json:"repository"`
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
}
