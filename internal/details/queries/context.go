package queries

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/senago/sputnik/internal/details/db"
)

func resolveConn(ctx context.Context) db.Executor {
	if conn, ok := db.GetConnContext(ctx); ok {
		return conn.Resolver(ctx)
	}

	panic(errors.AssertionFailedf("insufficient context to resolve conn"))
}
