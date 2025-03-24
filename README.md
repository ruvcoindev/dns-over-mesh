# DNS-over-MESH

DNS-over-MESH is a project that provides decentralized DNS resolution using a mesh network. It includes features such as DNS-over-TLS for secure queries and a web interface for monitoring.

## Features

- Decentralized DNS resolution using a mesh network.
- DNS record encryption using Ed25519.
- DNS-over-TLS support for secure queries.
- Web interface for monitoring the DNS server.

## Getting Started

### Prerequisites

- Go 1.18 or later
- Docker (optional, for containerized deployment)

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/ruvcoindev/dns-over-mesh.git
   cd dns-over-mesh

Build the project:

./scripts/build.sh

Run the project:

./dns-over-mesh

Configuration
Edit the config/config.yaml file to configure the mesh network nodes and other settings.

Monitoring
Access the monitoring web interface at http://localhost:8080.

Testing
Run the tests using:

go test ./...

Contributing
Contributions are welcome! Please open an issue or submit a pull request.

License
This project is licensed under the MIT License.
