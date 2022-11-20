package main

import (
	"fmt"
	cine "go_cineparse/pkg/cineparser"
	"go_cineparse/pkg/dbase"
	"log"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

/*
- парсим cinemis.us (https://vhost.fastserv.com/~cinemix/samPHPweb/web/playing.php?buster=06090525471)
- cохраняем результаты (например) в sqlite
- с базой работаем через GORM (тренируемся)
*/

const (
	sqlitePath      = "./cinemix.db"
	minutesInterval = 5
)

func main() {
	dbSqlite, err := gorm.Open(sqlite.Open(sqlitePath), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// new database service
	dbase1 := dbase.NewService(dbSqlite)
	dbase1.Init()

	// new cineparser service
	cine1 := cine.NewService("123")

	// prepare loop
	d := time.NewTicker(minutesInterval * time.Minute)
	myChan := make(chan bool)

	for {
		select {
		case <-myChan:
			fmt.Println("Completed!")
			return
		case tm := <-d.C:
			fmt.Println("-->", tm)
			// -------------------------------
			// cine - parse data
			rows, err := cine1.GetDataShit()
			if err != nil {
				log.Println(err)
			}
			for _, v := range rows {
				fmt.Printf("parsed: %+v\n", v)
				flgFound := dbase1.CheckRowInDB(v)
				fmt.Printf("result is %t\n", flgFound)
				if !flgFound {
					dbase1.InsertCineRow(v)
				}
				fmt.Println("-----------------------------")
			}
			// -------------------------------

		}
	}

}
