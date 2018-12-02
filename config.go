package edtiir

type Config struct {
	size    Size
	padding Padding
	spacing int
}

type Size struct {
	width  int
	height int
}

type Padding struct {
	top  int
	left int
}
