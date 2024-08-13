# nginx-remote-signal
Send nginx signals via http request

# Usage

`GET /signal/<reload|start|stop>` to send the appropriate signal to nginx.

Default port is 5000. May be changed by setting the `NRS_PORT` environment variable or `--port` argument.

See [example docker-cmpose file](./docs/docker-compose.example.yml) with certbot and nginx.