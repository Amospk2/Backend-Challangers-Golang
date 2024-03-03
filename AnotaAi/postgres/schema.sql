
CREATE TABLE IF NOT EXISTS users (
    id uuid UNIQUE NOT NULL,
    name varchar(255) NOT NULL,
    email varchar(255) UNIQUE NOT NULL,
    age varchar(15) NOT NULL,
    password varchar(255) NOT NULL
);


CREATE TABLE IF NOT EXISTS Products (
    id uuid UNIQUE NOT NULL,
    title varchar(255) NOT NULL,
    description text NOT NULL,
    price DECIMAL NOT NULL,
    category varchar(255),
    ownerID varchar(255)
);


CREATE TABLE IF NOT EXISTS Categorys (
    id uuid UNIQUE NOT NULL,
    title varchar(255) NOT NULL,
    description text NOT NULL,
    ownerID uuid
);
