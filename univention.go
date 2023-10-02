package itswizard_m_uploadrest

import "github.com/jinzhu/gorm"

type UniventionRestSetup struct {
	gorm.Model
	FirstRun       bool   `gorm:"-" json:"first_run"`
	Uuid           string `gorm:"unique" json:"uuid"`
	OrganisationID string `gorm:"unique" json:"-"`
	InstitutionID  string `json:"-"`
	AESkey         []byte `gorm:"VARCHAR 500" json:"aeskey"`
	AuthKey        string `gorm:"VARCHAR 500" json:"auth_key"`
}

func GetUniventionRestSetup(uuid string, db *gorm.DB) (setup UniventionRestSetup, err error) {
	err = db.Where("uuid = ?", uuid).Last(&setup).Error
	return setup, err
}

func (p *UniventionRestSetup) Initial(AESKey []byte) *UniventionRestSetup {
	p.AESkey = AESKey
	return p
}

func (p *UniventionRestSetup) Save(db *gorm.DB) error {
	err := db.Save(p).Error
	return err
}
