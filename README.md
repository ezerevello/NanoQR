# NanoQR

[![Go Version](https://img.shields.io/github/go-mod/go-version/ezerevello/NanoQR?style=flat-square&logo=go)](https://go.dev/) [![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg?style=flat-square)](https://opensource.org/licenses/MIT) [![Docker Ready](https://img.shields.io/badge/Docker-Ready-2496ED?style=flat-square&logo=docker&logoColor=white)](https://www.docker.com/)

**NanoQR** is a high-performance, **headless** QR code generator service built with Go.

I created it because I thought the idea was quite fun and an interesting project: an (almost) over engineered headless api qr generator and extensible software for multiple handlers.

Designed to be lightweight and easy to self-host, NanoQR provides a clean API to generate QR codes as Base64-encoded PNGs (or any other encoding you want).


-----

## 🚀 Key Features

  * **⚡ High Performance:** Built 100% in Go for minimal latency.
  * **🐳 Docker Ready:** Multi-stage Dockerfile optimized for Alpine Linux.
  * **🛠 Modular Design:** Decoupled architecture using interfaces, making it easy to swap handlers for when there are more (CLI, gRPC, etc.).
  * **📦 Zero Bloat:** Headless by design—perfect for microservices.

-----

## 📂 Project Structure

```text
.
├── cmd/
│   └── api/
│       └── main.go              # Application entry point
├── internal/
│   ├── handlers/
│   │   └── http_handler.go      # REST API transport
│   ├── model/
│   │   └── model.go             # Data transfer objects (JSON)
│   └── service/                 # Core business logic
│       ├── qr_service.go        # Domain interfaces
│       └── default_qr_service.go # Core QR generation implementation
├── Dockerfile                   # Minimal production image
├── go.mod
├── go.sum
└── README.md
```

-----

## 🛠 Quick Start

### Run with go
To try it out quickly, you can run the binary directly with go in your terminal
```bash
go run cmd/api/main.go
```
Or

### Run with Docker
```bash
docker build -t nanoqr .
docker run -p 8080:8080 nanoqr
```

The QR generator service will be listening at: `http://localhost:8080/api/qr`

➡️ For a quick test in terminal, you can run (linux):
```bash
curl -X POST http://localhost:8080/api/qr \
-H "Content-Type: application/json" \
-d '{"input": "https://google.com", "size": 256, "recoverLevel": "medium"}'\
```

➡️ For windows (cmd):
```cmd
curl -X POST http://localhost:8080/api/qr -H "Content-Type: application/json" -d "{\"input\": \"https://google.com\", \"size\": 256, \"recoverLevel\": \"medium\"}" 
```

-----

## 📑 API Documentation

### Generate QR Code

**Endpoint:** `POST /api/qr`

#### Request Body

| Parameter | Type | Required | Default | Description |
| :--- | :--- | :--- | :--- | :--- |
| `input` | `string` | **Yes** | - | URL or anything to encode. |
| `size` | `int` | No | `256` | Image size (41px to 2048px). |
| `recoverLevel` | `string` | No | `medium` | Error correction: `low`, `medium`, `high`, `highest`. |

#### Example Request

```json
{
  "input": "https://github.com/",
  "size": 512,
  "recoverLevel": "high"
}
```

#### Example Response

```json
{
  "info": {
    "input": "https://github.com/",
    "size": 512,
    "recoverLevel": "high"
  },
  "qr": "iVBORw0KGgoAAAANSUhEUgAA...",
  "status": "success"
}
```

----

### CORS Middleware (HTTP)

NanoQR includes a CORS middleware to allow secure cross-origin requests from your web applications.

#### How it works

First, the `Origin` header of incoming requests (outside of the middleware) is checked:
```go
var allowedOrigins = []string {
	"https://nanoqr-web.vercel.app",
	"http://localhost:5173",
	"https://another-web.com",
}

func isOriginAllowed(origin string) bool {

	for _, o := range allowedOrigins {
		if o == origin {
			return true
		}
	}
	return false

}
```
_The URLs are examples; simply replace them._

- If the origin is allowed, it sets the headers.
- For preflight (OPTIONS) requests, it responds with status 204 and the appropriate headers, without invoking the main handler.
- For POST requests, it allows the request to proceed to the QR handler as usual.

#### Func:
```go
func CORSMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Add CORS headers
		origin := r.Header.Get("Origin")
		if isOriginAllowed(origin) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
			w.Header().Set("Content-Type", "application/json")
		}

		// Preflight
        if r.Method == http.MethodOptions {
            w.WriteHeader(http.StatusNoContent)
            return
        }

		// Call next handler (QRhandler)
		next.ServeHTTP(w, r)
	})

}
```

-----

## 🔧 Extending NanoQR

NanoQR uses a `QRService` interface, allowing you to build your own handlers (CLI, Lambda, gRPC) effortlessly.

```go
type QRService interface {
    Generate(input string, size int, recoverLevel string) ([]byte, error)
}
```

### Creating a Custom CLI Handler

You can inject the service into your functions to maintain testability and flexibility:

```go
func CLIHandler(service QRService, input string) {
    qr, err := service.Generate(input, 256, "medium")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(base64.StdEncoding.EncodeToString(qr))
}

// Usage
qrService := &DefaultQRService{}
CLIHandler(qrService, "https://google.com")
```

-----

## 🐱 Contributing

Contributions are what make the open-source community such an amazing place to learn, inspire, and create. Any contributions you make are greatly appreciated.


-----

## ⚖️ License

Distributed under the **MIT License**.
