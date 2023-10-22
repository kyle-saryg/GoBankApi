# GoBank API

## Summary:
Mock banking RESTful api written in Go with Postgresql. Supports account creation (with password encryption), and utilizes JWT sessions for authorized users (issued upon login). 

## Usage:
1. Build gobank-api-image dockerfile (Must be in gobankapi top directory)\
> docker build -t gobank-api-image

2. Compose docker images file
> docker-compose up

3. Can now send http requests to localhost:3000/<ENDPOINT>


# EndPoints:
## /login
### Supported Methods:
 - POST (Response provides JWT)
    > Body: {\
    "number": int64,\
    "password": string  
    }\
    Response: {\
    "number": int64\
    "token": string\
    }

## /account
### Supported Methods:
 - GET (fetches all accounts)
    > Response:\
    [{\
    "id": int,\
    "firstName": string,\
    "lastName": string,\
    "number": int64,\
    "encryptedPassword": string,\
    "balance": int,\
    "createdAt": string (UTC format)\
    }, ... ]
 - POST (Creates a new account)
    > Body: {\
    "firstName": string,\
    "lastName": string,\
    "password": string  
    }\
    Response:
    {\
    "id": int,\
    "firstName": string,\
    "lastName": string,\
    "number": int64,\
    "encryptedPassword": string,\
    "balance": int,\
    "createdAt": string (UTC format)\
    }

## /account/\<id>
### Supported Methods:
 - GET (gets account by id, needs JWT in header)
    > Header: {\
    "x-jwt-token": string  
    }\
    Response:
    {\
    "id": int,\
    "firstName": string,\
    "lastName": string,\
    "number": int64,\
    "encryptedPassword": string,\
    "balance": int,\
    "createdAt": string (UTC format)\
    }

 - DELETE (needs JWT)
    > Header: {\
    "x-jwt-token": string  
    }\
    Response:
    {\
    "deleted": int (id of deleted account)\
    }
