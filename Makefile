.PHONY: build-image
image-build:
	docker build -t snowguard-linux-arm64 \
		--build-arg SUPERSET_ADMIN_KEY=$(SUPERSET_ADMIN_KEY) \
		-f .docker/snowguard.Dockerfile .

	docker build -t snowguard-linux-arm64 \
		--platform linux/arm64 \
		--build-arg SUPERSET_ADMIN_KEY=$(SUPERSET_ADMIN_KEY) \
		-f .docker/snowguard.Dockerfile .

.PHONY: build
build:
	./scripts/build.sh

.PHONY: build-superset
build-superset:
	./scripts/build-superset.sh
