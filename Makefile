# Alex Mackay 2024

example-go:
	cd ./go && go run .

example-python:
	cd ./python && python3 example.py

test-python:
	cd ./python && python3 -m unittest

test-go:
	cd ./go && go test -v -cover ./...

.PHONY: test-python test-go