package controller

import (
	"encoding/json"
	"fmt"

	"io/ioutil"
	"net/http"

	service "github.com/LeandroAlcantara-1997/controller/service"
	"github.com/LeandroAlcantara-1997/model/email"
	"github.com/LeandroAlcantara-1997/model/entity"
	"github.com/LeandroAlcantara-1997/model/repository"
)

func HomeCliente(w http.ResponseWriter, r *http.Request) {
	service.ExecuteTemplate(w, "recuperarSenha.html", nil)
	veiculos, err := repository.GetCarrosCadastrados()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
	}

	service.GetSecao(r)
	w.WriteHeader(http.StatusFound)
	encoder := json.NewEncoder(w)
	encoder.Encode(veiculos)
	
}

func GetCadastraCliente(w http.ResponseWriter, r *http.Request) {
	service.ExecuteTemplate(w, "cadastroCliente.html", nil)
}

func PostCadastraCliente(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		service.ReponseError(w, 400, "Erro ao cadastrar cliente", err)
		return
	}
	var novocadastro entity.Cliente
	err = json.Unmarshal(body, &novocadastro)
	if err != nil {
		service.ReponseError(w, 400, "Erro ao cadastrar cliente", err)
		return
	}
	err = novocadastro.ValidaCliente()
	if err != nil {
		service.ReponseError(w, 400, "Erro ao cadastrar cliente", err)
		return
	}
	err = novocadastro.Contato.ValidaContato()
	if err != nil {
		service.ReponseError(w, 400, "Erro ao cadastrar cliente", err)
		return
	}

	err = repository.InsertCliente(&novocadastro)
	if err != nil {
		service.ReponseError(w, 400, "Erro ao cadastrar cliente", err)
		return
	}
	
	http.Redirect(w, r, "/cliente/home", http.StatusFound)
}

func DeletaCadastro(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")

	err := repository.DeletaCliente(id)
	if err != nil {
		service.ReponseError(w, 400, "Erro ao deletar cadastro", err )
		return 
	}
	
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func GetLoginCliente(w http.ResponseWriter, r *http.Request) {
	service.ExecuteTemplate(w, "home.html", nil)
}

func PostLoginCliente(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	senha := r.FormValue("senha")

	err := repository.Logar(email, senha)
	if err != nil {
		service.ReponseError(w, 401, "Login inválido", err )
		return
	}

	err = service.GeraCookie(r, w, email)
	if err != nil {
		service.ReponseError(w, 400, "Erro ao gerar cookie", err )
		return
	}
	http.Redirect(w, r, "/cliente/home", http.StatusFound)
}

func GetRestauraSenha(w http.ResponseWriter, r *http.Request) {
	service.ExecuteTemplate(w, "recuperarSenha.html", nil)
}

func PostRestauraSenha(w http.ResponseWriter, r *http.Request) {
	emailvalue := r.FormValue("email")
	err := repository.GetEmailToSenha(emailvalue)
	if err != nil {
		service.ReponseError(w, 400, "Erro adquirir email", err )
		return
	}

	err = email.RecuperarSenha(emailvalue)
	if err != nil {
		service.ReponseError(w, 400, "Erro ao recuperar senha", err )
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Enviado com sucesso")
	return

}

func GetCarrosAlugados(w http.ResponseWriter, r *http.Request) {
	return
}


