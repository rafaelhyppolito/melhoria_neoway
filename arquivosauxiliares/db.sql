create table base (
    id SERIAL PRIMARY KEY,
    Cpf varchar(18) not null,
	Priv int null,
	Incompleto int null,
	DtUltCompra date null, 
	TicketMedio numeric(9,2),
	TicketUltCompra numeric(9,2),
	LojMaisFrequente varchar(18),
	LojUltCompra varchar(18)
);

create table basetmp (
    id SERIAL PRIMARY KEY,
    Cpf varchar(255) not null,
	Priv varchar(255) null,
	Incompleto varchar(255) null,
	DtUltCompra varchar(255) null, 
	TicketMedio varchar(255) null,
	TicketUltCompra varchar(255) null,
	LojMaisFrequente varchar(255) null,
	LojUltCompra varchar(255) null
);