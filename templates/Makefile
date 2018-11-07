APP={{ .Name }}
IMAGE={{ .Image }}
DOCKER_ROOT={{ .DockerRoot }}
NAMESPACE={{ .Namespace }}

FQTAG=$(DOCKER_ROOT)/$(IMAGE)

test:
	go test ./...

docker: test
	docker build -t $(FQTAG) . 
	docker push $(FQTAG)

deploy: docker
	kubectl apply -f k8s.yaml
	kubectl delete pod --namespace $(NAMESPACE) -l app=$(APP)
