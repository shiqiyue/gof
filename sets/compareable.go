package sets

type Compareable interface {
	CompareTo(Compareable) int8
}
