package controller

import (
	"encoding/json"
	"fmt"

	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/LeandroAlcantara-1997/model/email"
	"github.com/LeandroAlcantara-1997/model/entity"
	"github.com/LeandroAlcantara-1997/model/repository"
)

var templates *template.Template
var tem = template.Must(template.ParseGlob("view/*.html"))

func CadastraCliente(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Erro ao cadastrar Cliente ", err)
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
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Cadastrado realidado com sucesso!")
	return
}

func LoginCliente(w http.ResponseWriter, r *http.Request) {
	tem.ExecuteTemplate(w, "index.html", nil)
	email := r.FormValue("email")
	senha := r.FormValue("senha")
	fmt.Println(email, senha)
	err := repository.Logar(email, senha)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, err)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	fmt.Fprint(w, "Logado")
	return
}

func LoginAdmin(w http.ResponseWriter, r *http.Request) {
	var admin entity.Admin
	admin.User = r.FormValue("user")
	admin.Senha = r.FormValue("senha")
	err := admin.ValidaAdmin()
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Erro ao validar admin", err)
		return
	}

	err = repository.InsertAdmin(&admin)
	if err != nil {
		fmt.Fprint(w, "Erro ao inserir admin: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	fmt.Fprint(w, "Admin logado")
	//http.Redirect()
	return
}
func RestauraSenha(w http.ResponseWriter, r *http.Request) {
	var cliente entity.Cliente
	err := email.RecuperarSenha(cliente.Contato.Email)
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

func GetAluguel(w http.ResponseWriter, r *http.Request) {
	c := entity.Veiculo{
		Id:     1,
		Modelo: "Corsa",
	}
	encoder := json.NewEncoder(w)
	encoder.Encode(c)
	return
}
