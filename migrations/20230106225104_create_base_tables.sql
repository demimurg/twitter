-- +goose Up
-- there already exist postgresql table user,
-- so we need plural form users
CREATE TABLE users (
    id         bigserial PRIMARY KEY,
    full_name  text      NOT NULL,
    email      text      NOT NULL UNIQUE,
    caption    text      NOT NULL,
    birth_date date      NOT NULL,
    created_at timestamp NOT NULL DEFAULT now(),
    deleted_at timestamp
);

CREATE TABLE follower (
    follower_id bigserial REFERENCES users (id),
    followee_id bigserial REFERENCES users (id),

    created_at  timestamp NOT NULL DEFAULT now(),
    deleted_at  timestamp
);

CREATE TABLE tweet (
    id         bigserial PRIMARY KEY,
    user_id    bigserial REFERENCES users (id),
    text       text      NOT NULL,
    likes      integer   NOT NULL DEFAULT 0,
    created_at timestamp NOT NULL DEFAULT now(),
    deleted_at timestamp
);

CREATE TABLE comment (
    id         bigserial PRIMARY KEY,
    user_id    bigserial REFERENCES users (id),
    tweet_id   bigserial REFERENCES tweet (id),
    text       text      NOT NULL,
    likes      integer   NOT NULL DEFAULT 0,
    created_at timestamp NOT NULL DEFAULT now(),
    deleted_at timestamp
);

-- +goose Down

DROP TABLE users;
DROP TABLE follower;
DROP TABLE tweet;
DROP TABLE comment;