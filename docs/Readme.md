# Bank Application (DEV DOCS)

## Database Design and Setup

Design the database schema diagram using `dbdiagram.io` and convert the diagram to `SQL` script.

![DB Schema](./img/Bank_Application.png)

The `dbml` file used to generate the sql script is available [here](./db_diagram_io.dbml).

We will use `Docker` to set up the `PostgreSQL` database. Here we will use the `Postgres:17-alpine` image from `Docker Hub`.

1. Pull the `PostgreSQL` image from `Docker Hub`.

```bash
docker pull postgres:17-alpine
```

2. Run the `PostgreSQL` container.

```bash
docker run --name bank_app_db -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secretpassword -d postgres:17-alpine
```

> The above command will create a container named `bank_app_db` with the `Postgres:17-alpine` image. We map the host port `5432` to the container port `5432`. We also set the `POSTGRES_USER` and `POSTGRES_PASSWORD` environment variables. The container will run in detached mode (`-d` means the container will run in the background).

3. Connect to the `PostgreSQL` database.

```bash
docker exec -it bank_app_db psql -U root
```

> The above command will connect to the `PostgreSQL` database using the `psql` command-line tool. We use the `-U` flag to specify the user (`root` in this case).

### Create the migration files 

We will use `golang-migrate` to create the migration files. The migration files are used to create the database schema and seed the database with initial data.

