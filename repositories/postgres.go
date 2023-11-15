package repositories

import (
  "fmt"
  "gorm.io/gorm"
  "gorm.io/gorm/logger"
  "gorm.io/driver/postgres"
)

type PostgresCon struct {
  DB *gorm.DB
}

func (p *PostgresCon) Connect() (error) {
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
  db, err := gorm.Open(postgres.Open(dns), &gorm.Config{
    Logger: logger.Default.LogMode(logger.Silent),
  })
  p.DB = db
  return err
}

func (p *PostgresCon) Close() {
  sqlDB, err := p.DB.DB()
  if err != nil { panic(err) }
  sqlDB.Close()
}
