package module

type ScreenshotModule struct {
	Path  string  `json:"path"  gorm:"column:path"`
	Name  string  `json:"name"  gorm:"column:name"`
	Scene []Scene `json:"children"`
}

type Scene struct {
	Path string `json:"path"  gorm:"column:path"`
	Name string `json:"name" db:"scene"  gorm:"column:name"`
}
