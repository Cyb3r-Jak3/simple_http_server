PHONY: coverage lint

coverage:
	go test -coverprofile="c.out"
	go tool cover -html="c.out"

lint:
	go vet .
	golint -set_exit_status .

dry-release:
	goreleaser --snapshot --skip-publish --rm-dist