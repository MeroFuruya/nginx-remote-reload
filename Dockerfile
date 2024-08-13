ARG GO_VERSION=1.22.5
FROM golang:${GO_VERSION}-alpine AS build

COPY go.mod go.sum ./
RUN go mod download -x

COPY main.go .
RUN go build -o /nginx-remote-signal

FROM nginx:alpine AS nginx
COPY /nrs-docker-entrypoint.sh /docker-entrypoint.d/10-nrs-docker-entrypoint.sh
RUN chmod +x /docker-entrypoint.d/10-nrs-docker-entrypoint.sh
COPY --from=build /nginx-remote-signal /usr/local/bin/nginx-remote-signal
