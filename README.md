uth0 - 嘘
---

a dumb and fake JWT authentication server for use with [Hasura](https://docs.hasura.io/1.0/graphql/manual/auth/authentication/jwt.html)

## Run

    USER=nyan PASS=hoyo SECRET=ultrasupersecret go run main.go

for example

    curl -X POST -d 'user=nyan' -d 'pass=hoyo' http://localhost:8080/auth

    > POST /auth HTTP/1.1
    > Host: localhost:8044
    > User-Agent: curl/7.54.0
    > Accept: */*
    > Content-Length: 17
    > Content-Type: application/x-www-form-urlencoded
    >

    < HTTP/1.1 200 OK
    < Content-Type: application/json; charset=UTF-8
    < Date: Tue, 24 Sep 2019 07:02:26 GMT
    < Content-Length: 293
    <
    {"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJodHRwczovL2hhc3VyYS5pby9qd3QvY2xhaW1zIjp7IngtaGFzdXJhLWFsbG93ZWQtcm9sZXMiOlsidXNlciJdLCJ4LWhhc3VyYS1kZWZhdWx0LXJvbGUiOiJ1c2VyIn0sImV4cCI6MTU2OTM1MTc0NiwiaWF0IjoxNTY5MzA4NTQ2LCJpc3MiOiJ1c28vMC4xIn0._MSaZ1pbkJYviwD1SitpR_kS8FqIvddx0RpPMS_8NIA"}

this claims:

    {
      "https://hasura.io/jwt/claims": {
        "x-hasura-allowed-roles": [
          "user"
        ],
        "x-hasura-default-role": "user"
      },
      "exp": 1569351746,
      "iat": 1569308546,
      "iss": "uso/0.1"
    }

