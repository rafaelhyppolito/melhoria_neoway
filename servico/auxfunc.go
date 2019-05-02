package servico

import (
	"database/sql"
	"github.com/rafaelhyppolito/melhoria_neoway/repo"
	"fmt"
) 

var Conexao *sql.DB
func init(){
	Conexao = repo.Connect()
}


/*
------------------------ TRATAMENTO DE ARQUIVO E IMPORTAÇÃO PARA O BANCO ----------------------------------------
*/

//Funcao que realiza o carregamento do arquivo e sua importação para o banco de dados
func ImportaArquivoCSV(caminho string)  {
	//Limpa a tabela temporária
	repo.ExecSQL(repo.TruncateTmpSQL(), Conexao)

	//Importa o conteudo do arquivo informado para dentro da base
	repo.ExecSQL(repo.ImportaArquivoSQL(caminho), Conexao)

	//Remove as virgulas e troca por pontos, nos campos de valores
	repo.ExecSQL(repo.RemoveVirgulasSQL(), Conexao)

	//Seta o formato de data para evitar erros
	repo.ExecSQL(repo.DateStyleSQL(), Conexao)

	//Insere na tabela final, convertendo os dados para os formatos corretos
	repo.ExecSQL(repo.InsertFinalSQL(), Conexao)
	fmt.Println("Arquivo carregado com sucesso!")
}