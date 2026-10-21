package main

import (
	"context"
	"fmt"
	"log"
	"postgres/data"

	"github.com/jackc/pgx/v5"
	"github.com/k0kubun/pp"
)

func main() {
	ctx := context.Background()
	conn, err := data.CheckConnection(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if err := data.CreateTable(ctx, conn); err != nil {
		log.Fatal(err)
	}
	/*
		for i := 0; i < 50; i++ {
			newBook := book.NewBook(fmt.Sprintf("Book %d", i), "1234", "норм книга", true, time.Now())
			if err := data.InsertTable(ctx, conn, newBook); err != nil {
				log.Fatal(err)
			}
		}
	*/

	/*Получение данных из таблицы.
	books, err := data.GetTable(ctx, conn)

	if err != nil {
		log.Fatal(err)
	}
	for _, book := range books {
		fmt.Println(book)
	}
	*/
	/* Обновление.
	newBook := book.NewBook("Бес", "1234", "норм книга", true, time.Now())

	if err := data.UpdateByIdTable(ctx, conn, 1, newBook); err != nil {
		log.Fatal(err)
	}
	*/

	/*Удаление.
	var ids []int
	ids = append(ids, 1)
	ids = append(ids, 2)
	ids = append(ids, 3)

	if err := data.DeleteByIds(ctx, conn, ids); err != nil {
		log.Fatal(err)
	}
	*/

	/*
		books, err := data.SelectLimitTable(ctx, conn)
		if err != nil {
			log.Fatal(err)
		}
		pp.Println(books)
	*/

	count, err := data.GetCountRow(ctx, conn)
	if err != nil {
		log.Fatal(err)
	}
	if err := ListPages(5, count, ctx, conn); err != nil {
		log.Fatal(err)
	}

	fmt.Println("OK")
}

func ListPages(n, count int, ctx context.Context, conn *pgx.Conn) error {
	limit := n
	offset := 0
	page := 1
	for offset < count {
		fmt.Println("///////////////////////////page", page)
		books, err := data.SelectLimitTable(ctx, conn, limit, offset)
		if err != nil {
			return err
		}
		pp.Println(books)
		fmt.Println("########################")
		offset += limit
		page++
	}
	return nil

}
