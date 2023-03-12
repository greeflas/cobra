.PHONY: run
run:
	go run main.go

.PHONY: version
version:
	go run main.go version

.PHONY: version_flag
version_flag:
	go run main.go --version

.PHONY: version_short
version_short:
	go run main.go version -s

.PHONY: hello
hello:
	go run main.go hello World

.PHONY: verbose
verbose:
	go run main.go hello World -v

.PHONY: doc
doc:
	go run doc.go
