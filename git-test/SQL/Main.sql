create database sepulsa_reporting;

-- using db and creating table
use sepulsa_reporting;

create table users(
	id int primary key
);


-- showing table content 

select*from users;

/*adding column */ 

alter table users 
add column (name varchar(30),
address varchar(100),
phone varchar(30),
ssn varchar(12)
);

alter table users drop name;
-- -------------------------------------------------------------------

-- creating table, adding and dropping multiple columns
create table customers(
	id int primary key
    );
    
    
alter table customers 
add column (name varchar(30),
address varchar(100),
phone varchar(30),
ssn varchar(12)
);

alter table customers 
drop column ssn,
drop column phone;

alter table customers 
add column (
phone varchar(30),
ssn varchar(12)
);
-- ---------------------------------------------------- 
   
rename table test to real_keys;

select*from customers;    

-- showing db and tb-- -------------------------------------------
show databases;
show tables;

/*dropping to delete*/ 
drop table test;
drop database sepulsa_reporting;
----------------------------------


create table product_types(
id int primary key,
type_name varchar(30)
);
select*from product_types;

--  INSERT INTO `sepulsa_reporting`.`product_types` (`id`, `type_name`) VALUES ('5', 'PDAM');
-- DELETE FROM `sepulsa_reporting`.`products` WHERE (`id` = '1');


create table products(
id int primary key,
type_id int,
foreign key (type_id) references product_types(id),
name varchar(50),
label varchar(50)
);
select*from products;

select*
from products p
left join product_types pt on p.type_id=pt.id;
