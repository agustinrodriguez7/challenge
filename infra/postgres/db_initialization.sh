docker-compose up -d
docker-compose run db bash
psql  --host=db --username=dbuser --dbname=todoapp