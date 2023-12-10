# Makefile for Your Project

# Directories
BIN_DIR = bin
PROTO_DIR = proto
SERVER_DIR = server
CLIENT_DIR = client

# OS-specific configurations
UNAME := $(shell uname -s)
VERSION_AND_ARCH = $(shell uname -rm)
OS = linux ${VERSION_AND_ARCH}
PACKAGE = $(shell head -1 go.mod | awk '{print $$2}')
CHECK_DIR_CMD = test -d $@ || (echo "\033[31mDirectory $@ doesn't exist\033[0m" && false)
HELP_CMD = grep -E '^[a-zA-Z_-]+:.*?\#\# .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?\#\# "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
RM_F_CMD = rm -f
RM_RF_CMD = ${RM_F_CMD} -r
SERVER_BIN = ${SERVER_DIR}
CLIENT_BIN = ${CLIENT_DIR}

# Default goal
.DEFAULT_GOAL := help

# Phony targets (targets that are not files)
.PHONY: all greet calculator blog test clean clean_greet clean_calculator clean_blog rebuild bump about help

# Subprojects
project := greet calculator blog

# Default target to build all subprojects
all: $(project) ## Generate Pbs and build

# Build each subproject
greet: $@ ## Generate Pbs and build for greet
calculator: $@ ## Generate Pbs and build for calculator
blog: $@ ## Generate Pbs and build for blog

$(project):
	@${CHECK_DIR_CMD}
	protoc -I$@/${PROTO_DIR} --go_opt=module=${PACKAGE} --go_out=. --go-grpc_opt=module=${PACKAGE} --go-grpc_out=. $@/${PROTO_DIR}/*.proto
	go build -o ${BIN_DIR}/$@/${SERVER_BIN} ./$@/${SERVER_DIR}
	go build -o ${BIN_DIR}/$@/${CLIENT_BIN} ./$@/${CLIENT_DIR}

# Target to run tests
test: all ## Launch tests
	go test ./...

# Clean targets
clean: clean_greet clean_calculator clean_blog ## Clean generated files
	${RM_F_CMD} ssl/*.crt
	${RM_F_CMD} ssl/*.csr
	${RM_F_CMD} ssl/*.key
	${RM_F_CMD} ssl/*.pem
	${RM_RF_CMD} ${BIN_DIR}

clean_greet: ## Clean generated files for greet
	${RM_F_CMD} greet/${PROTO_DIR}/*.pb.go

clean_calculator: ## Clean generated files for calculator
	${RM_F_CMD} calculator/${PROTO_DIR}/*.pb.go

clean_blog: ## Clean generated files for blog
	${RM_F_CMD} blog/${PROTO_DIR}/*.pb.go

# Target to rebuild the entire project
rebuild: clean all ## Rebuild the whole project

# Target to update packages version
bump: all ## Update packages version
	go get -u ./...

# Target to display build-related information
about: ## Display info related to the build
	@echo "OS: ${OS}"
	@echo "Shell: ${SHELL} ${SHELL_VERSION}"
	@echo "Protoc version: $(shell protoc --version)"
	@echo "Go version: $(shell go version)"
	@echo "Go package: ${PACKAGE}"
	@echo "Openssl version: $(shell openssl version)"

# Target to show help
help: ## Show this help
	@${HELP_CMD}
