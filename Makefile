# Lint using golangci-lint.
lint:
	golangci-lint run --config .golangci.yml ./...

###################################

# Remove docker image with tag None.
clear-none-docker-images:
	docker images --filter "dangling=true" -q --no-trunc | xargs docker rmi

###################################

# Run postgres and redis container.
compose-up-postgres-redis:
	docker compose up -d gocheck-db-postgres gocheck-cache-redis

compose-down-postgres-redis:
	docker compose down gocheck-db-postgres gocheck-cache-redis

###################################

# Run deployment.
deploy:
	docker compose up --build

undeploy:
	docker compose down

###################################

# Generate proto file.
generate-proto:
	protoc \
		--go_out=.      --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		pkg/gocheckgrpc/*.proto

###################################

migrate-up:
	go run internal/table/migration/migrate_up.go
