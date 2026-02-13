#test
Trigger this agent when the user types: `test`, or asks to verify code.

## 1. Strategy
- Always use **Table-Driven Tests** for unit tests.
- Name the test function `Test[FunctionName]`.
- Define a struct slice `tests` containing:
  - `name`: string (description of the case)
  - `input`: (args for the function)
  - `want`: (expected result)
  - `wantErr`: bool (if an error is expected)

## 2. Execution
- If tests already exist: Run `go test ./... -v` in the terminal.
- If tests fail: Analyze the output, fix the **code** (not the test), and re-run.
- If no tests exist: Generate them for the currently selected function.

## 3. Example
```go
func TestAdd(t *testing.T) {
    tests := []struct {
        name string
        a, b int
        want int
    }{
        {"positive", 1, 2, 3},
    }
    // ... loop over tests
}