DROP TABLE IF EXISTS money_changes;
DROP TABLE IF EXISTS transactions;
DROP TABLE IF EXISTS accounts;
DROP TABLE IF EXISTS clients;


CREATE TABLE clients (
	id SERIAL PRIMARY KEY NOT NULL,
	name varchar(255) NULL,
	email varchar(255) NULL,
	phone varchar(255) NULL
);

CREATE TABLE accounts (
  id SERIAL PRIMARY KEY NOT NULL,
  client_id integer REFERENCES clients (id),
  balance integer DEFAULT 0
);

CREATE TABLE transactions(
  id SERIAL PRIMARY KEY NOT NULL,
  comment varchar(255) NOT NULL,
  success boolean DEFAULT FALSE
);

CREATE TABLE money_changes(
  transaction_id integer REFERENCES transactions (id),
  account_id integer REFERENCES accounts (id),
  diff integer
)
