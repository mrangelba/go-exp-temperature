run:
	docker-compose -f deployments/docker-compose.yml up

build-docker:
	docker build -f Dockerfile . -t service
