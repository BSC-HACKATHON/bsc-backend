package utils

func Zero[T any]() T {
	v := new(T)
	return *v
}
