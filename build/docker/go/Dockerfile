FROM golang:1.22.1-alpine
WORKDIR /src

COPY ../../../Makefile .

RUN apk update \
    && apk add --no-cache make

RUN make go/install/tools
RUN make ent/gen
RUN make oapi/codegen/app
RUN go run ./cmd/cli/db/main.go -query migrate
RUN go run ./cmd/cli/db/main.go -query seed

CMD ["air", "-c", "./cmd/app/air.toml"]
