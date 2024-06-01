.PHONY: all
all: graph migrate sqlc

.PHONY: graph
graph:
	go generate ./graph

.PHONY: migrate
migrate:
	atlas migrate diff \
		--dir "file://db/migrations" \
		--to "file://db/schema.sql" \
		--dev-url "docker://postgres/16/dev"

.PHONY: sqlc
sqlc:
	sqlc generate -f db/sqlc.yml
