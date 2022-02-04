COPYRIGHT=Copyright (c) 2022 Cyber Home Lab authors
DEBUG_MSG=$$(tput setaf 6)DEBUG  : \t$$(tput sgr0)
INFO_MSG=$$(tput setaf 2)INFO   : \t$$(tput sgr0)
WARNING_MSG=$$(tput setaf 3)WARNING: \t$$(tput sgr0)
ERROR_MSG=$$(tput setaf 1)ERROR  : \t$$(tput sgr0)
SHELL=/bin/bash
GO111MODULE=on
BINARY_NAME=telegram.bin
CYBERHOMELAB_CONFIG=$(shell echo $$(pwd)/config.toml)

export GO111MODULE
export CYBERHOMELAB_CONFIG

.DEFAULT_GOAL := help

.ONESHELL:
init:
	@git config --local core.hooksPath .githooks/
	go version || ( echo "$(ERROR_MSG)go not found, please install it https://go.dev/doc/install" ; exit 2 )
	golangci-lint version || ( echo "$(ERROR_MSG)golangci-lint not found, please install it https://golangci-lint.run/usage/install/" ; exit 2 )
	if ! grep -q CYBERHOMELAB_CONFIG ~/.bashrc
	then
		echo "export CYBERHOMELAB_CONFIG=$(CYBERHOMELAB_CONFIG)" >> ~/.bashrc
		export CYBERHOMELAB_CONFIG=$(CYBERHOMELAB_CONFIG)
	fi

clean:
	@go clean
	echo -e "$(INFO_MSG)Cleanup was executed"

build:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME) main.go

run-go:
	@go run main.go

run-binary: build
	@./$(BINARY_NAME)

run: run-go

tests:
	@go test -v ./... ; echo $?

check:
	@.githooks/pre-push
	echo -e "$(INFO_MSG)All check(s) passed"

check-pre-push:
	@cat /tmp/pre-push.err /tmp/pre-push.log

check-coverage:
	@go test -cover ./... -coverprofile=coverage.out >/dev/null
	go tool cover -func=coverage.out | grep -vE '^total:|init|(7|8|9|10)[0-9].[0-9]%$$'
	go tool cover -html=coverage.out -o coverage.html
	echo -e "\n * For more info, check coverage.html or visit https://go.dev/blog/cover\n"

check-todos:
	@find . -type f -name '*.go' -exec grep -n TODO {} +

.ONESHELL:
check-copyright:
	@for file in $$(find . -type f -name "*.go")
	do
		if head -3 "$${file}" | tail -1 | grep -v $(COPYRIGHT)
		then
			echo -e "$(ERROR_MSG)File $${file} doesn't have the copyright set"
		else
			echo -e "$(INFO_MSG)File $${file} has the copyright set"
		fi
	done

.ONESHELL:
help:
	@echo -e "
	================================================
	\t\tCyber Home Lab
	\tTelegram client written in Go
	================================================
	
	Commands available:
	\tmake init
	\tmake clean
	\tmake build
	\tmake check
	\tmake check-pre-push
	\tmake check-coverage
	\tmake check-todos
	\tmake check-copyright
	\tmake tests
	\tmake run
	\tmake run-go
	\tmake run-binary
	"
