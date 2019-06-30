# echo-sample

This is an sample of todo app.

## Feature

- Implemented in the Go language.
- Using [labstack/echo](https://github.com/labstack/echo) as Web API Framework.
- Architecture adopts CleanArchitecture.

## Run

### Run command

```sh
docker-compose up -d
go run *go
```

### Startup confirmation

Access to `localhost:1323/` after above execution.  
Then you will see `Hello world!` in your browser.

### Exapmple

```sh
curl localhost:1323/todos
curl localhost:1323/todos/1
curl -i -H "Accept: application/json" -H "Content-type: application/json" -X POST -d '{"ID": 4, "Title": "todo4","Content": "sample-todo","Status": "0"}' localhost:1323/todos
```

## License

MIT
