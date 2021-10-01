# DOT backend go v2
___
## article api demo using [go-clean-arch](https://github.com/bxcodec/go-clean-arch) pattern.
### API
| name 	| method 	| endpoint 	| params 	| query 	| body 	|
|---	|---	|:---:	|---	|---	|---	|
| create user 	| POST 	| /v1/user 	|  	|  	| {      "name" :  "user2" } 	|
| create article 	| POST 	| /v1/article 	|  	|  	| {      "title" : "test article 4" ,      "subtitle" :  "this is test article 4" ,      "author_id" :  1 ,      "content" :  "test." } 	|
| get all article 	| GET 	| /v1/article 	|  	| title, page, limit 	|  	|
| get article detail 	| GET 	| /v1/article/detail/:id 	| article id 	|  	|  	|
| update article 	| PATCH 	| /v1/article/:id 	| article id 	|  	| {      "title" :  "ada badak laut" ,      "subtitle" :  "subtitle" ,      "content" :  "test" ,      "author_id" :  1 } 	|
| delete article 	| DELETE 	| /v1/article/:id 	| article id 	|  	|  	|
| add comment 	| POST 	| /v1/comment 	|  	|  	| {      "user_id" :  1 ,      "article_id" :  15 ,      "content" :  "tidak tau" } 	|
| delete comment 	| DELETE 	| /v1/comment/:id 	| comment id 	|  	|  	|

___
### why using [go-clean-arch](https://github.com/bxcodec/go-clean-arch) pattern?
in my opinion, this pattern is good for both small project and big project because of its scalability. a little bit bloated and not too DRY, but its prevent for tight coupling for every layer (domain, usecase, repository, delivery/handler, etc). and rather clean because of the dependency injection. every layer has its own contract to communicate with each other. so if there are any changes on any layer, it will not effect/break other layer as long as not breaking their contract.
___
### command
make sure makefile, docker and docker-compose are installed and env variable adapted to your system.
- #### deploying
    ``` 
    make docker.start env=local 
    ```
- #### perform integration and e2e test
  ```
  make docker.up
  make test.e2e
  ```
    


