package mysql_client

import (
	"context"
	"github.com/RandySteven/Idolized/queries"
)

func initMigrations() []queries.MigrationQuery {
	return []queries.MigrationQuery{
		queries.CreateAccountTable,
		queries.CreateRoleTable,
		queries.CreateUserTable,
		queries.CreateGroupTable,
		queries.CreateTalentTable,
		queries.CreateGroupTalentTable,
	}
}

func (m *mysqlClient) Migration(ctx context.Context) error {
	migrations := initMigrations()

	for _, query := range migrations {
		_, err := m.db.ExecContext(ctx, query.String())
		if err != nil {
			return err
		}
	}
	return nil
}
