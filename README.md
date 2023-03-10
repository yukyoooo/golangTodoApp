# golangTodoApp

## 実行関数
`go run main.go`
`go build -o main main.go``

Linux環境を指定してコンパイル
`GOOS=linux GOARCH=amd64 go build -o main main.go`

マルチプラットフォーム：コンパイルすることで異なる環境でも同一に動作させることが可能

### Goコマンド
- `go fmt` ソースコードの整形
- `go build` プログラムのビルド
- `go install パッケージ名` パッケージや実行ファイルをビルドした結果をGOPATH内にインストールする
- `go get パッケージ` 外部パッケージのダウンロードとインストールをまとめて実行する
- `go test パッケージ名` テストの実行
- `go test -v ./...` すべてのパッケージのテスト実行
    - -v 詳細
    - -cover テストのカバー率


### コマンド
- `sqlite3 example.sql` SQLモード

- `postgres -D /usr/local/var/postgres` postgresSQLサーバの起動
- `psql -l` データベース一覧を取得


## ディレクトリ構成
- app/models/
    - base.go DB関連の共通処理
    - users.go ユーザーのCRUD
- config/
    - config.go 初期設定(config, logの読込み)
- utils/
    - logging.go log機能追加
- config.ini
- main.go
- webapp.log 
- webapp.sql

## sqlite3コマンド
- `.talbe`
- `.exit`