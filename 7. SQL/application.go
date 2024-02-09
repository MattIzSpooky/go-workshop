package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"log"
	authors_db "workshop/sql/authors-db"
)

func main() {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "postgres://postgres:postgres@localhost:5432/workshop_db?sslmode=disable")

	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	queries := authors_db.New(conn)

	fmt.Println("Inserting author..")
	// create an author
	insertedAuthor, err := queries.CreateAuthor(ctx, authors_db.CreateAuthorParams{
		Name: "Matthijs Kropholler",
		Bio:  pgtype.Text{String: "Some guy from Valkenswaard", Valid: true},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Done!")
	log.Println(insertedAuthor)

	fmt.Println("Fetching author..")
	// get the author we just inserted
	fetchedAuthor, err := queries.GetAuthor(ctx, insertedAuthor.ID)
	if err != nil {
		panic(err)
	}
	fmt.Println("Done!")
	log.Println(fetchedAuthor)

	fmt.Println("Creating book")
	insertedBook, err := queries.CreateBook(ctx, authors_db.CreateBookParams{
		AuthorID: fetchedAuthor.ID,
		Name:     "Book McBookFace",
		Summary:  pgtype.Text{String: "Once upon a time there was a... snooooozeee", Valid: true},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Done!")
	log.Println(insertedBook)

	fmt.Println("Fetching books by author")
	fetchedBook, err := queries.GetBook(ctx, insertedAuthor.ID)
	if err != nil {
		panic(err)
	}
	fmt.Println("Done!")
	log.Println(fetchedBook)

	fmt.Println("Fetching author books")
	authorBooks, err := queries.ListAuthorBooks(ctx, insertedAuthor.ID)
	if err != nil {
		panic(err)
	}
	fmt.Println("Done!")
	log.Println(authorBooks)

	fmt.Println("Clearing database..")
	queries.ClearAuthorsTable(ctx)
	fmt.Println("Database cleared!")
	fmt.Println("Bye!")
}
