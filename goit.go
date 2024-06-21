package goit

type iterator[E any] interface {
	Next() (E, bool)
}

func OverStruct[E any](it iterator[E]) func(func(E) bool) {
	return func(yield func(E) bool) {
		for {
			res, ok := it.Next()
			if !yield(res) || !ok {
				return
			}
		}
	}
}
