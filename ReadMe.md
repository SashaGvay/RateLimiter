# rate_limiter

`rate_limiter` — is a lightweight wrapper around `time.Ticker` to limit the frequency of operations in Go


### install:
```
go get github.com/yourusername/rate_limiter
```

### use:
```go
// only 5 request in second
func Foo(ctx context.Context, requests []Request) {
    limiter := rate_limiter.NewRateLimiter(time.Second / 5)
    defer limiter.Stop()

    for _, request := range requests {
        limiter.Wait()
        request.Do()
    }
}
```