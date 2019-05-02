package model

//Estrutura da tabela principal
type Base struct {
	Cpf string `json:"cpf" db:"cpf"`
	Private string `json:"Private" db:"Priv"`
	Incompleto string `json:"Incompleto" db:"Incompleto"`
	DtUltCompra string `json:"DtUltCompra" db:"DtUltCompra"`
	TicketMedio string `json:"TicketMedio" db:"TicketMedio"`
	TicketUltCompra string `json:"TickertUltCompra" db:"TickertUltCompra"`
	LojMaisFrequente string `json:"LojMaisFrequente" db:"LojMaisFrequente"`
	LojUltCompra string `json:"LojUltCompra" db:"LojUltCompra"`
}