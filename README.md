# nginx-remote-reload
Reload nginx via network request from other service

# usage


```yml
services:
  nginx:
    image: merofuruya/nginx-remote-reload:latest
    volumes:
      - ./nginx.conf:/etc/nginx/nginx.conf
      - ./nginx.conf.d:/etc/nginx/conf.d
      - ./certbot/etc:/etc/letsencrypt
      - ./certbot/var:/var/lib/letsencrypt
    ports:
      - 80:80
    networks:
      - nginx
  
  certbot:
    build:
      context: .
      dockerfile_inline: |
        FROM certbot/certbot
        RUN apk --no-cache add wget
    volumes:
      - ./certbot/etc:/etc/letsencrypt
      - ./certbot/var:/var/lib/letsencrypt
    networks:
      - nginx
    command: certonly --webroot --webroot-path=/var/www/html --email --post-hook "wget -O - http://nginx:5000/signal/reload" -d example.com
```