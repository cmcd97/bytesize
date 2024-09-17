# Authentication Package

The `auth` directory contains all code related to user authentication. This module is intentionally kept separate from the rest of the application logic, as it is generally the only section that can be accessed without prior authentication.

The directory includes the following files:

- **`login.go`**: Manages user authentication routes. It defines the `LoginFormValue` struct for capturing login form data and includes validation logic. The `RegisterLoginRoutes` function sets up the `/login` and `/logout` routes:

  - **`/login` GET**: Renders the login page.
  - **`/login` POST**: Processes login form submissions, validates credentials, and handles authentication. If authentication fails, it re-renders the login form with error messages.
  - **`/logout` POST**: Clears the authentication cookie and redirects the user to the login page.

  This file integrates with Echo for routing, Pocketbase for user management, and Templ for rendering components.

- **`login.templ`**: Contains the HTML/HTMX code for the login page.

- **`register.go`**: Manages user registration routes. It defines the `RegisterFormValue` struct for capturing registration form data and includes validation logic. The `RegisterRegisterRoutes` function sets up the `/register` route:

  - **`/register` GET**: Renders the registration page.
  - **`/register` POST**: Processes registration form submissions, validates the input, and creates a new user. If validation or registration fails, it re-renders the form with error messages.

  This file integrates with Echo for routing, Pocketbase for user management, and Templ for rendering components.

- **`register.templ`**: Contains the HTML/HTMX code for the registration page.
