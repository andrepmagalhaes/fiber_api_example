docker stop q2bank_postgres
docker rm q2bank_postgres
rm -rf /home/andre/Dev/q2bank_test/PostgreSQL
docker run --name q2bank_postgres -e "POSTGRES_PASSWORD=123456" -p 5432:5432 -v /home/andre/Dev/q2bank_test/PostgreSQL:/var/lib/postgresql/data -d postgres
sleep 5
tables="$(cat db_schema.sql)"
docker exec -it q2bank_postgres psql -U postgres -c "CREATE DATABASE q2bank
    WITH
    OWNER = postgres
    ENCODING = 'UTF8'
    CONNECTION LIMIT = -1
    IS_TEMPLATE = False;"

docker exec -it q2bank_postgres psql -U postgres -d q2bank -c "$tables"