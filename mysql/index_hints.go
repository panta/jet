package mysql

import (
	"fmt"

	"github.com/go-jet/jet/v2/internal/jet"
)

// IndexHint provides a way to optimize query execution per-statement basis
type IndexHint = jet.IndexHint

// USE_INDEX adds a "USE INDEX" clause
func USE_INDEX(indexName string) IndexHint {
	return IndexHint(fmt.Sprintf("USE INDEX(%s)", indexName))
}

// FORCE_INDEX adds a "FORCE INDEX" clause
func FORCE_INDEX(indexName string) IndexHint {
	return IndexHint(fmt.Sprintf("FORCE INDEX(%s)", indexName))
}
