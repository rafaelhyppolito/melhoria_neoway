package repo

//Funcao com srcipt SQL para limpar a tabela temporária
func TruncateTmpSQL() string{
	return "TRUNCATE TABLE basetmp"
}

//Funcao com srcipt SQL para importar o arquivo txt para dentro da base SQL
func ImportaArquivoSQL(caminho string) string {
	return "COPY basetmp (cpf,priv,incompleto,dtultcompra,ticketmedio,ticketultcompra,"+
		   "lojmaisfrequente,lojultcompra) FROM '" + caminho + "' USING DELIMITERS ';'"
}

//Funcao com srcipt SQL para remover todas as virgulas dos valores e trocar por pontos
func RemoveVirgulasSQL() string{
	return "UPDATE basetmp SET ticketmedio = replace(ticketmedio,',','.'), "+
		   "ticketultcompra = replace(ticketultcompra,',','.')"
}

//Funcao com script SQL para setar o formato de data e evitar erros
func DateStyleSQL() string{
	return "SET datestyle = dmy"
}

//Funcao com script SQL para inserção final dos dados já higienizados
func InsertFinalSQL() string {
	return "insert into base(cpf,priv,incompleto,dtultcompra,ticketmedio,ticketultcompra,"+
		   "lojmaisfrequente,lojultcompra) select replace(replace(cpf,'.',''),'-',''), CAST"+
		   "(priv AS INTEGER), CAST(incompleto AS INTEGER), CASE WHEN dtultcompra = 'NULL' "+
		   "THEN NULL ELSE CAST(dtultcompra AS DATE) END AS dtultcompra, CASE WHEN ticketmedio"+
		   " = 'NULL' THEN NULL ELSE CAST(ticketmedio AS NUMERIC(9,2)) END AS ticketmedio, CASE"+
		   " WHEN ticketultcompra = 'NULL' THEN NULL ELSE CAST(ticketultcompra AS NUMERIC(9,2)) "+
		   "END AS ticketultcompra, CASE WHEN lojmaisfrequente = 'NULL' THEN NULL ELSE "+
		   "replace(replace(replace(lojmaisfrequente,'.',''),'-',''),'/','') END AS lojmaisfrequente, "+
		   "CASE WHEN lojultcompra = 'NULL' THEN NULL ELSE replace(replace(replace(lojultcompra,'.',''),"+
		   "'-',''),'/','') END AS lojultcompra from basetmp"
}