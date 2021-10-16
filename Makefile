build:
	docker-compose build todo-backend

run:
	docker-compose up todo-backend

stop:
	docker-compose stop

migrate:
	migrate -path ./schema -database 'postgres://postgres:qwerty@0.0.0.0:5436/postgres?sslmode=disable' up