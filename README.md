# GalenAI

# run locally

# run using docker
  run command -- 
    docker compose up --build     --> this will create container and run it (postgres and app container)
    docker exec -it {postgres container name} psql -U postgres   --> this will land inside postgres container
      \c cruddb  --> run this inside postgres container --> land inside cruddb database

      CREATE TABLE employees (
          id BIGSERIAL PRIMARY KEY,
          full_name VARCHAR(255) NOT NULL,
          job_title VARCHAR(150) NOT NULL,
          country VARCHAR(100) NOT NULL,
          salary NUMERIC(12,2) NOT NULL,
          created_at TIMESTAMP NOT NULL DEFAULT NOW(),
          updated_at TIMESTAMP NOT NULL DEFAULT NOW()
      );

      CREATE INDEX idx_employees_country ON employees(country);
      CREATE INDEX idx_employees_job_title ON employees(job_title);

      CREATE TABLE users (
          id SERIAL PRIMARY KEY,
          user_name VARCHAR(50) UNIQUE NOT NULL,
          password TEXT NOT NULL,
          role VARCHAR(20) DEFAULT 'user',
          created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
          updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
      );

      CREATE INDEX idx_users_username ON users(user_name);


      run commands one by one will create tables and indexes

      \dt   --> will show all tables  {employees, and users}
    
  

  POST: http://localhost:8080/v1/users/register        {"user_name":"sukvij","password":"1234","role":"admin"}      create user
  POST: http://localhost:8080/v1/users/login     {"user_name":"sukvij","password":"1234"}       registratiion --> send jwt token--use token for request

  POST: http://localhost:8080/v1/employees   -- create employee   { "full_name": "Vijendra Sukariya", "job_title": "Senior Software Engineer", "country": "india", "salary": 150000.00 }   send jwt token also

  GET: http://localhost:8080/v1/employees    -- get all employees
  GET: http://localhost:8080/v1/employees/1   -- get employee by id
  GET: http://localhost:8080/v1/salary-metrics/country/{india}     country wise salary metrics

  GET: http://localhost:8080/v1/salary-metrics/job_title/Senior Software Engineer    job title wise salary metrics
  GET: http://localhost:8080/v1/salary-calculation/3    salary calculation by id


# api overflow
  registration --> login (jwt creation) --> create employee --> {salary-calculation, salary-metrics}

# auth flow
  registration (create user)
  login (create token)

# further implementation
  database migration 









some dummy employees datas

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


