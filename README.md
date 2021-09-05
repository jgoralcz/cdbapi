# Character API (Bongo bot)
This is a character API based of characters submitted from various users in my [Discord Channel](https://discord.gg/dfajqcZ). You will be heavily rate limited in the beta version, so you can test it out, but it is mainly for [patrons](https://www.patreon.com/bongobot) and my bot so the server won't be overloaded.

You can use the live instance at `beta-${URL}/v1`

Docs are available at: `${URL}/swagger/index.html`

Side Note: Images, characters, and series are updated every day. We could use your help. Please join the official Bongo Bot support server for help: https://discord.gg/dfajqcZ

<br/>
<br/>
<br/>

### local development
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

<br/>

### Swagger Docs
Follow instructions here: https://github.com/swaggo/echo-swagger
1. `go get -u github.com/swaggo/swag/cmd/swag`
1. `go get -u github.com/swaggo/echo-swagger`
1. `swag init`

You can then view the docs at `${URL}/swagger/index.html`

<br>

### godocs
If you want godocs on an api, clone the repo and run like so:

1. `go get golang.org/x/tools/cmd/godoc`
1. `godoc -http=:6060`

<br/>
<br/>
<br/>
<br/>

### References and useful tips
1. Go Version Manager
    1. `bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)`
    1. gvm listall
    1. gvm install go1.16.2
    1. gvm use go1.16.2 [--default]
    1. https://jimkang.medium.com/install-go-on-mac-with-homebrew-5fa421fc55f5
1. Secure golang container
    1. secure, small golang container: https://medium.com/@chemidy/create-the-smallest-and-secured-golang-docker-image-based-on-scratch-4752223b7324
    1. docker-slim command
1. updating docs
    1. https://github.com/swaggo/echo-swagger
1. Useful commands for debugging/cleaning
    1. `go env`
    1. `go mod tidy`
1. Interesting way to scrape godocs
    1. `https://pkg.go.dev/golang.org/x/tools/cmd/godoc?tab=overview`
    1. `godoc -http=:6060 & wget -e robots=off -r -np -N -E -p -k`
1. Stress Testing
    1. https://medium.com/@harrietty/load-testing-an-api-with-apache-benchmark-or-jmeter-24cfe39d3a23
