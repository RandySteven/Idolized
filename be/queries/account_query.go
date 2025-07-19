package queries

const (
	InsertAccount GoQuery = `
		INSERT INTO accounts (account_id, parent_account_id, account_role)
		VALUES 
		    (?, ?, ?)
	`

	SelectAccountByAccountID GoQuery = `
		SELECT id, account_id, parent_account_id, account_role, is_verified, verified_time, created_at, updated_at, deleted_at
		FROM
		    accounts
		WHERE 
		    account_id = ?
	`
)
