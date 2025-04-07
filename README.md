# Timeline Service with gRPC and GraphQL

A timeline service that provides both gRPC and GraphQL APIs for fetching user timelines. Built with Go, gRPC, and GraphQL.

## Features

- GraphQL API for fetching user timelines
- gRPC API for timeline operations
- In-memory store with mock data
- Docker support for easy deployment

## API Endpoints

### GraphQL

- Endpoint: `http://localhost:8080`
- Sample Query:

```graphql
query GetTimeline($userId: ID!) {
  getTimeline(userId: $userId) {
    id
    content
    author
    timestamp
  }
}
```

### gRPC

- Endpoint: `localhost:50051`

## Getting Started

### Prerequisites

- Docker installed on your system

### Running with Docker

1. Pull the Docker image:

```bash
docker pull pawan063/timeline-grpc-go
```

2. Run the container:

```bash
docker run -p 8080:8080 -p 50051:50051 pawan063/timeline-grpc-go
```

This will start the service with:

- GraphQL server on port 8080
- gRPC server on port 50051

### Testing the API

1. GraphQL:

   - Open your browser and navigate to `http://localhost:8080`
   - Use the following query to test:

   ```graphql
   query GetTimeline($userId: ID!) {
     getTimeline(userId: $userId) {
       id
       content
       author
       timestamp
     }
   }
   ```

   - Variables:

   ```json
   {
     "userId": "1"
   }
   ```

2. gRPC:
   - Use a gRPC client to connect to `localhost:50051`

## Mock Data

The service comes with pre-populated mock data:

- 5 users (IDs: 1-5)
- 20 posts per user
- Follow relationships between users

## Project Structure

```
.
├── server/
│   ├── graph/          # GraphQL implementation
│   ├── grpc/           # gRPC implementation
│   └── store/          # In-memory data store
└── Dockerfile          # Docker configuration
```

## License

MIT License
