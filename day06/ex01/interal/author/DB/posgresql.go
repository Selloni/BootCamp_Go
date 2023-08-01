//package DB
//
////import (
////	"context"
////	"log"
////	"web/ex01/interal/author"
////	"web/ex01/pkg/db/psql"
////)
////
////type repository struct {
////	client psql.Client
////}
//
////func formalQuery(q string) string {
////	return strings.ReplaceAll(strings.ReplaceAll(q, "\t", ""), "\n", " ")
////}
//
////func (r repository) Create(ctx context.Context, author *author.Author) error {
////	q := `insert into post (post) values ($1)`
////	if err := r.client.QueryRow(ctx, q, author.Text); err != nil {
////		log.Fatal("dont creat SQL")
////	}
////	return nil
////}
////
////func (r repository) FindAll(ctx context.Context) (u []author.Author, err error) {
////	q := `select id, post from post`
////	query, err := r.client.Query(ctx, q)
////	if err != nil {
////		return nil, err
////	}
////	text := make([]author.Author, 0)
////	for query.Next() {
////		var ath author.Author
////		err = query.Scan(&ath.ID, &ath.Text)
////		if err != nil {
////			return nil, err
////		}
////		text = append(text, ath)
////	}
////	return text, nil
////}
//
////func (r repository) FindOne(ctx context.Context, id string) (author.Author, error) {
////	//TODO implement me
////	panic("implement me")
////}
//
////func (r repository) Update(ctx context.Context, user author.Author) error {
////	//TODO implement me
////	panic("implement me")
////}
////
////func (r repository) Delete(ctx context.Context, id string) error {
////	//TODO implement me
////	panic("implement me")
////}
//
//func NewRepository(client psql.Client) *repository {
//	return &repository{
//		client: client,
//	}
//}
