.PHONY: migrate-up
migrate-up:
	go run cmd/migrations/main.go -dir up


# $GOPATH/bin/sqlc --experimental generate

# special chars to use
’
