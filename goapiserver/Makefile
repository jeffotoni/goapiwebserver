# Makefile

.EXPORT_ALL_VARIABLES:	

GO111MODULE=on
PORT_SERVER = 5002

all:
	echo $$GOPATH
	##go get -d
	##go run *.go

test:
	@echo "\n"
	@echo "\033[0;33m################################## go tests ##################################\033[0m"

build:
	@echo "\n"
	@echo "--------------------------------------------------------------------------------------"
	@echo "\033[0;33m################################## build goapiserver ##################################\033[0m"
	go build -o goapiserver && ./goapiserver
	@echo "\033[0;32mfrontend generated\033[0m \033[0;33m[ok]\033[0m \n"

prod:
	@echo "\033[0;33m################################## build go-bindata ##################################\033[0m"
	go-bindata -o pkg/assets/binhtml.go --nocompress --pkg=assets templates/... 
	@echo "\033[0;32mbindata generated\033[0m  \033[0;33m[ok]\033[0m "
	@echo "\n"
	@echo "--------------------------------------------------------------------------------------"
	@echo "\033[0;33m################################## build goapiserver ##################################\033[0m"
	go build -o goapiserver && ./goapiserver
	@echo "\033[0;32mfrontend generated\033[0m \033[0;33m[ok]\033[0m \n"
