FROM golang:1.23.3-bookworm AS build
WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /opt/go-docker

FROM alpine:latest
WORKDIR /app
COPY --from=build /opt/go-docker /opt/go-docker
EXPOSE 8080
ENTRYPOINT ["/opt/go-docker"]