# Melhoria do Testeneoway
Melhoria do Teste de manipulação de dados e persistência em base de dados relacional

# Requisitos
Possuir GO instalado;

Possuir Banco de dados PostgreSQL instalado;

Possuir Git instalado.

# Executando
1 - Abra o pgAdmin4 e dentro da base "postgres" execute os comandos contidos no arquivo "arquivoauxiliares/db.sql";

2 - Realize o download do projeto (ou clone com git) para dentro do repositório de projetos do Go (C:\Go\src\github.com\rafaelhyppolito);

3 - Abra o Git na pasta do projeto e execute o comando "GOOS=windows GOARCH=386 go build -o inicio.exe" (build Windows);

4 - Execute o arquivo "inicio.exe".

# Forma de utilizar
1 - Após executar o arquivo "inicio.exe", o serviço ficará escutando na porta 8080, então basta acessar no browser o seguinte endereço: "localhost:8080";

2 - Carregará uma página com um campo para informar o caminho completo do arquivo a ser importado;

3 - Depois de informar o caminho completo, clique em "Enviar" e aguardo a finalização da importação que será confirmada com uma mensagem em tela.

Obs.: Em "arquivoauxiliares/base_teste.txt", contém o arquivo para teste.