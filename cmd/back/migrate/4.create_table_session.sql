-- up
create table sessions
(
    id         uuid                  not null,
    token      text                  not null,
    ip         text                  not null,
    user_agent text                  not null,
    user_id    uuid references users not null,
    created_at timestamp             not null default now(),
    updated_at timestamp             not null default now(),

    unique (token),
    primary key (id)
);

-- down
drop table sessions;
