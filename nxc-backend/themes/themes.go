package themes

type Themes struct {
	Id   int    `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Name string `json:"name"`
}

func (Themes) TableName() string {
	return "x_themes"
}
