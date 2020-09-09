FROM golang:1.15.1-alpine3.12

# may use make later
RUN apk --no-cache --no-progress add git ca-certificates make \
  && update-ca-certificates \
  && rm -rf /var/cache/apk/*

WORKDIR /go/src/app

# Get Reflex for live reload in dev
RUN go get github.com/cespare/reflex

COPY go.mod .
COPY go.sum .

RUN GO111MODULE=on GOPROXY=https://proxy.golang.org go mod download

COPY . .

# Rerun app on changes to any go file
ENTRYPOINT reflex -r "\.go$" -s -- sh -c "go run ./"

EXPOSE 8080