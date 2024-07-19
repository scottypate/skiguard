.PHONY: build-image
build-image:
	docker build -t scalecraft/snowguard:$(VERSION)-linux-arm64 -t scalecraft/snowguard:latest \
		--platform linux/arm64 \
		--build-arg SUPERSET_ADMIN_KEY=$(SUPERSET_ADMIN_KEY) \
		-f .docker/snowguard.Dockerfile .

	docker build -t scalecraft/snowguard:$(VERSION)-linux-amd64 -t scalecraft/snowguard:latest \
		--platform linux/amd64 \
		--build-arg SUPERSET_ADMIN_KEY=$(SUPERSET_ADMIN_KEY) \
		-f .docker/snowguard.Dockerfile .

.PHONY: push-image
push-image:
	docker push scalecraft/snowguard:$(VERSION)-linux-arm64
	docker push scalecraft/snowguard:$(VERSION)-linux-amd64

	docker manifest create scalecraft/snowguard:$(VERSION) \
		--amend scalecraft/snowguard:$(VERSION)-linux-arm64 --amend scalecraft/snowguard:$(VERSION)-linux-amd64

	docker manifest push scalecraft/snowguard:$(VERSION)

	docker manifest create scalecraft/snowguard:latest \
		--amend scalecraft/snowguard:$(VERSION)-linux-arm64 --amend scalecraft/snowguard:$(VERSION)-linux-amd64

	docker manifest push scalecraft/snowguard:latest

.PHONY: build
build:
	./scripts/build.sh

.PHONY: build-superset
build-superset:
	./scripts/build-superset.sh
