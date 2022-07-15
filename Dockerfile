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

ENV HOST=database
ENV PORT=5432
ENV USER=test
ENV PASSWORD=test
ENV DBNAME=eth

CMD ["/feesService", "--port", "8081", "--gateway_port", "8080"]