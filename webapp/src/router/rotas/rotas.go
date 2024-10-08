package rotas

import (
	"net/http"
	"webapp/src/middlewares"

	"github.com/gorilla/mux"
)

type Rota struct {
	Uri                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

func Configurar(router *mux.Router) *mux.Router {
    rotas := RotasLogin
    rotas = append(rotas, RotasUsuarios...)
    rotas = append(rotas, RotaPaginaPrincipal)
    rotas = append(rotas, RotasPublicacoes...)
    
    for _, rota := range rotas {
		if rota.RequerAutenticacao {
			router.HandleFunc(rota.Uri, middlewares.Logger(middlewares.Autenticar(rota.Funcao))).Methods(rota.Metodo)
		} else {
			router.HandleFunc(rota.Uri, middlewares.Logger(rota.Funcao)).Methods(rota.Metodo)
		}
	}
	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))
    return router
}
