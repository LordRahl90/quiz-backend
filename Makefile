GO=go

start:
	${GO} run main.go

build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o .docker/main .
	docker-compose up --build

kill:
	docker-compose kill

stack:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o .docker/main .
	# docker swarm leave
	# docker swarm init
	docker build -f .docker/Dockerfile -t lordrahl/little-quiz .
	docker stack deploy -c docker-compose.yml little-quiz