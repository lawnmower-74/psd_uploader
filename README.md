## PSDのアップロード
```bash
$ docker-compose run app go run upload.go ./upload/xxxxx.psd
```

## PSDのダウンロード
```bash
# 1 = id
$ docker-compose run app go run retrieve.go 1 ./uploaded
```