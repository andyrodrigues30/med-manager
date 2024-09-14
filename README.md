# Med Manager

A simple self-hosted application for managing medications and tracking health such as glucose levels, heart rate, blood pressure.
This application stores the data in a PostgreSQl database running in a docker container (all services run in a docker container) with simple authentication.

## The Services
- Database

- Frontend Web UI

- Account API Service
- Health Tracker API Service
- Analytics API Service

## Running the Application
To run all services:
1. Ensure all the `.env` files are created and located in the correct directories.
2. In the project root directory run: `docker compose up -d`

Alternatively run services seperately follow the following steps.

### Step 1: The database
#### Environment Variables
Ensure the `.env` file has been created in the root directory with the following environment variables (ensure you add the postgres user name and password):

```
POSTGRES_USER=<POSTGRES_USER>
POSTGRES_PASSWORD=<POSTGRES_PASSWORD>

POSTGRES_DB=medManagerDB
DATABASE_URL=postgres://"$POSTGRES_USER":"$POSTGRES_PASSWORD"@localhost/"$POSTGRES_DB"?sslmode=disable

GOOSE_MIGRATION_DIR=./migrations
GOOSE_DRIVER=postgres
GOOSE_DBSTRING="$DATABASE_URL"
```

#### Database Migrations
Make `goose.sh` executeable: `chmod +x goose.sh`
Apply the database migrations with `./goose.sh up`

Check migration status with `./goose.sh status`

#### Run The Database ONLY
To run the Database only run the following docker command: `docker compose up med-manager-db -d`.

### Step 2: The Account API Service
#### Environment Variables
Ensure the `.env` file has been created in the `account-service` directory with the following environment variables (ensure you add the postgres user name and password):

```
PORT=8000
```