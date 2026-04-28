package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func OpenDB(dsn string) (*gorm.DB, error) {
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

// burda db'yi açıyoruz ve gorm.Config ile config'i ayarlıyoruz.
// gorma biz mysql kullanuyoruz diyoruz.
// gorm.Open gerçek bağlantıyı kurar ve sana bir DB instance verir
// *gorm.DB bu ise bize query,insert ve update yapmak için kullanıyoruz.

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&User{},
		&Movie{},
		&UserMovie{},
		&Quote{},
	)
}

//Migrate ile de  structurelere bakarak tabloyu oluşturur.
//AutoMİgrate tablo yoksa tablo, kolon yoksa kolon oluşturur.
//buna safe migration denir.
//Uygulama her açıldığında database otomatik hazır hale gelir
