# CluIns

CluIns is a lightweight Go tool that streams real-time CPU and memory usage metrics over HTTP using Server-Sent Events (SSE). It is designed for easy integration into cluster-based environments, such as Kubernetes, for basic resource monitoring and diagnostics.

## Features
- Streams live CPU and memory statistics via HTTP endpoint
- Uses dynamic port allocation and detects the host's local IP
- Ready for containerization and deployment in clusters

## Getting Started

### 1. Clone the Repository
```sh
git clone https://github.com/ABAlosaimi/CluIns.git
cd CluIns
```

### 2. Build and Run Locally
Make sure you have Go installed (1.25.1 recommended):

```sh
go build -o cluins ./main
./cluins
```
The server will print the TCP/IP address. Access the metrics endpoint at:
```
http://<host-ip>:<port>/metric/resources
```

### 3. Docker image
#### You can pull the image with:
```sh
docker pull afalosaimi/cluins 
```
#### Run the Docker container:
```sh
docker run afalosaimi/cluins:latest
```
> Note: The application uses dynamic port allocation.

### 4. Usage in Cluster-Based Environments (e.g., Kubernetes)
- Deploy the Docker image as a Pod or Deployment.
- Use a Service to expose the metrics endpoint within the cluster.
- For dynamic port allocation, consider using a headless Service or configure the app to use a fixed port for easier service discovery.

## Metrics Endpoint
- **URL:** `<host-ip>:<prot>/metric/resources`
- **Method:** `GET`
- **Response:** Server-Sent Events (SSE) streaming CPU and memory usage metricies

## Author
ABAlosaimi