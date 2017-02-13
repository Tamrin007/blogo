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
