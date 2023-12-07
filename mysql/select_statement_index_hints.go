package mysql

func (s *selectStatementImpl) INDEX_HINTS(hints ...IndexHint) SelectStatement {
	s.From.IndexHints = hints
	return s
}
