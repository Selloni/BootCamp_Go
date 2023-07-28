package psql

import (
	"context"
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

//
//func OpenSql(text string) {
//	connStr := "user=q password=q dbname=mypost sslmode=disable"
//	db, err := sql.Open("postgres", connStr)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer db.Close()
//
//	result, err := db.Exec("insert into ", text)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Println(result.RowsAffected())
//}
//

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
	BeginTxFunc(ctx context.Context, txOptions pgx.TxOptions, f func(pgx.Tx) error) error
}

func NewClient(ctx context.Context, name, pass, host, port, database string) {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", name, pass, host, port, database)
}
