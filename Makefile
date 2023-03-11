.PHONY: run
run:
	go run . help.PHONY: run

.PHONY: verbose
verbose:
	go run . -v

.PHONY: version
version:
	go run . version
