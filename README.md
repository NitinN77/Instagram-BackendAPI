# Instagram-BackendAPI

A Golang REST API to handle users and posts for a simple instagram backend. Uses MongoDB as the database. Tested using golang-testing and Postman.

## External Dependencies

- mongo-driver 
- godotenv (to handle secrets using environment variables, not required)

## Directory Structure

```
├── go.mod
├── go.sum
├── helper
│   └── connect.go
├── main.go
├── main_test.go
└── models
    └── models.go
```
- ``` go.mod, go.sum ``` Manages dependencies
- ``` connect.go ``` Establishes a connection to the MongoDB clusters
- ``` models.go ``` Defines the model objects
- ``` main.go ``` Primary file to handle http requests 
- ``` main_test.go ``` Unit testing script 

## Endpoints

#### ``` /api/posts ``` 
creates a post using data from the POST request's body

#### ``` /api/posts/<id> ``` 
fetches post details for the given id

#### ``` /api/posts/users/<id>?limit={}&lastid={} ``` 
fetches posts of the user with given id within a limit and with ids greater than lastid (if lastid is not null)

#### ``` /api/users ```
creates a user and encrypts the password using ciphers before storing

#### ``` /api/users/<id> ``` 
fetches user details of given id

## Pagination

The API employs an id-based pagination for the  ``` /api/posts/users/<id> ``` endpoint. 
On the initial call, the limit parameter must be passed along with the request. The API returns the requsted number of posts and the id of the last post. Subsequent calls 
from the front-end must also contain a ``` lastid ``` parameter with the id that was returned on the previous call. This ensures that the second call returns posts starting 
with ids greater than the ``` lastid ``` , and that the API response doesn't explode in size.

![](https://i.imgur.com/DOhl8Pc.png)

## Testing

![](https://i.imgur.com/vv4nVgj.png)

## Usage

1. clone the repo ``` git clone https://github.com/VoidlessVoid7/Instagram-BackendAPI . ```
2. ``` go build main.go ```
3. ``` go run ./ ```

## Screenshots

- Creating a post

![](https://i.imgur.com/3X7lqst.png)

- Fetching a post

![](https://i.imgur.com/n6H3WjF.png)

- Create a user

![](https://i.imgur.com/f1EPfqO.png)

- Fetch a user

![](https://i.imgur.com/GlOGzhu.png)

- Fetch posts created by a user

![](https://i.imgur.com/SKz48ya.png)
