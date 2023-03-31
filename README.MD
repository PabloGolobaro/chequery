# chequery

## Check generation service
### Simple example microservice application. Main components:
- client listener
- api server

### Functionality:
#### Client:
Listens api server for new check appearing.
If one appears client gets it's pdf file and sends it to dedicated printers, then updates checks status.

#### Server:
Starts asynchronous task to generate pdf files for new checks and then updates check status in DB.
Sends generated check pdf files to client listener and updates status of generated ones.

## Features:
- Swagger documentation
- Rest API
- Generating pdf output from incoming orders using WkHtmlToPdf
- Clear architecture pattern (I tried)
- Postgres as DB
- concurrency patterns
- JWT authorization
- simple Web UI
- Swagger v2.0 documentation

## How to start up project:
- Start application with all services:
``make server``
``make client``

#### All environment variables are set in Makefile:
     DB_PORT = 5432
     DB_HOST = localhost
     DB_USER = golobar
     DB_PASS = password
     DB_NAME = golo
     SERVER_HOST = localhost
     SERVER_PORT = 80
     JWT_SECRET = secret

### Usage:
#### You can use app functionality via web UI at ***http://localhost:80*** or by using api documented in swagger 2.0 format.
#### Example built in parameters to get Auth token
- username = "JohnSnow"
- password = "password"
- ``http://localhost/docs/  swagger documentation endpoint``




### Api endpoint example:
#####create order
``curl -X POST --location "http://localhost/api/v1/order" \
-H "Authorization: Bearer <your generated token>" \
     -d "{
     \"point_id\": 1,
     \"products\": [
     {
     \"name\": \"Meat\",
     \"price\": 145,
     \"quantity\": 3
     },
     {
     \"name\": \"vegetables\",
     \"price\": 32,
     \"quantity\": 2
     },
     {
     \"name\": \"Juice\",
     \"price\": 48,
     \"quantity\": 1
     }
     ]
     }"``

