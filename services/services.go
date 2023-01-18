package services

import (
	"log"

	"github.com/jmscatena/ecert-back-go/database"
	"github.com/jmscatena/ecert-back-go/interfaces"
)

func New[T interfaces.Tables](o interfaces.PersistenceHandler[T]) (int64, error) {
	db, err := database.Init()
	if err != nil {
		log.Fatalln(err)
		return -1, err
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

func Update[T interfaces.Tables](o interfaces.PersistenceHandler[T], uid uint64) (*T, error) {
	db, err := database.Init()
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	rec, err := o.Update(db, uid)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return rec, nil
}

func Get[T interfaces.Tables](o interfaces.PersistenceHandler[T], uid uint64) (*T, error) {
	db, err := database.Init()
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	rec, err := o.Find(db, uid)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return rec, nil
}

func GetAll[T interfaces.Tables](o interfaces.PersistenceHandler[T]) (*[]T, error) {
	db, err := database.Init()
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	var rec *[]T
	rec, err = o.List(db)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return rec, nil
}
