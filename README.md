# SETUP APP

go get -u github.com/joho/godotenv
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
go get -u github.com/dgrijalva/jwt-go

# RUN APP

go run main.go


# DDD PATTERN

myapp/
├── cmd/
│   └── main.go
├── config/
│   └── config.go
├── domain/
│   └── user/
│       ├── user.go
│       ├── repository.go
│       └── service.go
├── infrastructure/
│   ├── db/
│   │   └── db.go
│   └── router/
│       └── router.go
├── interfaces/
│   └── api/
│       └── user_handler.go
├── repository/
│   └── user/
│       └── user_repo.go
└── usecase/
    └── user/
        └── user_usecase.go
