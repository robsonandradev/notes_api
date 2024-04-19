run:
	go run .

run-autoreload:
	nodemon -w . -e go -x go run . --signal SIGTERM

test:
	go test -v ./use_cases/**

test-cover:
	go test -cover ./use_cases/**

test-report:
	go test -coverprofile=report.out ./use_cases/**
	go tool cover -html report.out

db-migrations:
	go run migrations/main.go
