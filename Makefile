.PHONY: help build up down logs clean migrate

help:
	@echo "Available commands:"
	@echo "  make build    - Build Docker images"
	@echo "  make up       - Start all services"
	@echo "  make down     - Stop all services"
	@echo "  make logs     - View logs"
	@echo "  make clean    - Remove everything"
	@echo "  make migrate  - Run database migrations"

build:
	docker-compose build

up:
	docker-compose up -d
	@echo "✓ Frontend: http://localhost:3000"
	@echo "✓ Backend: http://localhost:3001"
	@echo "✓ Hasura: http://localhost:8080/console"

down:
	docker-compose down

logs:
	docker-compose logs -f

clean:
	docker-compose down -v
	docker system prune -f

migrate:
	@for file in backend/migrations/*.sql; do \
		docker-compose exec -T postgres psql -U postgres -d events_db < $$file; \
	done
	@echo "✓ Migrations complete"
