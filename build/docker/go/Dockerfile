FROM golang:1.22.1-alpine
WORKDIR /src

COPY ../../../Makefile .

RUN apk update \
    && apk add --no-cache make

RUN make go/install/tools

CMD ["air", "-c", "./cmd/app/air.toml"]
