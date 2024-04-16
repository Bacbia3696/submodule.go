# Simplify Service Lifecycle Management in Go with Submodule

**Effortlessly manage the lifecycle of your services in Go with Submodule**, a lightweight and versatile library designed to streamline service management. Say goodbye to complex dependency handling, configuration management, and testing challenges.

## Clear Structure, Easy Management

* **Organize services with ease:** Wrap your service creation functions and chain them together for a clear and organized structure.
* **Simplify complex dependencies:** Manage and understand intricate dependencies between services more effectively.
* **Seamlessly integrate with your framework:** Bring Submodule into your existing frameworks and utilize its power wherever you have an async function.
* **Serverless:** Initialize what function needs, not what framework wants

## Streamlined Testing

* **Flexible testing environment:** Easily change dependencies for testing purposes, promoting testability and isolation.
* **Testable code chunks:** Organize your code into smaller, testable units, facilitating robust testing.
* **Controlled lifecycle management:** Implement unit tests and integration tests with ease by controlling the lifecycle of services.

## Lightweight and Simple

* **No unnecessary abstractions:** Built in Typescript with zero dependencies, Submodule embraces the fundamental concept of functions without complex magic.
* **Quick to understand and adopt:** Experience the simplicity and elegance of Submodule, even for developers new to the library.

Discover a painless way to manage the lifecycle of your services in Go with Submodule. Enhance your development workflow, improve code maintainability, and simplify testing processes.

## 💡 Usage

You can import `submodule` using:

```go
import (
    "github.com/submodule-org/submodule.go"
)
```

Then create a submodule like this:

```go
var ConfigMod = submodule.Create(func(ctx context.Context) (*Config, error) {
    return &Config{}, nil
})

var RedisMod = submodule.Derive(func(ctx context.Context, config *Config) (*Redis, error) {
    db := 0
    if config.RedisConfig.DB != nil {
        db = *config.RedisConfig.DB
    }
    return newRedis(config.RedisConfig.Address, config.RedisConfig.Password, db, config.RedisConfig.MaxIdle), nil
}, ConfigMod)

// execute the module
redis, err := RedisMod.Get(ctx)
```

## 📚 Documentation
see [godoc](https://pkg.go.dev/github.com/submodule-org/submodule.go)
more examples in [submodule_test.go](submodule_test.go)
