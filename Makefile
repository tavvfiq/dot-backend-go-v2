ENV_LOCAL_TEST=\
				HOST_POSTGRE=localhost \
				PORT_POSTGRE=5432 \
				USER_POSTGRE=postgres \
				PASSWORD_POSTGRE=password \
				DBNAME_POSTGRE=dot \
				REDIS_ADDR=localhost:6379 \
				REDIS_PASS= \
				REDIS_DB=0 \
				PORT=:8080

docker.build:
	ENV=${env} docker-compose build
docker.up:
	ENV=${env} docker-compose up -d
docker.down:
	ENV=${env} docker-compose down
docker.article-api.stop: 
	docker stop dot_article_api
test.e2e: docker.article-api.stop
	$(ENV_LOCAL_TEST) \
	go test -p 1 -tags=integration ./internal/it -v -count=1
docker.start: docker.build docker.up