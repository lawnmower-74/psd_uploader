## PSDのアップロード
```bash
$ docker-compose run app go run cmd/upload/main.go ./upload/example.psd
```

## PSDのダウンロード
```bash
# 1 = id
$ docker-compose run app go run cmd/retrieve/main.go 1 ./retrieved
```