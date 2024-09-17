# Lib Package

The `lib` directory contains reusable code essential for the application.

- **`auth.go`**: Manages user authentication, including:

  - `Login`: Validates user credentials.
  - `Register`: Ensures unique usernames and matching passwords before creating a new user.
  - `setAuthToken`: Generates and sets an authentication token as an HTTP cookie.

- **`base.templ`**: Defines the base HTML layout, including:

  - Essential metadata, stylesheets, and scripts.
  - Uses Tailwind CSS, DaisyUI, and HTMX.
  - `{ children... }` placeholder for dynamic content.

- **`htmx.go`**: Provides utilities for handling HTMX requests, including:

  - Checking if a request is an HTMX request.
  - Rendering templates based on request type.
  - Handling HTMX-specific redirects.

- **`render.go`**: Renders Templ components in an Echo context, including:
  - Setting the HTTP status code.
  - Rendering the Templ component to the response writer.
  - Returning a 500 Internal Server Error if rendering fails.
