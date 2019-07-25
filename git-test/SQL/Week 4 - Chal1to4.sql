-- alter table operators 
-- modify column created_at timestamp default CURRENT_TIMESTAMP
-- ;

use ata_online_shop;

-- CHAL1
-- NOMOR 1

alter table operators 
modify column created_at timestamp default CURRENT_TIMESTAMP,
modify column updated_at timestamp default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
;
Select*from operators;

INSERT INTO operators(id, name) VALUES 
('1', 'TELKOMSEL'),
('2', 'ISAT'),
('3', 'XL'),
('4', 'AXIS'),
('5', 'PLN');

Update operators set name = 'PLNPREPAID' where id = 4;

DELETE FROM operators WHERE id BETWEEN 1 AND 6;


-- NOMOR 2
select*from product_types;

alter table product_types 
modify column created_at timestamp default CURRENT_TIMESTAMP,
modify column updated_at timestamp default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
;

INSERT INTO product_types(id, name) VALUES 
('1', 'PULSA'),
('2', 'DATA'),
('3', 'PLN')
;

-- NOMOR 3, 4,5,6
select*from product_descriptions;

alter table product_descriptions
modify column created_at timestamp default CURRENT_TIMESTAMP,
modify column updated_at timestamp default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
;

INSERT INTO product_descriptions(id, description) VALUES 
('1','XL 50K'),
('2','XL 10K');


select*from products;

alter table products
modify column created_at timestamp default CURRENT_TIMESTAMP,
modify column updated_at timestamp default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
;

INSERT INTO products(id, product_type_id, operator_id, product_desc_id, code ,name , status) VALUES 
('1', '1','3','1','4','XL 50K',009),
('2', '1','3','2','5','XL 10K',009);


INSERT INTO products(id, product_type_id, operator_id, product_desc_id, code ,name , status) VALUES 
('3', '2','1','3','4','TSEL 50GB DATA',009),
('4', '2','1','4','5','TSEL 10GB DATA',088),
('5', '2','1','5','6','TSEL 9GB',010)
;

INSERT INTO products(id, product_type_id, operator_id, product_desc_id, code ,name , status) VALUES 
('6', '3','4','6','3','PLN 20K',009),
('7', '3','4','7','1','PLN 100K',009),
('8', '3','4','8','2','PLN 200K',009);

INSERT INTO product_descriptions(id, description) VALUES 
('3','TSEL 50GB'),
('4','TSEL 10GB'),
('5','TSEL 9GB')
;

INSERT INTO product_descriptions(id, description) VALUES 
('6','PLN 20K'),
('7','PLN 100K'),
('8','PLN 200K')
;

-- NOMOR 7
select*from payment_method_desc;


INSERT INTO payment_method_desc(id, name) VALUES 
('1','VA'),
('2','CC');


select*from payment_methods;

alter table payment_methods
modify column created_at timestamp default CURRENT_TIMESTAMP,
modify column updated_at timestamp default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
;

INSERT INTO payment_methods(id, name, status,desc_id) VALUES 
('1','CIMB',098,'1'),
('2','BCA',077,'1'),
('3','VISA',618,'2')
;

-- NOMOR 8

select*from address;

INSERT INTO address(id,address) VALUES 
('1','HOME'),
('2','OFFICE')
;


select*from users;

alter table users
modify column created_at timestamp default CURRENT_TIMESTAMP,
modify column updated_at timestamp default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
;

INSERT INTO users(id, status, dob,gender,address_id) VALUES 
('1','002',19951212,'m','1'),
('2','001',19961223,'m','1'),
('3','001',19970108,'f','2'),
('4','002',19971121,'f','2'),
('5','002',19930606,'f','2')
;

-- NOMOR 9
select*from transactions;

alter table transactions
modify column created_at timestamp default CURRENT_TIMESTAMP,
modify column updated_at timestamp default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
;

INSERT INTO transactions(id, user_id, payment_method_id,status,total_qty,total_price) VALUES 
('1','1','1','002','3',150000),
('2','1','1','002','3',150000),
('3','1','1','002','3',150000),
('4','2','1','002','3',150000),
('5','2','1','002','3',150000),
('6','2','1','002','3',150000),
('7','3','1','002','3',150000),
('8','3','1','002','3',150000),
('9','3','1','002','3',150000),
('10','4','2','002','3',150000),
('11','4','2','002','3',150000),
('12','4','2','002','3',150000),
('13','5','2','002','3',150000),
('14','5','2','002','3',150000),
('15','5','3','002','3',150000)
;

-- NOMOR 10
select*from transaction_details;

alter table transaction_details
modify column created_at timestamp default CURRENT_TIMESTAMP,
modify column updated_at timestamp default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
;

INSERT INTO transaction_details(transaction_id, product_id,status,qty,price) VALUES 
('1','1','002','1',50000),
('1','2','002','1',50000),
('1','3','002','1',50000),

('2','4','002','1',50000),
('2','5','002','1',50000),
('2','2','002','1',50000),

('3','4','002','1',50000),
('3','2','002','1',50000),
('3','1','002','1',50000),

('4','2','002','1',50000),
('4','6','002','1',50000),
('4','7','002','1',50000),


('5','2','002','1',50000),
('5','4','002','1',50000),
('5','5','002','1',50000),


('6','1','002','1',50000),
('6','6','002','1',50000),
('6','7','002','1',50000),

('7','3','002','1',50000),
('7','5','002','1',50000),
('7','1','002','1',50000),


('8','4','002','1',50000),
('8','6','002','1',50000),
('8','7','002','1',50000),

('9','2','002','1',50000),
('9','7','002','1',50000),
('9','5','002','1',50000),

('10','2','002','1',50000),
('10','7','002','1',50000),
('10','5','002','1',50000),

('11','2','002','1',50000),
('11','7','002','1',50000),
('11','5','002','1',50000),

('12','2','002','1',50000),
('12','7','002','1',50000),
('12','5','002','1',50000),

('13','2','002','1',50000),
('13','7','002','1',50000),
('13','5','002','1',50000),

('14','2','002','1',50000),
('14','7','002','1',50000),
('14','5','002','1',50000),

('15','2','002','1',50000),
('15','7','002','1',50000),
('15','5','002','1',50000)
;

-- CHAL 2
-- NOMOR 1
select*from users;


alter table users
add column name varchar(90)
;

Update users set name="BOY" where id=1;

Update users set name="DAVID" where id=2;

Update users set name="LIS" where id=3;

Update users set name="JEN" where id=4;

Update users set name="ROS" where id=5;

Select name from users where gender="m";

-- NOMOR 2

Select * from products where id=3;

-- NOMOR 3

Select * from users where created_at > (now() - interval 7 Day) and name like '%a%';

-- NOMOR 4

select count(1) from users where gender = "f";


-- NOMOR 5
select * from users order by name asc;

-- NOMOR 6
Select * from transaction_details where product_id=3 limit 5;


-- CHALLENGE 3
-- NOMOR 1
select*from products;

Update products set name="product dummy" where id=1;

-- NOMOR 2
select*from transaction_details;

Update transaction_details set qty=3 where product_id=1;

-- CHALLENGE 4
-- NOMOR 1

DELETE FROM transaction_details WHERE product_id=1;
DELETE FROM products WHERE id=1;
select*from products;

-- NOMOR 2

DELETE FROM transaction_details WHERE product_id=2;
DELETE FROM products WHERE product_type_id=1;
select*from products;
