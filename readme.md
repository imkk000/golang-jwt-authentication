# Generate Token
```sh
# start jwt authentication server
go run generate_token/main.go
# get token
curl localhost:8080
```

```sh
curl -H "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzdWIiOiIxNTQ2NTg5OTQ1IiwiZGF0YSI6eyJ1c2VybmFtZSI6ImRlYnVnZ2luZyJ9LCJpYXQiOiIxNTQ2NTg5OTQ1In0=.sCRlyTUP4wTySr6bzqzjnskY0Js5N4gb3Xwpy69x5o4=" localhost
```