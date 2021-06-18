PHONY: coverage-html coverage-report lint test dry-release scan

coverage-html: test
	go tool cover -html="c.out"

coverage-report: test
	go tool cover -func="c.out"

test:
	go test -race -covermode=atomic -v -coverprofile="c.out"

lint:
	go vet .
	golint -set_exit_status .

dry-release:
	goreleaser --snapshot --skip-publish --rm-dist

scan:
	gosec -no-fail -fmt sarif -out results.sarif ./...