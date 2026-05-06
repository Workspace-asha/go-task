# Task API (Go + PostgreSQL + JWT + Docker + HTTPS)

A small, production-style REST API demonstrating:

* Clean architecture (Handler → Service → Repository)
* JWT authentication
* PostgreSQL (relational DB)
* Dockerized local setup
* Kubernetes-ready manifests
* **HTTPS enabled (self-signed certificates)** 

## Prerequisites
Docker (Desktop or Engine)


##Generate HTTPS Certificates (one-time)

mkdir certs

openssl req -x509 -newkey rsa:4096 \
  -keyout certs/server.key \
  -out certs/server.crt \
  -days 365 -nodes \
  -subj "/CN=localhost"

# BUILD AND TEST SERVICE

1) Build & Run

docker compose up --build

This starts:

* db → PostgreSQL (port 5432)
* api → Go service (HTTPS on port 8080)

2) Initialize Database (one-time)

docker exec -it $(docker ps -qf "ancestor=postgres:16") \
psql -U postgres -d tasks -c "
CREATE TABLE IF NOT EXISTS projects (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  description TEXT
);"

3) Get JWT Token (HTTPS)

TOKEN=$(curl -sk -X POST https://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"password"}' \
  | jq -r .token)


OR below command to get the jwt token

curl -k -X POST https://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"password"}'


4) End-to-End Test (CRUD over HTTPS)

### Create Project

curl -k -X POST https://localhost:8080/projects \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"name":"Demo","description":"Test project"}'

### List Project (with pagination)
curl -k "https://localhost:8080/projects?page=1&limit=10" \
  -H "Authorization: Bearer $TOKEN"


### Get by ID

curl -k https://localhost:8080/projects/1 \
  -H "Authorization: Bearer $TOKEN"


### Update Project

curl -k -X PUT https://localhost:8080/projects/1 \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"name":"Updated","description":"Updated desc"}'


### Delete Project

curl -k -X DELETE https://localhost:8080/projects/1 \
  -H "Authorization: Bearer $TOKEN"


# Health Check (HTTPS)
curl -k https://localhost:8080/health

