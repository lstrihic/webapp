CURRENT_DIR=$(shell pwd)
SERVICE_NAME=webapp
VERSION=$(shell cat ${CURRENT_DIR}/VERSION)

.PHONY: generate_schema
generate_schema:
	@go run -mod=mod entgo.io/ent/cmd/ent generate --target "$(CURRENT_DIR)/adapter/db/entity" "$(CURRENT_DIR)/adapter/db/schema"

.PHONY: create_schema
create_schema:
	@go run -mod=mod entgo.io/ent/cmd/ent init --target "$(CURRENT_DIR)/adapter/db/schema" $(name)

.PHONY: describe_schema
describe_schema:
	@go run -mod=mod entgo.io/ent/cmd/ent describe "$(CURRENT_DIR)/adapter/db/schema"

