FindOne(ctx context.Context, {{.lowerStartCamelPrimaryKey}} {{.dataType}}) (*{{.upperStartCamelObject}}, error)
FindAll(ctx context.Context, Current int64, PageSize int64) (*[]{{.upperStartCamelObject}}, error)
Count(ctx context.Context) (int64,error)