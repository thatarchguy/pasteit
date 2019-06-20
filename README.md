# Paste.it

Inspiration from [https://ix.io](https://ix.io)

```
$ echo -n "<h1>lol</h1>" | curl -F 'f=@-' localhost:8080
http://localhost:8080/abl
```
```
$ cat README.md | curl -F 'f=@-' localhost:8080
http://localhost:8080/aFa
```

.bashrc alias
```
alias pasteit='curl -F 'f=@-' localhost:8080'
# cat file | pasteit
```

## Deploying

Databases supported are **Mysql**, **Postgres**, and **Sqlite3**

Postgres:
I use a quick docker container for this:
```
$ docker run --rm -it --publish 0.0.0.0:5432:5432 --name pg -e POSTGRES_PASSWORD=postgres postgres:alpine

$ ./pasteit -database postgres://postgres:postgres@localhost/postgres?sslmode=disable
```

Sqlite:
```
$ ./pasteit -database sqlite://:memory:
```

### CLI & Environment Variables

| Variable  | Purpose  | Default  |
|-----------|----------|----------|
| ADDRESS   | Address server listens on  | 0.0.0.0  |
| PORT      | Post server listens on  | 8080  |
| HOSTNAME  | DNS of server  | "address:port"  |
| DATABASE  | Database Address | sqlite://:memory: |
