package queries

import (
	"context"
	"fmt"

	"github.com/senago/sputnik/internal/details/db"
)

func resolveConn(ctx context.Context) db.Executor {
	if conn, ok := db.GetConnContext(ctx); ok {
		return conn.Resolver(ctx)
	}

	panic(fmt.Errorf("insufficient context to resolve conn"))
}
