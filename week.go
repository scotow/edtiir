package edtiir

import (
	"fmt"
	"strconv"
)

type Week struct {
	mainLink string
	index    int
	config   Config
}

func (w *Week) Link() string {
	return fmt.Sprintf("%s&range=%s", w.mainLink, w.cellsRange())
}

func (w *Week) cellsRange() string {
	upper := w.config.padding.top + w.index*(w.config.size.height+w.config.spacing)

	upperLeft := toColumnLetter(w.config.padding.left) + toRowNumber(upper)
	lowerRight := toColumnLetter(w.config.padding.left+w.config.size.width-1) + toRowNumber(upper+w.config.size.height-1)

	return fmt.Sprintf("%s:%s", upperLeft, lowerRight)
}

func toColumnLetter(i int) string {
	return string('A' + i)
}

func toRowNumber(i int) string {
	return strconv.Itoa(i + 1)
}
