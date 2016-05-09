default: builddocker

setup:

buildgo:
	CGO_ENABLED=0 GOOS=linux go build -ldflags "-s" -a -installsuffix cgo -o echo-server

builddocker:
	docker build -t echo-build -f Dockerfile.build .

	docker run -t echo-build /bin/true
	docker cp `docker ps -q -n=1`:/echo-server .
	chmod 755 ./echo-server
	docker build --rm=true -t echo-server -f Dockerfile.scratch .

run: builddocker
	docker run -p 8080:8080 echo-server
