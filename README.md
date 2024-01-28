# URL shortener implementation with golang

## How to Use

- Clone and run the main file with `go run main.go`
- To create short URL, go to http://localhost:9808/create-short-url with json data

```{
  "long_url": "https://github.com/akbarsigit",
  "user_id" : "123"
  }
```

- To check the redirection function, you can go to the shortened URL that will redirect to the initial URL.

`"http://localhost:9808/ACXyeAfc"`
