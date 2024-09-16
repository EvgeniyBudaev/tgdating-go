CREATE TABLE IF NOT EXISTS profiles
(
    id               BIGSERIAL    NOT NULL,
    session_id       VARCHAR      NOT NULL UNIQUE PRIMARY KEY,
    display_name     VARCHAR(255) NOT NULL,
    birthday         DATE         NOT NULL,
    gender           VARCHAR(100),
    location         TEXT,
    description      TEXT,
    height           REAL         NOT NULL DEFAULT 0.0 CHECK (height >= 0),
    weight           REAL         NOT NULL DEFAULT 0.0 CHECK (weight >= 0),
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