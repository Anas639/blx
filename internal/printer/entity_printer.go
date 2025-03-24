package printer

type EntityPrinter[T any] interface {
	PrintSingle(task *T)
	PrintMany(tasks []*T)
}
