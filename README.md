This is repository helps us to understand how to use Wire (Dependency Injection) inspired from this [article](https://hellokoding.com/crud-restful-apis-with-go-modules-wire-gin-gorm-and-mysql/)

## Helpful resources

1. [Wire Go DI](https://github.com/google/wire/blob/master/_tutorial/README.md)

A very common issue, where we forget about adding inject command [here](https://github.com/google/wire/issues/178)

## Commands

```go
go run main.go
```

Above command will not work with since wire.go and wire_gen.go of same package will have same function names. To avoid duplicate function name error we will use

```go
// +build
```

This informs Go compiler to ignore specific file package thus avoiding duplicate error

## Use following command to run project

For Development
```go
go run .
```

For Production
```go
go build .
```