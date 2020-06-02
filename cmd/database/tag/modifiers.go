package tag

import (
	"checkaem_server/cmd/database"
	"checkaem_server/cmd/entities/tag"
	"context"
)

func Insert(t *tag.Tag) (*tag.Tag, error) {
	err := database.Connection.QueryRow(context.Background(), insertQuery, t.Name).Scan(&t.Name, &t.Count)

	if err != nil {
		return nil, err
	}

	return t, nil
}
