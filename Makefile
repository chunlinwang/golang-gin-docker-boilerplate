ALL_SERVICES := ${CORE_SERVICES} 

COMPOSE_FILE := -f docker-compose.yml
SERVICE=app

# --------------------------

.PHONY: help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: ## build services
	@docker-compose ${COMPOSE_FILE} up -d --build ${ALL_SERVICES}

down: ## down services
	@docker-compose ${COMPOSE_FILE} down

stop: ## stop services
	@docker-compose ${COMPOSE_FILE} stop ${ALL_SERVICES}

restart: ## restart services
	@docker-compose ${COMPOSE_FILE} restart ${ALL_SERVICES}

rm:
	@docker-compose $(COMPOSE_FILE) rm -f ${ALL_SERVICES}

logs: ## Check logs
	@docker-compose $(COMPOSE_FILE) logs --follow --tail=1000 ${ALL_SERVICES}

images:
	@docker-compose $(COMPOSE_FILE) images ${ALL_SERVICES}

clean: ## Remove all Containers and Delete Volume Data
	@docker-compose ${COMPOSE_FILE} down -v

cli: ## Connect to a docker container
	@docker-compose exec ${SERVICE} sh