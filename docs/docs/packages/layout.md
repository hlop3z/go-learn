# New Projects Layout

Go emphasizes simplicity and modularity. With frameworks like **Gin**, **Echo**, or plain Go, Go projects often follow a structure inspired by **Clean Architecture** or **Domain-Driven Design**.

## Standard Go Project Layout

- [Golang Standards](https://github.com/golang-standards/project-layout/)

## Ready to `GO` Layout

- [Custom Layout](https://github.com/hlop3z/project-layout-go)

**Create Layout:**

```bash
mkdir -p build/{ci,package} internal/{app,pkg} web/{app,middleware,static,template} \
assets cmd config deployments docs examples pkg scripts tests tools
```

**Clone Layout:**

```bash
git clone https://github.com/hlop3z/project-layout-go.git && rm -rf project-layout-go/.git
```

## Example: Go Multiple Directories Layout (Clean Architecture)

```plaintext
{your_project}/
├── assets/                     # ~ Static assets such as images, fonts, etc.
├── build/                      # ~ Build-related files and configurations
│   ├── ci/                     # Continuous Integration configurations (e.g., GitHub Actions, CircleCI)
│   └── package/                # Packaging configurations (e.g., Docker, tarballs)
├── cmd/                        # ~ Entry points for commands (main package)
│   └── {your_app}/             # Application command for starting the app, handling migrations, etc.
├── config/                     # ~ Configuration files for different environments (e.g., dev, prod)
├── deployments/                # ~ Deployment-related configurations and scripts
├── docs/                       # ~ Documentation files for the project
├── examples/                   # ~ Example configurations or use cases for developers
├── internal/                   # ~ Internal application logic (not to be exposed publicly)
│   ├── app/                    # Internal application logic for handling requests, business rules
│   │   └── {your_app}/         # Specific logic for your application’s core functionality
│   └── pkg/                    # Internal reusable packages (e.g., shared utilities, helpers)
│       └── {your_private_lib}/ # Private library with internal functions and methods
├── pkg/                        # ~ Public packages that can be reused by other applications
│   └── {your_public_lib}/      # Reusable public library with functionality exposed for external use
├── scripts/                    # ~ Utility scripts for tasks like database migration, deployment, etc.
├── tests/                      # ~ Unit and integration tests for the application
├── tools/                      # ~ Developer tools and utilities (e.g., code generators, linters)
├── web/                        # ~ Web application-specific files (server-side)
│   ├── app/                    # Core web application logic (controllers, views)
│   ├── static/                 # Static files served by the web server (CSS, JS, images)
│   └── template/               # Template files for rendering views (HTML, JSX, etc.)
├── go.mod                      # ~ (Package/Module) => github.com/YOUR_USER_OR_ORG_NAME/YOUR_REPO_NAME
└── {etc ...}                   # All other files and directories
```

## Example: Go Single Directory Layout

```plaintext
module/
├── main.go             # Entry point of the application
├── types.go            # Defines data structures, constants, and interfaces
├── config.go           # Handles configuration loading and parsing
├── handlers.go         # HTTP handlers or main API endpoints
├── services.go         # Business logic and application services
├── repository.go       # Handles database or external data access
├── middleware.go       # Defines HTTP middleware (e.g., auth, logging)
├── routes.go           # Sets up and manages HTTP routes
├── utils.go            # Utility/helper functions
├── helpers.go          # Helper/utility functions
└── errors.go           # Custom error definitions and error handling utilities
```
