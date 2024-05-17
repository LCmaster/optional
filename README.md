# Optional
An optional type implementation for Go

```go
opt := optional.From(42)
empty := optional.Empty[int]()
```

## Package

```go
import (
    "github.com/LCmaster/optional"
)
```
## Usage

This package adds a single `Optional` type representing a value or the lack of it.
It removes the risk of null pointer exceptions by leveraging generics to implement type-safe optional values.

```go
    opt := optional.From(42)
	if opt.IsPresent() {
        ...
    }
```

Extracting the value conforms to the comma ok return type:

```go
    opt := optional.From(42)
    if value, ok := opt.Get(); ok {
        ...
    }
```

A default value can be provided and will be returned in case of an empty variable:

```go
    opt1 := optional.From(42)
    value1 := opt1.OrElse(1337) // will return 42
    opt2 := optional.Empty[int]()
    value2 := opt2.OrElse(1337) // Will return 1337
```

A function can be passed as provider of a default value as well:

```go
    opt1 := optional.From(42)
    value1 := opt1.OrElseGet(func() int { return 1337}) // will return 42
    opt2 := optional.Empty[int]()
    value2 := opt2.OrElse(func() int { return 1337}) // Will return 1337
```

And finally a consumer function can be passed as argument, and will be invoked if the underlying value is present:

```go
    opt := optional.From(42)
    opt.IfPresent(func(answer int) { 
        fmt.Printf("The Answer to the Ultimate Question of Life, the Universe, and Everything is %d\n", answer)
    })
```

## License
Provided under the MIT license.