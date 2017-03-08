package main

import (
	"database/sql"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

//Company tabla con dos relaciones
type Company struct {
	ID        int        `sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Name      string     `sql:"size:255;unique;index"`
	Employees []Employee // one-to-many relationship
	Address   Address    // one-to-one relationship
}

//Employee tabla que almacena info de empleados
type Employee struct {
	FirstName        string    `sql:"size:255;index:name_idx"`
	LastName         string    `sql:"size:255;index:name_idx"`
	SocialSecurityNo string    `sql:"type:varchar(100);unique" gorm:"column:ssn"`
	DateOfBirth      time.Time `sql:"DEFAULT:current_timestamp"`
	Address          *Address  // one-to-one relationship
	Deleted          bool      `sql:"DEFAULT:false"`
}

//Address tabla que almacena direcciones
type Address struct {
	Country  string `gorm:"primary_key"`
	City     string `gorm:"primary_key"`
	PostCode string `gorm:"primary_key"`
	Line1    sql.NullString
	Line2    sql.NullString
}

func main() {

	db, err := gorm.Open("postgres", "postgresql://postgres:angel0505@localhost:5432/dbgorm")

	//"postgresql://usuario:contraseña@localhost:puerto/baseDatos"

	if err != nil {
		panic(err)
	}

	// Ping function checks the database connectivity
	err = db.DB().Ping()
	if err != nil {
		panic(err)
	}

	//create models
	//db.CreateTable(&Address{})
	//db.CreateTable(&Company{})
	//db.CreateTable(&Employee{})

	defer db.Close()

	//fmt.Println("creando modelos...!")

	//company de prueba
	sampleCompany :=
		Company{
			Name: "Liftit",
		}

	sampleAddres :=
		Address{
			Country:  "USA",
			City:     "Moutain View",
			PostCode: "1600",
		}

	sampleEmploye :=
		Employee{
			FirstName:        "John",
			LastName:         "Doe",
			SocialSecurityNo: "00-000-0000",
		}

	//Crea un registro único de la Compañía y todas las asociaciones (Dirección y Empleados)
	db.Create(&sampleCompany)
	db.Create(&sampleAddres)
	db.Create(&sampleEmploye)

	//completar luego:
	//  http://jinzhu.me/gorm/crud.html#create
	//  http://blog.ralch.com/tutorial/golang-object-relation-mapping-with-gorm/

}
