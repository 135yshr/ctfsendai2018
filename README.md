# ctfsendai2018

[仙台CTF 2018 セキュリティ技術競技会（CTF）](https://www.sendai-ctf.org/event/ctf2018-security-kaisaiyoko-2/) で使用したWebの問題です。

## 問題

### Quest 1

RESTful APIサーバーのセキュリティ診断をしたところ、制限無しでユーザー情報を取得できてしまう問題が見つかりました。
RESTful APIサーバーにアクセスしてフラグを探してください。

  1. RESTful API のバージョンは、1
  2. 一般的には、/api/v1/users や /api/v1/products といったURLで情報を操作します。

*診断対象サーバ*
http://localhost:8080

### Quest 2

セキュリティ診断を続けたところ、どうやら、テストということでjoeアカウントが使われているアカウントがあるようです。

joeアカウントが使われているアカウント見つけてログインしてください。

  1. joeアカウントとは、ユーザーのIDとパスワードが同じアカウントのことを指します。
  2. /api/v1/login にメールアドレスとパスワードをJSON形式でPOSTするとログインすることができます。
  3. JSONはこんな形式です。{"email":"メールアドレス","password":"パスワード"}

*診断対象サーバ*
http://localhost:8080/api/v1/login

### Quest 3

ログイン時に生成されたトークンの作り方に問題があるようです。
Pen.2で見つけたユーザーアカウントを使ってログインした後、取得したトークンを解析して、管理者権限でデータを登録してください。
取得したトークンは、Authorizationに設定します。

*診断対象サーバ*
http://localhost:8080/api/v1/users

## 問題サーバーの使い方

事前に[Docker](https://www.docker.com/products/docker-desktop)をインストールしてください

### Docker イメージの作成

```
docker build ./ -t sendaictf2018
```

### 作成したイメージの作成

```
docker run -it --rm -p 8080:8080 --name sendaictf2018 sendaictf2018
```

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
curl http://180.42.14.119:8080/api/v1/users -X POST -d '{"id":"aG9nZUBob2dlLmNvbQ","email":"hoge@hoge.com","name":"hoge hoge","age":18}' -H 'Authorization:eyJhbGciOiJub25lIiwidHlwIjoiSldUIn0.eyJhdXRoIjoxLCJleHAiOjE5MjQ5NTIzOTksInVzZXIiOiJLb21hdHN1TW9yaXRvbW8ifeKAiw.'
```
