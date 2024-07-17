.PHONY: build-image
image-build:
	docker build -t snowguard \
		--build-arg SUPERSET_ADMIN_KEY=$(SUPERSET_ADMIN_KEY) \
		-f .docker/snowguard.Dockerfile .

.PHONY: build
build:
	./scripts/build.sh

.PHONY: build-superset
build-superset:
	./scripts/build-superset.sh
