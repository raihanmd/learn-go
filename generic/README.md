# Hello! this is my repo for generics in golang

## What is generics in golang?

generic is a way of writing reusable code that can work with different types

## How to use generics in golang?

```go
package main

func Foo[T any](n T) T {
  return n
}
```

it might be different like for typescript with `<>` bracket
