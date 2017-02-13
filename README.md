# Go ハンズオン

## 簡易 CLI ブログ

- 動作例

```text
./blogo create --title "Go blog CLI" --body "Sample text"

./blogo read --id 1
Title: Go blog CLI
Body: Sample text

./blogo update --id 1 --title "Go Blog CLI" --body "Sample text."

./blogo read --id 1
Title: Go Blog CLI
Body: Sample text.

./blogo delete --id 1

./blogo read --id 1
Error: sql: no rows in result set
```

### Step1. Hello, world.

リポジトリを `go get` して、 `Hello, world.` を確認して下さい。

- go get

```
go get github.com/Tamrin007/blogo
blogo
Hello, world.
```

### Step2. Parse flag

`flag` パッケージを用いて、サブコマンド・オプションを作っていきます。

```go
package main

import (
	"flag"
	"fmt"
)

var (
  createFlag *flag.FlagSet

	id    int
	title string
	body  string
)

func init() {
  createFlag = flag.NewFlagSet("create", flag.ExitOnError)
	createFlag.StringVar(&title, "title", "", "記事のタイトル")
}

func run() error {
  switch os.Args[1] {
	case "create":
		createFlag.Parse(os.Args[2:])
		if title == "" {
			return fmt.Errorf("The title must not be empty")
		}
    fmt.Println("title: ", title)
  }

  return nil
}

func main() {
  err = run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		os.Exit(1)
	}
}
```

`go run main.go` で実行し、確かめてみましょう。

#### ex.

- `body` フラグを追加しましょう。
- `read`, `update`, `delete` もそれぞれ書いてみましょう。
  - `read`: id
  - `update`: id, title, body
  - `delete`: id がそれぞれ必要です。

### Step3. DB の準備

Docker の mysql コンテナを立ち上げ、 bash にログインします。

```sh
docker pull mysql
docker run --name go-seminar -e MYSQL_ROOT_PASSWORD=< お好きなパスワード > -d -p 3306:3306 mysql
docker exec -it go-seminar /bin/bash
```

MySQL にログインします

```sh
mysql -u root -p
```

テーブルを作成しておきます。

```sql
CREATE TABLE `articles` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `title` varchar(255) NOT NULL COMMENT 'title',
  `body` varchar(255) NOT NULL COMMENT 'body',
  `created` timestamp NOT NULL DEFAULT NOW() COMMENT 'when created',
  `updated` timestamp NOT NULL DEFAULT NOW() ON UPDATE CURRENT_TIMESTAMP COMMENT 'when last updated',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='list of articles';
```

ここでレコードを入れておいてもいいでしょう。

### Step4. Go での DB 接続

"database/sql" と "github.com/go-sql-driver/mysql" をインポートすることで、 MySQL への接続が可能となります。

```go
import (
  ...
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)
```

`main()` の中で

```go
datasource := `root:< 設定したパスワード >@tcp(localhost:3306)/go?parseTime=true&collation=utf8_general_ci&interpolateParams=true`
```

を定義しましょう。

#### ex.

- "database/sql" パッケージのドキュメントを読み、 MySQL に接続してみましょう。
  - 必ずエラーチェックして下さい。
  - defer 文について調べて下さい。また、 defer 文を用いて接続を閉じる処理を書いて下さい。

### Step5. DB 操作用のパッケージを作成する

`db/` ディレクトリを用意し、以下の内容の `article.go` を作成しましょう。

```go
package article

import (
  "database/sql"
  "time"
)

type Article struct {
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

// InsertArticle inserts article to DB.
func InsertArticle(db *sql.DB, title string, body string) {

}
```

- `Article` 構造体に `ID`, `Title`, `Body` フィールドを追加して下さい。
- DB にレコードを追加する `InsertArticle()` を完成させて下さい。
- `main` パッケージで "github.com/Tamrin007/blogo/db" をインポートして下さい。
- `run()` 内で create 処理の際に `InsertArticle` を呼び出して下さい。
  - レコードが挿入されたか MySQL で確認して下さい。
- `ReadArticle()`, `UpdateArticle()`, `DeleteArticle()` もそれぞれ作成して下さい。
  - `ReadArticle()` は `Article` 構造体を返しましょう。
  - どの関数もエラーは必ず返すようにしましょう。
  - 作成した関数を read, update, delete 処理の際に呼び出し、結果を出力して下さい。
