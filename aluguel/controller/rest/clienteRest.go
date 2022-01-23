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
		service.ReponseError(w, 400, "Erro ao retornar veiculos", err)
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
	
	http.Redirect(w, r, "/cliente/home", http.StatusFound)
}

func DeletaCadastro(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	
	if err := repository.DeletaCliente(id); err != nil {
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

	
	if err := repository.Logar(email, senha); err != nil {
		service.ReponseError(w, 401, "Login inválido", err )
		return
	}

	
	if err := service.GeraCookie(r, w, email); err != nil {
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
	
	if err := repository.GetEmailToSenha(emailvalue); err != nil {
		service.ReponseError(w, 400, "Erro adquirir email", err )
		return
	}

	
	if err := email.RecuperarSenha(emailvalue); err != nil {
		service.ReponseError(w, 400, "Erro ao recuperar senha", err )
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Enviado com sucesso")

}

func GetCarrosAlugados(w http.ResponseWriter, r *http.Request) {
	return
}


