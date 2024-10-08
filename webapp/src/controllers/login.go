package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/modelos"
	"webapp/src/respostas"
)

func FazerLogin(w http.ResponseWriter, r *http.Request) {
	var credenciais map[string]string
	erro := json.NewDecoder(r.Body).Decode(&credenciais)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.Error{Erro: "Dados inválidos"})
		return
	}

	usuario, erro := json.Marshal(credenciais)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.Error{Erro: "Erro ao criar JSON"})
		return
	}
	
	url := fmt.Sprintf("%s/login", config.ApiUrl)
	response, erro := http.Post(url, "application/json", bytes.NewBuffer(usuario))
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.Error{Erro: "Erro ao fazer requisição ao backend"})
		return
	}
	defer response.Body.Close()
	
	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	var dadosAutenticacao modelos.DadosAutenticacao
	if erro = json.NewDecoder(response.Body).Decode(&dadosAutenticacao); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.Error{Erro: erro.Error()})
		return
	}

	if erro = cookies.Salvar(w, dadosAutenticacao.ID, dadosAutenticacao.Token); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.Error{Erro: erro.Error()})
	}

	respostas.JSON(w, http.StatusOK, dadosAutenticacao)
}