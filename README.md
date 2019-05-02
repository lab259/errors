<p align="center"> <img src="https://storage.googleapis.com/gopherizeme.appspot.com/gophers/d9349f6fc662ca7d42020f9ec82710751caf5c7a.png" width="250"/></p>

# [WIP] errors [![CircleCI](https://circleci.com/gh/lab259/errors.svg?style=shield&circle-token=c4509080576acf88fa313e2bb5ccabb4787a44fb)](https://circleci.com/gh/lab259/errors)

**THIS DOCUMENT IS A WORK IN PROGRESS**

Errors have always been a problem in rthe development of our servers in Go. This package aims to address the main problems, centralizing a solution.

## Scenario

Let's start by creating our scenario: The user Astrobervaldo already signed in
system S. Astrobervaldo tries to create a new TODO item which fails because of an
invalid JSON was sent as input.

### API

In APIs, the probable endpoints would be implemented in a POST to `/v1/todo`. On
this endpoint, one of the first things that would happen is to parse the request
body in order to extract what Astrobervaldo sent. As we can image, it probably
would use a `json.Unmarshal` and it would return an `error`.

What should we say to the user? What structure will be used? Is it centralized?

### GraphQL

TODO

## Solution

The idea is to wrap errors with simple structures that will aggregate more and
more little pieces of information about the error.

Consider the following code:

```go
err := json.Unmarshal(data, &todoItemInput)
if err != nil {
    return errors.WrapHttp(err, 400) // BadRequest
}
```

The code above shows an error being wrapped as an HttpError. Upwards, it can
be checked and the status code, inside of it, can be applied to a possible
HTTP response structure.

To improve the usage, a `Wrap` function was created:

```go
err := json.Unmarshal(data, &todoItemInput)
if err != nil {
    return errors.Wrap(err, Http(400)) // BadRequest
}
```

This example has the same effect of the example shown previously. The difference
is that `Wrap` is a variadic function. Yes, it can receive multiple arguments
and make multiple wraps:

```go
err := json.Unmarshal(data, &todoItemInput)
if err != nil {
    return errors.Wrap(err, Http(400), Code("invalid-json")) // BadRequest
}
```

In this example, `err` was wrapped twice. With a `HttpError` and a
`ErrorWithCode`. In other words, the error returned is a `ErrorWithCode`
with a `HttpError` as `Reason`, this `HttpError` has the original err as `Reason`.

### How do the pieces of info get together?

For that, an interface `type ErrorResponse interface` was implemented.

An `ErrorResponse` aggregates all extra information about an error. The idea is
that an `ErrorResponse` will be serialized and sent whom is using the system.

### How aggregation occurs?

Some errors, those that need to add information to the `ErrorResponse`, implement
another interface `ErrorResponseAggregator`. This interface has a
`AppendData(ErrorResponse)` which get called for each `Reason` recursively:

```go
err := json.Unmarshal(data, &todoItemInput)
if err != nil {
    return errors.WrapHttp(err, 400) // BadRequest
}
```

Here, the returned error would have its `AppendData` method called, and only
this one because its `Reason` is not an `ErrorResponseAggregator` (it is the
original unmarshaling error).

```go
err := json.Unmarshal(data, &todoItemInput)
if err != nil {
    return errors.Wrap(err, Http(400), Code("invalid-json")) // BadRequest
}
```

In this case, a `ErrorWithCode` would have its `AppendData` called, a
`HttpError` would have its `AppendData` called afterwards, and then the original
the unmarshaling error would be reached and nothing more would happen.

### Error Types

#### ErrorWithMessage

This error adds the `Message` information to the error. The message should be humam readable.

**Wrap**: `Message(message string)`: It returns a default `ErrorWithMessage`
implementation which is also a `ErrorResponseAggregator`.

**Usage**:

```go
// using plain string
errors.Wrap(err, "failed on such task")

// using wrapper
errors.Wrap(err, errors.Message("failed on such task"))
```

#### HttpError

This error adds the `StatusCode` information to the error.

**Wrap**: `Http(statusCode int)`: It returns a default `HttpError` implementation
which is also a `ErrorResponseAggregator`.

**Usage**:

```go
// using plain int
errors.Wrap(err, http.StatusNotFound)

// using wrapper
errors.Wrap(err, errors.Http(http.StatusNotFound))
```

#### ModuleError

This error adds the `Module` information to the error. The module specifies what
area of the system the error is emerging from.

**Wrap**: `Module(module string)`: It returns a default `ModuleError` implementation
which is also a `ErrorResponseAggregator`.

**Usage**:

```go
// using predefined module
var (
    ModuleCatalog = errors.Module("catalog")
)

errors.Wrap(err, ModuleCatalog)
```

#### ErrorWithCode

This error adds the `Code` information to the error. The code refers to a (documented)
known error.

**Wrap**: `Code(code string)`: It returns a default `CodeError` implementation
which is also a `ErrorResponseAggregator`.

**Usage**:

```go
// using predefined code
var (
    CodeProductNotFound = errors.Code("product-not-found")
)

errors.Wrap(err, CodeProductNotFound)
```

#### ValidationError

A special case error for dealing with validation errors from the
[`go-playground/validator`](https://github.com/go-playground/validator) package.

**Wrap**: `Validation()`: It returns a default `ValidationError` implementation
which is also a `ErrorResponseAggregator`.

**Usage**:

```go
// using wrapper
errors.Wrap(err, errors.Validation())

// add a custom message
errros.Wrap(err, "custom message", errors.Validation())
```

### Combining Errors

In the following example, we show how combining is possible:

```go
// modules.go
var (
    ModuleService = errors.Module("services")
    // another modules...
)

// services/codes.go
var (
    UserNotFound = errors.Combine(errors.Code("users.notFound"), myapp.ModuleService)
)

// services/users/find.go
func Find(id string) (*User, error) {
    // something goes here...

    if err != nil {
        // instead of `errors.Wrap(err, errors.Code("users.notFound"), errors.Module("services"))`
        return nil, errors.Wrap(err, codes.UserNotFound)
    }

    // more...
}
```
