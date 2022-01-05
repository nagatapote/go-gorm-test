setup: 
	docker-compose up --build
start:
	docker-compose up -d
	go run main.go