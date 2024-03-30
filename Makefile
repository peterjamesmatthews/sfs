.PHONY: develop
develop:
	docker compose --env-file .env up -d --build --force-recreate 

.PHONY: clean
clean:
	docker compose down -v
