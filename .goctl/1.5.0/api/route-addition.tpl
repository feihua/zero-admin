
	server.AddRoutes(
		{{.routes}} {{.jwt}}{{.signature}} {{.prefix}} {{.timeout}} {{.maxBytes}}
	)
