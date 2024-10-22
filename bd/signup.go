package bd

import (
	"fmt"

	"example.com/m/v2/models"
	_ "github.com/go-sql-driver/mysql"
)

func SigUp(sig models.SignUp) error {
	fmt.Println("Comienza el registro")

	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	sentencia := "INSERT INTO usuarios.users ('USER_UUID', 'USER_Email') VALUES ('" + sig.UserUUID + "','" + sig.UserEmail + "')"
	fmt.Println(sentencia)
	_, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("Sig Up > Ejecución Exitosa")
	return nil

}
