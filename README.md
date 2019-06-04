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

Postgres:
You will need a postgres server. I use a quick docker container for this:
```
docker run --rm -it --publish 0.0.0.0:5432:5432 --name pg -e POSTGRES_PASSWORD=postgres postgres:alpine
```

Run the application
```
./pasteit
```

### Environment Variables

| Variable  | Purpose  | Default  |
|-----------|----------|----------|
| ADDRESS   | Address server listens on  | 0.0.0.0  |
| PORT      | Post server listens on  | 8080  |
| HOSTNAME  | DNS of server  | "address:port"  |
| PG_ADDR   | Postgress Address  | 127.0.0.1  |
| PG_PORT   | Postgres Port  | 5432  |
| PG_USER   | Postgres User  | postgres |
| PG_PASS   | Postgres Pass  | postgres  |
| PG_DBNAME | Postgres Database Name  | postgres  |
| PG_SSL    | Postgres SSL Mode | disabled  |
