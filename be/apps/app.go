package apps

import (
	mysql_client "github.com/RandySteven/Idolized/pkg/mysql"
	redis_client "github.com/RandySteven/Idolized/pkg/redis"
)

type (
	App struct {
		MySQL mysql_client.MySQL
		Redis redis_client.Redis
	}
)
