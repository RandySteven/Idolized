package queries

const (
	CreateAccountTable MigrationQuery = `
		CREATE TABLE IF NOT EXISTS accounts (
		    id BIGNUMBER PRIMARY KEY AUTO_INCREMENT NOT NULL,
		    account_id VARCHAR(16) UNIQUE NOT NULL,
		    parent_account_id VARCHAR(16) NULLABLE,
		    account_role VARCHAR(16) NOT NULL,
		    is_verified BOOLEAN DEFAULT FALSE,
		    verified_time TIMESTAMP DEFAULT NULL,
		    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		    deleted_at TIMESTAMP DEFAULT NULL
		)
	`

	CreateUserTable MigrationQuery = `
		CREATE TABLE IF NOT EXISTS users (
		    id BIGNUMBER PRIMARY KEY AUTO_INCREMENT NOT NULL,
		    account_id VARCHAR(16) NOT NULL,
		    email VARCHAR(244) UNIQUE NOT NULL,
		    password VARCHAR(244) NOT NULL,
		    full_name VARCHAR(64) NOT NULL,
		    user_name VARCHAR(64) NOT NULL UNIQUE,
		    date_of_birth CHAR(10) NOT NULL,
		    status VARCHAR(16) NOT NULL,
		    gender VARCHAR(6) NOT NULL,
		    profile_picture VARCHAR(244) NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		    deleted_at TIMESTAMP DEFAULT NULL,
		    FOREIGN KEY (account_id) REFERENCES accounts (account_id) ON DELETE CASCADE
		)
	`

	CreateTalentTable MigrationQuery = `
		CREATE TABLE IF NOT EXISTS talents (
		    id BIGNUMBER PRIMARY KEY AUTO_INCREMENT NOT NULL,
		    account_id VARCHAR(16) NOT NULL,
		    email VARCHAR(244) UNIQUE NOT NULL,
		    password VARCHAR(244) NOT NULL,
		    stage_name VARCHAR(64) NOT NULL UNIQUE,
		    talent_photo_url VARCHAR(244) NOT NULL,
		    relation_status VARCHAR(16) NOT NULL,
		    date_of_birth CHAR(6) NOT NULL,
		    hide_date_of_birth BOOLEAN NOT NULL,
		    is_verified BOOLEAN DEFAULT FALSE,
		    verified_time TIMESTAMP DEFAULT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		    deleted_at TIMESTAMP DEFAULT NULL,
		    FOREIGN KEY (account_id) REFERENCES accounts (account_id) ON DELETE CASCADE
		)
	`

	CreateGroupTable MigrationQuery = `
		CREATE TABLE IF NOT EXISTS groups (
		    id BIGNUMBER PRIMARY KEY AUTO_INCREMENT NOT NULL,
		    account_id VARCHAR(16) NOT NULL,
		    name VARCHAR(64) NOT NULL UNIQUE,
		    email VARCHAR(244) UNIQUE NOT NULL,
		    password VARCHAR(244) NOT NULL,
		    description VARCHAR(244) NOT NULL,
		    picture_url VARCHAR(244) NOT NULL,
		    is_verified BOOLEAN DEFAULT FALSE,
		    verified_time TIMESTAMP DEFAULT NULL,
		    verified_by VARCHAR(16) NOT NULL,
		    debut_date DATE NOT NULL,
		    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		    deleted_at TIMESTAMP DEFAULT NULL,
		    FOREIGN KEY (account_id) REFERENCES accounts (account_id) ON DELETE CASCADE	
		)
	`

	CreateGroupTalentTable MigrationQuery = `
		CREATE TABLE IF NOT EXISTS group_talents (
		    id BIGNUMBER PRIMARY KEY AUTO_INCREMENT NOT NULL,
		    group_account_id VARCHAR(16) NOT NULL,
		    talent_account_id VARCHAR(16) NOT NULL,
		    is_verified BOOLEAN DEFAULT FALSE,
		    verified_time TIMESTAMP DEFAULT NULL,
		    member_status VARCHAR(16) NOT NULL,
		    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		    deleted_at TIMESTAMP DEFAULT NULL,
		    FOREIGN KEY (group_account_id) REFERENCES accounts (account_id) ON DELETE CASCADE,
		    FOREIGN KEY (talent_account_id) REFERENCES accounts (account_id) ON DELETE CASCADE
		)
	`

	CreateRoleTable MigrationQuery = `
		CREATE TABLE IF NOT EXISTS roles (
		    id BIGNUMBER PRIMARY KEY AUTO_INCREMENT NOT NULL,
		    role VARCHAR(24) NOT NULL,
		    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		    deleted_at TIMESTAMP DEFAULT NULL
		)
	`
)
