package module

type ScreenshotModule struct {
	Path  string  `json:"path"`
	Name  string  `json:"name"`
	Scene []Scene `json:"children"`
}

type Scene struct {
	Path string `json:"path"`
	//Sence string 				  `json:"scene" db:"scene"`
	Name string `json:"name" db:"scene"`
}
