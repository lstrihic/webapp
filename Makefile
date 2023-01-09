CURRENT_DIR=$(shell pwd)
SERVICE_NAME=webapp
VERSION=$(shell cat ${CURRENT_DIR}/VERSION)

.PHONY: generate_schema
generate_schema:
	@go run -mod=mod entgo.io/ent/cmd/ent generate --template "$(CURRENT_DIR)/adapter/db/template" --target "$(CURRENT_DIR)/adapter/db/entity" "$(CURRENT_DIR)/adapter/db/schema"

.PHONY: create_schema
create_schema:
	@go run -mod=mod entgo.io/ent/cmd/ent init --target "$(CURRENT_DIR)/adapter/db/schema" $(name)

.PHONY: describe_schema
describe_schema:
	@go run -mod=mod entgo.io/ent/cmd/ent describe "$(CURRENT_DIR)/adapter/db/schema"

.PHONY: generate_sql
generate_sql:
	@go run -mod=mod github.com/kyleconroy/sqlc/cmd/sqlc generate -f ./adapter/sql/sqlc.yaml
