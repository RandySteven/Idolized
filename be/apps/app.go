package apps

import (
	"github.com/RandySteven/Idolized/configs"
	mysql_client "github.com/RandySteven/Idolized/pkg/mysql"
	nsq_client "github.com/RandySteven/Idolized/pkg/nsq"
	redis_client "github.com/RandySteven/Idolized/pkg/redis"
)

type (
	App struct {
		MySQL mysql_client.MySQL
		Redis redis_client.Redis
		NSQ   nsq_client.Nsq
	}
)

func NewApps(config *configs.Config) (*App, error) {
	mysql, err := mysql_client.NewMySQLClient(config)
	if err != nil {
		return nil, err
	}

	redis, err := redis_client.NewRedisClient(config)
	if err != nil {
		return nil, err
	}

	nsq, err := nsq_client.NewNsqClient(config)
	if err != nil {
		return nil, err
	}

	return &App{
		MySQL: mysql,
		Redis: redis,
		NSQ:   nsq,
	}, nil
}
