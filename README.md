# ctfsendai2018

## access command

### register

```
curl http://localhost:8080/api/v1/register -X POST -d '{"email":"hoge@hoge.com","password":"pass"}'
```

### login

```
curl http://localhost:8080/api/v1/login -X POST -d '{"email":"hoge@hoge.com","password":"pass"}'
```

### read users

```
curl http://localhost:8080/api/v1/users
```

### add user

	EMail    string `json:"email"`
	Name     string `json:"name"`
	Age      uint8  `json:"age"`
	Note     string `json:"note"`
	Auth     uint8  `json:"auth"`

```
curl http://localhost:8080/api/v1/users -X POST -d '{"id":"aG9nZUBob2dlLmNvbQ","email":"hoge@hoge.com","name":"hoge hoge","age":18}' -H 'Authorization:eyJhbGciOiJub25lIiwidHlwIjoiSldUIn0.eyJhdXRoIjoyLCJleHAiOjE5MjQ5NTIzOTksInVzZXIiOiJob2dlIGhvZ2UifQ.'
```
