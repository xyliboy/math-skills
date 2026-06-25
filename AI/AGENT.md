AGENTS.md – Development & Review Guidelines (Statistics Project)
This document defines how development, testing, and review should be performed for this statistics CLI project.

1. Coding Rules
1.1 Standard Library Only
Only the Go standard library is allowed. No external packages should be used.

1.2 Code Style & Structure
All code must be formatted using:
gofmt -w .
Functions must be small (ideally <30 lines) and focused on a single responsibility.
Avoid complex nested logic. Prefer clear and readable code.
1.3 Naming Conventions
Variables: camelCase (e.g., numbers, filePath, sum)
Functions: verb-based (e.g., calculateAverage, readFile, computeVariance)
Types/Constants: PascalCase (e.g., StatsResult)
2. Project Goal (IMPORTANT)
The program must:

Read numbers from a file (one number per line)
Compute:
Average
Median
Variance
Standard Deviation
Output results as rounded integers:
Example:

Average: 35
Median: 4
Variance: 5
Standard Deviation: 65
3. Program Rules
3.1 Input Handling
Input is a file path passed as CLI argument:
go run . data.txt
Each line = one number
Ignore empty lines (if any)
Handle invalid lines safely (decide policy: error or skip)
3.2 Output Rules
Output must match exact format
Values must be rounded integers
Order must be:
Average
Median
Variance
Standard Deviation
4. Testing Rules
4.1 Test-Driven Development (TDD)
Always write tests before implementation
Red → Green → Refactor cycle
4.2 Unit Tests
Each function must be testable independently:

Functions to test:

readNumbers()
calculateAverage()
calculateMedian()
calculateVariance()
calculateStdDev()
4.3 Golden Tests
Golden tests must verify full program output.

Example:

Input file:

1
2
3
Expected:

Average: 2
Median: 2
Variance: 1
Standard Deviation: 1
5. Core Implementation Rules
5.1 Average
sum / count
5.2 Median
Sort numbers
Odd → middle value
Even → average of two middle values
5.3 Variance
variance = Σ(x - mean)^2 / N
5.4 Standard Deviation
sqrt(variance)
5.5 Rounding
All final outputs must be rounded using standard rounding.

6. Edge Cases (CRITICAL)
You MUST handle:

Empty file
File with 1 number
Even vs odd dataset size
Negative numbers
Large numbers
Invalid input (non-numeric lines)
Division by zero (empty dataset)
Define behavior explicitly:

Error OR skip invalid values
7. Project Structure
.
├── main.go
├── internal/
│   ├── io.go
│   ├── stats.go
├── tests/
│   ├── unit/
│   ├── testdata/
8. Using AI in Development
8.1 Allowed
Use AI only when:

You tried solving the problem yourself
You want to understand WHY something works
You explore alternative implementations
8.2 NOT Allowed
Do NOT copy full solutions
Do NOT let AI design the entire program
Do NOT skip understanding math formulas
Do NOT rely on AI without testing
8.3 AI Usage Validation
You are using AI correctly if:

You can explain your code without help
You can reimplement logic alone
You understand each formula
9. Code Review Rules
Before committing:

Code is formatted (gofmt)
No unused variables
All tests pass
Edge cases covered
Output format EXACT
10. Tooling & CI
Before every commit run:

gofmt -w .
go vet ./...
go test ./...
If any fails → DO NOT commit.