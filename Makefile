.PHONY: build-image
build-image:
	docker build -t scalecraft/skiguard:$(VERSION)-linux-arm64 -t scalecraft/skiguard:latest \
		--platform linux/arm64 \
		--build-arg SUPERSET_ADMIN_KEY=$(SUPERSET_ADMIN_KEY) \
		-f .docker/skiguard.Dockerfile .

	docker build -t scalecraft/skiguard:$(VERSION)-linux-amd64 -t scalecraft/skiguard:latest \
		--platform linux/amd64 \
		--build-arg SUPERSET_ADMIN_KEY=$(SUPERSET_ADMIN_KEY) \
		-f .docker/skiguard.Dockerfile .

.PHONY: push-image
push-image:
	docker push scalecraft/skiguard:$(VERSION)-linux-arm64
	docker push scalecraft/skiguard:$(VERSION)-linux-amd64

	docker manifest create scalecraft/skiguard:$(VERSION) \
		--amend scalecraft/skiguard:$(VERSION)-linux-arm64 --amend scalecraft/skiguard:$(VERSION)-linux-amd64

	docker manifest push scalecraft/skiguard:$(VERSION)

	docker manifest create scalecraft/skiguard:latest \
		--amend scalecraft/skiguard:$(VERSION)-linux-arm64 --amend scalecraft/skiguard:$(VERSION)-linux-amd64

	docker manifest push scalecraft/skiguard:latest

.PHONY: build
build:
	./scripts/build.sh

.PHONY: build-superset
build-superset:
	./scripts/build-superset.sh
