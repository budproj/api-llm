include .env

run:
		templ generate
		PORT=$(PORT) go run entrypoint/controllers/http/main.go