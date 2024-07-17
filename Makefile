.PHONY: image-build
image-build:
	docker build -t snowguard \
		--build-arg SUPERSET_ADMIN_KEY=$(SUPERSET_ADMIN_KEY) \
		-f .docker/snowguard.Dockerfile .

.PHONY: build
build:
	./scripts/build.sh
