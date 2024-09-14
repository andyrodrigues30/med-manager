goose create 01_create_users sql

GOOSE_DRIVER=postgres GOOSE_DBSTRING=postgres://postgres:UeGR570V2VdQgTkjdPJA@localhost/taskDB?sslmode=disable GOOSE_MIGRATION_DIR=migrations goose status
GOOSE_DRIVER=postgres GOOSE_DBSTRING=postgres://postgres:UeGR570V2VdQgTkjdPJA@localhost/taskDB?sslmode=disable GOOSE_MIGRATION_DIR=migrations goose up

GOOSE_MIGRATION_DIR=migrations goose create 02_add_meds sql
GOOSE_MIGRATION_DIR=migrations goose up