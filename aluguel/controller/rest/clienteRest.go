package controller

import (
	"encoding/json"
	"fmt"

	"io/ioutil"
	"net/http"

	utils "github.com/LeandroAlcantara-1997/controller/utils"
	"github.com/LeandroAlcantara-1997/model/email"
	"github.com/LeandroAlcantara-1997/model/entity"
	"github.com/LeandroAlcantara-1997/model/repository"
)

func HomeCliente(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "recuperarSenha.html", nil)
	veiculos, err := repository.GetCarrosCadastrados()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
	}

	utils.GetSecao(r)
	w.WriteHeader(http.StatusFound)
	encoder := json.NewEncoder(w)
	encoder.Encode(veiculos)
	
}

func GetCadastraCliente(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "cadastroCliente.html", nil)
}

func PostCadastraCliente(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	var novocadastro entity.Cliente
	err = json.Unmarshal(body, &novocadastro)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Erro ao converter json para cliente: ", err)
		return
	}
	err = novocadastro.ValidaCliente()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Erro ao cadastrar cliente: ", err)
		return
	}
	err = novocadastro.Contato.ValidaContato()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Dados de contato inválidos: ", err)
		return
	}

	err = repository.InsertCliente(&novocadastro)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Erro ao inserir cadastro cliente ", err)
		return
	}
	
	http.Redirect(w, r, "/homeCliente", http.StatusFound)
	return
}

func DeletaCadastro(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	err := repository.DeletaCliente(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return 
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Cadastro deletado com sucesso")
	return 
}

func GetLoginCliente(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "home.html", nil)
}

func PostLoginCliente(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	senha := r.FormValue("senha")

	err := repository.Logar(email, senha)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, err)
		return
	}

	err = utils.GeraCookie(r, w, email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	http.Redirect(w, r, "/homeCliente", http.StatusFound)
	return
}

func GetRestauraSenha(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "recuperarSenha.html", nil)
}

func PostRestauraSenha(w http.ResponseWriter, r *http.Request) {
	emailvalue := r.FormValue("email")
	err := repository.GetEmailToSenha(emailvalue)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	err = email.RecuperarSenha(emailvalue)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Enviado com sucesso")

	return

}

func GetCarrosAlugados(w http.ResponseWriter, r *http.Request) {
	return
}


