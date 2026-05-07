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

- [x] `14loops/` — for loops, range loops, break/continue
- [x] `15functions/` — Multiple return values, named returns, variadic args
- [x] `16closures/` — Anonymous functions, closures, functions as values
- [x] `17defer/` — `defer`, `panic`, `recover`, execution order
- [ ] `18errors/` — `error` interface, custom errors, `fmt.Errorf`, `errors.Is/As`
- [ ] `19interfaces/` — Implicit implementation, empty interface, type assertions
- [ ] `20type_switch/` — `switch x.(type)` pattern
- [ ] `21stringer/` — `fmt.Stringer`, custom `String()` methods
- [ ] `22embedding/` — Struct embedding, composition over inheritance

---

## Phase 2 — Concurrency (Go's Superpower)

- [ ] `23goroutines/` — `go` keyword, lightweight threads, data races
- [ ] `24channels/` — Unbuffered vs buffered, range over channels, closing, channel direction (`chan<-`, `<-chan`)
- [ ] `25select/` — Multiplexing channels with `select`
- [ ] `26sync/` — `WaitGroup`, `Mutex`, `RWMutex`
- [ ] `27atomic/` — `sync/atomic` operations
- [ ] `28context/` — Cancellation, timeouts, request-scoped values

---

## Phase 3 — Standard Library Essentials

- [x] `29file_io/` — `os`, `io`, `bufio` — reading and writing files
- [ ] `30io_interfaces/` — `io.Reader`, `io.Writer`, `io.Closer` — composable I/O interfaces
- [ ] `31json/` — `encoding/json`, marshal/unmarshal, struct tags
- [x] `32http/` — `net/http`, `http.Handler`/`HandlerFunc` pattern, building servers and clients
- [ ] `33strings/` — `strings` and `strconv` packages
- [ ] `34regexp/` — Regular expressions with `regexp` package
- [ ] `35time_advanced/` — Timers, tickers, `time.After`, `time.Tick`
- [ ] `36slog/` — `log/slog`, structured logging, log levels, custom handlers

---

## Phase 4 — Modules & Tooling

- [ ] `37modules/` — `go.mod`, `go.sum`, `go get`, `go tidy`, versioning
- [ ] `38packages/` — Exported vs unexported, organizing multi-package projects
- [ ] `39testing/` — `testing` package, `go test`, table-driven tests, benchmarks
- [ ] `40linting/` — `go vet`, `golangci-lint`, code formatting with `gofmt`
- [ ] `41build_tags/` — Conditional compilation with build constraints

---

## Phase 5 — Intermediate Patterns

- [ ] `42generics/` — Type parameters, constraints (Go 1.18+)
- [ ] `43functional_options/` — Idiomatic config/builder pattern
- [ ] `44middleware/` — HTTP middleware chaining
- [ ] `45worker_pool/` — Bounded concurrency with goroutines + channels
- [ ] `46errgroup/` — `golang.org/x/sync/errgroup` for concurrent error handling

---

## Phase 6 — Advanced Topics

- [ ] `47reflection/` — `reflect` package, inspecting types at runtime
- [ ] `48profiling/` — `pprof`, CPU and memory profiling, `go tool pprof`
- [ ] `49database/` — `database/sql`, drivers, connection pooling
- [ ] `50grpc/` — Protocol Buffers, gRPC server and client
- [ ] `51unsafe/` — `unsafe` package, memory layout (when and why to avoid)

---

## Resources

- [Official Go Tour](https://go.dev/tour/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go by Example](https://gobyexample.com/)
- [Go Standard Library Docs](https://pkg.go.dev/std)
- [Go Playground](https://go.dev/play/)
- [Go Blog](https://go.dev/blog/)
- [Go Wiki](https://github.com/golang/go/wiki)
