package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Pool struct {
	pools []*pgxpool.Pool
}

func NewPool() *Pool {
	return &Pool{}
}

func (c *Pool) GetConn() *pgxpool.Pool {
	conn := c.pools[0]
	return conn
}

func (c *Pool) Connect(ctx context.Context, strConf string) error {
	var err error
	var pool *pgxpool.Pool
	cfg, err := pgxpool.ParseConfig(strConf)
	if err != nil {
		return err
	}
	if pool, err = pgxpool.NewWithConfig(ctx, cfg); err != nil {
		return err
	}
	for i := 0; i < 10; i++ {
		if err = pool.Ping(ctx); err == nil {
			c.pools = append(c.pools, pool)
			return nil
		}
		time.Sleep(time.Second * 3)
	}
	return err
}
