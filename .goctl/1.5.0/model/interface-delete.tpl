Delete(ctx context.Context, {{.lowerStartCamelPrimaryKey}} {{.dataType}}) error
DeleteTx(ctx context.Context,session sqlx.Session, {{.lowerStartCamelPrimaryKey}} {{.dataType}}) error