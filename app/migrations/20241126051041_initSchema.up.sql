CREATE SCHEMA IF NOT EXISTS dating;
CREATE EXTENSION IF NOT EXISTS postgis;

CREATE TABLE IF NOT EXISTS dating.profiles
(
    id               BIGSERIAL    NOT NULL,
    session_id       VARCHAR(255) NOT NULL UNIQUE PRIMARY KEY,
    display_name     VARCHAR(255) NOT NULL,
    birthday         DATE         NOT NULL,
    gender           VARCHAR(100) NOT NULL,
    location         TEXT,
    description      TEXT,
    height           REAL         NOT NULL,
    weight           REAL         NOT NULL,
    is_frozen       BOOL         NOT NULL,
    is_blocked       BOOL         NOT NULL,
    is_premium       BOOL         NOT NULL,
    is_show_distance BOOL         NOT NULL,
    is_invisible     BOOL         NOT NULL,
    created_at       TIMESTAMP    NOT NULL,
    updated_at       TIMESTAMP    NOT NULL,
    last_online      TIMESTAMP    NOT NULL
);

CREATE TABLE IF NOT EXISTS dating.profile_images
(
    id         BIGSERIAL    NOT NULL PRIMARY KEY,
    session_id VARCHAR(255) NOT NULL,
    name       VARCHAR(255),
    url        VARCHAR,
    size       BIGINT,
    is_blocked BOOL         NOT NULL,
    is_primary BOOL         NOT NULL,
    is_private BOOL         NOT NULL,
    created_at TIMESTAMP    NOT NULL,
    updated_at TIMESTAMP    NOT NULL,
    CONSTRAINT fk_profile_images_session_id FOREIGN KEY (session_id) REFERENCES dating.profiles (session_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS dating.profile_navigators
(
    id         BIGSERIAL             NOT NULL PRIMARY KEY,
    session_id VARCHAR(255)          NOT NULL,
    location   geometry(Point, 4326) NOT NULL,
    created_at TIMESTAMP             NOT NULL,
    updated_at TIMESTAMP             NOT NULL,
    CONSTRAINT fk_profile_navigators_session_id FOREIGN KEY (session_id) REFERENCES dating.profiles (session_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS dating.profile_filters
(
    id            BIGSERIAL    NOT NULL PRIMARY KEY,
    session_id    VARCHAR(255) NOT NULL,
    search_gender VARCHAR(100) NOT NULL,
    looking_for   VARCHAR(100) NOT NULL,
    age_from      INTEGER      NOT NULL,
    age_to        INTEGER      NOT NULL,
    distance      REAL         NOT NULL,
    page          INTEGER      NOT NULL,
    size          INTEGER      NOT NULL,
    created_at    TIMESTAMP    NOT NULL,
    updated_at    TIMESTAMP    NOT NULL,
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
    language_code      VARCHAR(50),
    allows_write_to_pm BOOL         NOT NULL,
    query_id           TEXT,
    created_at         TIMESTAMP    NOT NULL,
    updated_at         TIMESTAMP    NOT NULL,
    CONSTRAINT fk_profile_telegram_session_id FOREIGN KEY (session_id) REFERENCES dating.profiles (session_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS dating.profile_likes
(
    id               BIGSERIAL    NOT NULL PRIMARY KEY,
    session_id       VARCHAR(255) NOT NULL,
    liked_session_id VARCHAR(255) NOT NULL,
    is_liked         BOOL         NOT NULL,
    created_at       TIMESTAMP    NOT NULL,
    updated_at       TIMESTAMP    NOT NULL,
    CONSTRAINT fk_profile_likes_profile_id FOREIGN KEY (session_id) REFERENCES dating.profiles (session_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS dating.profile_blocks
(
    id                      BIGSERIAL    NOT NULL PRIMARY KEY,
    session_id              VARCHAR(255) NOT NULL,
    blocked_user_session_id VARCHAR(255) NOT NULL,
    is_blocked              BOOL         NOT NULL,
    created_at              TIMESTAMP    NOT NULL,
    updated_at              TIMESTAMP    NOT NULL,
    CONSTRAINT fk_profile_blocks_session_id FOREIGN KEY (session_id) REFERENCES dating.profiles (session_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS dating.profile_complaints
(
    id                  BIGSERIAL    NOT NULL PRIMARY KEY,
    session_id          VARCHAR(255) NOT NULL,
    criminal_session_id VARCHAR(255) NOT NULL,
    reason              VARCHAR(255) NOT NULL,
    created_at          TIMESTAMP    NOT NULL,
    updated_at          TIMESTAMP    NOT NULL,
    CONSTRAINT fk_profile_complaints_session_id FOREIGN KEY (session_id) REFERENCES dating.profiles (session_id) ON DELETE CASCADE
);