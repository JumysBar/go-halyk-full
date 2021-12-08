package postgres

import (
	"context"
	"example_project/api/repository"
	"example_project/domain"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

type postgresDbInterfase struct {
	db *pgxpool.Pool
}

func (p *postgresDbInterfase) GetUser(name string) (*domain.User, error) {
	resultUser := &domain.User{}

	if err := p.db.QueryRow(context.Background(), "SELECT ID, Name, Age FROM users WHERE Name=$1", name).Scan(&resultUser.ID, &resultUser.Name, &resultUser.Age); err != nil {
		return nil, err
	}
	return resultUser, nil
}

func (p *postgresDbInterfase) UpdateUserAge(name string, newAge int64) error {
	if _, err := p.db.Exec(context.Background(), "UPDATE users SET age=$1 WHERE Name=$2", newAge, name); err != nil {
		return err
	}
	return nil
}

func NewPostgresDbInterface(host string, port int, db string) (repository.DBInterface, error) {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		"postgres", "postgres", host, port, db,
	)

	config, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		return nil, err
	}
	config.MaxConns = 25
	config.MaxConnLifetime = 5 * time.Minute
	config.MaxConnIdleTime = 2 * time.Minute
	// явно оборачивать каждый запрос в транзакцию
	config.ConnConfig.PreferSimpleProtocol = true
	// экранирование каждого из параметров
	config.ConnConfig.RuntimeParams = map[string]string{
		"standard_conforming_strings": "on",
	}

	_db, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	return &postgresDbInterfase{
		db: _db,
	}, nil
}
