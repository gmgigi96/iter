# Go iterators

The `iter` package provides a set of utilities for creating and manipulating iterators in Go. It leverages Go's generics feature to provide type-safe operations on iterators.

## Features
- *Basic Iteration*: Define and work with iterators using the `Iter` interface.
- *Generics*: Use Go's generics to create type-safe iterators.
- *Transformation*: Apply `Map`, `Filter`, and other transformations on iterators.
- *Infinite Iterators*: Create infinite iterators using `Count`, `Repeat` and `Cycle`.
- *Map Iterators*: Create iterators from Go maps and extract keys or values.
- *Utilities*: Various utility functions to convert from/to Go built-in data structures.
- *Advanced Operations*: Functions like `Accumulate`, `Reduce`, `Chain, and more for advanced iterator operations.

## Usage
To use the `iter` package, simply import it in your Go code:

```go
import "github.com/gmgigi96/iter"
```
