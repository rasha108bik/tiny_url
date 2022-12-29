package postgres

import (
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// const (
// 	defaultMaxPoolSize  = 1
// 	defaultConnAttempts = 2
// 	defaultConnTimeout  = time.Second
// )

// type Postgres struct {
// 	maxPoolSize  int
// 	connAttempts int
// 	connTimeout  time.Duration

// 	Pool *pgxpool.Pool
// }

// func New(ctx context.Context, url string, opts ...Options) (*Postgres, error) {
// 	pg := &Postgres{
// 		maxPoolSize:  defaultConnAttempts,
// 		connAttempts: defaultConnAttempts,
// 		connTimeout:  defaultConnTimeout,
// 	}

// 	for _, opt := range opts {
// 		opt(pg)
// 	}

// 	poolConfig, err := pgxpool.ParseConfig(url)
// 	if err != nil {
// 		return nil, fmt.Errorf("postgres - NewPostgres - pgxpool.ParseConfig: %w", err)
// 	}

// 	poolConfig.MaxConns = int32(pg.maxPoolSize)

// 	for pg.connAttempts > 0 {
// 		pg.Pool, err = pgxpool.ConnectConfig(ctx, poolConfig)
// 		if err == nil {
// 			break
// 		}

// 		log.Printf("Postgres is trying to connect, attempts left: %d", pg.connAttempts)

// 		time.Sleep(pg.connTimeout)
// 		pg.connAttempts--
// 	}

// 	if err != nil {
// 		return nil, fmt.Errorf("postgres - NewPostgres - connAttempts == 0: %w", err)
// 	}

// 	return pg, nil
// }

// func (p *Postgres) Close() {
// 	if p.Pool != nil {
// 		p.Pool.Close()
// 	}
// }

type Postgres struct {
	Postgres *sql.DB
}

func New(dns string) (*Postgres, error) {
	db, err := sql.Open("postgres", dns)
	if err != nil {
		return nil, err
	}

	return &Postgres{
		Postgres: db,
	}, nil
}

func (p *Postgres) Close() {
	p.Postgres.Close()
}