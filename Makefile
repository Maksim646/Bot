# Определяем переменные
DOCKER_COMPOSE = docker-compose
SERVICE_NAME = bot

# Сообщения
BUILD_MSG = Building the $(SERVICE_NAME) service...
START_MSG = Starting the $(SERVICE_NAME) service...
STOP_MSG = Stopping the $(SERVICE_NAME) service...
REMOVE_MSG = Removing the $(SERVICE_NAME) service...
DB_START_MSG = Starting the PostgreSQL database...
DB_STOP_MSG = Stopping the PostgreSQL database...
DB_REMOVE_MSG = Removing the PostgreSQL database...

# Команды
.PHONY: all build up down restart logs db-start db-stop db-remove

# Сборка образа
build:
	@echo "$(BUILD_MSG)"
	@$(DOCKER_COMPOSE) build

# Запуск контейнеров
up: build
	@echo "$(START_MSG)"
	@$(DOCKER_COMPOSE) up -d

# Остановка контейнеров
down:
	@echo "$(STOP_MSG)"
	@$(DOCKER_COMPOSE) down

# Перезапуск контейнеров
restart: down up

# Просмотр логов
logs:
	@$(DOCKER_COMPOSE) logs -f

# Запуск базы данных
db-start:
	@echo "$(DB_START_MSG)"
	@$(DOCKER_COMPOSE) up -d postgres

# Остановка базы данных
db-stop:
	@echo "$(DB_STOP_MSG)"
	@$(DOCKER_COMPOSE) stop postgres

# Удаление базы данных
db-remove:
	@echo "$(DB_REMOVE_MSG)"
	@$(DOCKER_COMPOSE) rm -f postgres

# Запуск тестов (если у вас есть тесты в проекте)
test:
	@echo "Running tests..."
	@docker exec -it $$(docker ps -aqf name=$(SERVICE_NAME)) go test ./...

# Запуск интерактивной сессии с базой данных
run-sql:
	@echo "Running interactive session with PostgreSQL..."
	@docker exec -it $$(docker ps -aqf name=$(SERVICE_NAME)-postgres) psql -U postgres