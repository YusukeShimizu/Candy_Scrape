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

# Getting started
Build image.
```sh
$ docker build . -t gcr.io/<PROJECT_NAME>/candy_scrape:latest --no-cache
```

Push image.
```sh
docker push gcr.io/<PROJECT_NAME>/candy_scrape:latest
```

Set .env as secret.

```
cp .env.sample .env
kubectl create secret generic candy-scrape-secret --from-env-file=.env
```

Apply yamls.

```sh
$ sed -i  "bak" "s/your-project-ID/<PROJECT_NAME>/g" StatefulSet.yaml
$ kubectl apply -f StatefulSet.yaml --record
```
Wait 5 minutes☕️ and then login to container.