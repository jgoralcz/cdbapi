go build
./cdbapi

Building a secure golang container: https://medium.com/@chemidy/create-the-smallest-and-secured-golang-docker-image-based-on-scratch-4752223b7324

go mod tidy


## godocs
https://pkg.go.dev/golang.org/x/tools/cmd/godoc?tab=overview

# godoc -http=:6060 & wget -e robots=off -r -np -N -E -p -k 

go get -u github.com/swaggo/http-swagger
go get github.com/swaggo/swag/cmd/swag


Use Chi or Mux for a better handler that isn't used purely for performance. Negroni is also a good additional middleware option.


## swagger docs
docker build -f Dockerfilego -t go_test . && docker rm -f go_test || true && docker run -d -p 8443:8443 -v $PWD/api.json:/Users/Josh/Documents/GitHub/cdbapi/api.json --name go_test go_test

Follow instructions here: https://github.com/swaggo/gin-swagger
go get -u github.com/swaggo/swag/cmd/swag
swag init

copy files to docs/
