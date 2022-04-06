FROM golang:1.17-alpine AS build
ENV NAME sept
WORKDIR /$NAME
COPY . .
RUN    go env -w CGO_ENABLED=0 \
    && go env -w GO111MODULE=on
RUN    go build -v
RUN    go test ./... -cover -v

FROM alpine:latest
RUN apk add --no-cache tzdata
CMD [ "/$NAME" ]
COPY --from=build /$NAME/$NAME  /$NAME
