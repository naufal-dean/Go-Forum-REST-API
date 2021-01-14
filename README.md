# Onboarding Dean

## Description
This repository contains a Golang RESTful API used for a forum activities. It provides authentication using JWT, CRUD for Thread, and CRUD for Post.

## Author
Naufal Dean Anugrah - [naufal.dean@pinhome.id](naufal.dean@pinhome.id)

## Dependencies
1. Golang (tested in `go version go1.15.6 windows/amd64`)
2. PostgreSQL
3. Docker (optional). Used to containerize the app and the database.
4. Postman (optional). Used to test the API endpoints.
5. Browser (optional). Used to see the API documentation.

## How to run the app
### Preparation
1. Copy file `.env.example` to `.env`.
2. Set the environment variable value in `.env` if necessary.
3. Install Golang dependency using command `go mod tidy`.
### Run
#### Option 1 (using docker)
1. Run command `docker-compose up`.
#### Option 2
1. Run the PostgreSQL.
2. Run command `go build -o main.exe && main.exe`.
### Access
1. The app is hosted in `localhost:8080`.

## How to run the test
### Preparation
1. Copy file `.env.test.example` to `.env.test`.
2. Set the environment variable value in `.env.test` if necessary.
3. (If not yet done in the app run preparation) Install Golang dependency using command `go mod tidy`.
### Run
1. Run command `go test ./... -count 1 -p 1`.

## API routes
After the app is running. See the routes documentation in [localhost:8080/api/docs/](localhost:8080/api/docs/)

## Testing design
The test design is following the Golang convention. Test file is placed in the same package with the tested file
using `*_test.go` name. It is table-driven test, which one table entry is full set of the test case.\
There are some advantages of creating test file in the same package with the tested file:
1. It is clearer which file is being tested.
2. Unexported object can be accessed from the test file.

## Project structure design
```
├───app  # the code for the app
│   ├───controller  # controller/handler implementations
│   │   └───v1
│   │       ├───auth
│   │       ├───posts
│   │       ├───threads
│   │       └───users
│   ├───core        # the app dependency wrapper
│   ├───lib         # custom utilities implementations
│   │   ├───auth
│   │   ├───hash
│   │   └───util
│   ├───log         # log file saved here
│   ├───middleware  # middleware implementations
│   ├───model       # model definitions (orm, custom error, etc)
│   │   ├───cerror
│   │   └───orm
│   ├───response    # response function and general response data
│   │   ├───data    # general response data struct definitions
│   │   │
│   │   ├───json.go
│   │   ├───error.go
│   │   └───success.go
│   ├───route       # route definition and its handler
│   │   └───v1
│   ├───seed        # database seeder
│   ├───static      # static file
│   │   └───swaggerui    # the swagger-ui
│   │       └───statik   # compiled swagger-ui
│   └───test        # global test dependency helper (app, db, etc)
│
└───main.go
```

## Routing and handler design
The route setup is inside the `/app/route` folder. It is wrapped in `api/v{version}` prefix. Then each resources
are assigned to different prefix (some plural nouns), in this case `users`, `threads`, and `posts` prefix. For
each resource generally there are 5 routes (create, get all, get one, update, delete). Each use some specific http
method and some using id path parameter. The details can be seen in the API documentation.

The handler is inside the `/app/controller`. The dependency (database connection, validate object) is wrapped in
`*core.App` struct and passed to the handler via dependency injection. That process is executed in the 
`app/init.go` file.

## DTO vs domain models
DTO (Domain Transfer Object) is used for transferring the data and do not have logic. In this app the DTO is the json
used for the input (in create and update request for example) and output/response to the client. The DTO can be seen in
the API documentations in the `Schemas` part.

Domain models is can implements some logic, for example that implemented as methods. The domain models in this app is
the Golang structure used for gorm for example. In the `orm.User` model there are `PasswordValid` method to check is
the supplied password valid and `BeforeCreate` hook that hash the user password before saved to the database.
