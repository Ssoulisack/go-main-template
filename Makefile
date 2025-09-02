# Go parameters
GOCMD=go
GORUN=$(GOCMD) run

SWAG=swag
SWAG_DOC=$(SWAG) init -g ./bootstrap/swagger.go

GIT_WORKFLOW=./git_workflow.sh

.PHONY: git run swag

git:
	$(GIT_WORKFLOW)

run:
	$(GORUN) main.go

swag:
	$(SWAG_DOC)