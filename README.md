
# GoSS

Performance thought session store, stand-alone server built using Go and Redis.
 


| Stage             | Status                                                                |
| ----------------- | ------------------------------------------------------------------ |
| Development | ![#00b48a](https://via.placeholder.com/10/00b48a?text=+)  |
| Staging | ![#00b48a](https://via.placeholder.com/10/00b48a?text=+)    |
| Production | ![#cc1d01](https://via.placeholder.com/10/cc1d01?text=+)  |


## Run (Docker)

Clone the project

```bash
  git clone https://github.com/himanshu-sah/goss.git
```

Go to the project directory

```bash
  cd goss
```

Run

```bash
  docker-compose up
```


  
## API Reference

#### Create Session

```http
  POST /create
```

| Body | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `session` | `json` | **Required** |
| `ttl`     | `int`  | In seconds(Optional) |

#### Get Session

```http
  GET /get
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `sessionId`      | `string` | **Required**. Session ID received on creation of session. |

#### Delete Session

```http
  DELETE /delete
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `sessionId`      | `string` | **Required**. Session ID received on creation of session. |

#### Truncate Store (Delete All Sessions)

```http
  DELETE /truncate
```