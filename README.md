# wolf_bro_example

## Usage
### Init the project
```
go mod tidy
```

### Build docker image
```
make docker
```

### Deploy locally
```
make up
```

### Undeploy
```
make down
```

## API endpoints
### Users
```

$curl -H "Authorization: token" -i http://127.0.0.1:8000/users
{"users":["User1","User2","User3"]}

$curl -H "Authorization: token" -i http://127.0.0.1:8000/users/1
{"users":"User1"}
```

### Banner
```
curl http://127.0.0.1:8000/banner/WOLF
 _       ______  __    ______
| |     / / __ \/ /   / ____/
| | /| / / / / / /   / /_
| |/ |/ / /_/ / /___/ __/
|__/|__/\____/_____/_/
```
