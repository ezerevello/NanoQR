# NanoQR

**NanoQR** is a high-performance, **headless** QR code generator service built with Go.

I created it because I thought the idea was quiet fun and an interesting project: an (almost) over engineered headless api qr generator and extensible software for multiple handlers.

Designed to be lightweight and easy to self-host, NanoQR provides a clean API to generate QR codes as Base64-encoded PNGs (or any other encoding you want).


-----

## 🚀 Key Features

  * **⚡ High Performance:** Built 100% in Go for minimal latency.
  * **🐳 Docker Ready:** Multi-stage Dockerfile optimized for Alpine Linux.
  * **🛠 Modular Design:** Decoupled architecture using interfaces, making it easy to swap handlers (CLI, gRPC, etc.).
  * **📦 Zero Bloat:** Headless by design—perfect for microservices.

-----

## 📂 Project Structure

```text
.
├── cmd/
│   └── api/             # Application entry point
├── internal/
│   ├── model/           # Data transfer objects (JSON)
│   └── service/         # Core business logic
│       ├── qr_service.go      # Domain interfaces
│       ├── http_handler.go    # REST API transport
│       └── ...                # Implementations
└── Dockerfile           # Minimal production image
```

-----

## 🛠 Quick Start

### Run with go
To try it out quickly, you can run the binary directly with go in your terminal
```bash
go run cmd/api/main.go
```
or

### Run with Docker
```bash
docker build -t nanoqr .
docker run -p 8080:8080 nanoqr
```

The QR generator service will be listening at: `http://localhost:8080/api/qr`

-----

## 📑 API Documentation

### Generate QR Code

**Endpoint:** `POST /api/qr`

#### Request Body

| Parameter | Type | Required | Default | Description |
| :--- | :--- | :--- | :--- | :--- |
| `input` | `string` | **Yes** | - | Text or URL to encode. |
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
  "status": "success",
  "info": {
    "input": "https://github.com/",
    "size": 512,
    "recoveryLevel": "high"
  },
  "qr": "iVBORw0KGgoAAAANSUhEUgAA..."
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
CLIHandler(qrService, "Hello World")
```

-----

## 🤝 Contributing

Contributions are what make the open-source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1.  Fork the Project
2.  Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3.  Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4.  Push to the Branch (`git push origin feature/AmazingFeature`)
5.  Open a Pull Request

-----

## ⚖️ License

Distributed under the **MIT License**. See `LICENSE` for more information.