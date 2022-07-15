FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /fees


COPY ./go.mod ./go.sum ./
RUN go mod download

COPY ./ ./

WORKDIR /fees/cmd

RUN GOOS=linux GOARCH=amd64  CGO_ENABLED=0 go build -ldflags="-w -s" -installsuffix 'static' -o feesService

FROM scratch

COPY --from=builder fees/cmd/feesService feesService

EXPOSE 8081
EXPOSE 8080

CMD ["/feesService", "--port", "8081", "--gateway_port", "8080"]