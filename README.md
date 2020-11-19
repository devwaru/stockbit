# stockbit
Answer of stockbit coding test

### Answer For Question No 1 (simple database query)

Answer for this question can is 

```sql
SELECT u.id, u.user_name, p.user_name AS parent_user_name
FROM users u
LEFT JOIN users p NO p.id = u.parent
``` 

or you can access on [this file](https://github.com/muhammadaser/stockbit/blob/main/first_answer/query.sql)


### Answer For Question No 2 (write a microservice for search movies)

This answer can access on [second_answer](https://github.com/muhammadaser/stockbit/tree/main/second_answer/movie) folder.
this project contain 2 services, client and server. 
steps to run this project is :

#### Server

1. after clone stockbit repo, change your directory to `second_answer/movie/server`.
2. build this project using command `go build`
3. after success build binary for this project run binary `./server`(for unix)

#### Client

1. after clone stockbit repo, change your directory to `second_answer/movie/client`.
2. build this project using command `go build`
3. after success build binary for this project run binary `./client`(for unix). make sure server running
4. this client have 2 endpoints, `/v1/movies` and `/v1/movie/{movie_id}`.
5. endpoint `/v1/movies` need 2 parameters, `s` for search word and `page` for page.
6. this sample for full endpoint `http://localhost:8080/v1/movies?s=batman&page=1`
7. for movie detail endpoint you can get id from movies endpoint.
8. this sample for full endpoint `http://localhost:8080/v1/movie/tt0372784`


### Answer For Question No 3 (refactor code)

this answer can access on [third_answer](https://github.com/muhammadaser/stockbit/tree/main/third_answer) folder.

to run this project, change directory to `third_answer`. build this project with command `go build`.
run binary `./third_answer` (for unix).


### Answer For Question No 4 (anagram string)

this answer can access on [fourth_answer](https://github.com/muhammadaser/stockbit/tree/main/fourth_answer) folder.

to run this project, change directory to `fourth_answer`. build this project with command `go build`.
run binary `./fourth_answer` (for unix).
