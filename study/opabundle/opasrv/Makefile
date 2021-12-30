bundleTar:
	opa build rbac.rego -o rbac.tar.gz

build: bundleTar
	go build -o bin/ main.go

run: build
	nohup ./bin/main &