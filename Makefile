build:
	docker-compose build todo-app

run:
	docker-compose up todo-app

stop:
	docker-compose stop

migrate:
	migrate -path ./schema -database 'postgres://postgres:qwerty@0.0.0.0:5436/postgres?sslmode=disable' up
