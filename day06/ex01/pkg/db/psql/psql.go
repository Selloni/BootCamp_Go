package psql

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"strconv"
)

type Post struct {
	Id   int
	Text string
}

type Config struct {
	Name, Pass, Host, Port, Database string
}

func ConnectPsql() (*pgxpool.Pool, error) {
	cnf := Config{
		Name:     "grandpat",
		Pass:     "grandpat",
		Host:     "localhost",
		Port:     "5432",
		Database: "postgres",
	}
	psqlClient, err := NewClient(context.Background(), cnf)
	if err != nil {
		return nil, err
	}
	return psqlClient, nil
}

func NewClient(ctx context.Context, con Config) (*pgxpool.Pool, error) {
	var pool *pgxpool.Pool
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", con.Name, con.Pass, con.Host,
		con.Port, con.Database)
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatal(err)
	}
	pool, err = pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		log.Fatal(err)
	}
	return pool, nil
}

func InsertText(pool *pgxpool.Pool, text string) error {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return err
	}
	defer conn.Release()
	_, err = conn.Exec(context.Background(), "INSERT INTO post (post) VALUES ($1)", text)
	if err != nil {
		return err
	}
	return nil
}

func PostText(pool *pgxpool.Pool, str uint) (buff []Post) {
	min := strconv.Itoa(int(str))
	max := strconv.Itoa(int(str + 2))
	rows, err := pool.Query(context.Background(), fmt.Sprintf("SELECT id, post FROM post WHERE id BETWEEN %s AND %s", min, max))
	if err != nil {
		log.Fatal("fatal get text out BD :", err)
	}
	defer rows.Close()
	for rows.Next() {
		var tmp Post
		err := rows.Scan(&tmp.Id, &tmp.Text)
		if err != nil {
			log.Fatal(err)
		}
		buff = append(buff, tmp)
	}
	return
}

func OpenPagePost(pool *pgxpool.Pool, vars map[string]string) (pp Post) {
	rows, err := pool.Query(context.Background(), fmt.Sprintf("SELECT id, post FROM post WHERE id = %s", vars["id"]))
	if err != nil {
		log.Fatal("fatal get post ", err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&pp.Id, &pp.Text)
		if err != nil {
			log.Fatal(err)
		}
	}
	return
}
