DROP TABLE IF EXISTS clients;
DROP TABLE IF EXISTS accounts;
DROP TABLE IF EXISTS transactions;

CREATE TABLE clients (
	id SERIAL PRIMARY KEY NOT NULL,
	name varchar(255) NULL,
	email varchar(255) NULL,
	phone varchar(255) NULL
	);

CREATE TABLE accounts (
  id SERIAL PRIMARY KEY NOT NULL,
  cleint_id integer REFERENCES clients (id),
  balance integer DEFAULT 0
);

CREATE TABLE transactions(
  id SERIAL PRIMARY KEY NOT NULL,
  comment varchar(255) NOT NULL,
  success boolean DEFAULT FALSE
);
