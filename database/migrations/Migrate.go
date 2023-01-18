package migrations

import (
	"github.com/jmscatena/ecert-back-go/models"
	"gorm.io/gorm"
)

func RunMigrate(db *gorm.DB) {

	db.AutoMigrate(&models.Certificado{})
	db.AutoMigrate(&models.Evento{})
	db.AutoMigrate(&models.Usuario{})
	db.AutoMigrate(&models.Instituicao{})
	db.AutoMigrate(&models.CertVal{})
}
