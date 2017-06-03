## サンプル
- 起動
  - docker run -d -p 8080:8080 tarent/loginsrv -jwt-secret my_secret -simple bob=secret
- ブラウザ
  - http://localhost:8080/login
  - user: bob, password: secret
- JWT取得
  - default
    - `curl -i --data "username=bob&password=secret" http://127.0.0.1:8080/login`
  - json
    - `curl -i -H 'Content-Type: application/json'  --data '{"username": "bob", "password": "secret"}' http://127.0.0.1:6789/login`
  - web based
    - `curl -i -H 'Accept: text/html' --data "username=bob&password=secret" http://127.0.0.1:6789/login`


- help
  - `docker run --rm tarent/loginsrv --help`

## githubサンプル
- 起動
  - `docker run -p 8080:8080 tarent/loginsrv -github client_id=xxx,client_secret=yyy`



## 連携サンプル
- 準備
```
go get github.com/dgrijalva/jwt-go
go get github.com/mattn/goreman
go get github.com/mattn/gorem
```
- 起動
  - goreman start


## JWTについて
- 参考
  - http://dev.classmethod.jp/server-side/json-signing-jws-jwt-usecase
