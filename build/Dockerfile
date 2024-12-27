# 使用 golang:1.21.0 作为基础镜像
FROM golang:1.21.0 AS build

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o exporter ./cmd/main.go

FROM alpine:3.20
LABEL authors="zark"
COPY --from=build /app/exporter /
RUN chmod +x /exporter
EXPOSE 2112
CMD ["/exporter"]
