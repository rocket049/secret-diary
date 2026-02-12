release:
	go mod vendor
	qtdeploy build
dev:
	go mod vendor
	go build
