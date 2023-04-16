package collections

// Generic interface that all collections must implement.
type Collection[T CollectionElement] interface {
	// Returns the type of the given collection.
	Type() CollectionType
	// Returns the size of the given collection.
	Size() int
	// Returns the description of the given collection.
	String() string
}

// A Generic interface that must be implemented by collection elements.
// By default, all primitive types and structs containing primitive types implement this interface.
type CollectionElement interface {
	Equatable
}

// Typed constant that helps in determining the collection type.
type CollectionType int

const (
	LIST    CollectionType = 0
	SET                    = 1
	HASHMAP                = 2
)

// Interface with comparable constraint (== & != operator supportable).
type Equatable interface {
	comparable
}

// A Generic number interface that supports all basic number types.
type Number interface {
	int64 | float64 | int32 | int16 | int8 | int | float32
}

type GenericType interface {
	Number | string
}
