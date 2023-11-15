package repositories

import (
  "fmt"
  "gorm.io/gorm"
  "gorm.io/gorm/logger"
  "gorm.io/driver/postgres"
)

type DBConnection interface {
  Connect() (*gorm.DB, error)
  Close(*gorm.DB)
}

type PostgresCon struct {}

func (p *PostgresCon) Connect() (*gorm.DB, error) {
  //TODO: Receive connection data from env vars
  //dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
  host     := "localhost"
  dbuser   := "postgres"
  dbpasswd := "postgres"
  dbname   := "notes"
  dbport   := "5432"
  dns := fmt.Sprintf(
    "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo",
    host, dbuser, dbpasswd, dbname, dbport,
  )
  return gorm.Open(postgres.Open(dns), &gorm.Config{
    Logger: logger.Default.LogMode(logger.Silent),
  })
}

func (p *PostgresCon) Close(db *gorm.DB) {
  sqlDB, err := db.DB()
  if err != nil { panic(err) }
  sqlDB.Close()
}
