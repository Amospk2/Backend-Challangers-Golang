CREATE TABLE IF NOT EXISTS planets (
    id uuid NOT NULL,
    nome varchar(255) UNIQUE NOT NULL,
    terreno varchar(255) NOT NULL,
    clima varchar(255) NOT NULL,
    filmes int default 0
);