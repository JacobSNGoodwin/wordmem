FROM golang:1.15.1-alpine3.12 as builder

WORKDIR /go/src/app

# Get Reflex for live reload in dev
ENV GO111MODULE=on
RUN go get github.com/cespare/reflex

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./run .

# Rerun app on changes to any go file
# CMD reflex -r "\.go$" -s -- sh -c "go run ./"

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

#Copy executable from builder
COPY --from=builder /go/src/app/run .

EXPOSE 8080
CMD ["./run"]
