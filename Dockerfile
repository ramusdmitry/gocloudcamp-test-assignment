FROM golang:1.19-buster as builder

WORKDIR /music-player
COPY . .

WORKDIR /music-player/cmd
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o ../app

FROM alpine:3.15.4
WORKDIR /music-player

ENV PORT="8000"
ENV DBNAME="blog-db"
ENV DBUSERNAME="admin"
ENV DBPASSWORD="admin"
ENV DBHOST="postgres"
ENV DBPORT="5432"
ENV SSLMODE="disable"
ENV SECRET_JWT="test_salt"

EXPOSE 8000

COPY --from=builder /music-player/app .
CMD ["./app"]