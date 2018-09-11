# cagoproxy

HTTPリクエストをcurlコマンドに変換しつつリバースプロキシする君

# Start ProxyServer

```
$ cagoproxy --verbose -listen=8000 -proxy_pass='http://localhost:8080' -logfile=/tmp/cagoproxy.log
```

# Tail log

```
$ tail -f /tmp/hoge.log| jq .
```

```
{
  "level": "info",
  "msg": "curl -X 'GET' -d '' -H 'Accept: */*' -H 'User-Agent: curl/7.54.0' 'http://localhost:8080/README.md'",
  "time": "2018-09-11T19:42:02+09:00"
}
{
  "level": "info",
  "msg": "curl -X 'GET' -d '' -H 'Accept: */*' -H 'User-Agent: curl/7.54.0' 'http://localhost:8080/README.md'",
  "time": "2018-09-11T19:42:03+09:00"
}
{
  "level": "info",
  "msg": "curl -X 'GET' -d '' -H 'Accept: */*' -H 'User-Agent: curl/7.54.0' 'http://localhost:8080/README.md'",
  "time": "2018-09-11T19:42:03+09:00"
}
```

# for Developer
```
$ dep ensure
$ go build
TODO: multi platform build support...
```
