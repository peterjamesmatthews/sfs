.PHONY: develop
develop:
	docker compose up -d 

.PHONY: clean
clean:
	docker compose down 
