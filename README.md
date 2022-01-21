They're neck-and-neck when parsing ~2700 Go files:

```
$ go run main.go
ctags 1.618385084s
tree-sitter 1.64496675s
```

That's ~1600 files parsed per second, or 0.6ms each.
