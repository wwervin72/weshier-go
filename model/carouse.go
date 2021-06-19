package model

// CarouseModel carousel model
type CarouseModel struct {
	BaseModel
	Name   string `zh:"名称" json:"name" gorm:"column:name"`
	Jump   string `zh:"跳转" json:"jump" gorm:"column:jump"`
	Desc   string `zh:"描述" json:"desc" gorm:"column:desc"`
	Title  string `zh:"标题" json:"title" gorm:"column:title"`
	Img    string `zh:"图片地址" json:"url" validate:"-" gorm:"img"`
	Status uint   `zh:"状态" json:"status" gorm:"status"`
}

// TableName carousel table name
func (cm *CarouseModel) TableName() string {
	return "ws_carousel"
}

// Create create carousel
func (cm *CarouseModel) Create() error {
	return DB.Self.Model(cm).Create(&cm).Error
}
