package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

var RotasUsuarios = []Rota{
	{
		Uri:                "/criar-usuario",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarTelaDeCadastro,
		RequerAutenticacao: false,
	},
	
	{
		Uri:                "/usuario",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarUsuario,
		RequerAutenticacao: false,
	},
	{
		Uri:                "/cadastrar",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarTelaDeCadastro,
		RequerAutenticacao: false,
	},
}
