### panic的处理

```go
	defer func() {
		if err := recover(); err != nil {
			log.Logger.Error(err)
			fmt.Printf("panic recover! err: %v", err)
			debug.PrintStack()
		}
	}()
```