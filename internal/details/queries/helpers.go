package queries

import (
	"github.com/cockroachdb/errors"
	"github.com/jackc/pgx/v5"
	"github.com/lib/pq"
)

func nest[T any, U any](values []T, selector func(T) U) any {
	n := make([]U, 0, len(values))

	for _, value := range values {
		n = append(n, selector(value))
	}

	return pq.Array(n)
}

func scanRows[T any](rows pgx.Rows, scanner func(row pgx.Row, item *T) error) ([]T, error) {
	defer rows.Close()

	items := make([]T, 0, rows.CommandTag().RowsAffected())

	for rows.Next() {
		var item T
		if err := scanner(rows, &item); err != nil {
			return nil, errors.Wrap(err, "scanner")
		}
		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "rows.Err")
	}

	return items, nil
}
