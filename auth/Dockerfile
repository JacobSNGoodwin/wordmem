FROM golang:1.15.1-alpine3.12

# Get Reflex for live reload in dev

WORKDIR /go/src/app
COPY . .

RUN go get github.com/cespare/reflex

EXPOSE 8080

# Rerun app on changes to any go file
ENTRYPOINT reflex -r "\.go$" -s -- sh -c "go run ./"