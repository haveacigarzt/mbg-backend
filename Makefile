include .envrc
# ==================================================================================== #
# HELPERS
# ==================================================================================== #
## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]
# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #
## run/api: run the cmd/api application
.PHONY: run/api
run/api:
	@go run ./cmd/api -db-dsn=${MBG_DB_DSN} -cors-trusted-origins="http://localhost:5173 http://localhost:9001 http://192.168.1.5:3000 http://192.168.1.8:5173"
## db/psql: connect to the database using psql
.PHONY: db/psql
db/psql:
	psql ${MBG_DB_DSN}
## db/migrations/new name=$1: create a new database migration
.PHONY: db/migrations/new
db/migrations/new:
	@echo 'Creating migration files for ${name}...'
	migrate create -seq -ext=.sql -dir=./migrations ${name}
## db/migrations/up: apply all up database migrations
.PHONY: db/migrations/up
db/migrations/up: confirm
	@echo 'Running up migrations...'
	migrate -path ./migrations -database ${MBG_DB_DSN} up
## db/migrations/force version=$1: usually when got an error on migrate, force migrations to previous version
.PHONY: db/migrations/force
db/migrations/force: confirm
	@echo 'Running up migrations...'
	migrate -path ./migrations -database ${MBG_DB_DSN} force ${version}
# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #
## tidy: format all .go files, and tidy and vendor module dependencies
.PHONY: tidy
tidy:
	@echo 'Tidying module dependencies...'
	go mod tidy
	@echo 'Verifying and vendoring module dependencies...'
	go mod verify
	go mod vendor
	@echo 'Formatting .go files...'
	go fmt ./...
## audit: run quality control checks
.PHONY: audit
audit:
	@echo 'Checking module dependencies'
	go mod tidy -diff
	go mod verify
	@echo 'Vetting code...'
	go vet ./...
	staticcheck ./...
	@echo 'Running tests...'
	go test -race -vet=off ./...

# ==================================================================================== #
# BUILD
# ==================================================================================== #
## build/api: build the cmd/api application
.PHONY: build/api
build/api:
	@echo 'Building cmd/api...'
	go build -o=./bin/api ./cmd/api