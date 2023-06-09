create table cities (
     id_city   int primary key,
     name   varchar (50)
);

create table items (
     id_item    int primary key ,
     name  varchar(50)
);

create table orders (
     id       int primary key generated by default as identity,
     city_id   int references cities(id_city),
     item_id   int references items(id_item),
     quantity varchar(50),
     price    varchar(50)
)
