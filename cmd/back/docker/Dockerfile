FROM alpine:3.20.3 AS builder

RUN apk update && apk add --no-cache ca-certificates

FROM scratch

WORKDIR /bin

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY /bin/back back
COPY /cmd/back/migrate migrate/

ENTRYPOINT ["back"]
