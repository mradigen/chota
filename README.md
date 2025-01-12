# chota

A URL shortener written in Go. Supporting both in-memory and PostgreSQL storage backends.

## TODO

- [ ] Modify in memory storage to use `sync.RWMutex` instead of maps
- [ ] Create a frontend under `/website`
- [ ] Return saved slug if URL already exists in Storage
- [ ] Ability to choose custom slug
- [ ] Debug logging
- [ ] Add a `deploy/docker-compose.yml`

## Table of Contents

- [Introduction](#introduction)
- [Usage](#usage)
- [Installation](#installation)
- [Configuration](#configuration)
- [Deployment](#deployment)
- [Development](#development)
- [Contributing](#contributing)

## Usage

The following endpoints are available:

- `GET /shorten?url=:URL`: Shorten a new URL.
- `GET /:slug`: Redirect to the original URL.

Example:

```sh
curl 'http://localhost:8080/shorten?url=https://example.com'
```

Response:

```json
{ "short_url": "pvog" }
```

## Installation

Clone the repository:

```sh
git clone https://github.com/mradigen/chota
cd chota
```

### Docker

1. Build the Docker image:

    ```sh
    make docker
    ```

2. Copy the `.env.example` file and edit it as needed (refer [configuration](#configuration)):

    ```sh
    cp .env.example .env
    ```

3. Run the Docker container:
    ```sh
    docker run -p 8080:8080 --env-file=.env chota
    ```

### Native

1. Build the project:

    ```sh
    make build
    ```

2. Run:
    ```sh
    ./chota
    ```

## Configuration

The application supports environment-based configuration:

| Variable                | Default          | Description                                                            |
| ----------------------- | ---------------- | ---------------------------------------------------------------------- |
| `PORT`                  | `8080`           | Port for the HTTP server.                                              |
| `STORAGE_MODE`          | `memory`         | Storage backend: `memory` or `postgres`.                               |
| `DATABASE_URL`          | `postgres://...` | Connection string for PostgreSQL. Used only if `STORAGE_MODE=postgres` |
| `BIND_ADDRESS`          | `127.0.0.1`      | Address to listen on.                                                  |
| `DEBUG` (unimplemented) | `false`          | Log debug activity.                                                    |

## Deployment

### Kubernetes

Use the provided `kubernetes.yml` file to deploy the application to a Kubernetes cluster:

```sh
make kubernetes
```

or

```sh
kubectl apply -f deploy/kubernetes.yml
```

## Development

### Run

```sh
make run
```

### Testing:

Currently there is a very basic test setup. Run the unit and integration tests using:

```sh
make test
```

## Contributing

Contributions are highly encouraged! Please open an issue or submit a pull request with your changes.
