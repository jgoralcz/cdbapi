# watch -n 1 

# characters
ab -r -c 50 -n 500 http://localhost:8443/v1/characters/random
ab -r -c 10 -n 100 http://localhost:8443/v1/characters/?name=Kurisu

# series
ab -r -c 50 -n 500 http://localhost:8443/v1/series/random?limit=100
ab -r -c 10 -n 100 http://localhost:8443/v1/series/?name=Ste

