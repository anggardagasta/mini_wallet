#### Register
###### URL [POST]
```go
  http://localhost:8082/v1/users/register
```
###### Body
```go
{
	"username": "anggarda",
	"email": "anggarda@gmail.com",
	"password": "123"
}
```
#### Auth
###### URL [POST]
```go
  http://localhost:808/v1/users/auth
```
###### Body
```go
{
	"username": "anggarda",
	"password": "123"
}

OR

{
	"email": "anggarda@gmail.com",
	"password": "123"
}
```