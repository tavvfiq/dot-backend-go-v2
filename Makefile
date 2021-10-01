docker.build:
	ENV=${env} docker-compose build
docker.up:
	ENV=${env} docker-compose up -d
docker.down:
	ENV=${env} docker-compose down
docker.start: docker.build docker.up