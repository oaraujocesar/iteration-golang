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
