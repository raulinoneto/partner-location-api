build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/createpartner cmd/lambda/createpartner/*.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/getpartner cmd/lambda/getpartner/*.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/searchpartners cmd/lambda/searchpartners/*.go
	chmod 0755 bin/* -v

clean:
	rm -rf ./bin -v

test:
	go test -coverpkg=./... ./...

deploy:
	docker exec -i partners-app serverless deploy

build-and-deploy:
	make build;
	make deploy

install:
	docker-compose build

run:
	docker-compose up

