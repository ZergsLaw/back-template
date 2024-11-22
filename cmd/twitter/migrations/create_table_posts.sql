create table posts
(
    id              uuid      primary key,
    post            text      not null,
    created_at      timestamp not null default now(),
    updated_at      timestamp not null default now(),
    is_published    bool      not null default true,
    user_id         uuid      not null
);