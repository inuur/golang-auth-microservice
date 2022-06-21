# Authentication microservice on golang
## Endpoints

The service supports the following endpoints:

`POST /users/` - create new user

*Request Body example*:

```json
{
  "username":"user",
  "password":"password",
  "email":"email@mail.ru"
}
```
*Response Body example*:
```json
{
    "id": "62b1e8157fd90928aceb9f92",
    "username": "user",
    "email":"email@mail.ru"
}
```
---

`GET /users/:id` - get user by `id`

*Response Body example*:

```json
{
    "id": "62b1e8157fd90928aceb9f92",
    "username": "user",
    "email":"email@mail.ru"
}
```

---

`POST /token/:id` - generates access and refresh tokens for the user with `id` ID

*Response Body example*:

```json
{
    "access_token": "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOiI2MmIxZTgxNTdmZDkwOTI4YWNlYjlmOTIiLCJVc2VybmFtZSI6ImludXVya2EiLCJleHAiOjE2NTU4Mjg3MzIsImp0aSI6ImNhODk2NDNlLTMxZWItNDc4OC1hNjdjLWIwNjhkNzlhNWE5ZSJ9.WEaOMMZmKaVFe8cNcQhR3K-yxiQTO0SGCTxzgZCywfrvVNilepn929DyjJphqenAUKjHC8eSI-UjX3tfQPI06A",
    "refresh_token": "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOiIiLCJVc2VybmFtZSI6IiIsImV4cCI6MTY1NTgzMDcxMn0.DRaDz6iq3vxM9z7yNcyygIdw3oDHGk0_iOMSnXzbVSWJAix5jKOMQvLDRlUEJzmsa-KnhhjhH0fn7vAmjJfeRQ"
}
```
---

`POST /token/refresh` - refresh access and refresh token

*Request Body example*:

```json
{
    "access_token": "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOiI2MmIxZTgxNTdmZDkwOTI4YWNlYjlmOTIiLCJVc2VybmFtZSI6ImludXVya2EiLCJleHAiOjE2NTU4Mjg3MzIsImp0aSI6ImNhODk2NDNlLTMxZWItNDc4OC1hNjdjLWIwNjhkNzlhNWE5ZSJ9.WEaOMMZmKaVFe8cNcQhR3K-yxiQTO0SGCTxzgZCywfrvVNilepn929DyjJphqenAUKjHC8eSI-UjX3tfQPI06A",
    "refresh_token": "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOiIiLCJVc2VybmFtZSI6IiIsImV4cCI6MTY1NTgzMDcxMn0.DRaDz6iq3vxM9z7yNcyygIdw3oDHGk0_iOMSnXzbVSWJAix5jKOMQvLDRlUEJzmsa-KnhhjhH0fn7vAmjJfeRQ"
}
```

*Response Body example*:

```json
{
    "access_token": "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOiI2MmIxZTgxNTdmZDkwOTI4YWNlYjlmOTIiLCJVc2VybmFtZSI6ImludXVya2EiLCJleHAiOjE2NTU4Mjg3ODAsImp0aSI6ImYxOTEwZDZjLTBmZjgtNDRmNC1hZDAyLWUxNWY1NjE3MGUxOCJ9.sNevYNr5ptncdtRaOoyS1cC66EspYgEXSuO6tsZd7gWT8ZNIv704LaVZVkUGd-T3HBRibZ910duQz4irjcdbXA",
    "refresh_token": "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOiIiLCJVc2VybmFtZSI6IiIsImV4cCI6MTY1NTgzMDc2MH0.EuU0wtbW_xYU0WRJ5TXMe-9zPcTaGBaK-HNo7RGrTB4e0F9DWsCJel09xx5ehU4bCQyAWVbKGT0YIoBQUI2HDw"
}
```
---

## Configuration

```golang
type Config struct {
	IsDebug bool `env:"IS_DEBUG" env-default:"true"`
	Listen  struct {
		BindIP string `env:"BIND_IP" env-default:"0.0.0.0"`
		Port   string `env:"PORT" env-default:"8080"`
	}
	MongoDB struct {
		Username string `env:"MONGO_USERNAME"`
		Password string `env:"MONGO_PASSWORD"`
		Host     string `env:"MONGO_HOST" env-required:"true"`
		Port     string `env:"MONGO_PORT" env-required:"true"`
		Database string `env:"MONGO_DATABASE" env-required:"true"`
	}
	JWT struct {
		SecretKey           string `env:"JWT_SECRET"`
		AccessTokenExpTime  int    `env:"ACCESS_TOKEN_EXP_MINUTES" env-default:"15"`
		RefreshTokenExpTime int    `env:"REFRESH_TOKEN_EXP_HOURS" env-default:"48"`
	}
	Bcrypt struct {
		Cost string `env:"BCRYPT_COST" env-default:"10"`
	}
}
```

