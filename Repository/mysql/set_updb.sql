CREATE TABLE users {
    id int primary key AUTO_INCREAMENT,
    name varchar(255) not null
    phone_number varchar(255) not null unique,
    password varchar(255) not null,
    created_at datatime DEFAULT CURRENT_TIMESTAMP
};