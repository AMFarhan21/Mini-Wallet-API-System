-- docker compose up -d
-- docker exec -it postgres psql -U {username} -d {db name}

CREATE TABLE users (
    id serial primary key,
    name varchar(255),
    email varchar(255) unique
);

CREATE TABLE wallets (
    id serial primary key,
    user_id int not null references users(id),
    balance bigint not null default 0
);

CREATE TABLE transactions (
    id serial primary key,
    wallet_id int not null references wallets(id),
    type varchar(20),
    amount bigint,
    reference_id varchar(255),
    created_at timestamp default current_timestamp
);

INSERT INTO users (name, email) VALUES
('andi', 'andi@gmail.com'),
('farhan', 'farhan@gmail.com'),
('asep', 'asep@gmail.com'),
('budi', 'budi@gmail.com');

INSERT INTO wallets (user_id, balance) VALUES
(1, 0),
(2, 0),
(3, 0),
(4, 0);