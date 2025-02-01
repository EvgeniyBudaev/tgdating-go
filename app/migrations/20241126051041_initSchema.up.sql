CREATE SCHEMA IF NOT EXISTS dating;
CREATE EXTENSION IF NOT EXISTS postgis;

CREATE TABLE IF NOT EXISTS dating.profiles
(
    id               BIGSERIAL    NOT NULL PRIMARY KEY,
    telegram_user_id VARCHAR(255) NOT NULL UNIQUE,
    display_name     VARCHAR(255) NOT NULL,
    age              BIGINT       NOT NULL,
    gender           VARCHAR(50)  NOT NULL,
    description      TEXT,
    created_at       TIMESTAMP    NOT NULL,
    updated_at       TIMESTAMP    NOT NULL,
    last_online      TIMESTAMP    NOT NULL
);

CREATE TABLE IF NOT EXISTS dating.profile_payments
(
    id               BIGSERIAL    NOT NULL PRIMARY KEY,
    telegram_user_id VARCHAR(255) NOT NULL,
    price            VARCHAR(255) NOT NULL,
    currency         VARCHAR(255) NOT NULL,
    tariff           VARCHAR(255) NOT NULL,
    created_at       TIMESTAMP    NOT NULL,
    available_until       TIMESTAMP    NOT NULL,
    CONSTRAINT fk_profile_payments_telegram_user_id FOREIGN KEY (telegram_user_id) REFERENCES dating.profiles (telegram_user_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS dating.profile_statuses
(
    id                 BIGSERIAL    NOT NULL PRIMARY KEY,
    telegram_user_id   VARCHAR(255) NOT NULL UNIQUE,
    is_blocked         BOOL         NOT NULL,
    is_frozen          BOOL         NOT NULL,
    is_hidden_age      BOOL         NOT NULL,
    is_hidden_distance BOOL         NOT NULL,
    is_invisible       BOOL         NOT NULL,
    is_left_hand       BOOL         NOT NULL,
    created_at         TIMESTAMP    NOT NULL,
    updated_at         TIMESTAMP    NOT NULL,
    CONSTRAINT fk_profile_statuses_telegram_user_id FOREIGN KEY (telegram_user_id) REFERENCES dating.profiles (telegram_user_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS dating.profile_images
(
    id               BIGSERIAL    NOT NULL PRIMARY KEY,
    telegram_user_id VARCHAR(255) NOT NULL,
    name             VARCHAR(255),
    url              VARCHAR,
    size             BIGINT,
    created_at       TIMESTAMP    NOT NULL,
    updated_at       TIMESTAMP    NOT NULL,
    CONSTRAINT fk_profile_images_telegram_user_id FOREIGN KEY (telegram_user_id) REFERENCES dating.profiles (telegram_user_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS dating.profile_image_statuses
(
    id         BIGSERIAL NOT NULL PRIMARY KEY,
    image_id   BIGINT    NOT NULL,
    is_blocked BOOL      NOT NULL,
    is_primary BOOL      NOT NULL,
    is_private BOOL      NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    CONSTRAINT fk_profile_image_statuses_id FOREIGN KEY (image_id) REFERENCES dating.profile_images (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS dating.profile_navigators
(
    id               BIGSERIAL             NOT NULL PRIMARY KEY,
    telegram_user_id VARCHAR(255)          NOT NULL UNIQUE,
    country_code     VARCHAR(255),
    country_name     VARCHAR(255),
    city             VARCHAR(255),
    location         geometry(Point, 4326) NOT NULL,
    created_at       TIMESTAMP             NOT NULL,
    updated_at       TIMESTAMP             NOT NULL,
    CONSTRAINT fk_profile_navigators_telegram_user_id FOREIGN KEY (telegram_user_id) REFERENCES dating.profiles (telegram_user_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS dating.profile_filters
(
    id               BIGSERIAL    NOT NULL PRIMARY KEY,
    telegram_user_id VARCHAR(255) NOT NULL UNIQUE,
    search_gender    VARCHAR(50)  NOT NULL,
    age_from         INTEGER      NOT NULL,
    age_to           INTEGER      NOT NULL,
    distance         REAL         NOT NULL,
    page             INTEGER      NOT NULL,
    size             INTEGER      NOT NULL,
    created_at       TIMESTAMP    NOT NULL,
    updated_at       TIMESTAMP    NOT NULL,
    CONSTRAINT fk_profile_filters_telegram_user_id FOREIGN KEY (telegram_user_id) REFERENCES dating.profiles (telegram_user_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS dating.profile_telegrams
(
    id                 BIGSERIAL    NOT NULL PRIMARY KEY,
    user_id            VARCHAR(255) NOT NULL UNIQUE,
    username           VARCHAR(255) NOT NULL,
    first_name         VARCHAR(255),
    last_name          VARCHAR(255),
    language_code      VARCHAR(50),
    allows_write_to_pm BOOL         NOT NULL,
    query_id           TEXT,
    created_at         TIMESTAMP    NOT NULL,
    updated_at         TIMESTAMP    NOT NULL,
    CONSTRAINT fk_profile_telegram_telegram_user_id FOREIGN KEY (user_id) REFERENCES dating.profiles (telegram_user_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS dating.profile_likes
(
    id                     BIGSERIAL    NOT NULL PRIMARY KEY,
    telegram_user_id       VARCHAR(255) NOT NULL,
    liked_telegram_user_id VARCHAR(255) NOT NULL,
    is_liked               BOOL         NOT NULL,
    created_at             TIMESTAMP    NOT NULL,
    updated_at             TIMESTAMP    NOT NULL,
    CONSTRAINT fk_profile_likes_telegram_user_id FOREIGN KEY (telegram_user_id) REFERENCES dating.profiles (telegram_user_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS dating.profile_blocks
(
    id                       BIGSERIAL    NOT NULL PRIMARY KEY,
    telegram_user_id         VARCHAR(255) NOT NULL,
    blocked_telegram_user_id VARCHAR(255) NOT NULL,
    is_blocked               BOOL         NOT NULL,
    created_at               TIMESTAMP    NOT NULL,
    updated_at               TIMESTAMP    NOT NULL,
    CONSTRAINT fk_profile_blocks_telegram_user_id FOREIGN KEY (telegram_user_id) REFERENCES dating.profiles (telegram_user_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS dating.profile_complaints
(
    id                        BIGSERIAL    NOT NULL PRIMARY KEY,
    telegram_user_id          VARCHAR(255) NOT NULL,
    criminal_telegram_user_id VARCHAR(255) NOT NULL,
    reason                    VARCHAR(255) NOT NULL,
    created_at                TIMESTAMP    NOT NULL,
    updated_at                TIMESTAMP    NOT NULL,
    CONSTRAINT fk_profile_complaints_telegram_user_id FOREIGN KEY (telegram_user_id) REFERENCES dating.profiles (telegram_user_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS dating.profile_settings
(
    id                        BIGSERIAL    NOT NULL PRIMARY KEY,
    telegram_user_id          VARCHAR(255) NOT NULL,
    measurement               VARCHAR(255) NOT NULL,
    created_at                TIMESTAMP    NOT NULL,
    updated_at                TIMESTAMP    NOT NULL,
    CONSTRAINT fk_profile_settings_telegram_user_id FOREIGN KEY (telegram_user_id) REFERENCES dating.profiles (telegram_user_id) ON DELETE CASCADE
);