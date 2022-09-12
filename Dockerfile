FROM golang:1.17-alpine AS build
WORKDIR /sept
COPY . .
RUN    go env -w CGO_ENABLED=0 \
    && go build -v \
    && go test ./... -cover -v

FROM alpine:latest
RUN apk add --no-cache tzdata
ENTRYPOINT [ "/sept" ]
COPY --from=build /sept/sept /sept
