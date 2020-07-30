package filter

const (
	// Ignore ignore the message
	Ignore = iota
	// Delete delete the message
	Delete
	// Ban ban the user
	Ban
)

// DeleteHandler a message delete wrapper
type DeleteHandler func()

// BanHandler a ban user wrapper
type BanHandler func()

// Filter basic filter interface
type Filter interface {
	ShouldApply(bool, bool) bool
	Apply(string) int
}
