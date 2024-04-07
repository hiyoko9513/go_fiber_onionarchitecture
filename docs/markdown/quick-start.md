# Quick Start
最終更新日: 2024/3/30  

1. dockerでgoを管理する場合
2. localでgoを管理する場合

# 共通初期設定
## env file作成
localかdockerはそれぞれの環境に合わせて
```shell
$ cp ./cmd/app/.env.<local|localdocker> ./cmd/app/.env
```

# 1. dockerバージョン

## コンテナ起動
```shell
$ make docker/up
```

## アクセスURL
サインアップのエンドポイント
http://localhost:8080/signup

# 2. localバージョン

## コンテナ起動
```shell
$ make docker/up/db
```

## ツールのインストールとAirサーバー起動
```shell
$ go mod tidy
$ make go/install/tools
$ make ent/gen
$ make oapi/codegen/app
$ go run ./cmd/cli/db/main.go -query migrate
$ go run ./cmd/cli/db/main.go -query seed
$ air -c ./cmd/app/air.toml
```
※ airが利用できない場合は、ターミナル再起動

## アクセスURL
サインアップのエンドポイント
http://localhost:8080/signup


