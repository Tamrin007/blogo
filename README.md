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

- go get

```
go get github.com/Tamrin007/blogo
blogo
Hello, world.
```
