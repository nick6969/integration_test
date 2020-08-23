
docker-compose-up:
	docker-compose up -d

docker-compose-down:
	docker-compose down

integrationTest: docker-compose-up
	if ! [ -d "./database/mysql/migrate_test" ]; then mkdir ./database/mysql/migrate_test; fi
	cp ./database/mysql/migrateTest/* ./database/mysql/migrate_test
	cp ./database/mysql/migrations/* ./database/mysql/migrate_test
	go test --tags=integration -count=1 -v ./database/mysql
	rm -rf ./database/mysql/migrate_test
	@make docker-compose-down
