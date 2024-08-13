ARG GO_VERSION=1.22.5
FROM golang:${GO_VERSION}-alpine AS base

RUN --mount=type=cache,target=/go/pkg/mod/ \
  --mount=type=bind,source=go.sum,target=go.sum \
  --mount=type=bind,source=go.mod,target=go.mod \
  go mod download -x

FROM base AS build
RUN --mount=type=cache,target=/go/pkg/mod/ \
  --mount=type=bind,target=. \
  go build -o /nginx-remote-signal

FROM nginx:alpine AS nginx
COPY --from=build /nginx-remote-signal /usr/local/bin/nginx-remote-signal

CMD ["/bin/sh", "-c", "nginx-remote-signal -p 5000", "&", "nginx", "-g", "daemon off;"]

