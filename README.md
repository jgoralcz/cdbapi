go build
./go_cdbapi

Building a secure golang container: https://medium.com/@chemidy/create-the-smallest-and-secured-golang-docker-image-based-on-scratch-4752223b7324

go mod tidy


## godocs
https://pkg.go.dev/golang.org/x/tools/cmd/godoc?tab=overview

# godoc -http=:6060 & wget -e robots=off -r -np -N -E -p -k 

go get -u github.com/swaggo/http-swagger
go get github.com/swaggo/swag/cmd/swag


Use Chi or Mux for a better handler that isn't used purely for performance. Negroni is also a good additional middleware option.
