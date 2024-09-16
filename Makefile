tests: unittests integration

unittests:
	go test -v -short "./..."

integration:
	go test -v -run Integration "./..."
