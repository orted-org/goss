# for development hot reload using air 
watch:
	./bin/air
test: 
	go test -v -cover ./app/service...
.PHONY: watch