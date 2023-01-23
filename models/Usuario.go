package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Usuario struct {
	ID           uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Nome         string    `gorm:"size:255;not null;unique" json:"nome"`
	Email        string    `gorm:"size:100;not null;unique" json:"email"`
	Senha        string    `gorm:"size:100;not null;"`
	Apresentador bool      `gorm:"default:False;not null;" json:"apresentador"`
	Ativo        bool      `gorm:"default:True;not null;" json:"ativo"`
	Admin        bool      `gorm:"default:False;not null;"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"criado_em"`
	UpdatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"atualizado_em"`
}

func (u *Usuario) Create(db *gorm.DB) (int64, error) {
	if verr := u.Validate("insert"); verr != nil {
		return -2, verr
	}
	u.Prepare()
	err := db.Debug().Create(&u).Error
	if err != nil {
		return 0, err
	}
	return int64(u.ID), nil
}

func (u *Usuario) Update(db *gorm.DB, uid uint64) (*Usuario, error) {

	if verr := u.Validate("insert"); verr != nil {
		return nil, verr
	}
	u.Prepare()

	db = db.Debug().Model(&Usuario{}).Where("id = ?", uid).Take(&Usuario{}).UpdateColumns(
		map[string]interface{}{
			"Senha":        u.Senha,
			"Nome":         u.Nome,
			"Email":        u.Email,
			"Apresentador": u.Apresentador,
			"updated_at":   time.Now(),
		},
	)
	if db.Error != nil {
		return &Usuario{}, db.Error
	}
	err := db.Debug().Model(&Usuario{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &Usuario{}, err
	}
	return u, nil
}

func (u *Usuario) List(db *gorm.DB) (*[]Usuario, error) {
	Usuarios := []Usuario{}
	err := db.Debug().Model(&Usuario{}).Limit(100).Find(&Usuarios).Error
	if err != nil {
		return nil, err
	}
	return &Usuarios, err
}

func (u *Usuario) Find(db *gorm.DB, uid uint64) (*Usuario, error) {
	err := db.Debug().Model(Usuario{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &Usuario{}, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &Usuario{}, errors.New("Usuario Inexistente")
	}
	return u, err
}

func (u *Usuario) FindBy(db *gorm.DB, param string, uid ...interface{}) (*[]Usuario, error) {
	Usuarios := []Usuario{}
	params := strings.Split(param, ";")
	uids := uid[0].([]interface{})
	if len(params) != len(uids) {
		return nil, errors.New("condição inválida")
	}
	result := db.Where(strings.Join(params, " AND "), uids...).Find(&Usuarios)
	if result.Error != nil {
		return nil, result.Error
	}
	return &Usuarios, nil
}

func (u *Usuario) Delete(db *gorm.DB, uid uint64) (int64, error) {

	db = db.Debug().Model(&Usuario{}).Where("id = ?", uid).Take(&Usuario{}).Delete(&Usuario{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

func (u *Usuario) DeleteBy(db *gorm.DB, cond string, uid interface{}) (int64, error) {
	result := db.Delete(&Usuario{}, cond+" = ?", uid)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

func (u *Usuario) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Nome == "" {
			return errors.New("obrigatório: nome")
		}
		if u.Senha == "" {
			return errors.New("obrigatório: senha")
		}
		if u.Email == "" {
			return errors.New("obrigatório: email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("inválido: email")
		}
		return nil
	case "login":
		if u.Senha == "" {
			return errors.New("obrigatório: senha")
		}
		if u.Email == "" {
			return errors.New("obrigatório: email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("inválido: email")
		}
		return nil
	default:
		if u.Nome == "" {
			return errors.New("obrigatório: nome")
		}
		if u.Senha == "" {
			return errors.New("obrigatório: senha")
		}
		if u.Email == "" {
			return errors.New("obrigatório: email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("inválido: email")
		}
		return nil
	}
}

func Hash(Senha string) []byte {
	hash, _ := bcrypt.GenerateFromPassword([]byte(Senha), bcrypt.DefaultCost)
	return hash
}

func VerifySenha(hashedSenha, Senha string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedSenha), []byte(Senha))
}

func (u *Usuario) Prepare() {
	u.Nome = html.EscapeString(strings.TrimSpace(u.Nome))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.Senha = string(Hash(u.Senha))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}
