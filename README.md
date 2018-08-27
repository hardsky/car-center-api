
# Car Center API

REST api for car_center_web

We use *dep* as dependency manager.

env variables:

- CR_DEBUG - turn on debvug logs
- CR_ADDR - service address (for example localhost:8080)
- CR_DB_ADDR - db address (for example :5432)
- CR_DB_USER - db user
- CR_DB_PSW - db user password
- CR_DB_DATABASE - service database

How to use:
1. create database for application (let's name it 'car_center')

2. `cd migration`

3. `dep ensure`

4. apply migrations from */migrations* directory, for example
```
cd migrations/
env CR_DB_USER=postgres CR_DB_PSW=<password_for_db_user> CR_DB_DATABASE=car_center CR_DB_ADDR=:5432 go run *.go up
```
(see readme in migrations/ dir.)

5. restore api dependencies with `dep ensure`

6. run service with command
```
env CR_ADDR=:8080 CR_DB_USER=postgres CR_DB_PSW=<password_for_db_user> CR_DB_DATABASE=duplicator CR_DB_ADDR=:5432 go run main.go
```
or with debug logs
```
env CR_DEBUG=1 CR_ADDR=:8080 CR_DB_USER=postgres CR_DB_PSW=<password_for_db_user> CR_DB_DATABASE=duplicator CR_DB_ADDR=:5432 go run main.go
```
