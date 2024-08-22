create table records
(
    id         serial primary key,
    name       varchar(255) not null unique,
    price      int          not null default 0,
    params     varchar(255) not null default '',
    is_deleted boolean      not null default true,
    created_at timestamp    not null default now(),
    updated_at timestamp    not null default now()
);

insert into records (name, price) values ('Name', 1000);
