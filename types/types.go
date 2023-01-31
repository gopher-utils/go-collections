package types

type Number interface {
	int64 | float64 | int32 | int16 | int8 | int | float32
}

type GenericType interface {
	Number | string
}
