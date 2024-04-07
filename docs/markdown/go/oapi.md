# oapi

## bundle & validate
oapiのbundle
- oapiのドキュメントを修正した時に実施
```shell
$ make oapi/gen/app
$ make oapi/validate/app
```

## ui
uiをwebで確認
```shell
$ make oapi/run/app
```
アクセスURL  
http://localhost:8081  

## codegen
下記のファイルを自動生成  
./internal/presentation/http/app/oapi/codegen.go  
```shell
$ make oapi/codegen/app
```
