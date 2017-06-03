
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE posts (
    uid serial NOT NULL,
    title text,
    body text,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    CONSTRAINT posts_pkey PRIMARY KEY (uid)
);

INSERT INTO "public"."posts"("title", "body") VALUES('post 1', 'body 1') RETURNING "uid", "title", "body", "created_at", "updated_at";
INSERT INTO "public"."posts"("title", "body") VALUES('post 2', 'body 2') RETURNING "uid", "title", "body", "created_at", "updated_at";


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE posts;