CREATE TABLE IF NOT EXISTS tb_account (
    id                  INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY
    , created_at        DATETIME NOT NULL DEFAULT NOW()
    , updated_at        DATETIME NOT NULL DEFAULT NOW() ON UPDATE NOW()
    , deleted_at        DATETIME NOT NULL DEFAULT '0000-00-00 00:00:00'
    , uuid              CHAR(36) NOT NULL
    , token             CHAR(36) NOT NULL

    , UNIQUE INDEX `uuid_uidx` (`uuid` ASC)
    , UNIQUE INDEX `token_uidx` (`token` ASC)
);