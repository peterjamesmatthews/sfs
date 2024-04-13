.PHONY: develop
develop:
	docker compose --env-file .env up -d --build --force-recreate 

.PHONY: migrate
migrate:
	atlas migrate diff \
		--dir "file://server/db/migrations" \
		--to "file://server/db/schema.sql" \
		--dev-url "docker://postgres/16/dev"

.PHONY: dump
dump:
	docker exec database pg_dump -U postgres sfs > server/db/seed/seed.sql

.PHONY: clean
clean:
	docker compose down -v
