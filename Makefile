# ------------------ Golang commands ------------------
.PHONY: clean install update test build run
clean:
	@go clean -cache
	@go clean -modcache
	@go clean -testcache
	@rm go.sum
install:
	@go mod download
update:
	@go mod tidy
	@go get -u ./...
	@go mod tidy
test:
	@go test -v ./cmd/test/...
build:
	@go build -ldflags "-s -w" -o ./.build/main ./cmd/main.go
run:
	@./.build/main

# Third party libraries commands
.PHONY: scan
scan:
	@go install golang.org/x/vuln/cmd/govulncheck@latest
	@echo "Running vulnerability scan..."
	@govulncheck ./...


# ------------------ Docker commands ------------------
.PHONY: docker-cdn-build docker-cdn
docker-cdn-build:
	@docker compose up --build --no-deps -d cdn-build
docker-cdn:
	@echo "Launching minio..."
	@docker compose up --build --no-deps -d minio
	@echo ""
	@echo "Launching cdn..."
	@docker compose up --build --no-deps -d cdn

.PHONY: docker-ghcr-login
docker-ghcr-login:
	@echo "" ;\
	echo "Login - GitHub Docker Registry" ;\
	read -p "Enter your Github username: " gUsername ;\
	read -p "Enter your Github personal access token: " gPass ;\
	echo "" ;\
	echo $$gPass | docker login ghcr.io -u $$gUsername --password-stdin ;\

.PHONY: docker-ghcr-push-specific docker-ghcr-pull-specific
docker-ghcr-push-specific:
	@echo "" ;\
	echo "Tag - GitHub Docker Registry" ;\
	gCorp="4kpros" ;\
	gRepo="cdn" ;\
	read -p "Enter your package name(cdn): " gPackage; gTag=$${gPackage:-"cdn"} ;\
	read -p "Enter your tag(default is 1): " gTag; gTag=$${gTag:-"1"} ;\
	docker tag cdn-$$gPackage ghcr.io/$$gCorp/$$gRepo/$$gPackage:$$gTag ;\
	echo "" ;\
	echo "Pushing $$gPackage - GitHub Docker Registry" ;\
	docker push ghcr.io/$$gCorp/$$gRepo/$$gPackage:$$gTag;
docker-ghcr-pull-specific:
	@echo "" ;\
	echo "Tag - GitHub Docker Registry" ;\
	gCorp="4kpros" ;\
	gRepo="cdn" ;\
	read -p "Enter your package name(cdn): " gPackage; gTag=$${gPackage:-"cdn"} ;\
	read -p "Enter your tag(default is 1): " gTag; gTag=$${gTag:-"1"} ;\
	echo "" ;\
	echo "Pulling $$gPackage - GitHub Docker Registry" ;\
	docker pull ghcr.io/$$gCorp/$$gRepo/$$gPackage:$$gTag
