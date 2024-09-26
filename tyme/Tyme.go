package tyme

// Tyme Tyme
type Tyme interface {
	Culture

	// Next 推移
	Next(n int) Tyme
}
