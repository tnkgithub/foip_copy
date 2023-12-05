# Build Container
FROM golang:1.19-alpine AS build-env
WORKDIR /go/src/
COPY ./ ./
RUN apk --no-cache add make \
  && go install github.com/swaggo/swag/cmd/swag@latest \
  && go get -u github.com/swaggo/swag/cmd/swag@latest \
  && make

# Running Container
FROM alpine:latest
RUN apk add --no-cache --update ca-certificates
COPY --from=build-env /go/src/bin/server /usr/local/bin/server
COPY --from=build-env /go/src/config/config.yaml /etc/server/config.yaml

EXPOSE 8080
CMD ["/usr/local/bin/server", "--config-path", "/etc/server/config.yaml"]