DOCKER_COMPOSE_INMEMORY := docker-compose.inmemory.yml

DOCKER_COMPOSE_POSTGRES := docker-compose.postgres.yml

up-inmemory:
	docker compose -f docker-compose.inmemory.yml up --build

up-postgres:
	docker compose -f docker-compose.postgres.yml up --build

down:
	docker compose down
