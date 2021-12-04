<div align="center">
  <h1>
    <img src="https://raw.githubusercontent.com/ashleymcnamara/gophers/master/GopherLink.png" width="192" /><br />
    Abstract
  </h1>
</div>

## For Loops

They're really simple:

```go
 for i := 0; i < 5; i++ {
   // do whatever you want here
 }
```

there are no parenthesis here but `{}` are required.

Other examples:

```go
 for i <= 3{
   // do whatever you want here
 }
```

```go
 for {
   // do whatever you want here
   break
 }
```

```go
  for n := 0; n <= 5; n++ {
    if n % 2 == 0 {
      continue
    }
    fmt.Println(n)
  }
```

## Benchmarks

To create benchmarks in go you need to write a function like this one:


```go
func Benchmarkxxxxxx(b *testing.B) {
  for i := 0; i < b.N; i++ {
    // call your function
  }
}
```

`testing.B` give us access to `b.N` which is the number of times this piece of code is going to run.

By running `go test -bench=.` on my terminal, this was the result:

```
  goos: darwin
  goarch: amd64
  pkg: github.com/oaraujocesar/iteration-golang
  cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
  BenchmarkRepeat-12      11480233               106.3 ns/op
  PASS
  ok      github.com/oaraujocesar/iteration-golang        2.699s
```

`11480233` is the number of times my function got executed

`106.3 ns/op` means that every operation takes 106.3 nanoseconds to execute

Benchmarks in Go are executed in sequence by default.
