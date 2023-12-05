FROM golang:1.17.5

EXPOSE 5000:5000

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
ENV PORT 5000
ENV GIN_MODE release
RUN go build -v -o /usr/local/bin/app ./cmd/server/main.go

CMD ["app"]