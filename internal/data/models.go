package data

import "database/sql"

type Models struct {
	Articles ArticleModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Articles: ArticleModel{DB: db},
	}
}
