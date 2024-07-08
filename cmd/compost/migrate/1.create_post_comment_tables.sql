-- up

CREATE TABLE post_table
(
    id               serial,
    title            TEXT    not null,
    content          TEXT    not null,
    author_id        int     not null,
    comments_allowed boolean not null,
    created_at       timestamp default now(),

    primary key (id)
);

CREATE TABLE comment_table
(
    id                serial,
    post_id           int  not null,
    content           TEXT not null,
    author_id         int  not null,
    parent_comment_id int,
    created_at        timestamp default now(),

    primary key (id)
);

-- down

drop table post_table;
drop table comment_table;