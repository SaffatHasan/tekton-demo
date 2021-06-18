# retrieved from https://github.com/tektoncd/pipeline/releases
TEKTON_VERSION=v0.25.0

create-cluster:
	kind create cluster --name tekton

install-tekton:
	kubectl apply -f https://storage.googleapis.com/tekton-releases/pipeline/previous/$(TEKTON_VERSION)/release.yaml

install-tasks:
	kubectl apply \
		-f https://raw.githubusercontent.com/tektoncd/catalog/main/task/git-clone/0.4/git-clone.yaml	 \
		-f https://raw.githubusercontent.com/tektoncd/catalog/main/task/golang-test/0.2/golang-test.yaml \
		-f https://raw.githubusercontent.com/tektoncd/catalog/main/task/golangci-lint/0.2/golangci-lint.yaml

