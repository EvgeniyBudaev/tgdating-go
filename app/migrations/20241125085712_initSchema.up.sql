CREATE SCHEMA IF NOT EXISTS dating;

CREATE TABLE IF NOT EXISTS dating.profiles
(
    id               BIGSERIAL    NOT NULL,
    session_id       VARCHAR(255) NOT NULL UNIQUE PRIMARY KEY,
    display_name     VARCHAR(255) NOT NULL,
    birthday         DATE         NOT NULL,
    gender           VARCHAR(100) NOT NULL,
    location         TEXT,
    description      TEXT,
    height           REAL         NOT NULL DEFAULT 0.0 CHECK (height >= 0),
    weight           REAL         NOT NULL DEFAULT 0.0 CHECK (weight >= 0),
    is_deleted       BOOL         NOT NULL DEFAULT false,
    is_blocked       BOOL         NOT NULL DEFAULT false,
    is_premium       BOOL         NOT NULL DEFAULT false,
    is_show_distance BOOL         NOT NULL DEFAULT true,
    is_invisible     BOOL         NOT NULL DEFAULT false,
    created_at       TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at       TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    last_online      TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP CHECK (last_online >= created_at)
    );

CREATE TABLE IF NOT EXISTS dating.profile_images
(
    id         BIGSERIAL    NOT NULL PRIMARY KEY,
    session_id VARCHAR(255) NOT NULL,
    name       VARCHAR(255),
    url        VARCHAR,
    size       BIGINT,
    is_deleted BOOL         NOT NULL DEFAULT false,
    is_blocked BOOL         NOT NULL DEFAULT false,
    is_primary BOOL         NOT NULL DEFAULT false,
    is_private BOOL         NOT NULL DEFAULT false,
    created_at       TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at       TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_profile_images_session_id FOREIGN KEY (session_id) REFERENCES dating.profiles (session_id) ON DELETE CASCADE
    );

CREATE EXTENSION IF NOT EXISTS postgis SCHEMA dating;

CREATE TABLE IF NOT EXISTS dating.profile_navigators
(
    id         BIGSERIAL    NOT NULL PRIMARY KEY,
    session_id VARCHAR(255) NOT NULL,
    location   geometry(Point, 4326),
    created_at       TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at       TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_profile_navigators_session_id FOREIGN KEY (session_id) REFERENCES dating.profiles (session_id) ON DELETE CASCADE
    );

CREATE TABLE IF NOT EXISTS dating.profile_filters
(
    id            BIGSERIAL    NOT NULL PRIMARY KEY,
    session_id    VARCHAR(255) NOT NULL,
    search_gender VARCHAR(100) NOT NULL,
    looking_for   VARCHAR(100) NOT NULL,
    age_from      INTEGER      NOT NULL DEFAULT 18,
    age_to        INTEGER      NOT NULL DEFAULT 100,
    distance      REAL         NOT NULL DEFAULT 0,
    page          INTEGER      NOT NULL DEFAULT 1,
    size          INTEGER      NOT NULL DEFAULT 1,
    created_at       TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at       TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_profile_filters_session_id FOREIGN KEY (session_id) REFERENCES dating.profiles (session_id) ON DELETE CASCADE
    );

CREATE TABLE IF NOT EXISTS dating.profile_telegrams
(
    id                 BIGSERIAL    NOT NULL PRIMARY KEY,
    session_id         VARCHAR(255) NOT NULL UNIQUE,
    user_id            BIGINT       NOT NULL,
    username           VARCHAR(255) NOT NULL,
    first_name         VARCHAR(255),
    last_name          VARCHAR(255),
    language_code      VARCHAR(255),
    allows_write_to_pm BOOL                  DEFAULT false,
    query_id           TEXT,
    created_at       TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at       TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_profile_telegram_session_id FOREIGN KEY (session_id) REFERENCES dating.profiles (session_id) ON DELETE CASCADE
    );

CREATE TABLE IF NOT EXISTS dating.profile_likes
(
    id               BIGSERIAL    NOT NULL PRIMARY KEY,
    session_id       VARCHAR(255) NOT NULL,
    liked_session_id VARCHAR(255) NOT NULL,
    is_liked         BOOL         NOT NULL DEFAULT false,
    created_at       TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at       TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_profile_likes_profile_id FOREIGN KEY (session_id) REFERENCES dating.profiles (session_id) ON DELETE CASCADE
    );

CREATE TABLE IF NOT EXISTS dating.profile_blocks
(
    id                      BIGSERIAL    NOT NULL PRIMARY KEY,
    session_id              VARCHAR(255) NOT NULL,
    blocked_user_session_id VARCHAR(255) NOT NULL,
    is_blocked              BOOL         NOT NULL DEFAULT false,
    created_at       TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at       TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_profile_blocks_session_id FOREIGN KEY (session_id) REFERENCES dating.profiles (session_id) ON DELETE CASCADE
    );

CREATE TABLE IF NOT EXISTS dating.profile_complaints
(
    id                  BIGSERIAL    NOT NULL PRIMARY KEY,
    session_id          VARCHAR(255) NOT NULL,
    criminal_session_id VARCHAR(255) NOT NULL,
    reason              VARCHAR(255) NOT NULL,
    created_at       TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at       TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_profile_complaints_session_id FOREIGN KEY (session_id) REFERENCES dating.profiles (session_id) ON DELETE CASCADE
    );