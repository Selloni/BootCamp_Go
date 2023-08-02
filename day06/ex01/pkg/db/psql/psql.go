package psql

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

type Post struct {
	Id   int
	Text string
}

type Config struct {
	Name, Pass, Host, Port, Database string
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

func PostText(pool *pgxpool.Pool) (buff []Post) {
	rows, err := pool.Query(context.Background(), "SELECT id, post FROM post")
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
	fmt.Println(vars["id"])
	rows, err := pool.Query(context.Background(), fmt.Sprintf("SELECT id, post FROM post WHERE id = %s", vars["id"]))
	//rows, err := pool.Query(context.Background(), fmt.Sprintf("SELECT id, post FROM post WHERE id = 2"))
	if err != nil {
		log.Fatal("fatal get post ", err)
	}
	fmt.Println(rows)
	//defer rows.Close()
	//for rows.Next() {
	//	err := rows.Scan(&pp.Id, &pp.Text)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}
	return
}
