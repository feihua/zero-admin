func (m *default{{.upperStartCamelObject}}Model) FindAll(ctx context.Context, Current int64, PageSize int64) (*[]{{.upperStartCamelObject}}, error) {
	{{if .withCache}}{{.cacheKey}}
	var resp []{{.upperStartCamelObject}}
	err := m.QueryRowCtx(ctx, &resp, {{.cacheKeyVariable}}, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query :=  fmt.Sprintf("select %s from %s limit {{if .postgreSql}}$1{{else}}?{{end}},{{if .postgreSql}}$2{{else}}?{{end}}", {{.lowerStartCamelObject}}Rows, m.table)
		return conn.QueryRowsCtx(ctx, v, query, (Current-1)*PageSize, PageSize)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}{{else}}query := fmt.Sprintf("select %s from %s limit {{if .postgreSql}}$1{{else}}?{{end}},{{if .postgreSql}}$2{{else}}?{{end}}", {{.lowerStartCamelObject}}Rows, m.table)
	var resp []{{.upperStartCamelObject}}
	err := m.conn.QueryRowsCtx(ctx, &resp, query, (Current-1)*PageSize, PageSize)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}{{end}}
}
