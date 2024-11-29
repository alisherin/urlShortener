begin;

create table users (
    id bigserial primary key,
    name varchar(255) not null,
    password varchar(255) not null
);

commit;

begin;

INSERT INTO users (name, password) VALUES ('tester', '123456');

commit;

begin;

create table links (
    id bigserial primary key,
    url varchar(500) not null,
    short varchar(1000) not null,
    users_id bigint not null references users,
    expired_at timestamp not null,
    counter int default (0),
    last_used_at timestamp null,
    unique (short)
);

commit;