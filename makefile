set-up:
	docker-compose up --build
attach-app:
	docker-compose exec app sh
test:
	docker-compose run --rm app go test ./...