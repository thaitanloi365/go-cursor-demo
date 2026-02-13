# Project Architecture

## Design Philosophy
We follow a simple Layered Architecture.

1. **Models (`models.go`)**: 
   - Contains all struct definitions (e.g., `Todo`).
   - Includes JSON tags.
   - Pure data, no logic.

2. **Handlers (`handlers.go`)**:
   - Contains HTTP handler functions.
   - Responsible for parsing requests and writing responses.

3. **Routing (`main.go`)**:
   - Initializes the router.
   - Maps endpoints to Handlers.
   - Starts the server.

## API Standards
- **POST** requests must return `201 Created`.
- **GET** requests must return `200 OK`.