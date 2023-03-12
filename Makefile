.PHONY: run
run:
	go run .

.PHONY: version
version:
	go run . version

.PHONY: version_flag
version_flag:
	go run . --version

.PHONY: version_short
version_short:
	go run . version -s

.PHONY: hello
hello:
	go run . hello World

.PHONY: verbose
verbose:
	go run . hello World -v
