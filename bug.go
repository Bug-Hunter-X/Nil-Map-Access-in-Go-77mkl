```go
func main() {
    var m map[string]int
    fmt.Println(m["foo"]) // This will not panic, it will return 0.
}
```