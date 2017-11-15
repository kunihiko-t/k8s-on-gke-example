run:
	go run server.go
restart-app:
	docker-compose stop app
	docker-compose rm -f app
	docker-compose up -d app
make-image:
	docker build -t asia.gcr.io/<GCP-PROJECT-ID>/app .
