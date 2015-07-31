package linter

//import (
//	"golang.org/x/net/context"
//)

// Iterator contains documents and provides them in fixed-sized batches.
type Iterator interface {
	Name() string
	Count() int
	Next(step int) []*Checker
}

//func Name(ctx context.Context) string {
//	return IteratorFromContext(ctx).Name()
//}

func Name(iterator Iterator) string {
	return iterator.Name()
}

//func Count(ctx context.Context) int64{
//	return IteratorFromContext(ctx).Count()
//}

func Count(iterator Iterator) int{
	return iterator.Count()
}

//func Next(ctx context.Context, step int) []*Checker {
//	return IteratorFromContext(ctx).Next(step)
//}

func Next(iterator Iterator, step int) []*Checker {
	return iterator.Next(step)
}