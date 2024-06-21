# Goit

Because this feels so much simpler than func(func(func))...

## Warning

This is an experimental feature of Go1.23. To use this, `GOEXPERIMENT=rangefunc` must be set in your environment variables.

Read more here: https://go.dev/wiki/RangefuncExperiment

## Example

```go
package goit_test

import (
	"testing"

	"github.com/acheong08/goit"
)

type forIterator struct {
	count int
	max   int
}

func (f *forIterator) Next() (c int, ok bool) {
	ok = true
	f.count++
	if f.count >= f.max {
		ok = false
	}
	return f.count, ok
}

func TestBasicLoop(t *testing.T) {
	it := forIterator{
		max: 10,
	}
	i := 0
	for n := range goit.OverStruct[int](&it) {
		i = n
	}
	if i != 10 {
		t.Errorf("Expected 10, got %d", i)
	}
}
```
