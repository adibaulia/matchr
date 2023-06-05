package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

// unexported struct will be ignored.
type test struct {
	id  int
	Xxx string
	Ttt int
}

// generate code
func main() {

	cfg := gen.Config{
		OutPath: "./domain/generated",
		//Mode:          gen.WithDefaultQuery | gen.WithQueryInterface,
		Mode:             gen.WithDefaultQuery,
		ModelPkgPath:     "domain/generated",
		FieldWithTypeTag: true,
		FieldNullable:    true,
	}

	g := gen.NewGenerator(cfg)

	dsn := "host=localhost user=postgres password=mysecretpassword dbname=matchr port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	g.UseDB(db)
	g.GenerateAllTable()
	g.Execute()
}
