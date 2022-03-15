package rest

import (
	"encoding/json"

	"io/ioutil"
	"net/http"

	service "github.com/LeandroAlcantara-1997/service"
	"github.com/gorilla/mux"

	"github.com/LeandroAlcantara-1997/entity"
	"github.com/LeandroAlcantara-1997/repository"
)

func PostCadastraCliente(w http.ResponseWriter, r *http.Request) {
	var novocadastro entity.Cliente
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		service.ReponseError(w, 400, "Erro ao cadastrar cliente", err)
		return
	}

	if err = json.Unmarshal(body, &novocadastro); err != nil {
		service.ReponseError(w, 400, "Erro ao cadastrar cliente", err)
		return
	}

	if err = novocadastro.ValidaCliente(); err != nil {
		service.ReponseError(w, 400, "Erro ao cadastrar cliente", err)
		return
	}

	if err = repository.InsertCliente(&novocadastro); err != nil {
		service.ReponseError(w, 400, "Erro ao cadastrar cliente", err)
		return
	}
	service.JsonResponse(w, 201, novocadastro)
	return
}

func DeletaCadastro(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)

	if err := repository.DeletaCliente(id["id"]); err != nil {
		service.ReponseError(w, 400, "Erro ao deletar cadastro", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	return
}

func PostLoginCliente(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	senha := r.FormValue("senha")

	if err := repository.Logar(email, senha); err != nil {
		service.ReponseError(w, 401, "Login inválido", err)
		return
	}

	if err := service.GeraCookie(r, w, email); err != nil {
		service.ReponseError(w, 400, "Erro ao gerar cookie", err)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	return
}

func PostRestauraSenha(w http.ResponseWriter, r *http.Request) {
	emailvalue := r.FormValue("email")

	if err := repository.GetEmailToSenha(emailvalue); err != nil {
		service.ReponseError(w, 400, "Erro adquirir email", err)
		return
	}

	if err := service.RecuperarSenha(emailvalue); err != nil {
		service.ReponseError(w, 400, "Erro ao recuperar senha", err)
		return
	}
	service.JsonResponse(w, 200, "Enviado com sucesso")
	return
}

func GetCarrosAlugados(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)
	veiculo, err := repository.GetAlugadosCliente(id["id"])
	if err != nil {
		service.ReponseError(w, 400, "Erro ao recuperar veiculo", err)
		return
	}

	service.JsonResponse(w, 302, veiculo)
	return
}
