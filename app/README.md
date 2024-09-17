# App Package

This directory contains most of the application logic and is organized into four main subdirectories:

- **`components`**: Standalone Templ components that can be reused across multiple places, such as custom buttons, tables, and navbars.
- **`css`**: The output directory where Tailwind writes its custom CSS classes.
- **`handlers`**: Custom HTTP handlers.
- **`views`**: Similar to the `components` directory, but contains Templ components that represent entire pages. These pages are typically used in a single location within the app and often utilize components from the `components` directory.
