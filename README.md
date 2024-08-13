# nginx-remote-signal
Send nginx signals via http request

# Usage

Default port is 5000.
GET `/signal/<reload|start|stop>` to send the appropriate signal to nginx.

Set environement variable `NRS_PORT` to change the port.

See [example docker-cmpose file](./docs/docker-compose.example.yml) with certbot and nginx.