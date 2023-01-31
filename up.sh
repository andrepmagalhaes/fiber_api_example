export $(grep -v '^#' .env | xargs)

path="$(pwd)"
rm -rf $path/PostgreSQL

docker compose down
docker compose up -d

sleep 2

tables="$(cat db_schema.sql)"
docker exec -it q2bank_test-db-1 psql -U $POSTGRES_USER -d $POSTGRES_DB -c "$tables"