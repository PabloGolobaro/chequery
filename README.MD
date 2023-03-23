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
Gets new orders by endpoint PUT:http://localhost/api/v1/order.
Starts asynchronous task to generate pdf files for new checks and then updates check status in DB.
Sends generated check pdf files to client listener and updates status of generated ones.

## Features:
- Swagger documentation
- Rest API
- Generating pdf output from incoming orders using WkHtmlToPdf
- Clear architecture pattern (I tried)
- Postgres as DB
- concurrency patterns (I tried)

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

### Api endpoint example:
- create order
     curl -X POST --location "http://localhost/api/v1/order?point_id=1" \
     -H "Content-Type: application/json" \
     -d "{
     \"order\": \"{\\\"meat\\\": 1} \"
     }"