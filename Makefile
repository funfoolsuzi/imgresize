
build:
	docker build -t imgresize .

up:
	docker-compose up -d

down:
	docker-compose down

mockgen:
	~/go/bin/mockgen -source=repo/image_repo.go -destination=mock/image_repo.go -package=mock ImageRepo
	~/go/bin/mockgen -source=imaginaryclient/client.go -destination=mock/resizer.go -package=mock Resizer