# Quick Start
最終更新日: 2024/3/30  

1. dockerでgoを管理する場合
2. localでgoを管理する場合

# 共通初期設定
## env file作成
localかdockerはそれぞれの環境に合わせて
```shell
$ cp ./cmd/app/.env.<local|docker> ./cmd/app/.env
```

## cognitoコンテナ起動
```shell
$ docker-compose up -d mock-cognito-server --build
$ docker-compose up -d mock-cognito-cli --build
```

## cognito設定
`./build/docker/aws/mock-cognito-cli/output/mock-cognito.properties`
```text
AWS_COGNITO_USER_POOL_ID=xxx
AWS_COGNITO_USER_CLIENT_ID=xxx
```
↓  
`./cmd/app/.env`
```text
AWS_COGNITO_USER_POOL_ID=
AWS_COGNITO_USER_CLIENT_ID=
```
コピーする

# 1. dockerバージョン

## go コンテナ起動
```shell
$ docker-compose up -d go --build
```

## アクセスURL
http://localhost:8080


# 2. localバージョン

## ツールのインストールとAirサーバー起動
```shell
$ go mod tidy
$ make go/install/tools
$ air -c ./cmd/app/air.toml
```
※ airが利用できない場合は、ターミナル再起動

## アクセスURL
http://localhost:8080


