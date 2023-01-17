package services

import (
	"log"

	"github.com/jmscatena/ecert-back-go/database"
	"github.com/jmscatena/ecert-back-go/interfaces"
)

func New[T interfaces.Tables](o interfaces.PersistenceHandler[T]) (uint64, error) {
	db, err := database.Init()
	if err != nil {
		log.Fatalln(err)
		return 0, err
	}
	recid, err := o.Create(db)
	if err != nil {
		log.Fatalln(err)
		return 0, err
	}
	if recid != 0 {
		return recid, nil
	}
	return 0, nil
}
