SHELL := /bin/bash

# ---- Config ----
IMAGE_REPO ?= ghcr.io/your-username/go-chomp
IMAGE_TAG  ?= v0.1.0
IMAGE      := $(IMAGE_REPO):$(IMAGE_TAG)
CHART_DIR  := deploy/go-chomp
RELEASE    := chomp
NAMESPACE  := games

# Optional: for kind
KIND_CLUSTER ?= kind

# ---- Local build/run ----
.PHONY: build run
build:
	go build -buildvcs=false -o chomp .

run:
	./chomp -w 6 -h 5

run-cpu:
	./chomp -w 6 -h 5 -cpu

# ---- Docker ----
.PHONY: docker-build docker-run docker-push
docker-build:
	docker build -t $(IMAGE) .

docker-run:
	docker run --rm -it $(IMAGE) -w 6 -h 5

docker-push:
	docker push $(IMAGE)

# ---- kind helpers (optional) ----
.PHONY: kind-load
kind-load:
	kind load docker-image $(IMAGE) --name $(KIND_CLUSTER)

# ---- Helm (Kubernetes) ----
.PHONY: helm-install helm-upgrade helm-uninstall k-attach k-logs ns
ns:
	kubectl get ns $(NAMESPACE) >/dev/null 2>&1 || kubectl create ns $(NAMESPACE)

helm-install: ns
	helm upgrade --install $(RELEASE) $(CHART_DIR) \
	  --namespace $(NAMESPACE) \
	  --set image.repository=$(IMAGE_REPO) \
	  --set image.tag=$(IMAGE_TAG)

helm-upgrade:
	helm upgrade $(RELEASE) $(CHART_DIR) \
	  --namespace $(NAMESPACE)

helm-uninstall:
	helm uninstall $(RELEASE) -n $(NAMESPACE)

# Attach an interactive shell to the pod (so you can play)
k-attach:
	kubectl -n $(NAMESPACE) exec -it deploy/$(RELEASE)-go-chomp -- /app/chomp -w 6 -h 5

k-logs:
	kubectl -n $(NAMESPACE) logs -f deploy/$(RELEASE)-go-chomp
