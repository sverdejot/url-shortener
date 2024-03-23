# `url-shortener`: an url shortener written in go

To run the server:

```bash
make
./url-shortener
```

You can try via:

```bash
curl -X POST http://localhost:8080/shorten \
    -H "Content-Type: application/json" \
    -d '{"long_url": "https://www.sverdejot.dev"}'
```
