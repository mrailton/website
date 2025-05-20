package data

import (
	"database/sql"
	"time"
)

type Article struct {
	ID          int64
	Title       string
	Slug        string
	Content     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	PublishedAt time.Time
	Version     int
}

type ArticleModel struct {
	DB *sql.DB
}

func (m *ArticleModel) Latest() []Article {
	query := `
		SELECT slug, title, published_at
		FROM articles
		WHERE published_at < NOW()
		ORDER BY published_at DESC
		LIMIT 3;
	`

	rows, err := m.DB.Query(query)
	if err != nil {
		return nil
	}

	defer rows.Close()

	var articles []Article

	for rows.Next() {
		var article Article

		err := rows.Scan(&article.Slug, &article.Title, &article.PublishedAt)

		if err != nil {
			return nil
		}

		articles = append(articles, article)
	}

	return articles
}
