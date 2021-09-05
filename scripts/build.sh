docker build -t go_test . && docker rm -f go_test || true && docker run -d -p 8443:8443 -v $PWD/api.json:/usr/go/api.json --name go_test go_test
