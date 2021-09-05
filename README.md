
### Rerferences and useful tips
1. Go Version Manager
    1. `bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)`
    1. gvm listall
    1. gvm install go1.16.2
gvm use go1.16.2 [--default]
    1. https://jimkang.medium.com/install-go-on-mac-with-homebrew-5fa421fc55f5
1. `go mod tidy`
2. Secure golang container
    1. secure, small golang container: https://medium.com/@chemidy/create-the-smallest-and-secured-golang-docker-image-based-on-scratch-4752223b7324
    1. docker-slim command


## Running locally
`go build` \
`./cdbapi`

## godocs
https://pkg.go.dev/golang.org/x/tools/cmd/godoc?tab=overview

# godoc -http=:6060 & wget -e robots=off -r -np -N -E -p -k 

go get -u github.com/swaggo/http-swagger
go get github.com/swaggo/swag/cmd/swag

## local dev

Create `api.json` like so (change any values you need to...):
```js
{
  "user": "pg_user",
  "host": "ip_address",
  "database": "db_name",
  "password": "password",
  "maxConnections": 10,
  "connectionTimeoutMillis": 30000,
  "idleTimeoutMillis": 5000,
  "port": 5432
}
```

```
docker build -f go.Dockerfile -t go_test . && docker rm -f go_test || true && docker run -d -p 8443:8443 -v $PWD/api.json:/usr/go/api.json --name go_test go_test
```

### Swagger Docs
Follow instructions here: https://github.com/swaggo/echo-swagger
go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/echo-swagger
swag init

go env

copy files to docs/

You can then view the docs at `${URL}/swagger/index.html`
