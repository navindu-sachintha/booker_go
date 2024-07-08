                     
<h1 align="center" style="font-weight: bold;">Booker 📚</h1>

<p align="center">
<a href="#tech">| Technologies |</a>
<a href="#started">Getting Started |</a>
<a href="#routes">API Endpoints |</a> 
</p>


<p align="center">Backend API for Booksore powered with golang</p>


<p align="center">
<a href="https://github.com/navindu-sachintha/booker_go">Visit the Project</a>
</p>
 
<h2 id="technologies">💻 Technologies</h2>

- Go toolchain
- goose (for database migration)
- SQLC (For generate Go codes to queries)
- POSTGRE SQL
- Docker (For database instance & containorize app)
- GIN for web framework
 
<h2 id="started">🚀 Getting started</h2>

Here you describe how to run your project locally
 
<h3>Prerequisites</h3>

Here you list all prerequisites necessary for running your project. For example:

- [Go Toolchain](https://go.dev/)
- [Git](https://git-scm.com/)
- [Docker](https://docs.docker.com/)
- [goose](https://pressly.github.io/goose/)
 
<h3>Cloning</h3>


```bash
git clone https://github.com/navindu-sachintha/booker_go.git
```
 
<h3>Config .env variables</h2>

Use the `.env.example` as reference to create your configuration file `.env` with your environment variables in root

```yaml
DB_URL=postgresql://<your_username>:<your_password>@localhost:5432/<your_db_name>?sslmode=disable
```
<h3>Edit docker-compose.yml file</h3>
Make sure to replace `docker-compose.yml` with your database info

```yaml
environment:
      POSTGRES_USER: your_username
      POSTGRES_PASSWORD: your_password
      POSTGRES_DB: your_database_name
```
 
<h3>Starting</h3>

How to start your project

```bash
cd booker_go

# Execute docker postgres instance
docker compose run db

cd /sql/schema

# Run Database migration
goose postgres DB_URL up

cd ../../

# Download dependencies
go mod download
go mod tidy && go mod vendor

cd /cmd/main

# Build go Application
go go build -o ../../bookstoreapp

cd ../../
./bookstoreapp
```
 
<h2 id="routes">📍 API Endpoints</h2>

​
| route               | description                                          
|----------------------|------------------------------
| <kbd>POST /book</kbd>     | Create book in database 
| <kbd>GET /books</kbd>     | Get all available books 
| <kbd>GET /book/:id</kbd> | Get Book by its ID
|<kbd>PUT /book/:id</kbd>| Update Book name according to id
|<kbd>DELETE /book</kbd>| Delete book by it's id

`Since it using docker volume data will be not removed even after `