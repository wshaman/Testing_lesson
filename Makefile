TEST_OUT_FILE = /tmp/c.out

test:
	go test -v ./...

cover:
	@[ -f $(TEST_OUT_FILE) ] && rm $(TEST_OUT_FILE) || true
	go test -v -coverprofile=/tmp/c.out ./... && go tool cover -html=$(TEST_OUT_FILE)

bench:
	go test -bench=. -benchmem -v ./...

pprof:
	go test -cpuprofile cpu.prof -memprofile mem.prof -bench=. ./app/utils/naming
	go tool pprof -web cpu.prof
	go tool pprof -web mem.prof

integration-test:
	@echo "Starting integration test"
	go test --tags=integration -v ./app/integration
	@echo "Done"
