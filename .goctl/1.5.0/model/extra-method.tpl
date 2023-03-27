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

func (m *default{{.upperStartCamelObject}}Model) Count(ctx context.Context) (int64, error) {
	{{if .withCache}}{{.cacheKey}}
	var count int64
	err := m.QueryRowCtx(ctx, &resp, {{.cacheKeyVariable}}, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query :=  fmt.Sprintf("select count(*) as count from %s", m.table)
		return conn.QueryRowCtx(ctx, v, query)
	})
	switch err {
	case nil:
		return count, nil
	case sqlc.ErrNotFound:
		return 0, ErrNotFound
	default:
		return 0, err
	}{{else}}query := fmt.Sprintf("select count(*) as count from %s", m.table)
	var count int64
	err := m.conn.QueryRowCtx(ctx, &count, query)
	switch err {
	case nil:
		return count, nil
	case sqlc.ErrNotFound:
		return 0, ErrNotFound
	default:
		return 0, err
	}{{end}}
}
