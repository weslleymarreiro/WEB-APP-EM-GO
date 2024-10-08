package rotas

import (
	"net/http"
	"webapp/src/controllers"
	
)


var RotasPublicacoes = []Rota {
	{
		Uri: "/publicacoes",
		Metodo: http.MethodPost,
		Funcao: controllers.PostarPublicacao,
		RequerAutenticacao: true,
	},
}