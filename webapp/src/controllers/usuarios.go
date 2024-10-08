package controllers

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/respostas"

	_ "github.com/go-sql-driver/mysql"
)

func emailJaCadastrado(email string) (bool, error) {
    db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/todolist")
    if err != nil {
        return false, fmt.Errorf("erro ao conectar ao banco de dados: %v", err)
    }
    defer db.Close()

    var count int
    query := "SELECT COUNT(*) FROM usuarios WHERE email = ?"
    err = db.QueryRow(query, email).Scan(&count)
    if err != nil {
        return false, fmt.Errorf("erro ao verificar e-mail: %v", err)
    }
    return count > 0, nil
}
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    err := r.ParseForm()
    if err != nil {
        respostas.JSON(w, http.StatusBadRequest, respostas.Error{Erro: "Unable to parse form"})
        fmt.Printf("Erro ao parsear o formulário: %v\n", err)
        return
    }

    email := r.FormValue("email")

    emailExistente, err := emailJaCadastrado(email)
    if err != nil {
        respostas.JSON(w, http.StatusInternalServerError, respostas.Error{Erro: "Error checking email"})
        fmt.Printf("Erro ao verificar email: %v\n", err)
        return
    }
    if emailExistente {
        respostas.JSON(w, http.StatusConflict, respostas.Error{Erro: "Email already registered"})
        return
    }

    usuario := map[string]string{
        "nome":  r.FormValue("nome"),
        "nick":  r.FormValue("nick"),
        "email": r.FormValue("email"),
        "senha": r.FormValue("senha"),
    }

    usuarioSeguro := map[string]string{
        "nome":  usuario["nome"],
        "nick":  usuario["nick"],
        "email": usuario["email"],
        "senha": "*****",
    }

    usuarioJSON, err := json.Marshal(usuario)
    if err != nil {
        respostas.JSON(w, http.StatusInternalServerError, respostas.Error{Erro: "Unable to marshal JSON"})
        fmt.Printf("Erro ao marshall JSON: %v\n", err)
        return
    }

    url := fmt.Sprintf("%s/usuarios", config.ApiUrl) 
    response, erro := http.Post(url, "application/json", bytes.NewBuffer(usuarioJSON))
    if erro != nil {
        respostas.TratarStatusCodeDeErro(w, response)
        fmt.Printf("Erro ao fazer POST para a API externa: %v\n", erro)
        return
    }
    defer response.Body.Close()

    if response.StatusCode >= 400 {
        respostas.TratarStatusCodeDeErro(w, response)
        return
    }

    fmt.Println("Resposta da API:", response.Status)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(usuarioSeguro)

    fmt.Println("Dados recebidos e enviados:", usuarioSeguro)
}
