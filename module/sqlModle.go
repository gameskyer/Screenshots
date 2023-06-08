package module

type Screenshot struct {
	Id       int    `json:"id" db:"id" gorm:"column:id"`
	Path     string `json:"path" db:"path"  gorm:"column:path"`
	Url      string `json:"url" db:"url"  gorm:"column:url"`
	Name     string `json:"name" db:"name"  gorm:"column:name"`
	Scene    string `json:"scene" db:"scene"  gorm:"column:scene"`
	Domain   string `json:"domain" db:"domain"  gorm:"column:domain"`
	FileName string `json:"fileName" db:"fileName" gorm:"column:fileName"`
}

func (Screenshot) TableName() string {
	return "screenshot"
}
