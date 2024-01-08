CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create table users(
    id uuid primary key default uuid_generate_v4() not null,
    fname varchar(50) not null,
    lname varchar(50) not null,
    email varchar(255) unique not null,
    salt varchar(255) not null,
    hash varchar(255) not null,
    created_at timestamp not null,
    updated_at timestamp
);
