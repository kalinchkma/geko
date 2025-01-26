# Geko

Wrapper for creating golang server native application using popular library 

## Package used
1. [Gin](https://gin-gonic.com/docs/): For http request handling
2. [Gomail](https://pkg.go.dev/gopkg.in/gomail.v2): For mailer
3. [Gorm](https://gorm.io/docs/): For ORM


## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

## MakeFile

Run build make command with tests
```bash
make all
```

Build the application
```bash
make build
```

Run the application
```bash
make run-http-server
```

Run the database up migration
```bash
make run-migration-up
```

Run the database down migration
```bash
make run-migration-down
```

Create DB container
```bash
make docker-run
```

Shutdown DB Container
```bash
make docker-down
```

DB Integrations Test:
```bash
make itest
```

Live reload the application:
```bash
make watch
```

Run the test suite:
```bash
make test
```

Clean up binary from the last build:
```bash
make clean
```
