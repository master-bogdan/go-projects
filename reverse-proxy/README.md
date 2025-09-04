# Reverse Proxy Project

## Overview
The Reverse Proxy project is a Go application designed to act as a load balancer and reverse proxy. It supports multiple backends and provides configurable balancing strategies.

## Project Structure
```
reverse-proxy/
├── config.yaml           # Configuration file for the reverse proxy
├── go.mod                # Go module file
├── go.sum                # Go module ~~dependencies~~
├── balancers/
│   ├── balancer.go       # Interface for load balancers
│   ├── least-conn.go     # Least connections balancing strategy
│   └── round-robin.go    # Round-robin balancing strategy
├── cmd/
│   └── main.go           # Entry point of the application
├── config/
│   └── config.go         # Configuration parsing and management
├── k8s/
│   ├── dummy-backend-deployment.yml # Kubernetes deployment for dummy backend
│   └── dummy-backend-service.yml    # Kubernetes service for dummy backend
└── proxy-server/
    └── proxy-server.go   # Core reverse proxy server logic
```

## Features
- **Load Balancing**: Supports round-robin and least-connections strategies.
- **Health Checks**: Configurable health checks for backend servers.
- **Retry Mechanism**: Retries failed requests with exponential backoff.
- **Timeouts**: Configurable timeouts for requests and connections.
- **Kubernetes Support**: Includes example Kubernetes manifests for deployment.

## Configuration
The application is configured using the `config.yaml` file. Below is an example configuration:
```yaml
listen: ":8080"
backends:
  - "http://192.168.49.2:30080"
balancer: "round_robin"    # or "least_conn"
health:
  path: "/health"
  interval: 2
  timeout: 500
  passiveFailuresForOpen: 5
  cooldown: 10
retry:
  max: 2
  backoff: 200
timeout:
  read: 5
  write: 30
  idle: 60
transport:
  dialTimeout: 500
  tlsHandshakeTimeout: 500
  maxIdlePerHost: 100
```

## Getting Started

### Prerequisites
- Go 1.20 or later

### Installation
1. Clone the repository:
   ```bash
   git clone <repository-url>
   ```
2. Navigate to the project directory:
   ```bash
   cd reverse-proxy
   ```
3. Install dependencies:
   ```bash
   go mod tidy
   ```

### Running the Application
To start the application, run:
```bash
go run cmd/main.go
```

### Kubernetes Deployment
1. Apply the Kubernetes manifests:
   ```bash
   kubectl apply -f k8s/dummy-backend-deployment.yml
   kubectl apply -f k8s/dummy-backend-service.yml
   ```
2. Update the `backends` field in `config.yaml` with the service URL.

## Contributing
Contributions are welcome! Feel free to open issues or submit pull requests.

## License
This project is licensed under the MIT License. See the LICENSE file for details.

## Acknowledgments
- Inspired by Go's `net/http` package and its extensibility.
- Special thanks to the Go community for their excellent resources and support.
