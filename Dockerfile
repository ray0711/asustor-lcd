FROM golang:1.22.1-alpine3.18 as golang

# RUN apk add -U tzdata
# RUN apk --update add ca-certificates

WORKDIR /app
COPY . .

RUN go mod download
RUN go mod verify

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /lcd .

FROM scratch

# COPY --from=golang /usr/share/zoneinfo /usr/share/zoneinfo
# COPY --from=golang /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# COPY --from=golang /etc/passwd /etc/passwd
# COPY --from=golang /etc/group /etc/group

COPY --from=golang /lcd .

# EXPOSE 8080

CMD ["/lcd"]