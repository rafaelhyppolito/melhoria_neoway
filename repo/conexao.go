package repo

import(
    "fmt"
    "database/sql"
	_ "github.com/lib/pq"
)

//Funcao que retorna uma forma de conexão com o banco
func Connect() *sql.DB {

	driverConfig := "user=postgres password=root port=5432 database=postgres sslmode=disable"
	connection, err := sql.Open("postgres", driverConfig)

    if err != nil {
        fmt.Printf("database.Connect ERROR: %s\r\n", err)
	}

	err = connection.Ping()

	if err != nil {
        fmt.Printf("Deu ruim: %s", err)
	}
	return connection
}

//Funcao que executa comandos em SQL. Não tem retorno
func ExecSQL(sql string, connection *sql.DB) {
	sqlStatement := sql
	_, err := connection.Exec(sqlStatement)
	if err != nil {
		  panic(err)
	}	
}