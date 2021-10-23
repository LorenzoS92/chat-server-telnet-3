start:
	@echo 'Starting chat server at port :10000'
	docker-compose up -d

validate:
	@echo 'Show and validate docker compose configuration'
	docker-compose config
