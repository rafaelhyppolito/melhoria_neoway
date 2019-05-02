package main

import (
	"github.com/rafaelhyppolito/melhoria_neoway/repo"
	"time"
	"fmt"
	"net/http"
)


//Funcao para tratar as requisicoes POST e GET
func index(w http.ResponseWriter, r *http.Request) {

    if r.URL.Path != "/" {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
    }
 
    switch r.Method {
    case "GET":     
         http.ServeFile(w, r, "front/index.html")
    case "POST":
        if err := r.ParseForm(); err != nil {
            fmt.Fprintf(w, "ParseForm() err: %v", err)
            return
		}
		arquivo := r.FormValue("arquivo")
		
		var start time.Time
    	start = time.Now()
		repo.ImportaArquivoCSV(arquivo)
		fmt.Fprintf(w, "Arquivo carregado com sucesso! \r\nAs informações estão no banco de dados!")
		fmt.Println("Tempo de execução: ",time.Since(start))

    default:
        fmt.Fprintf(w, "Desculpe, apenas métodos GET e POST são suportados.")
    }
}

//Funcao principal que é executada assim que o aplicação é iniciada
func main() {
	http.HandleFunc("/", index)
	fmt.Println("Serviço ativo e ouvindo na porta 8080.")
	http.ListenAndServe(":8080", nil)
}