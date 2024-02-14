# HTTP-Go

Go のテストコンテナ

```bash
podman build -t http-go .
podman run -p 8080:80 localhost/http-go:latest
curl http://localhost:8080
```
