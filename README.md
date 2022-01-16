# E-Voting System on Blockchain - (web)

## Installation :


1). Install GoLang 1.16+ , 
2).setup local mysql server with dbname=auth & tablename=user at port 3306

```console
mysql> use auth;
mysql> describe user;
+------------+------------------+------+-----+-------------------+-----------------------------------------------+
| Field      | Type             | Null | Key | Default           | Extra                                         |
+------------+------------------+------+-----+-------------------+-----------------------------------------------+
| id         | int unsigned     | NO   | PRI | NULL              | auto_increment                                |
| first_name | varchar(50)      | NO   |     | NULL              |                                               |
| last_name  | varchar(50)      | NO   |     | NULL              |                                               |
| email      | varchar(100)     | NO   | UNI | NULL              |                                               |
| password   | char(60)         | NO   |     | NULL              |                                               |
| status_id  | tinyint unsigned | NO   | MUL | 1                 |                                               |
| created_at | timestamp        | NO   |     | CURRENT_TIMESTAMP | DEFAULT_GENERATED                             |
| updated_at | timestamp        | NO   |     | CURRENT_TIMESTAMP | DEFAULT_GENERATED on update CURRENT_TIMESTAMP |
| deleted    | tinyint unsigned | NO   |     | 0                 |                                               |
+------------+------------------+------+-----+-------------------+-----------------------------------------------+
9 rows in set (0.00 sec)
```

3). Install all dependencies from go.mod file in the root folder .
by --->  $ go get <package_name> or go install <package_name>

```golang
module uio

go 1.16

require (
	github.com/boltdb/bolt v1.3.1 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/gorilla/context v1.1.1 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/gorilla/sessions v1.2.1 // indirect
	github.com/haisum/recaptcha v0.0.0-20170327142240-7d3b8053900e // indirect
	github.com/josephspurrier/csrfbanana v0.0.0-20170308132943-2c49e3597176 // indirect
	github.com/julienschmidt/httprouter v1.3.0 // indirect
	github.com/justinas/alice v1.2.0 // indirect
	github.com/mr-tron/base58 v1.2.0 // indirect
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519 // indirect
	golang.org/x/sys v0.0.0-20211025201205-69cdffdb9359 // indirect
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22 // indirect
)
```
4). Run/Build the project from main.go

```console
$ go run main.go
```

5). You should be able to see Application running on http-port::80 on Localhost on your system .
