FROM golang:1.17-alpine3.15

RUN apk update && apk upgrade && apk add --update alpine-sdk protoc git &&\
  go install github.com/go-delve/delve/cmd/dlv@latest

WORKDIR /docker_dlv
COPY . .
RUN go build -gcflags "all=-N -l" -o /bin/srv ./app/...
ENTRYPOINT [ "/go/bin/dlv", "exec", "/bin/srv", "--listen=:2345", "--headless",  "--api-version=2", "--continue", "--accept-multiclient" ]
