.PHONY: all test clean

run:
	go run main.go

test: 
	$(foreach file, $(wildcard ls test/*/*_test.go), go test $(file) -v;)