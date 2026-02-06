# Day 6 - Thread-Safe Cache

## Objective
Create a thread-safe key/value cache.

## Tasks
- Create a Cache struct with an internal map.
- Protect access using sync.RWMutex.
- Add optional expiration.
- Add unit tests.
- Add a Store interface.

## Go Concepts
Mutex, concurrency, testing, design.

## Extra Exercises
- Add automatic cleanup (goroutine + ticker).
- Add disk backend (JSON).
- Add benchmarks.
