install: gen
	go install

gen:
	go generate ./...
