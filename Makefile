.PHONY: run
run:
	go run main.go wire_gen.go --from=$(FROM) --to=$(TO) --archiveDir=$(DIR) --runMode=$(RUN_MODE)

.PHONY: test
test:
	go test -v ./...