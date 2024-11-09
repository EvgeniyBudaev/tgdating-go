CREATE TABLE IF NOT EXISTS profiles
(
    id               BIGSERIAL    NOT NULL,
    session_id       VARCHAR      NOT NULL UNIQUE PRIMARY KEY,
    display_name     VARCHAR(255) NOT NULL,
    birthday         DATE         NOT NULL,
    gender           VARCHAR(100),
    location         TEXT,
    description      TEXT,
    height           REAL,
    weight           REAL,
    is_deleted       BOOL         NOT NULL DEFAULT false,
    is_blocked       BOOL         NOT NULL DEFAULT false,
    is_premium       BOOL         NOT NULL DEFAULT false,
    is_show_distance BOOL         NOT NULL DEFAULT true,
    is_invisible     BOOL         NOT NULL DEFAULT false,
    created_at       TIMESTAMP             DEFAULT CURRENT_TIMESTAMP,
    updated_at       TIMESTAMP    NULL,
    last_online      TIMESTAMP             DEFAULT CURRENT_TIMESTAMP CHECK (last_online >= created_at)
);

CREATE TABLE IF NOT EXISTS profile_images
(
    id         BIGSERIAL NOT NULL PRIMARY KEY,
    session_id VARCHAR   NOT NULL,
    name       VARCHAR(255),
    url        VARCHAR,
    size       BIGINT,
    is_deleted BOOL      NOT NULL DEFAULT false,
    is_blocked BOOL      NOT NULL DEFAULT false,
    is_primary BOOL      NOT NULL DEFAULT false,
    is_private BOOL      NOT NULL DEFAULT false,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL,
    CONSTRAINT fk_profile_images_session_id FOREIGN KEY (session_id) REFERENCES profiles (session_id)
);

CREATE EXTENSION IF NOT EXISTS postgis;

CREATE TABLE IF NOT EXISTS profile_navigators
(
    id         BIGSERIAL NOT NULL PRIMARY KEY,
    session_id VARCHAR   NOT NULL,
    location   geometry(Point, 4326),
    is_deleted BOOL      NOT NULL DEFAULT false,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL,
    CONSTRAINT fk_profile_navigators_session_id FOREIGN KEY (session_id) REFERENCES profiles (session_id)
);

CREATE TABLE IF NOT EXISTS profile_filters
(
    id            BIGSERIAL NOT NULL PRIMARY KEY,
    session_id    VARCHAR   NOT NULL,
    search_gender VARCHAR(100),
    looking_for   VARCHAR(100),
    age_from      INTEGER,
    age_to        INTEGER,
    distance      REAL,
    page          INTEGER,
    size          INTEGER,
    is_deleted    BOOL      NOT NULL DEFAULT false,
    created_at    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at    TIMESTAMP NULL,
    CONSTRAINT fk_profile_filters_session_id FOREIGN KEY (session_id) REFERENCES profiles (session_id)
);

CREATE TABLE IF NOT EXISTS profile_telegrams
(
    id                 BIGSERIAL    NOT NULL PRIMARY KEY,
    session_id         VARCHAR      NOT NULL UNIQUE,
    user_id            BIGINT       NOT NULL,
    username           VARCHAR(255) NOT NULL,
    first_name         VARCHAR(255),
    last_name          VARCHAR(255),
    language_code      VARCHAR,
    allows_write_to_pm BOOL,
    query_id           TEXT,
    is_deleted         BOOL         NOT NULL DEFAULT false,
    created_at         TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at         TIMESTAMP    NULL,
    CONSTRAINT fk_profile_telegram_session_id FOREIGN KEY (session_id) REFERENCES profiles (session_id)
);

CREATE TABLE IF NOT EXISTS profile_likes
(
    id               BIGSERIAL NOT NULL PRIMARY KEY,
    session_id       VARCHAR   NOT NULL,
    liked_session_id VARCHAR   NOT NULL,
    is_liked         BOOL      NOT NULL DEFAULT false,
    is_deleted       BOOL      NOT NULL DEFAULT false,
    created_at       TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at       TIMESTAMP NULL,
    CONSTRAINT fk_profile_likes_profile_id FOREIGN KEY (session_id) REFERENCES profiles (session_id)
);

CREATE TABLE IF NOT EXISTS profile_blocks
(
    id                      BIGSERIAL NOT NULL PRIMARY KEY,
    session_id              VARCHAR   NOT NULL,
    blocked_user_session_id VARCHAR   NOT NULL,
    is_blocked              BOOL      NOT NULL DEFAULT false,
    created_at              TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at              TIMESTAMP NULL,
    CONSTRAINT fk_profile_blocks_session_id FOREIGN KEY (session_id) REFERENCES profiles (session_id)
);

CREATE TABLE IF NOT EXISTS profile_complaints
(
    id                  BIGSERIAL NOT NULL PRIMARY KEY,
    session_id          VARCHAR   NOT NULL,
    criminal_session_id VARCHAR   NOT NULL,
    reason              TEXT,
    is_deleted          BOOL      NOT NULL DEFAULT false,
    created_at          TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMP NULL,
    CONSTRAINT fk_profile_complaints_session_id FOREIGN KEY (session_id) REFERENCES profiles (session_id)
);