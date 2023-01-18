package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Certificado struct {
	ID             uint64    `gorm:"primary_key;auto_increment" json:"id"`
	EventoID       uint64    `gorm:"not null" json:"evento_id"`
	Evento         Evento    `json:"evento"`
	ParticipanteID uint64    `gorm:"not null" json:"participante_id"`
	Participante   Usuario   `json:"participante"`
	Validacao      string    `json:"validacao"`
	CreatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (p *Certificado) Validate() error {

	if p.EventoID < 1 {
		return errors.New("Obrigatório - Evento")
	}
	if p.ParticipanteID < 1 {
		return errors.New("Obrigatório - Apresentador")
	}
	return nil
}

func (p *Certificado) Create(db *gorm.DB) (int64, error) {
	if verr := p.Validate(); verr != nil {
		return -1, verr
	}
	err := db.Debug().Model(&Certificado{}).Create(&p).Error
	if err != nil {
		return 0, err
	}
	return int64(p.ID), nil
}

func (p *Certificado) ListAll(db *gorm.DB) (*[]Certificado, error) {
	var err error
	Certificados := []Certificado{}
	err = db.Debug().Model(&Certificado{}).Limit(100).Find(&Certificados).Error
	if err != nil {
		return &[]Certificado{}, err
	}
	return &Certificados, nil
}

func (p *Certificado) Find(db *gorm.DB, pid uint64) (*Certificado, error) {
	err := db.Debug().Model(&Certificado{}).Where("id = ?", pid).Take(&p).Error
	if err != nil {
		return &Certificado{}, err
	}
	return p, nil
}

func (p *Certificado) Update(db *gorm.DB) (*Certificado, error) {
	err := db.Debug().Model(&Certificado{}).Where("id = ?", p.ID).Take(&Certificado{}).UpdateColumns(
		map[string]interface{}{
			"EventoID":       p.EventoID,
			"Evento":         p.Evento,
			"ParticipanteID": p.ParticipanteID,
			"Participante":   p.Participante,
			"UpdatedAt":      time.Now()}).Error

	if err != nil {
		return &Certificado{}, err
	}
	return p, nil
}

func (p *Certificado) Delete(db *gorm.DB, pid uint64, uid uint64) (int64, error) {

	db = db.Debug().Model(&Certificado{}).Where("id = ? and apresentador_id = ?", pid, uid).Take(&Certificado{}).Delete(&Certificado{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
