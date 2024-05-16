# Optional Values in Go

[![GitHub release (with filter)](https://img.shields.io/github/v/release/BooleanCat/option?sort=semver&logo=Go&color=%23007D9C)](https://github.com/BooleanCat/option/releases) [![Actions Status](https://github.com/BooleanCat/option/workflows/test/badge.svg)](https://github.com/BooleanCat/option/actions) [![Go Reference](https://pkg.go.dev/badge/github.com/BooleanCat/option.svg)](https://pkg.go.dev/github.com/BooleanCat/option) [![Go Report Card](https://goreportcard.com/badge/github.com/BooleanCat/option)](https://goreportcard.com/report/github.com/BooleanCat/option) [![codecov](https://codecov.io/gh/BooleanCat/option/branch/main/graph/badge.svg?token=N2E43RSR14)](https://codecov.io/gh/BooleanCat/option)

Support user-friendly optional types in Go.

```go
value = option.Some(4)
no_value = option.None[int]()
```

_[Read the docs.](https://pkg.go.dev/github.com/BooleanCat/option)_

## Usage

This package adds a single type, the `Option`. `Option`'s are instantiated as
one of two variants. `Some` denotes the presence of a value and `None` denotes
the absence.

Historically pointers have been used to denote optional values, this package
removes the risk of null pointer exceptions by leveraging generics to implement
type-safe optional values.

`Options` can be tested for the presence of a value:

```go
two = option.Some(2)
if two.IsSome() {
    ...
}
```

Values can be extracted along with a boolean test for their presence:

```go
two = option.Some(2)
if value, ok := two.Value(); ok {
    ...
}
```

Optionals that you're sure have a value can be "unwrapped":

```go
two := option.Some(2)
two.Unwrap()  // returns 2
```

Accessing a value on a `None` variant will cause a runtime panic.

```go
none := option.None[int]()
none.Unwrap()  // panics
```

## Ergonomics

Use of a package like this may be pervasive if you really commit to it. This
package was inspired by Rust's options implemenation. It might be worth
considering dropping the repetative `option.` preceding the variants. Since
importing names into the global namespace is to be avoided, the following
import pattern may work for you:

```go
import (
    "fmt"

    "github.com/BooleanCat/option"
)

var (
    Some = option.Some
    None = option.None
)

func main() {
    two := Some(2)
    fmt.Println(two)
}
```
