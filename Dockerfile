# Server Builder
FROM golang:alpine AS go-builder

RUN adduser -D -g '' appuser

RUN mkdir /switchcraft
COPY . /switchcraft/

WORKDIR /switchcraft/

RUN go mod tidy

RUN  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o switchcraft -ldflags="-w -s" main.go

FROM alpine:latest AS ca-certs

RUN apk add -U --no-cache ca-certificates

# Final Image
FROM scratch

COPY --from=go-builder /switchcraft/switchcraft /switchcraft
COPY --from=ca-certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 8080

ENTRYPOINT [ "/switchcraft" ]
