-- +goose Up
CREATE TABLE "user"
(
    id         bigserial PRIMARY KEY,
    full_name  text        NOT NULL,
    email      text UNIQUE NOT NULL,
    caption    text        NOT NULL,
    birth_date date        NOT NULL,
    created_at timestamp   NOT NULL,
    deleted_at timestamp
);

CREATE TABLE "follower"
(
    follower_id bigserial REFERENCES "user" (id),
    followee_id  bigserial REFERENCES "user" (id),

    created_at  timestamp NOT NULL,
    deleted_at  timestamp
);

CREATE TABLE "tweet"
(
    id         bigserial PRIMARY KEY,
    user_id    bigserial REFERENCES "user" (id),
    text       text      NOT NULL,
    likes      integer   NOT NULL DEFAULT 0,
    created_at timestamp NOT NULL,
    deleted_at timestamp
);

CREATE TABLE "comment"
(
    id         bigserial PRIMARY KEY,
    user_id    bigserial REFERENCES "user" (id),
    tweet_id   bigserial REFERENCES "tweet" (id),
    text       text      NOT NULL,
    likes      integer   NOT NULL DEFAULT 0,
    created_at timestamp NOT NULL,
    deleted_at timestamp
);

-- +goose Down

DROP TABLE "user";
DROP TABLE "follower";
DROP TABLE "tweet";
DROP TABLE "comment";