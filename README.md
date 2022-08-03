# Trial go and redis pubsub

Repository for validating go and redis pubsub.

## Infrastructure

- [Go](https://go.dev/) v1.18
- [Redis](https://redis.io/) v7

## Use Packages

- [echo](https://echo.labstack.com/) v4
- [go-redis](https://redis.uptrace.dev/) v9
- [air](https://github.com/cosmtrek/air) v1.40.4

## Starting up

```
$ docker-compose build
$ docker-compose up -d
```

## Operation 

First, look at the docker-compose logs.

```
docker-compose logs -f
```

Next, start each sub-server.

```
# -- console1 --
$ curl -X POST -H "Content-Type: application/json" -d '{"channels":["publish-user1"]}' localhost:8091/subscribe
# -- console2 --
$ curl -X POST -H "Content-Type: application/json" -d '{"channels":["publish-user2"]}' localhost:8092/subscribe
# -- console3 --
# multiple designations possible.
$ curl -X POST -H "Content-Type: application/json" -d '{"channels":["publish-user1", "publish-user2"]}' localhost:8093/subscribe
```

Once subscribed, all that is left to do is to execute publish.

```
$ curl -X POST -H "Content-Type: application/json" -d '{"channel": "publish-user1", "user": {"name":"yusaer", "email":"yusaer@example.com"}}' localhost:8081/publish
```

The following message is output to the docker-compose log.

```
trial-go-redis-pubsub-sub1-1   | {Name:yusaer Email:yusaer@example.com}
trial-go-redis-pubsub-sub2-1   | {Name:yusaer Email:yusaer@example.com}
trial-go-redis-pubsub-sub3-1   | {Name:yusaer Email:yusaer@example.com}
```

## Redis

How to operate Redis.

```
$ docker exec -it [CONTAINER ID] /bin/bash
root@46f68f517bf1:/data# redis-cli
127.0.0.1:6379> keys *
```
