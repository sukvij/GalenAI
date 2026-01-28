# GalenAI


# migrations
    docker ps
    docker compose up --build
    docker ps
    docker exec -it d575b16daef1 sh      go inseide docker  and run ls to see all the folders named app     and   ls /app   inside app
    docker exec -it d399b68226a4 psql -U postgres    enter postgres
        \l     shows all the dbs
        \c cruddb  connect to db cruddb
        \dt      show tables


    docker compose run --rm migrate up  
    docker compose run --rm migrate down 1     -- last migration
    docker compose run --rm migrate down     -- all down
    docker compose run --rm migrate version      current version


    migrate -path ./database/migrations \
     -database "postgres://postgres:postgres@localhost:5432/cruddb?sslmode=disable" \
     up     --- locally migrate
    
    