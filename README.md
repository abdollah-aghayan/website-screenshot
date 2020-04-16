## Website Screenshot

Docker-powered stateless Go API for taking screenshot from provided website

### Installation

Before anything you need to have Docker and Golang installed

You can build a single container with following command 

```sh
make all
```
In order to increase the timeout do it as following

```sh
make all timeout="-e TIMEOUT=60"
```

### Usage 

Application will work on port 8080 so you can send a request following url

http://127.0.0.1:8080

example 

    http://127.0.0.1:8080/screenshot?url=http://google.com

### Scaling 

(this part not tested yet)

To scale the application you can use docker-compose scaling feature 
    
docker-compose file

```code

version: '3'

services:

  screenshot:
    image: screenshot:1.0
    environment:
      - TIMEOUT:30

  nginx:
    image: nginx:stable-alpine
    volumes:
      - ./nginx.conf:/etc/nginx/nginx.conf
    depends_on:
      - screenshot
    ports:
      - "8080:8080"

```

nginx config 

```code
user  nginx;

events {
    worker_connections   1000;
}
http {
        server {
              listen 8080;
              location / {
                proxy_pass http://localhost:8080/;
              }
        }
}
```

then you can do

```sh
make scale instance=the_number_of_instance
```