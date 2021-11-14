package been


type InsertSql struct {
	Path string `json:"path" db:"path"`
	Url string `json:"url" db:"url"`
	Name string `json:"name" db:"name"`
	Scene string `json:"scene" db:"scene"`
	Domain string `json:"domain" db:"domain"`
}
