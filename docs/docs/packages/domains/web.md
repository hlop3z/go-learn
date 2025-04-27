# Web Frameworks

If you’re looking for the **most lightweight** tool to create APIs in Go, without adding too much overhead, **Gin** and **Echo** are the top contenders. However, if you're looking for **something ultra-lightweight**, here's a comparison along with some recommendation:

---

## Summary of Options

| Framework    | Lightweight? | Features                                    | Use Case                  | Pros                     | Cons                                        |
| ------------ | ------------ | ------------------------------------------- | ------------------------- | ------------------------ | ------------------------------------------- |
| **Gin**      | Yes          | Routing, middleware, JSON                   | RESTful APIs              | Very fast, minimalistic  | Somewhat opinionated                        |
| **Echo**     | Yes          | Routing, middleware, WebSockets, validation | RESTful, web apps         | Fast, easy-to-use        | More features than Gin                      |
| **Chi**      | Yes          | Routing, middleware                         | Microservices, small apps | Minimalistic, extensible | Needs more manual work for features         |
| **net/http** | Yes          | Custom routing, manual setup                | Custom APIs               | No external dependencies | Requires manual handling of routing, errors |

---

## **Recommendation:**

If you want the **most lightweight solution**:

1. **`net/http`** if you want zero dependencies and control over everything.
2. **`Chi`** if you want **minimal routing** and to add features only as needed (best for microservices).
3. **`Gin`** if you want **performance** with an out-of-the-box feature set (great balance).

---

## 1. **net/http** (Pure Go)

- **Purpose**: Standard library package for HTTP servers.
- **Features**:
  - **No extra dependencies**.
  - Everything done by you: routing, request handling, middleware.

```go
import (
 "net/http"
)

func main() {
 http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Hello, World!"))
 })

 http.ListenAndServe(":8080", nil)
}
```

**Pros:**

- **No dependencies**.
- Ultra-lightweight and native Go.

**Cons:**

- **Manually handle routing, error handling**, etc.
- **Not as feature-rich** as frameworks like Gin or Echo.

---

## 2. **Gin**

- **Purpose**: Fast, small, and high-performance web framework for Go.
- **Features**:

  - Routing, middleware, JSON handling.
  - Built-in features for web APIs (JSON validation, rendering, etc.).
  - Very fast (close to `net/http` speed).
  - Easy-to-use, good documentation.

- **Use Case**: Perfect for building RESTful APIs, web services with minimal overhead.

```go
import (
 "github.com/gin-gonic/gin"
)

func main() {
 r := gin.Default()

 r.GET("/hello", func(c *gin.Context) {
  c.JSON(200, gin.H{
   "message": "Hello, World!",
  })
 })

 r.Run(":8080")
}
```

**Pros:**

- **Fast** and **simple**.
- Good balance between features and performance.

**Cons:**

- **Somewhat opinionated** (but still lightweight).

---

## 3. **Echo**

- **Purpose**: Fast and minimalist web framework for Go.
- **Features**:
  - Router, middleware, request validation, and rendering.
  - Also fast but with a little more structure than Gin.
  - Supports HTTP/2, WebSockets, and more advanced features.

```go
import (
 "github.com/labstack/echo/v4"
)

func main() {
 e := echo.New()

 e.GET("/hello", func(c echo.Context) error {
  return c.JSON(200, map[string]string{
   "message": "Hello, World!",
  })
 })

 e.Logger.Fatal(e.Start(":8080"))
}
```

**Pros:**

- **Fast**, easy routing.
- **Good docs** and slightly more structured than Gin.

**Cons:**

- **A bit more features** than Gin, though still very lightweight.

---

## 4. **Chi**

- **Purpose**: Ultra-lightweight router for Go.
- **Features**:
  - Focuses on **routing** and is **incredibly minimal**.
  - Small memory footprint, works well for microservices.
  - **Extensible** with middleware (but you get to decide what you want).

```go
import (
 "github.com/go-chi/chi"
 "net/http"
)

func main() {
 r := chi.NewRouter()

 r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Hello, World!"))
 })

 http.ListenAndServe(":8080", r)
}
```

**Pros:**

- **Minimal** — just routing, with extensible middleware.
- Great for very small services or microservices.
- Supports **middleware** and **context**.

**Cons:**

- **Minimal** by design, so you might need to implement some features manually.
