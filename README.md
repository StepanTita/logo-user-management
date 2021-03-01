# logo-user-management

**REQUESTS:** <br>
Base path: https://logo-user-management.herokuapp.com

* GET /logo/users/<username>/auth <br>
  Request:

```json 
{
  "data": {
    "password": "test1234"
  } 
}
```

Response:

```json
{
  "data": {
    "username": "test123",
    "name": "test",
    "surname": "test",
    "email": "example@gmail.com"
  }
}
```
-------------------------------------------------------------------------------
* POST /logo/users <br>
  Request:

```json 
{
  "data": {
    "username": "test123",
    "name": "test",
    "surname": "test",
    "email": "example@gmail.com",
    "password": "test1234"
  } 
}
```

Response:

```json
{
  "data": {
    "username": "test123",
    "name": "test",
    "surname": "test",
    "email": "example@gmail.com"
  }
}
```
-------------------------------------------------------------------------------
* PATCH /logo/users/<username> <br>
  Request:

```json 
{
  "data": {
    "username": "test123",
    "name": "test",
    "surname": "test",
    "email": "example@gmail.com",
    "password": "test1234"
  } 
}
```

Response:

```json
{
  "data": {
    "username": "test123",
    "name": "test",
    "surname": "test",
    "email": "example@gmail.com"
  }
}
```

-------------------------------------------------------------------------------
* DELETE /logo/users/<username> <br>
  Request:

```json 
{
  "data": {
    "username": "test123",
    "name": "test",
    "surname": "test",
    "email": "example@gmail.com",
    "password": "test1234"
  } 
}
```

Response:

status 204

# logo-user-management
