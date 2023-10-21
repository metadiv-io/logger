package types

// ILogger is an interface for a logger.
type ILogger interface {
	SetPrefix(prefix string)
	Println(v ...any)
}
