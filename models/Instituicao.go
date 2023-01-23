package models

import (
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"gorm.io/gorm"
)

type Instituicao struct {
	ID          uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Nome        string    `gorm:"size:155;not null;unique" json:"nome"`
	Responsavel string    `gorm:"size:50;not null" json:"responsavel"`
	Email       string    `gorm:"size:100;not null;unique" json:"email"`
	Tel         string    `gorm:"size:16;not null;unique" json:"telefone"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"criado_em"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"atualizado_em"`
}

func (i *Instituicao) Create(db *gorm.DB) (int64, error) {
	var verr error
	if verr = i.Validate("insert"); verr != nil {
		return -1, verr
	}
	err := db.Debug().Create(&i).Error
	if err != nil {
		return 0, err
	}
	return int64(i.ID), nil
}

func (i *Instituicao) Update(db *gorm.DB, uuid uint64) (*Instituicao, error) {
	db = db.Debug().Model(&Instituicao{}).Where("id=?", uuid).Take(&Instituicao{}).UpdateColumns(
		map[string]interface{}{
			"Nome":        i.Nome,
			"Responsavel": i.Responsavel,
			"Email":       i.Email,
			"Tel":         i.Tel,
			"UpdatedAt":   time.Now(),
		},
	)
	if db.Error != nil {
		return &Instituicao{}, db.Error
	}
	err := db.Debug().Model(&Instituicao{}).Where("id=?", uuid).Take(&Instituicao{}).Error
	if err != nil {
		return &Instituicao{}, err
	}
	return i, nil
}

func (i *Instituicao) List(db *gorm.DB) (*[]Instituicao, error) {
	Instituicoes := []Instituicao{}
	err := db.Debug().Model(&Instituicao{}).Limit(100).Find(&Instituicoes).Error
	if err != nil {
		return &[]Instituicao{}, err
	}
	return &Instituicoes, err
}

func (i *Instituicao) Find(db *gorm.DB, uuid uint64) (*Instituicao, error) {
	err := db.First(&i, uuid).Error
	if err != nil {
		return &Instituicao{}, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &Instituicao{}, errors.New("usuário: inexistente")
	}
	return i, err
}

func (i *Instituicao) FindBy(db *gorm.DB, param string, uid ...interface{}) (*[]Instituicao, error) {
	Instituicoes := []Instituicao{}
	params := strings.Split(param, ";")
	uids := uid[0].([]interface{})
	if len(params) != len(uids) {
		return nil, errors.New("condição inválida")
	}
	result := db.Where(strings.Join(params, " AND "), uids...).Find(&Instituicoes)
	if result.Error != nil {
		return nil, result.Error
	}
	return &Instituicoes, nil
}

func (i *Instituicao) Delete(db *gorm.DB, uuid uint64) (int64, error) {
	db = db.Debug().Model(&Instituicao{}).Where("id=?", uuid).Take(&Instituicao{}).Delete(&Instituicao{})
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

func (i *Instituicao) DeleteBy(db *gorm.DB, cond string, uid uint64) (int64, error) {
	result := db.Delete(&Instituicao{}, cond+" = ?", uid)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

func (i *Instituicao) Validate(action string) error {
	if i.Nome == "" {
		return errors.New("obrigatório: nome")
	}
	if i.Responsavel == "" {
		return errors.New("obrigatório: responsável")
	}
	if i.Email == "" {
		return errors.New("obrigatório: email")
	}
	if i.Tel == "" {
		return errors.New("obrigatório: telefone")
	}
	if err := checkmail.ValidateFormat(i.Email); err != nil {
		return errors.New("inválido: email")
	}
	return nil
}
