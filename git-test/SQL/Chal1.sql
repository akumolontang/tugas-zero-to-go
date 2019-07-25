-- CHAL 1 NOMOR 1
create database ata_online_shop;

use ata_online_shop;
------------------------------------- 

-- CHAL 1 NOMOR 2
create table users(
	id int(11) primary key,
    status smallint,
    dob date,
    gender char(1),
    created_at Timestamp,
    updated_at timestamp
);

create table product_types(
	id int(11) primary key,
    name varchar(255),
    created_at timestamp,
    updated_at timestamp
);


create table operators(
	id int(11) primary key,
    name varchar(255),
    created_at timestamp,
    updated_at timestamp
);

create table product_descriptions(
	id int(11) primary key,
    description text,
    created_at timestamp,
    updated_at timestamp
);


create table payment_methods(
	id int(11) primary key,
    name varchar(255),
    status smallint,
    created_at timestamp,
    updated_at timestamp
);


create table products(
	id int(11) primary key,
    product_type_id int(11),
    operator_id int(11),
    product_desc_id int(11),
    code varchar(50),
    name varchar(50),
    status smallint,
    created_at timestamp,
    updated_at timestamp,
    foreign key (product_type_id) references product_types(id),
    foreign key (operator_id) references operators(id),
    foreign key (product_desc_id) references product_descriptions(id)
);


create table transactions(
	id int(11) primary key,
    user_id int(11),
    payment_method_id int(11),
    status varchar(10),
    total_qty int(11),
    total_price numeric(25,2),
    created_at timestamp,
    updated_at timestamp,
	foreign key (user_id) references users(id),
    foreign key (payment_method_id) references payment_methods(id)
);


create table transaction_details(
	transaction_id int(11),
    product_id int(11),
    status varchar(10),
    qty int(11),
    price numeric(25,2),
    created_at timestamp,
    updated_at timestamp,
    primary key (transaction_id, product_id),
	foreign key (product_id) references products(id),
	foreign key (transaction_id) references transactions(id)
);


------------------------------------- 
-- CHAL 1 NOMOR 3
create table kurir(
	id int(11) primary key,
    name varchar(20),
    created_at timestamp,
    updated_at timestamp
);

-- CHAL 1 NOMOR 4
alter table kurir
add column (
ongkos_dasar int(12)
);

-- CHAL 1 NOMOR 5
rename table kurir to shipping;

-- CHAL 1 NOMOR 6
drop table shipping;

-------------------------------------

-- CHAL 1 NOMOR 7
-- a.)

create table payment_method_desc(
	id int(11) primary key,
    name varchar(255)
);

alter table payment_methods
add column (
desc_id int(11)
);

alter table payment_methods
add foreign key (desc_id) 
references payment_method_desc(id);


-- b.)
create table address(
	id int(11) primary key,
    address varchar(255)
);

alter table users
add column (
address_id int(11)
);

alter table users
add foreign key (address_id) 
references address(id);

-- c.)
create table user_payment_method_detail(
	user_id int(11),
    payment_method_id int(11),
    desc_detail varchar(255),
	primary key (user_id, payment_method_id),
    foreign key (user_id) references users(id),
    foreign key (payment_method_id) references payment_methods(id)
);




