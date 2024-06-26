include .env

run:
		templ generate
		PORT=$(PORT) go run main.go
