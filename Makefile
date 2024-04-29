build-docker-image: 
	docker build --no-cache -t invoices-manager .

# Copy code to container and build it there.
build: 
	docker cp cmd invoices-manager:/invoices-manager/
	docker cp pkg invoices-manager:/invoices-manager/
	docker cp go.mod invoices-manager:/invoices-manager/go.mod
	docker cp go.sum invoices-manager:/invoices-manager/go.sum
	docker exec invoices-manager go build -o /bin/invoices-manager cmd/*.go

build-and-restart: build 
	docker restart invoices-manager

run:
	docker-compose up --force-recreate

run-tests:
	newman run tests/TestJump.postman_collection.json