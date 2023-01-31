# export $(grep -v '^#' .env | xargs)

# tables="$(cat db_schema.sql)"
# docker exec -it q2bank_test-db-1 psql -U postgres -c "CREATE DATABASE $POSTGRES_DB
#     WITH
#     OWNER = postgres
#     ENCODING = 'UTF8'
#     CONNECTION LIMIT = -1
#     IS_TEMPLATE = False;"

# echo "docker exec -it q2bank_test-db-1 psql -U $POSTGRES_USER -d $POSTGRES_DB -c "$tables""