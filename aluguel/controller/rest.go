package controller

import (
	"encoding/json"
	"fmt"

	//"html/template"
	"io/ioutil"
	"net/http"

	"github.com/LeandroAlcantara-1997/model/email"
	"github.com/LeandroAlcantara-1997/model/entity"
	"github.com/LeandroAlcantara-1997/model/repository"
)

//var templates *template.Template
//var tem = template.Must(template.ParseGlob("view/*.html"))

//Utilizar bcrypt para gerar o cookie
//var store = sessions.NewCookieStore([]byte("t0p-s3cr3t"))

func HomeCliente(w http.ResponseWriter, r *http.Request) {
	//tem.ExecuteTemplate(w, "recuperarSenha.html", nil)
	fmt.Fprint(w, "Redirecionado!")
}

func CadastraCliente(w http.ResponseWriter, r *http.Request) {
	//tem.ExecuteTemplate(w, "cadastroCliente.html", nil)
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
	/*w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Cadastrado realidado com sucesso!")*/
	http.Redirect(w, r, "/homeCliente", http.StatusFound)

	return
}

func LoginCliente(w http.ResponseWriter, r *http.Request) {

	email := r.FormValue("email")
	senha := r.FormValue("senha")

	err := repository.Logar(email, senha)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, err)
		return
	}
	//var write http.ResponseWriter
	//tem.ExecuteTemplate(w, "home.html", nil)
	http.Redirect(w, r, "/homeCliente", http.StatusFound)

	return
}

func LoginAdmin(w http.ResponseWriter, r *http.Request) {
	//tem.ExecuteTemplate(w, "loginAdmin.html", nil)
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
	/*w.WriteHeader(http.StatusAccepted)
	fmt.Fprint(w, "Admin logado")*/
	http.Redirect(w, r, "/homeAdmin", http.StatusFound)

	return
}

func RestauraSenha(w http.ResponseWriter, r *http.Request) {
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

func GetAluguel(w http.ResponseWriter, r *http.Request) {
	c := entity.Veiculo{
		Id:     1,
		Modelo: "Corsa",
	}
	encoder := json.NewEncoder(w)
	encoder.Encode(c)
	return
}
