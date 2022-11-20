package dbase

import (
	"fmt"
	"go_cineparse/pkg/models"

	"gorm.io/gorm"
)

type Service struct {
	dbSqlite *gorm.DB
}

func NewService(dbSqlite *gorm.DB) *Service {
	return &Service{dbSqlite: dbSqlite}
}

/*
Информация

https://gorm.io/docs/models.html#gorm-Model
https://gorm.io/docs/transactions.html#Control-the-transaction-manually
https://stackoverflow.com/questions/23669720/meaning-of-interface-dot-dot-dot-interface

gorm youtube https://www.youtube.com/watch?v=9koLNdEcSR0
*/

// -----------------------------
func (s *Service) Echo() {
	fmt.Println("dbase service")
}

func (s *Service) Init() {
	s.dbSqlite.AutoMigrate(&models.CineRow{}, &models.RandomTable{})
}

func (s *Service) CheckRowInDB(paramRow *models.CineRow) bool {
	var cineRows []*models.CineRow
	s.dbSqlite.Where(&models.CineRow{
		AuthorTitle: paramRow.AuthorTitle,
		Album:       paramRow.Album,
		Duration:    paramRow.Duration,
	}).Find(&cineRows)

	if len(cineRows) == 0 {
		fmt.Println(len(cineRows), " records not found:")
		return false
	} else {
		fmt.Println("found records:")
		for _, v := range cineRows {
			fmt.Printf("%+v\n", v)
		}
		return true
	}
}

func (s *Service) InsertCineRow(paramRow *models.CineRow) {
	s.dbSqlite.Create(&models.CineRow{
		AuthorTitle: paramRow.AuthorTitle,
		Album:       paramRow.Album,
		Duration:    paramRow.Duration,
	})
	fmt.Println("cineRow inserted")
}

// ARCHIVE ------------------------------------------

func (s *Service) doSomething() {
	// insert
	if 1 == 2 {
		s.dbSqlite.Create(&models.CineRow{
			AuthorTitle: "Harald Kloser - The Day After Tomorrow",
			Album:       "The Day After Tomorrow",
			Duration:    "3:27",
		})
		s.dbSqlite.Create(&models.RandomTable{Name: "booo"})
	}

	// проверка есть ли уже записи
	if 1 == 2 {
		var cineRow []*models.CineRow
		//db.Where("author_title = ?", 1).Find(&cineRow)
		s.dbSqlite.Where(&models.CineRow{AuthorTitle: "12"}).Find(&cineRow)
		fmt.Println("cnt_rows=", len(cineRow))

		for _, v := range cineRow {
			fmt.Printf("%+v\n", v)
		}
	}
}
