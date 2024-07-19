.PHONY: build-image
build-image:
	docker build -t scalecraft/snowguard:linux-arm64 \
		--platform linux/arm64 \
		--build-arg SUPERSET_ADMIN_KEY=$(SUPERSET_ADMIN_KEY) \
		-f .docker/snowguard.Dockerfile .

	docker build -t scalecraft/snowguard:linux-amd64 \
		--platform linux/amd64 \
		--build-arg SUPERSET_ADMIN_KEY=$(SUPERSET_ADMIN_KEY) \
		-f .docker/snowguard.Dockerfile .

.PHONY: push-image
push-image:
	docker push scalecraft/snowguard:linux-arm64
	docker push scalecraft/snowguard:linux-amd64

.PHONY: build
build:
	./scripts/build.sh

.PHONY: build-superset
build-superset:
	./scripts/build-superset.sh
