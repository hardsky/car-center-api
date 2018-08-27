
This directory contains migrations for database.
Its are sql scripts injected in go files on top of go-pg ORM.
see https://github.com/go-pg/migrations

use command:  
`go run *.go up` - to apply all migrations on database  
run  
`go run *.go --help` - to read usage details.

We use *dep* as dependency manager.
After git clone, please, run :

```
dep ensure
```

env variables:

- CR_DB_ADDR
- CR_DB_USER
- CR_DB_PSW
- CR_DB_DATABASE
