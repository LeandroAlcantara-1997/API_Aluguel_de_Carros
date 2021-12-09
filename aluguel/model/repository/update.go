package repository

import (
	"fmt"

	"github.com/LeandroAlcantara-1997/model/entity"
)

func UpdateSenha(email, senha string) error {
	db, err := OpenSQL()
	if err != nil {
		return fmt.Errorf("Erro ao atualizar senha: %v", err)
	}
	token, err := entity.GeraToken(email + senha)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	result, err := db.Exec("UPDATE login "+
					"SET senha = '" + senha + "', token='" + token + 
					"' WHERE email = '" + email + "'")
	if err != nil {
		return fmt.Errorf("Erro ao fazer update %v", err)
	}
	fmt.Println(result.RowsAffected())
	return nil
}