# Go REST API
## echo v4 based

### Public routes

```
$ curl -v http://127.0.0.1:1323/
```
sign in to get a token
```
$ curl \
  -X POST \
  http://localhost:1323/login \
  -H 'Content-Type: application/json' \
  -d '{"name":"Ethien","email":"ethien.salinas@gmail.com", "password":"qwerty"}'
```
### Protected routes
GET user info
```
$ curl -v http://localhost:1323/users/123 \
  -H "Authorization: Bearer JWT_HEADER.JWT_PAYLOAD.JWT_VERIFY SIGNATURE"
```