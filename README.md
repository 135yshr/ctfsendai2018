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
