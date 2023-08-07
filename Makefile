.PHONY: migrate-up
migrate-up:
	go run cmd/migrations/main.go -dir up
