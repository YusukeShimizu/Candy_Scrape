This is the secretary of yotts.
Check keyakinet and notify useful messages using redis and line.



## Getting Started

#### Set envs.

```go
Env      string `required:"true" envconfig:"ENV"`
Pace     string `required:"true" envconfig:"PACE"`
Secret   string `required:"true" envconfig:"SECRET"`
Token    string `required:"true" envconfig:"TOKEN"`
ID       string `required:"true" envconfig:"ID"`
PUBLICID string `required:"true" envconfig:"ID"`
PORT     string `required:"true" envconfig:"PORT"`
```

#### Run on local.

```sh
docker-compose up
docker run candy-scrape go run main.go
```

#### Deploy

At local
```sh
docker build -t xxx/xxx .
docker push xxx/xxx .
```
At server
```sh
docker pull xxx/xxx
docker run --name redis -d redis redis-server --appendonly yes
docker run -d -p port:port --restart=always --env-file=env.txt --link <redisid>:redis <appid> go run main.go
```

## TODO
1. Book system.
