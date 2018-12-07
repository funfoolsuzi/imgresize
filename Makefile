
build:
	docker build -t imgresize .

up:
	docker-compose up -d

down:
	docker-compose down