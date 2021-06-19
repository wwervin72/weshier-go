package model

// ArticleAnnexModel article annex model
type ArticleAnnexModel struct {
	BaseModel
	AnnexID   uint64       `json:"annexId" gorm:"no null;column:annex_id"`
	Annex     AnnexModel   `json:"annex" gorm:"foreignKey:AnnexID"`
	ArticleID uint64       `json:"articleId" gorm:"no null;column:article_id"`
	Article   ArticleModel `json:"article" gorm:"foreignKey:ArticleID"`
}

func (aam *ArticleAnnexModel) TableName() string {
	return "ws_article_annex"
}

func (aam *ArticleAnnexModel) Create() error {
	return DB.Self.Model(aam).Create(aam).Error
}
