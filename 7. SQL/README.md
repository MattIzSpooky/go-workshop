## 7. SQL

TODO:

For this example we will use sqlc
Why SQLC? No external libraries, you write SQL. It generates Go code that uses a db driver.
It can also handle migration tools such as golang-migrate.

https://docs.sqlc.dev/en/stable/overview/install.html

Also use https://github.com/golang-migrate/migrate for migrations

example command:
migrate create -ext sql -dir sql/migrations -seq create_users_table

This will generate empty migration files.

TODO: Explain why migration files are useful compared to 1 fat SQL file.

TODO: add contents of author_table up, and authors_table down

TODO: In this example we will create a simple database.
Table 1: Authors
Table 2: Books 

TODO: ADD Database diagram

last: run sqlc generate to generate sql -> go

idea: you could wrap the generated code with a repository pattern, 
DTOs, etc. to decrease the dependency of your application code to the underlying database.