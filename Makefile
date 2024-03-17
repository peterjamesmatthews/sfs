.PHONY: develop
develop: 
	docker compose \
		--profile develop \
		up \
		--build \
		--detach

.PHONY: clean
clean: 
	docker compose \
		--profile "*" \
		down
	docker image prune -af
