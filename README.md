# Попова Людмила

## Description:

A simple http server. The server has an endpoint ```GET /users``` and return an array of users from a file users.json in json format. 

1. Download the repo

2. cd to the project directory

3. Execute next command ```go run *.go```

4. Try this:

```bash
curl --location --request GET 'http://localhost:3333/users' \
--header 'Content-Type: application/json'
``` 

It will show a result:

```json
[
  {"name": "user1", "email": "user1@mail.com"},
  {"name": "user2", "email": "user2@mail.com"},
]
```
