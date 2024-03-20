GO_AIR_VERSION=latest
GO_STATICCHECK_VERSION=latest

# go
go/install/tools:
	go install github.com/cosmtrek/air@$(GO_AIR_VERSION) &&\
	go install honnef.co/go/tools/cmd/staticcheck@$(GO_STATICCHECK_VERSION)

# staticcheck
go/lint:
	staticcheck ./...

# git
git/commit-template:
	cp ./.github/.gitmessage.txt.example ./.github/.gitmessage.txt &&\
    git config commit.template ./.github/.gitmessage.txt &&\
    git config --add commit.cleanup strip

# other
sleep:
	sleep 20
