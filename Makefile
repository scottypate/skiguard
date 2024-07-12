.PHONY: image-build
image-build:
	docker build -t snowguard \
		--build-arg SUPERSET_SECRET_KEY=$(SUPERSET_SECRET_KEY) \
		-f .docker/snowguard.Dockerfile .

.PHONY: build
build:
	./scripts/build.sh
