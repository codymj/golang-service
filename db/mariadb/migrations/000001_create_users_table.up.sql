create table users
(
    id           bigint auto_increment primary key,
    username     varchar(16)           not null,
    email        varchar(32)           not null,
    password     varchar(128)          not null,
    location     varchar(16)           null,
    is_validated bool     default false not null,
    created_at   bigint default round(unix_timestamp(utc_timestamp(4))*1000) not null,
    modified_at  bigint default round(unix_timestamp(utc_timestamp(4))*1000) not null,
    CONSTRAINT uc_users unique (username, email)
) AUTO_INCREMENT=1000;
