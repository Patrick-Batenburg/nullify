# Nullify - Thou Shalt Guard Against the Void

[![Build Status](https://github.com/Patrick-Batenburg/null/actions/workflows/main.yml/badge.svg?branch=master)](https://github.com/Patrick-Batenburg/null/actions/workflows/main.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/Patrick-Batenburg/null)](https://goreportcard.com/report/github.com/Patrick-Batenburg/null) [![PkgGoDev](https://pkg.go.dev/badge/github.com/Patrick-Batenburg/null)](https://pkg.go.dev/github.com/Patrick-Batenburg/null)

## `null` package

The `null` package provides some helpful types for dealing with nullable SQL and JSON values. All types will marshal
to JSON null if invalid or if the SQL source data is null. The following types are supported:

| Type           | Description          | Notes                                                                                                                                                                                                                                                                         |
| -------------- | -------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `null.Bool`    | Nullable `bool`      |                                                                                                                                                                                                                                                                               |
| `null.Byte`    | Nullable `byte`      |                                                                                                                                                                                                                                                                               |
| `null.Bytes`   | Nullable `[]byte`    | `[]byte{}` and `[]byte(nil)` input will not produce invalid Bytes. This should be used for storing binary data (bytea in PSQL for example) in the database.                                                                                                                   |
| `null.Float32` | Nullable `float32`   |                                                                                                                                                                                                                                                                               |
| `null.Float64` | Nullable `float64`   |                                                                                                                                                                                                                                                                               |
| `null.Int`     | Nullable `int`       |                                                                                                                                                                                                                                                                               |
| `null.Int8`    | Nullable `int8`      |                                                                                                                                                                                                                                                                               |
| `null.Int16`   | Nullable `int16`     |                                                                                                                                                                                                                                                                               |
| `null.Int32`   | Nullable `int32`     |                                                                                                                                                                                                                                                                               |
| `null.Int64`   | Nullable `int64`     |                                                                                                                                                                                                                                                                               |
| `null.JSON`    | Nullable `[]byte`    | Will marshal to JSON null if invalid. `[]byte{}` and `[]byte(nil)` input will not produce an Invalid JSON. This should be used for storing raw JSON in the database. Also has `null.JSON.Marshal` and `null.JSON.Unmarshal` helpers to marshal and unmarshal foreign objects. |
| `null.String`  | Nullable `string`    |                                                                                                                                                                                                                                                                               |
| `null.Time`    | Nullable `time.Time` | Marshals to JSON null if the SQL source data is null.                                                                                                                                                                                                                         |
| `null.Uint`    | Nullable `uint`      |                                                                                                                                                                                                                                                                               |
| `null.Uint8`   | Nullable `uint8`     |                                                                                                                                                                                                                                                                               |
| `null.Uint16`  | Nullable `uint16`    |                                                                                                                                                                                                                                                                               |
| `null.Uint32`  | Nullable `uint32`    |                                                                                                                                                                                                                                                                               |
| `null.Uint64`  | Nullable `uint64`    |                                                                                                                                                                                                                                                                               |
| `null.UUID`    | Nullable `uuid.UUID` | Marshals to JSON null if the SQL source data is null. Uses `uuid.UUID`'s marshaler, unmarshaler, scanner and valuer from `github.com/google/uuid`.                                                                                                                            |

### Extending with complex types

It's possible to extend types with this package. These complex types embed `NullableImpl[T]`. They should override `sql.Scanner`, `driver.Valuer`, `encoding.TextMarshaler`, `encoding.TextUnmarshaler`, `json.Marshaler` and `json.Unmarshaler` interfaces, unless the implementation given by `NullableImpl[T]` suffice your usecase.

#### `sql.Scanner`

The `sql.Scanner` interface is used for scanning and converting SQL database values into Go types. Any struct that implements this interface can read values from a database query result and convert them into the appropriate Go type.

```go
Scan(src interface{}) error
```

`Scan` takes a value from the database (e.g., a row in a SQL result) and assigns it to the struct, converting the data as needed. The `src` argument is usually a `[]byte`, but it could be any type depending on the database driver.

#### `driver.Valuer`

The `driver.Valuer` interface is the counterpart to `sql.Scanner`, used for converting Go values into a format suitable for database storage. Any struct that implements this interface can be stored in a database.

```go
Value() (driver.Value, error)
```

`Value` returns a value that can be stored in a SQL database (e.g., `int`, `float64`, `string`, `[]byte`). The returned value must be one of the types that the database driver understands.

#### `encoding.TextMarshaler`

The `encoding.TextMarshaler` interface is used to convert Go types to a textual representation, often for use in encoding data as plain text (e.g., XML, JSON).

```go
MarshalText() (text []byte, err error)
```

`MarshalText` converts the struct into a slice of bytes representing its textual form (usually UTF-8 encoded text). This is useful for converting the struct into a format like XML or JSON, where the value is expected to be text.

#### `encoding.TextUnmarshaler`

The `encoding.TextUnmarshaler` interface is used to convert a textual representation back into a Go type. This is often used when parsing data from formats like XML or JSON into a Go struct.

```go
UnmarshalText(text []byte) error
```

`UnmarshalText` takes a slice of bytes (usually UTF-8 text) and parses it into the struct's fields. This allows a Go struct to be populated from a textual representation.

#### `json.Marshaler`

The `json.Marshaler` interface is used for converting Go types into their JSON representation. Implementing this interface allows a struct to define custom JSON encoding behavior.

```go
MarshalJSON() ([]byte, error)
```

`MarshalJSON` converts the struct into a slice of bytes containing its JSON representation. This is called when the struct is passed to `json.Marshal`.

#### `json.Unmarshaler`

The `json.Unmarshaler` interface is used for parsing JSON data and converting it into a Go type. Implementing this interface allows a struct to define custom JSON decoding behavior.

```go
UnmarshalJSON(data []byte) error
```

`UnmarshalJSON` takes a slice of bytes containing JSON data and parses it into the struct's fields. This is called when the struct is passed to `json.Unmarshal`.

---

# Installation

To install Nullify, use `go get`:

```shell
go get github.com/Patrick-Batenburg/nullify
```

This will then make the following packages available to you:

```shell
go get github.com/Patrick-Batenburg/nullify/null
```

Import the `nullify/null` package into your code:

```go
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/Patrick-Batenburg/nullify/null"
)

// Define a type that can have nullable fields, and
// will be used as our database result.
type User struct {
	ID                 null.UUID   `json:"id"`
	FirstName          null.String `json:"firstName"`
	MiddleName         null.String `json:"middleName"`
	LastName           null.String `json:"lastName"`
	Age                null.Int8   `json:"age"`
	Email              null.String `json:"email"`
	CreatedAt          null.Time   `json:"createdAt"`
	SomeOptionalID     null.UUID   `json:"someOptionalId"`
	SomeOptionalNumber null.Int64  `json:"someOptionalNumber"`
	SomeOptionalTime   null.Time   `json:"someOptionalTime"`
}
```

See the [example](./example/main.go) for more.

---

# Staying up to date

To update Nullify to the latest version, use `go get -u github.com/Patrick-Batenburg/nullify`.

---

# Supported go versions

We currently support the most recent major Go versions from 1.23 onward.

---

# Contributing

Please feel free to submit issues, fork the repository and send pull requests!

When submitting an issue, we ask that you please include a complete test function that demonstrates the issue.

---

# License

This project is licensed under the terms of the MIT license.
