package jet

// IndexHint provides a way to optimize query execution per-statement basis
type IndexHint string

type indexHints []IndexHint

func (o indexHints) Serialize(statementType StatementType, out *SQLBuilder, options ...SerializeOption) {
	if len(o) == 0 {
		return
	}

	for i, hint := range o {
		if i > 0 {
			out.WriteByte(' ')
		}

		out.WriteString(string(hint))
	}
}
