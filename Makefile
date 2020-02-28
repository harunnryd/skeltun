.PHONY: default help

APP_NAME = skeltun

default: help

help:
	@echo 'These are common ${APP_NAME} commands used in varios situations:'
	@echo
	@echo 'Usage:'
	@echo '   make install          Install all project dependencies.'
	@echo	'   make migration-sql    Make migration sql.'
	@echo	'   make migrate-up       Migrate up tables.'

migration-sql:
	@echo "Create migration: ${NAME}.${EXT}"
	go run main.go make:migration ${NAME} ${EXT}

migrate-up:
	@echo "Migrate migration files"
	go run main.go migrate:up ${EXT}

migrate-down:
	@echo "Migrate down migration files"
	go run main.go migrate:down ${EXT}