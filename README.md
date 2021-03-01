# logo-user-management

**REQUESTS:** <br>
Base path: https://logo-user-management.herokuapp.com

* GET /logo/users/\<username>/auth <br>
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
    "email": "example@gmail.com",
    "image_url": "some url"
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
    "password": "test1234",
    "image_url": "some url"
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
    "email": "example@gmail.com",
    "image_url": "some url"
  }
}
```
-------------------------------------------------------------------------------
*Every field is optional here:
* PATCH /logo/users/\<username> <br>
  Request:

```json 
{
  "data": {
    "username": "test123",
    "name": "test",
    "surname": "test",
    "email": "example@gmail.com",
    "password": "test1234",
    "image_url": "some url"
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
    "email": "example@gmail.com",
    "image_url": "some url"
  }
}
```

-------------------------------------------------------------------------------
* DELETE /logo/users/\<username> <br>
  Request:

```json 
{
  "data": {
    "password": "test1234"
  } 
}
```

Response:

status 204

# logo-user-management
