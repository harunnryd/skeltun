.PHONY: default help

APP_NAME = skeltun

info: help

define HEADER
    __          _ __         
   / /_  ____  (_) /__  _____
  / __ \/ __ \/ / / _ \/ ___/
 / /_/ / /_/ / / /  __/ /    
/_.___/\____/_/_/\___/_/     PLATE

endef
export HEADER

default: help

help:
	@echo "$$HEADER"
	@echo 'These are common ${APP_NAME} commands used in varios situations:'
	@echo
	@echo 'Usage:'
	@echo '   make run                                        Run the project.'
	@echo '   make route-list                                 Route list,'
	@echo '   make install                                    Install all project dependencies.'
	@echo '   make migration-sql NAME=<option> EXT=<option>   Make migration sql.'
	@echo '   make migrate-up    EXT=<option>                 Migrate up tables.'
	@echo '   make migrate-down  EXT=<option>                 Migrate down tables.'
	@echo '   make docker-up                                  Starting docker.'
	@echo '   make docker-down                                Stopping docker.'
	@echo '   make run-qpool                                  Run queue pool.'

migration-sql:
	@echo "Create migration: ${NAME}.${EXT}"
	go run main.go make:migration ${NAME} ${EXT}

migrate-up:
	@echo "Migrate migration files"
	go run main.go migrate:up ${EXT}

migrate-down:
	@echo "Migrate down migration files"
	go run main.go migrate:down ${EXT}

run:
	@echo "Run the project"
	go run main.go

route-list:
	@echo "Route list"
	go run main.go route:list

docker-up:
	@echo "Docker up"
	docker-compose up -d

docker-down:
	@echo "Docker down"
	docker-compose stop && docker rmi $$(docker-compose images -q) --force

qpool:
	@echo "Queue Pool"
	go run main.go worker