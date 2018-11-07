APP={{ .Name }}
IMAGE={{ .Image }}
DOCKER_ROOT={{ .DockerRoot }}
NAMESPACE={{ .Namespace }}

FQTAG=$(DOCKER_ROOT)/$(IMAGE)

SHA=$(shell docker inspect --format "{{`{{ index .RepoDigests 0 }}`}}" $(1))

test:
	go test ./...

docker: test
	docker build -t $(FQTAG) . 
	docker push $(FQTAG)

deploy: docker
	kubectl apply -f k8s.yaml
	kubectl --namespace $(NAMESPACE) set image deploment/$(APP) $(APP)=$(call SHA,$(FQTAG))
