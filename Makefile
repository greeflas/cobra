.PHONY: run
run:
	go run . help.PHONY: run

.PHONY: verbose
verbose:
	go run . -v

.PHONY: version
version:
	go run . version

.PHONY: version_short
version_short:
	go run . version -s


.PHONY: hello
hello:
	go run . hello World
