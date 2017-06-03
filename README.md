## サンプル
- 起動
  - docker run -d -p 8080:8080 tarent/loginsrv -jwt-secret my_secret -simple bob=secret
- ブラウザ
  - http://localhost:8080/login
  - user: bob, password: secret
- JWT取得
  - `curl --data "username=bob&password=secret" 127.0.0.1:8080/login`
