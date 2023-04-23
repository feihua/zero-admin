Insert(ctx context.Context, data *{{.upperStartCamelObject}}) (sql.Result,error)
InsertTx(ctx context.Context,session sqlx.Session, data *{{.upperStartCamelObject}}) (sql.Result,error)