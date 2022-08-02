FROM golang:1.18-alpine

ENV TZ=Asia/Tokyo
ENV GOPATH /go
ENV GO111MODULE on
ENV ROOT=/go/src

RUN apk update && \
    apk --no-cache add git && \
    apk --no-cache add tzdata

WORKDIR ${ROOT}
COPY ./pub ${ROOT}
EXPOSE 8080

RUN go install github.com/cosmtrek/air@latest
CMD ["air", "-c", ".air.toml"]
