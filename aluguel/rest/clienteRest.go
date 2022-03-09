package rest

import (
	"encoding/json"
	"fmt"

	"io/ioutil"
	"net/http"
	"github.com/gorilla/mux"
	service "github.com/LeandroAlcantara-1997/service"


	"github.com/LeandroAlcantara-1997/entity"
	"github.com/LeandroAlcantara-1997/repository"
)

func PostCadastraCliente(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		service.ReponseError(w, 400, "Erro ao cadastrar cliente", err)
		return
	}
	var novocadastro entity.Cliente
	
	if err = json.Unmarshal(body, &novocadastro); err != nil {
		service.ReponseError(w, 400, "Erro ao cadastrar cliente", err)
		return
	}
	
	if err = novocadastro.ValidaCliente(); err != nil {
		service.ReponseError(w, 400, "Erro ao cadastrar cliente", err)
		return
	}
	
	if err = novocadastro.Contato.ValidaContato(); err != nil {
		service.ReponseError(w, 400, "Erro ao cadastrar cliente", err)
		return
	}

	
	if err = repository.InsertCliente(&novocadastro); err != nil {
		service.ReponseError(w, 400, "Erro ao cadastrar cliente", err)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	encoder := json.NewEncoder(w)
	encoder.Encode(novocadastro)
}

func DeletaCadastro(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)
	fmt.Println(id["id"])

	//id1 := r.FormValue("id")
	//fmt.Println(id1)
	if err := repository.DeletaCliente(id["id"]); err != nil {
		service.ReponseError(w, 400, "Erro ao deletar cadastro", err )
		return 
	}
	
	w.WriteHeader(http.StatusNoContent)
}


func PostLoginCliente(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	senha := r.FormValue("senha")

	
	if err := repository.Logar(email, senha); err != nil {
		service.ReponseError(w, 401, "Login inválido", err )
		return
	}

	
	if err := service.GeraCookie(r, w, email); err != nil {
		service.ReponseError(w, 400, "Erro ao gerar cookie", err )
		return
	}
	w.WriteHeader(http.StatusAccepted)
}

func PostRestauraSenha(w http.ResponseWriter, r *http.Request) {
	emailvalue := r.FormValue("email")
	
	if err := repository.GetEmailToSenha(emailvalue); err != nil {
		service.ReponseError(w, 400, "Erro adquirir email", err )
		return
	}

	
	if err := service.RecuperarSenha(emailvalue); err != nil {
		service.ReponseError(w, 400, "Erro ao recuperar senha", err )
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Enviado com sucesso")

}

func GetCarrosAlugados(w http.ResponseWriter, r *http.Request) {
	return
}


