# Notes API

API to save user notes, implementing JWT for auth and MongoDB as database.

## Run Locally

1. Install both:

   - [Go 1.15+](https://golang.org/dl/)
   - [MongoDB](https://www.mongodb.com/try/download/community)

   You will need to have MongoDB running on port 27017.

2. Clone the project:

   ```bash
   git clone https://github.com/AloisCRR/jwt-api-notes.git
   ```

3. Go to the project directory:

   ```bash
   cd jwt-api-notes
   ```

4. Install dependencies:

   ```bash
   go mod tidy
   ```

5. Start the server:

   ```bash
   go run main.go
   ```

   REST API will run in [http://localhost:8080](http://localhost:8080).

## API Reference

### Note

<p align="center">
  <img src="https://i.imgur.com/FXEH7HW.png" alt="Note entity" width="200">
</p>

#### Signup or register

```http
POST /signup
```

| Body       | Type     | Description                      |
| :--------- | :------- | :------------------------------- |
| `email`    | `string` | **Required**. User email address |
| `password` | `string` | **Required**. Account password   |

| Response  | Type     | Description                           |
| :-------- | :------- | :------------------------------------ |
| `message` | `string` | API message                           |
| `status`  | `number` | HTTP status code                      |
| `token`   | `string` | Auth token to access protected routes |

#### Sign in or login

```http
POST /login
```

| Body       | Type     | Description                      |
| :--------- | :------- | :------------------------------- |
| `email`    | `string` | **Required**. User email address |
| `password` | `string` | **Required**. Account password   |

| Response  | Type           | Description                           |
| :-------- | :------------- | :------------------------------------ |
| `message` | `string`       | API message                           |
| `status`  | `number`       | HTTP status code                      |
| `token`   | `Bearer token` | Auth token to access protected routes |

#### Note creation

```http
POST /notes
```

| Headers          | Type           | Description                                   |
| :--------------- | :------------- | :-------------------------------------------- |
| `Authentication` | `Bearer token` | **Required**. Jwt given on sign in or sign up |

| Body      | Type     | Description              |
| :-------- | :------- | :----------------------- |
| `title`   | `string` | **Required**. Note title |
| `content` | `string` | Note content             |

| Response  | Type     | Description      |
| :-------- | :------- | :--------------- |
| `message` | `string` | API message      |
| `status`  | `number` | HTTP status code |

#### Get all notes

```http
GET /notes
```

| Headers          | Type           | Description                                   |
| :--------------- | :------------- | :-------------------------------------------- |
| `Authentication` | `Bearer token` | **Required**. Jwt given on sign in or sign up |

| Response  | Type     | Description      |
| :-------- | :------- | :--------------- |
| `data`    | `Note[]` | API message      |
| `message` | `string` | API message      |
| `status`  | `number` | HTTP status code |

#### Note update

```http
PUT /notes/${id}
```

| Parameter | Type     | Description           |
| :-------- | :------- | :-------------------- |
| `id`      | `string` | **Required**. Note ID |

| Headers          | Type           | Description                                   |
| :--------------- | :------------- | :-------------------------------------------- |
| `Authentication` | `Bearer token` | **Required**. Jwt given on sign in or sign up |

| Response  | Type     | Description      |
| :-------- | :------- | :--------------- |
| `message` | `string` | API message      |
| `status`  | `number` | HTTP status code |

#### Get single note

```http
GET /notes/${id}
```

| Parameter | Type     | Description           |
| :-------- | :------- | :-------------------- |
| `id`      | `string` | **Required**. Note ID |

| Headers          | Type           | Description                                   |
| :--------------- | :------------- | :-------------------------------------------- |
| `Authentication` | `Bearer token` | **Required**. Jwt given on sign in or sign up |

| Response  | Type     | Description      |
| :-------- | :------- | :--------------- |
| `data`    | `Note`   | API message      |
| `message` | `string` | API message      |
| `status`  | `number` | HTTP status code |

#### Delete note

```http
DELETE /notes/${id}
```

| Parameter | Type     | Description           |
| :-------- | :------- | :-------------------- |
| `id`      | `string` | **Required**. Note ID |

| Headers          | Type           | Description                                   |
| :--------------- | :------------- | :-------------------------------------------- |
| `Authentication` | `Bearer token` | **Required**. Jwt given on sign in or sign up |

| Response  | Type     | Description      |
| :-------- | :------- | :--------------- |
| `message` | `string` | API message      |
| `status`  | `number` | HTTP status code |

## Screenshots

No token provided

![No token](https://i.imgur.com/1nMJK8A.png)

Basic input validation

![Input validation](https://i.imgur.com/gUgNIcz.png)

Sign up or register

![Register](https://i.imgur.com/8i0TDZy.png)

Sign in or login

![Login](https://i.imgur.com/R7o7IYv.png)

Note creation

![Login](https://i.imgur.com/f1XPmxX.png)

All notes

![Login](https://i.imgur.com/HDfhlFv.png)

Note update

![Login](https://i.imgur.com/lC6vbTH.png)

Get single note

![Login](https://i.imgur.com/Zh1CGxP.png)

Delete note

![Login](https://i.imgur.com/8LRoMbP.png)

## Tech Stack

| Name                                                    | Description                                        |
| ------------------------------------------------------- | -------------------------------------------------- |
| [MongoDB](https://github.com/mongodb/mongo-go-driver)   | Database                                           |
| [Gin](https://github.com/gin-gonic/gin)                 | HTTP Server                                        |
| [jwt-go](https://github.com/dgrijalva/jwt-go)           | Library to generate, parse, validate and more JWTs |
| [validator](https://github.com/go-playground/validator) | Input validation                                   |
| [Bcrypt](https://golang.org/pkg/crypto/)                | Algorithm used to hash passwords.                  |

## Roadmap

- [x] App functionality
- [ ] Testing
- [ ] Hosting, domain, etc.
- [ ] CI/CD
