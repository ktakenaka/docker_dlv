# docker_dlv
```
$ docker-compose up -d
$ dlv connect localhost:2345
(dlv) break /docker_dlv/app/main.go:17
Breakpoint 1 set at 0x6e405e for main.main.func1() /docker_dlv/app/main.go:17
```

In another terminal
```
curl http://localhost:8080/
```

Then...
```
(dlv) continue
> main.main.func1() /docker_dlv/app/main.go:17 (hits goroutine(8):1 total:1) (PC: 0x6e405e)
(dlv) print w
net/http.ResponseWriter(*net/http.response) *{
        conn: *net/http.conn {
```          
