.PHONY: help up down logs .EXPORT_ALL_VARIABLES

up: ##  Start build
	docker-compose -f docker-compose.local.yaml  up --detach --remove-orphans --build

down: ##  Stop an remove containers
	docker-compose -f docker-compose.local.yaml down --remove-orphans

logs: ##  View logs
	docker-compose -f docker-compose.local.yaml logs -f