
docker-compose-up:
	docker-compose up -d

docker-compose-down:
	docker-compose down

integrationTest: docker-compose-up
	cp ./database/mysql/migrateTest/* ./database/mysql/migrations
	go test --tags=integration -count=1 -v ./database/mysql
	@make docker-compose-down
