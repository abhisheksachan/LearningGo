# Learning Go — Roadmap

Track your progress by checking off each topic as you complete it.

---

## ✅ Phase 0 — Basics (Completed)

- [x] `01hello/` — Hello World, packages, main function
- [x] `02variables/` — Variables, constants, zero values, short declaration
- [x] `03userinput/` — Reading user input with `fmt.Scan` / `bufio`
- [x] `04conversion/` — Type conversion and casting
- [x] `05mytime/` — `time` package basics
- [x] `06mypointers/` — Pointers, `&` and `*`, pass by reference
- [x] `08array/` — Fixed-size arrays
- [x] `09slices/` — Slices, `append`, `copy`, slice tricks
- [x] `10maps/` — Maps, CRUD operations, iteration
- [x] `11structs/` — Structs, methods, value vs pointer receivers
- [x] `12ifelse/` — If/else, initializer statement
- [x] `13switch/` — Switch, fallthrough, expression-less switch

---

## Phase 1 — Core Language

- [ ] `14loops/` — for loops, range loops, break/continue
- [x] `15functions/` — Multiple return values, named returns, variadic args
- [ ] `16closures/` — Anonymous functions, closures, functions as values
- [x] `17defer/` — `defer`, `panic`, `recover`, execution order
- [ ] `18errors/` — `error` interface, custom errors, `fmt.Errorf`, `errors.Is/As`
- [ ] `19interfaces/` — Implicit implementation, empty interface, type assertions
- [ ] `20type_switch/` — `switch x.(type)` pattern
- [ ] `21stringer/` — `fmt.Stringer`, custom `String()` methods
- [ ] `22embedding/` — Struct embedding, composition over inheritance

---

## Phase 2 — Concurrency (Go's Superpower)

- [ ] `23goroutines/` — `go` keyword, lightweight threads, data races
- [ ] `24channels/` — Unbuffered vs buffered, range over channels, closing
- [ ] `25select/` — Multiplexing channels with `select`
- [ ] `26sync/` — `WaitGroup`, `Mutex`, `RWMutex`
- [ ] `27atomic/` — `sync/atomic` operations
- [ ] `28context/` — Cancellation, timeouts, request-scoped values

---

## Phase 3 — Standard Library Essentials

- [ ] `29file_io/` — `os`, `io`, `bufio` — reading and writing files
- [ ] `30json/` — `encoding/json`, marshal/unmarshal, struct tags
- [ ] `31http/` — `net/http`, building HTTP servers and clients
- [ ] `32strings/` — `strings` and `strconv` packages
- [ ] `33regexp/` — Regular expressions with `regexp` package
- [ ] `34time_advanced/` — Timers, tickers, `time.After`, `time.Tick`

---

## Phase 4 — Modules & Tooling

- [ ] `35modules/` — `go.mod`, `go.sum`, `go get`, `go tidy`, versioning
- [ ] `36packages/` — Exported vs unexported, organizing multi-package projects
- [ ] `37testing/` — `testing` package, `go test`, table-driven tests, benchmarks
- [ ] `38linting/` — `go vet`, `golangci-lint`, code formatting with `gofmt`
- [ ] `39build_tags/` — Conditional compilation with build constraints

---

## Phase 5 — Intermediate Patterns

- [ ] `40generics/` — Type parameters, constraints (Go 1.18+)
- [ ] `41functional_options/` — Idiomatic config/builder pattern
- [ ] `42middleware/` — HTTP middleware chaining
- [ ] `43worker_pool/` — Bounded concurrency with goroutines + channels
- [ ] `44errgroup/` — `golang.org/x/sync/errgroup` for concurrent error handling

---

## Phase 6 — Advanced Topics

- [ ] `45reflection/` — `reflect` package, inspecting types at runtime
- [ ] `46profiling/` — `pprof`, CPU and memory profiling, `go tool pprof`
- [ ] `47database/` — `database/sql`, drivers, connection pooling
- [ ] `48grpc/` — Protocol Buffers, gRPC server and client
- [ ] `49unsafe/` — `unsafe` package, memory layout (when and why to avoid)

---

## Resources

- [Official Go Tour](https://go.dev/tour/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go by Example](https://gobyexample.com/)
- [Go Standard Library Docs](https://pkg.go.dev/std)
- [Go Playground](https://go.dev/play/)
