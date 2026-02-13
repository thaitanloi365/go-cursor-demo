# create-pr
This is your "Release Manager". It handles the shipping workflow you requested.
 
## Workflow
You must follow this sequence strictly:

1.  **Pre-Flight Check**:
    - Run `go vet ./...` and `go test ./...`.
    - If any check fails, **STOP** and ask the user to fix the code first.

2.  **Generate PR Content**:
    - Analyze the difference between the current code and the previous state.
    - Generate a Markdown summary with the following sections:

### Title
- Format: `<type>(<scope>): <short description>`
- Example: `feat(api): add POST /todos handler`

### Description
- A detailed, bulleted list of exactly what changed.
- Mention any architectural changes (e.g., "Moved struct to models.go").

### Test Plan
- List specific scenarios that were verified.
- Example:
  - [x] Verified `POST /todos` creates a valid JSON item.
  - [x] Verified `GET /todos` returns list.
  - [x] Verified error handling for invalid JSON.
