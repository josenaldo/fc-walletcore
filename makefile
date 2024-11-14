up:
	docker compose up -d --build ;

down:
	docker compose down
	sudo rm -rf ./.docker/mysql*

.PHONY: up down
