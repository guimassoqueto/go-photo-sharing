# go-photo-sharing

## How to Run

Required Golang version: 1.21.5

1. Install dependencies
```shell
go get
```
2. Create the .env file from .env.sample
```shell
cp .env.sample .env
```

3. Install [Modd](https://github.com/cortesi/modd)

4. Start dev
```shell
make dev
```

## Accessing the Database
By default the database UI runs at localhost:5000, but if you have interest in changing it,
modify the value of the variable ADMINER_PORT in the .env


## Modifying database schema (migrations)
To change the scema of the tables or another database modification open the migration repository: https://github.com/guimassoqueto/migrations/
the go-photo-sharing folder takes care of the database migrations
