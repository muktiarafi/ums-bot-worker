FROM golang:alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /dist

COPY go.mod .
COPY go.sum .
RUN go mod download
RUN go mod verify
COPY . .
RUN go build main.go

FROM chromedp/headless-shell:latest

RUN apt-get update
RUN apt-get install tini -y

WORKDIR /app
COPY --from=builder /dist/main .

ENTRYPOINT ["tini", "--"]

CMD [ "./main" ]