# go fiber onion architecture
go version 1.22.1

```text
root
├── cmd: コマンドラインツール
│   ├── cli
│   │   ├── db
│   │   └── task
│   └── app
│
├── internal:
│   ├── interactor: ユースケースを操作するロジック
│   ├── presenter: presentation layer
│   ├── infrastructure: infrastructure
│   ├── application: application layer
│   │   ├── dto: data transfer object - 外部アプリまたはレイヤー間でのデータ移送のため
│   │   └── usecase: usecase
│   │
│   ├── domain: domain layer
│   │   ├── entities: entity
│   │   ├── value-objects: 値オブジェクト(不使用)
│   │   └── services: interface
│   │
│   └── pkg: プロジェクトの共有コンポーネント(このプロジェクト固有)
│
├── api: openapi等
├── build: パッケージングと継続的インテグレーション(dockerfile等)
├── configs: 設定ファイル
├── docs: ドキュメント(api docは除く)
├── pkg: プロジェクトの共有コンポーネント(他のプロジェクトでも利用可)
└── util: 言語特有のutil
```

# go run
## docker up && exec
```shell
$ docker-compose up -d --build
$ docker-compose exec -it go ash
```

## in container
### mod download and server run
```shell
$ go mod tidy
$ go run ./cmd/app/main.go
```

### for air server run
```shell
$ go mod tidy
$ make go/install/tools
$ air -c ./cmd/app/air.toml
```

### access
http://localhost:8080

### server stop
```shell
$ docker-compose down
```

# go lint
```shell
$ make go/lint
```

# docs
- [git commit rule](./docs/markdown/git/commit.md)
- [git branch rule](./docs/markdown/git/branch.md)
- [git release-drafter document](./docs/markdown/git/release-drafter.md)
