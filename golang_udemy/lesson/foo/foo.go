package foo

const (
	// public
	Max = 100
	// private
	min = 1
)

// public
func ReturnMin() int {
	return min
}
