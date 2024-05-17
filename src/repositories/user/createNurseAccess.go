package userRepository

import (
	"errors"
	entities "hello-nurse/src/entities/user"
	"log"
)

func (i *controllerUser) CreateNurseAccess(params *entities.NurseRegisterAccess) error {

	println("pass: ", params.Password)
	res, err := i.DB.Exec(`UPDATE users SET has_access = $3 AND password = $1 WHERE id = $2`,
		params.Password, params.UserId, true,
	)

	if err != nil {
		log.Printf("Error giving access to nurse: %s", err)
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error checking row affected: %s", err)
		return err
	}

	if rowsAffected == 0 {
		log.Printf("Error no row affected: %s", err)
		return errors.New("No Row Affected")

	}

	return nil
}
