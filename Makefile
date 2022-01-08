dc-build:
	docker-compose up --build -d
	docker-compose exec app sh ./setup/go_mod_init.sh
run:
	docker-compose exec app go run main.go
test:
	docker-compose exec app go test ./... -v