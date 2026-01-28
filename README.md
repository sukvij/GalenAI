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
    
    












{
  "full_name": "Vijendra Sukariya",
  "job_title": "Senior Software Engineer",
  "country": "India",
  "salary": 250000.00
}

{
  "full_name": "Aarav Sharma",
  "job_title": "Software Engineer",
  "country": "India",
  "salary": 120000.00
}

{
  "full_name": "Emily Johnson",
  "job_title": "Product Manager",
  "country": "USA",
  "salary": 135000.00
}


{
  "full_name": "Liam Brown",
  "job_title": "Data Analyst",
  "country": "UK",
  "salary": 90000.00
}


{
  "full_name": "Sophia Müller",
  "job_title": "UX Designer",
  "country": "Germany",
  "salary": 85000.00
}


{
  "full_name": "Noah Williams",
  "job_title": "DevOps Engineer",
  "country": "Canada",
  "salary": 110000.00
}


