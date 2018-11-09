# ctfsendai2018

[仙台CTF 2018 セキュリティ技術競技会（CTF）](https://www.sendai-ctf.org/event/ctf2018-security-kaisaiyoko-2/) で使用したWebの問題です。

## 問題

### Quest 1

自社で作ったRESTful APIサーバーにペネトレーションテストをしてもらったところ、制限無しでユーザー情報を取得できてしまう問題が見つかりました。http://localhost:8080 にアクセスしてフラグを探してください。

### Quest 2

どうやら、テストということで joeアカウント が使われているアカウントがあるみたいです。joeアカウントが使われているアカウント見つけてログインしてください。

### Quest 3

なんとログイン時に生成されたトークンの作り方に問題があるようです。2. で見つけたユーザーアカウントを使ってログインした後、取得したトークンを解析して、管理者権限で http://localhost:8080 にデータを登録してください。取得したトークンは、Authorizationに設定します。

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
