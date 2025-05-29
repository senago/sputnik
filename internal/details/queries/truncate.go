package queries

import (
	"context"
	"fmt"
)

const queryTruncateAll = `-- TruncateAll
truncate table orbit, satellite, satellite_position;
`

func TruncateAll(ctx context.Context) error {
	_, err := resolveConn(ctx).Exec(
		ctx,
		queryTruncateAll,
	)
	if err != nil {
		return fmt.Errorf("exec: %w", err)
	}

	return nil
}
