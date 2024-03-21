.PHONY: develop
develop:
	docker compose up -d --build

.PHONY: clean
clean:
	docker compose down 
