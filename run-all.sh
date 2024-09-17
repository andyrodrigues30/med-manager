# create network
docker network create med-network

# start the database container
docker compose -f ./db/compose.yml up -d

# start the account service api
docker compose -f ./account-service/compose.yml up