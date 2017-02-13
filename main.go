package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"

	"github.com/Tamrin007/blogo/db"

	_ "github.com/go-sql-driver/mysql"
)

var (
	createFlag *flag.FlagSet
	readFlag   *flag.FlagSet
	updateFlag *flag.FlagSet
	deleteFlag *flag.FlagSet

	id    int
	title string
	body  string
)

func init() {
	createFlag = flag.NewFlagSet("create", flag.ExitOnError)
	createFlag.StringVar(&title, "title", "", "記事のタイトル")
	createFlag.StringVar(&body, "body", "", "記事の本文")

	readFlag = flag.NewFlagSet("read", flag.ExitOnError)
	readFlag.IntVar(&id, "id", 0, "記事 id")

	updateFlag = flag.NewFlagSet("update", flag.ExitOnError)
	updateFlag.IntVar(&id, "id", 0, "記事 id")
	updateFlag.StringVar(&title, "Title", "", "記事のタイトル")
	updateFlag.StringVar(&body, "Body", "", "記事の本文")

	deleteFlag = flag.NewFlagSet("delete", flag.ExitOnError)
	deleteFlag.IntVar(&id, "id", 0, "記事 id")
}

func run(db *sql.DB) error {
	switch os.Args[1] {
	case "create":
		createFlag.Parse(os.Args[2:])
		if title == "" {
			return fmt.Errorf("The title must not be empty")
		}
		if body == "" {
			return fmt.Errorf("The body must not be empty")
		}
		return article.InsertArticle(db, title, body)

	case "read":
		readFlag.Parse(os.Args[2:])
		if id == 0 {
			return fmt.Errorf("The id must not be empty")
		}
		article, err := article.ScanArticle(db, id)
		fmt.Println(article.Title)

		if err != nil {
			return err
		}

	case "update":
		updateFlag.Parse(os.Args[2:])
		if title == "" {
			return fmt.Errorf("The title must not be empty")
		}
		if body == "" {
			return fmt.Errorf("The body must not be empty")
		}
		if id == 0 {
			return fmt.Errorf("The id must not be empty")
		}
		return article.UpdateArticle(db, id, title, body)

	case "delete":
		deleteFlag.Parse(os.Args[2:])
		if id == 0 {
			return fmt.Errorf("The id must not be empty")
		}
		return article.DeleteArticle(db, id)
	}

	return nil
}

func main() {
	datasource := `root:mysql@tcp(localhost:3306)/go?parseTime=true&collation=utf8_general_ci&interpolateParams=true`
	// datasource := os.Getenv("DSN")

	db, err := sql.Open("mysql", datasource)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = run(db)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		os.Exit(1)
	}
}
