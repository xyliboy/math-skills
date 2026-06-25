# math-skills

Small CLI program that reads one number per line from a file and prints:

- Average
- Median
- Variance
- Standard Deviation

All results are rounded integers.

## Run

```sh
go run . data.txt
```

Expected output format:

```txt
Average: 35
Median: 4
Variance: 5
Standard Deviation: 65
```

## Test

```sh
gofmt -w .
go vet ./...
go test ./...
```

## Structure

```txt
.
├── main.go
├── internal/
│   ├── io.go
│   └── stats.go
└── tests/
    ├── unit/
    └── testdata/
```

## Input Rules

- One number per line.
- Empty lines are ignored.
- Invalid non-numeric lines return an error.
- Empty files return an error.
