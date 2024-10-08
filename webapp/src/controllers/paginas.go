package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/modelos"
	"webapp/src/requisicoes"
	"webapp/src/respostas"
	"webapp/src/utils"
)

func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "login.html", nil)
}
func CarregarTelaDeCadastro(w http.ResponseWriter, r *http.Request) {
    utils.ExecutarTemplate(w, "cadastro.html", nil) 
}
func CarregarPaginaPrincipal(w http.ResponseWriter, r *http.Request) {
    url := fmt.Sprintf("%s/usuarios/2/publicacoes", config.ApiUrl)

    response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
    if erro != nil {
        respostas.JSON(w, http.StatusInternalServerError, respostas.Error{Erro: erro.Error()})
        return
    }
    defer response.Body.Close()

    if response.StatusCode >= 400 {
        respostas.TratarStatusCodeDeErro(w, response)
        return
    }

    var publicacoes []modelos.Publicacao
    if erro = json.NewDecoder(response.Body).Decode(&publicacoes); erro != nil {
        respostas.JSON(w, http.StatusUnprocessableEntity, respostas.Error{Erro: erro.Error()})
        return
    }

    cookie, _ := cookies.Ler(r)
    usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)

    fmt.Printf("Publicações carregadas: %+v\n", publicacoes)
    utils.ExecutarTemplate(w, "home.html", struct {
        Publicacoes []modelos.Publicacao
        UsuarioID uint64
    }{
        Publicacoes: publicacoes,
        UsuarioID: usuarioID,
    })
}

