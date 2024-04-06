GO_AIR_VERSION=latest
GO_STATICCHECK_VERSION=latest

# go
go/install/tools:
	go install github.com/cosmtrek/air@$(GO_AIR_VERSION) &&\
	go install honnef.co/go/tools/cmd/staticcheck@$(GO_STATICCHECK_VERSION)

# staticcheck
go/lint:
	staticcheck ./...

# ent
ent/gen:
	go run -mod=mod entgo.io/ent/cmd/ent generate --template glob="./internal/pkg/ent/template/*.tmpl" ./internal/pkg/ent/schema

# docker
docker/up:
	docker-compose --env-file ./cmd/app/.env up -d --build
docker/up/db:
	docker-compose --env-file ./cmd/app/.env up -d db --build

# git
git/commit-template:
	cp ./.github/.gitmessage.txt.example ./.github/.gitmessage.txt &&\
    git config commit.template ./.github/.gitmessage.txt &&\
    git config --add commit.cleanup strip

# other
sleep:
	sleep 20
