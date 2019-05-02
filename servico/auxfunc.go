package servico

import (
	"github.com/rafaelhyppolito/melhoria_neoway/servico"
	"database/sql"
	"github.com/rafaelhyppolito/melhoria_neoway/repo"
	"strings"
	"fmt"
	"bufio"
	"os"
	"bytes"
	"regexp"
	"strconv"
	"unicode"
) 

var Conexao *sql.DB
func init(){
	Conexao = repo.Connect()
}


//Funcao que realiza o carregamento do arquivo e sua importação para o banco de dados
func importaArquivoCSV(caminho string)  {
	//Limpa a tabela temporária
	repo.ExecSQL(repo.TruncateTmp(), Conexao)

	//Limpa a tabela temporária
	repo.ExecSQL(repo.TruncateTmp(), Conexao)

	//Limpa a tabela temporária
	repo.ExecSQL(repo.TruncateTmp(), Conexao)

	//Lê e higieniza os dados do arquivo inserindo-os na tabela temporária
	servico.LerTexto(caminho) 

	//Insere na tabela final, convertendo os dados para os formatos corretos
	repo.ExecSQL(repo.InsertFinal(), Conexao)
	fmt.Println("Arquivo carregado com sucesso!")
}





/*
------------------------ TRATAMENTO DE ARQUIVO E IMPORTAÇÃO PARA O BANCO ----------------------------------------
*/
//Funcao que verifica se existe espaço duplo ao longo da String e retorna TRUE ou FALSE
func existeespacoduplo(texto string) bool {
	var ant, atu int
    for _, v := range texto {
		atu = int(v)
            if v == 32 {
				if ant == atu{
					return true
				}
			}
		ant = int(v)
	}
	return false
}

//Funcao que remove todos os espaços duplos do texto e retorna com os separadores csv (;)
func removeespacosduplo(texto string) string {
	for existeespacoduplo(texto){
		texto = strings.ReplaceAll(texto,"  ", " ")
	}	
	texto = strings.ReplaceAll(texto,",", ".")
	texto = strings.ReplaceAll(texto," ", "','")
	texto = "INSERT INTO basetmp(cpf,priv,incompleto,dtultcompra,ticketmedio,ticketultcompra,lojmaisfrequente,lojultcompra) VALUES ('"+texto+"'); "

	connection := repo.Connect()
	repo.ExecSQL(texto, connection)
	connection.Close()
	return texto
}

// Funcao que le o conteudo do arquivo e retorna um slice the string com todas as linhas do arquivo
func LerTexto(caminhoDoArquivo string) ([]string, error) {
	// Abre o arquivo
	arquivo, err := os.Open(caminhoDoArquivo)
	// Caso tenha encontrado algum erro ao tentar abrir o arquivo retorne o erro encontrado
	if err != nil {
		return nil, err
	}
	// Garante que o arquivo sera fechado apos o uso
	defer arquivo.Close()

	// Cria um scanner que le cada linha do arquivo
	var linhas []string
	scanner := bufio.NewScanner(arquivo)

	fmt.Println("Dados sendo higienizados... Isso pode levars algum tempo! \r\nPor favor aguarde ou Tecle CTRL+C para cancelar!")
	for scanner.Scan() {
		linhas = append(linhas, removeespacosduplo(scanner.Text()))
	}

	// Retorna as linhas lidas e um erro se ocorrer algum erro no scanner
	//fmt.Printf("Linhas: %s\n\r",linhas)
	return linhas, scanner.Err()
}


/*
------------------------ TRATAMENTO DE CPF E CNPJ ----------------------------------------
Fonte: https://github.com/Nhanderu/brdoc/blob/master/cpfcnpj.go
*/
// IsCPF verifies if the string is a valid CPF
// document.
func IsCPF(doc string) bool {

	const (
		sizeWithoutDigits = 9
		position          = 10
	)

	return isCPFOrCNPJ(
		doc,
		ValidateCPFFormat,
		sizeWithoutDigits,
		position,
	)
}

// IsCNPJ verifies if the string is a valid CNPJ
// document.
func IsCNPJ(doc string) bool {

	const (
		sizeWithoutDigits = 12
		position          = 5
	)

	return isCPFOrCNPJ(
		doc,
		ValidateCNPJFormat,
		sizeWithoutDigits,
		position,
	)
}

// ValidateCPFFormat verifies if the CPF has a
// valid format.
func ValidateCPFFormat(doc string) bool {

	const (
		pattern = `^\d{3}\.?\d{3}\.?\d{3}-?\d{2}$`
	)

	return regexp.MustCompile(pattern).MatchString(doc)
}

// ValidateCNPJFormat verifies if the CNPJ has a
// valid format.
func ValidateCNPJFormat(doc string) bool {

	const (
		pattern = `^\d{2}\.?\d{3}\.?\d{3}\/?(:?\d{3}[1-9]|\d{2}[1-9]\d|\d[1-9]\d{2}|[1-9]\d{3})-?\d{2}$`
	)

	return regexp.MustCompile(pattern).MatchString(doc)
}

// isCPFOrCNPJ generates the digits for a given
// CPF or CNPJ and compares it with the original
// digits.
func isCPFOrCNPJ(doc string, validate func(string) bool, size int, position int) bool {

	if !validate(doc) {
		return false
	}

	// Removes special characters.
	clean(&doc)

	// Invalidates documents with all
	// digits equal.
	if allEq(doc) {
		return false
	}

	// Calculates the first digit.
	d := doc[:size]
	digit := calculateDigit(d, position)

	// Calculates the second digit.
	d = d + digit
	digit = calculateDigit(d, position+1)

	return doc == d+digit
}

// clean removes every rune that is not a digit.
func clean(doc *string) {

	buf := bytes.NewBufferString("")
	for _, r := range *doc {
		if unicode.IsDigit(r) {
			buf.WriteRune(r)
		}
	}

	*doc = buf.String()
}

// allEq checks if every rune in a given string
// is equal.
func allEq(doc string) bool {

	base := doc[0]
	for i := 1; i < len(doc); i++ {
		if base != doc[i] {
			return false
		}
	}

	return true
}

// calculateDigit calculates the next digit for
// the given document.
func calculateDigit(doc string, position int) string {

	var sum int
	for _, r := range doc {

		sum += int(r-'0') * position
		position--

		if position < 2 {
			position = 9
		}
	}

	sum %= 11
	if sum < 2 {
		return "0"
	}

	return strconv.Itoa(11 - sum)
}