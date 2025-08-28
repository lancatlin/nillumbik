package author

import (
	"github.com/lancatlin/nillumbik/internal/db"
)

func toAuthor(a db.Author) Author {
	return Author{
		a.ID, a.Name, a.Bio,
	}
}

func mapSlice[K any, V any](conv func(K) V) func([]K) []V {
	return func(inputs []K) []V {
		output := make([]V, len(inputs))
		for i, input := range inputs {
			output[i] = conv(input)
		}
		return output
	}
}

var toAuthorSlice = mapSlice(toAuthor)
