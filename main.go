package main

import (
	"context"
	"fmt"
	"log"
	"postgres/data"
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

	/*Добавление в бд.
	newBook := book.NewBook("23456", "1234", "норм книга", true, time.Now())
	fmt.Println(newBook.Title)
	if err := data.InsertTable(ctx, conn, newBook); err != nil {
		log.Fatal(err)
	}*/

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

	var ids []int
	ids = append(ids, 1)
	ids = append(ids, 2)
	ids = append(ids, 3)

	if err := data.DeleteByIds(ctx, conn, ids); err != nil {
		log.Fatal(err)
	}

	fmt.Println("OK")
}
