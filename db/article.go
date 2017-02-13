package article

import (
	"database/sql"
	"time"
)

// Article structure.
type Article struct {
	ID        int
	Title     string
	Body      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

// InsertArticle inserts article to DB.
func InsertArticle(db *sql.DB, title string, body string) error {
	_, err := db.Exec("INSERT INTO articles (title, body) VALUES (?, ?)", title, body)
	return err
}

// ScanArticle scans article from DB.
func ScanArticle(db *sql.DB, id int) (Article, error) {
	var article Article
	err := db.QueryRow("SELECT * FROM articles WHERE id = ? LIMIT 1", id).Scan(&article.ID, &article.Title, &article.Body, &article.CreatedAt, &article.UpdatedAt)
	return article, err
}

// UpdateArticle updates article.
func UpdateArticle(db *sql.DB, id int, title string, body string) error {
	_, err := db.Exec("UPDATE articles SET title = ?, body = ? WHERE id=?", title, body, id)
	return err
}

// DeleteArticle deletes article from DB.
func DeleteArticle(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM articles WHERE id=?", id)
	return err
}
