package common

import (
	"context"
	"database/sql"
	"time"

	"github.com/cihub/seelog"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// 慢查询次数
var (
	slowSqlCount = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "slow_sql_count",
			Help: "slow sql count",
		},
		[]string{"method"})

	slowExecSqlCount    = slowSqlCount.With(prometheus.Labels{"method": "exec"})
	slowPrepareSqlCount = slowSqlCount.With(prometheus.Labels{"method": "prepare"})
	slowQuerySqlCount   = slowSqlCount.With(prometheus.Labels{"method": "query"})
)

type DBConnect interface {
	Begin() (*sql.Tx, error)
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

func GetDB2DBConnect(getter func() *sql.DB) func() DBConnect {
	return func() DBConnect {
		return dbConnect{DB: getter()}
	}
}

func Tx2DBConnect(tx *sql.Tx) func() DBConnect {
	return func() DBConnect {
		return txDBConnect{tx}
	}
}

type dbConnect struct {
	*sql.DB
}

func (d dbConnect) Prepare(query string) (*sql.Stmt, error) {
	defer slowLog(3*time.Second, func(cost time.Duration) {
		seelog.Error("Slow prepare:", query, " cost:", cost)
		slowPrepareSqlCount.Add(1)
	})()
	return d.DB.Prepare(query)
}

func (d dbConnect) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	defer slowLog(3*time.Second, func(cost time.Duration) {
		seelog.Error("Slow prepare:", query, " cost:", cost)
		slowPrepareSqlCount.Add(1)
	})()
	return d.DB.Prepare(query)
}

func (d dbConnect) Exec(query string, args ...interface{}) (sql.Result, error) {
	defer slowLog(3*time.Second, func(cost time.Duration) {
		seelog.Error("Slow query:", query, args, " cost:", cost)
		slowExecSqlCount.Add(1)
	})()
	return d.DB.Exec(query, args...)
}
func (d dbConnect) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	defer slowLog(3*time.Second, func(cost time.Duration) {
		seelog.Error("Slow query:", query, args, " cost:", cost)
		slowExecSqlCount.Add(1)
	})()
	return d.DB.ExecContext(ctx, query, args...)
}

func (d dbConnect) Query(query string, args ...interface{}) (*sql.Rows, error) {
	defer slowLog(3*time.Second, func(cost time.Duration) {
		seelog.Error("Slow query:", query, args, " cost:", cost)
		slowQuerySqlCount.Add(1)
	})()
	return d.DB.Query(query, args...)
}

func (d dbConnect) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	defer slowLog(3*time.Second, func(cost time.Duration) {
		seelog.Error("Slow query:", query, args, " cost:", cost)
		slowQuerySqlCount.Add(1)
	})()
	return d.DB.QueryContext(ctx, query, args...)
}

func (d dbConnect) QueryRow(query string, args ...interface{}) *sql.Row {
	defer slowLog(3*time.Second, func(cost time.Duration) {
		seelog.Error("Slow query:", query, args, " cost:", cost)
		slowQuerySqlCount.Add(1)
	})()
	return d.DB.QueryRow(query, args...)
}

func (d dbConnect) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	defer slowLog(3*time.Second, func(cost time.Duration) {
		seelog.Error("Slow query:", query, args, " cost:", cost)
		slowQuerySqlCount.Add(1)
	})()
	return d.DB.QueryRowContext(ctx, query, args...)
}

// txDBConnect wrap sql tx
type txDBConnect struct {
	*sql.Tx
}

func (t txDBConnect) Begin() (*sql.Tx, error) {
	return t.Tx, nil
}

func (t txDBConnect) BeginTx(_ context.Context, _ *sql.TxOptions) (*sql.Tx, error) {
	return t.Tx, nil
}

func (t txDBConnect) Prepare(query string) (*sql.Stmt, error) {
	defer slowLog(3*time.Second, func(cost time.Duration) {
		seelog.Error("Slow prepare:", query, " cost:", cost)
		slowPrepareSqlCount.Add(1)
	})()
	return t.Tx.Prepare(query)
}

func (t txDBConnect) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	defer slowLog(3*time.Second, func(cost time.Duration) {
		seelog.Error("Slow prepare:", query, " cost:", cost)
		slowPrepareSqlCount.Add(1)
	})()
	return t.Tx.Prepare(query)
}

func (t txDBConnect) Exec(query string, args ...interface{}) (sql.Result, error) {
	defer slowLog(3*time.Second, func(cost time.Duration) {
		seelog.Error("Slow query:", query, args, " cost:", cost)
		slowExecSqlCount.Add(1)
	})()
	return t.Tx.Exec(query, args...)
}
func (t txDBConnect) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	defer slowLog(3*time.Second, func(cost time.Duration) {
		seelog.Error("Slow query:", query, args, " cost:", cost)
		slowExecSqlCount.Add(1)
	})()
	return t.Tx.ExecContext(ctx, query, args...)
}

func (t txDBConnect) Query(query string, args ...interface{}) (*sql.Rows, error) {
	defer slowLog(3*time.Second, func(cost time.Duration) {
		seelog.Error("Slow query:", query, args, " cost:", cost)
		slowQuerySqlCount.Add(1)
	})()
	return t.Tx.Query(query, args...)
}

func (t txDBConnect) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	defer slowLog(3*time.Second, func(cost time.Duration) {
		seelog.Error("Slow query:", query, args, " cost:", cost)
		slowQuerySqlCount.Add(1)
	})()
	return t.Tx.QueryContext(ctx, query, args...)
}

func (t txDBConnect) QueryRow(query string, args ...interface{}) *sql.Row {
	defer slowLog(3*time.Second, func(cost time.Duration) {
		seelog.Error("Slow query:", query, args, " cost:", cost)
		slowQuerySqlCount.Add(1)
	})()
	return t.Tx.QueryRow(query, args...)
}

func (t txDBConnect) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	defer slowLog(3*time.Second, func(cost time.Duration) {
		seelog.Error("Slow query:", query, args, " cost:", cost)
		slowQuerySqlCount.Add(1)
	})()
	return t.Tx.QueryRowContext(ctx, query, args...)
}

func slowLog(threshold time.Duration, cb func(cost time.Duration)) func() {
	begin := time.Now()
	return func() {
		cost := time.Now().Sub(begin)
		if cost > threshold {
			cb(cost)
		}
	}
}
