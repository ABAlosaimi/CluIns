# CluIns

CluIns is a lightweight Go application that streams real-time CPU and memory usage metrics over HTTP using Server-Sent Events (SSE). It is designed for easy integration into cluster-based environments, such as Kubernetes, for resource monitoring and diagnostics.

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
Make sure you have Go installed (1.21+ recommended):
```sh
go build -o cluins ./main
./cluins
```
The server will print the address and port it is listening on. Access the metrics endpoint at:
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
docker run --rm afalosaimi/cluins:latest
```
> Note: The application uses dynamic port allocation. To map the container's dynamic port to a host port, you may need to inspect the container logs to find the actual port, or modify the code to use a fixed port for easier mapping.

### 4. Usage in Cluster-Based Environments (e.g., Kubernetes)
- Deploy the Docker image as a Pod or Deployment.
- Use a Service to expose the metrics endpoint within the cluster.
- For dynamic port allocation, consider using a headless Service or configure the app to use a fixed port for easier service discovery.

## Metrics Endpoint
- **URL:** `<ip>:<prot>/metric/resources`
- **Method:** `GET`
- **Response:** Server-Sent Events (SSE) streaming CPU and memory usage

## Author
[ABAlosaimi](https://github.com/ABAlosaimi)