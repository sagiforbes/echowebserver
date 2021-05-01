DOCKER_IMAGE=sagiforbes/echowebserver
OUTFILE=echo-webserver
clean:
	$(shell [ -f $(OUTFILE) ] && rm $(OUTFILE))
build:
	go build -o $(OUTFILE)
rebuild: clean build

docker: rebuild
	docker build -t ${DOCKER_IMAGE}:latest .

docker-push: docker
	docker push ${DOCKER_IMAGE}:latest

