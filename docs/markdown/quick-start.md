# Quick Start
最終更新日: 2024/3/20  

1. dockerでgoを管理する場合
2. localでgoを管理する場合(コンテナ内実行をlocalで実行するのみ)

# 1. dockerバージョン
## コンテナ起動とコンテナ実行
```shell
$ docker-compose up -d --build
$ docker-compose exec -it go ash
```

## コンテナ内で実行
### ツールのインストールとAirサーバー起動
```shell
$ go mod tidy
$ make go/install/tools
$ air -c ./cmd/app/air.toml
```

### アクセスURL
http://localhost:8080

## コンテナ外で実行

```shell
$ docker-compose down
```
