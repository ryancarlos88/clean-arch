CREATE TABLE IF NOT EXISTS orders (
    id VARCHAR(255) PRIMARY KEY,
    price INT NOT NULL,
    tax INT NOT NULL,
    final_price INT NOT NULL
);

insert into orders (id, price, tax, final_price) values ('c8fc4279-344c-4f88-80c6-9a114cd8be96', 100, 10, 110);
insert into orders (id, price, tax, final_price) values ('d7750613-5059-471b-9636-9f51bab522c8', 300, 10, 330);
insert into orders (id, price, tax, final_price) values ('e6c869eb-300f-4a15-9d17-44c7e1d1bb6f', 500, 10, 550);
insert into orders (id, price, tax, final_price) values ('30b48805-f068-4353-a192-05489bb16064', 700, 10, 770);
insert into orders (id, price, tax, final_price) values ('3954689d-fa40-4fa6-ab44-6a963bbac514', 900, 10, 990);
insert into orders (id, price, tax, final_price) values ('ee6123e1-46fe-4b20-8869-8a626bf947dd', 2000, 20, 2400);