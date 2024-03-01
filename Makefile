.PHONY: migrate-up populate-db up start
migrate-up:
	go run cmd/migrations/main.go -dir up

populate-db:
	go run scripts/populate_db.go  

# $GOPATH/bin/sqlc --experimental generate

# special chars to use
#’

up:
	docker compose up -d 

start: up migrate-up populate-db
	go run cmd/main.go
