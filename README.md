# gocheck
Checkout Team Go Code Convention.

![code-structure-explain](README_asset/gocheck-goout.drawio.png)

# Code structure

```
├── internal/
│   ├── app/            Package app contains application starter.
│   ├── config/         Package config provides functionality for loading and accessing application configurations.
│   ├── dto/            Package dto provides data transfer objects (DTOs) for transferring data between different layers of the application.
│   ├── entity/         Package entity provides domain entities representing the core business objects.
│   │   ├── table/      Package table provides table schema.
│   ├── extapi/         Package extapi provides implementations for external APIs.
│   ├── repo/           Package repo provides repository implementations for data storage.
│   │   ├── cache/      Package cache provides cache storage implementations.
│   │   ├── db/         Package db provides database connection and management functionalities.
│   ├── repomw/         Package repomw provides repo middleware.
│   ├── transport/      Package transport provides the implementation for handling input.
│   │   ├── grpc/       Package grpc provides gRPC server implementation.
│   │   ├── http/       Package http provides HTTP server implementation.
│   ├── usecase/        Package usecase provides the business logic implementation.
│   ├── usecasemw/      Package usecasemw provides usecase middleware.
├── pkg/                Package pkg provides package that can be imported by other services.
├── go.mod
├── go.sum
├── main.go
├── README.md
```

# Feature

- [x] simulate real case scenario.
- [x] integration test.
- [ ] unit test, (easy to implement because seperated by interface between layer but not implement yet).
- [x] integrated trace id front to back..
- [x] grpc server.
    - [x] graceful shutdown.
- [x] http server.
    - [x] graceful shutdown.
- [x] centralize protobuf.
- [x] code documentation with godoc.
- [x] gorm connection pooling.
- [x] validate request with golang validator.
- [x] usecase middleware.
    - [x] usecase logger middleware.
- [x] repo middleware.
    - [x] repo logger middleware.
- [x] repo cache layer.
