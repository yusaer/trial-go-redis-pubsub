# Trial go and redis pubsub

Repository for validating go and redis pubsub.

## Infrastructure

- Go
- Redis

## Use Packages

- [echo](https://github.com/labstack/echo)
- [go-redis](https://github.com/go-redis/redis)
- [air](https://github.com/cosmtrek/air)

## Starting up

```
$ docker-compose build
$ docker-compose up -d
```

## Redis

```
$ docker exec -it [CONTAINER ID] /bin/bash
root@46f68f517bf1:/data# redis-cli
127.0.0.1:6379> keys *
```