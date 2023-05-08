package module

type Router struct {
	Path          string          `json:"path"`
	Name          string          `json:"name"`
	ChildrenRoute []ChildrenRoute `json:"children"`
}

type ChildrenRoute struct {
	Path   string `json:"path" db:"path"`
	Url    string `json:"url" db:"url"`
	Name   string `json:"name" db:"name"`
	Scene  string `json:"scene" db:"scene"`
	Domain string `json:"domain" db:"domain"`
}

type Query struct {
	Path   string `json:"path" db:"path"`
	Url    string `json:"url" db:"url"`
	Name   string `json:"name" db:"name"`
	Scene  string `json:"scene" db:"scene"`
	Domain string `json:"domain" db:"domain"`
}
