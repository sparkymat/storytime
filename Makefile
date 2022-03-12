all: docker


docker: storytime
	docker build -t sparkymat/storytime .

storytime:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-s -w -extldflags "-static"' -o storytime storytime.go

.PHONY: storytime
