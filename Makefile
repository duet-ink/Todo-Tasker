run:
	@docker-compose -f config/deploy.yml -p todo-tasker up --build -d
stop:
	@docker-compose -f config/compose.yml -p todo-tasker down --remove-orphans
dev:
	@PORT=8080 go run .

