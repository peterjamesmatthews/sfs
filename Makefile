.PHONY: develop
develop:
	docker compose --env-file .env up -d --build --force-recreate 

.PHONY: dump
dump:
	docker exec database pg_dump -U postgres sfs > server/db/seed/seed.sql

.PHONY: clean
clean:
	docker compose down -v
