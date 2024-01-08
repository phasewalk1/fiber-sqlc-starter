# fiber-sqlc-starter
This is a template repository for backend applications using Go, [Fiber](https://gofiber.io/), [sqlc](https://sqlc.dev), and PostgreSQL.

## Features
* [Fiber](https://gofiber.io): A lightning fast web framework
* [sqlc](https://sqlc.dev): Generates type-safe code from SQL
* [Air](https://github.com/cosmtrek/air): Hot reloading
* PostgreSQL: The best database (sorry).
* Docker-ready: Example Docker configuration built in

## Getting Started
### Template Usage
Navigate to the top right and select *use this template*, and then clone the repository
```shell
git clone https://github.com/<your-username>/<generated-repo-name>
```
### Setup Postgres
#### Environment
You'll need the following environment files set in a `.env` file:
```shell
DB_USER="your-db-user"
DB_PASSWORD="your-db-password"
DB_NAME="your-db-name"
DB_HOST="your-db-host"
DB_PORT="your-db-port"
```
#### Using `docker-compose`
You can scaffold a local Postgres instance for dev use by navigating to the root directory and running:
```
docker-compose -f docker/docker-compose.yml --env-file .env up
```
#### Running the Application w/ Air
If you don't have [Air](https://github.com/cosmtrek/air) installed, you can get it by running:
```shell
go install github.com/cosmtrek/air@latest
```
Then from the root directory, run the application with
```shell
air
```
