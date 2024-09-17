# Middleware Package

The `middleware` directory contains middleware functions that are used to handle various aspects of request processing in the application. Middleware functions are essential for tasks such as authentication, logging, error handling, and more. They act as a bridge between the incoming request and the final route handler, allowing you to execute code before the request reaches the route handler.

## Purpose

The primary purpose of the middleware functions in this directory is to enhance the functionality and security of the application by performing common tasks that need to be executed for multiple routes.

## Auth Middleware

One of the key middleware functions in this directory is the `AuthGuard` function, located in the `middleware/auth.go` file. This middleware is responsible for ensuring that only authenticated users can access certain routes.

### `AuthGuard` Middleware

The `AuthGuard` middleware checks if a user is authenticated before allowing them to proceed to the requested route. If the user is not authenticated, they are redirected to the login page.

#### Example Usage

Here is an example of how the `AuthGuard` middleware is used:

```go
package middleware

import (
    "github.com/labstack/echo/v5"
    "github.com/cmcd97/bytesize/apis"
)

// AuthGuard ensures that the user is authenticated before proceeding to the next handler.
func AuthGuard(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        record := c.Get(apis.ContextAuthRecordKey)

        if record == nil {
            return c.Redirect(302, "/auth/login")
        }

        return next(c)
    }
}

```
