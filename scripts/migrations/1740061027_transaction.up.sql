CREATE TABLE IF NOT EXISTS tb_transaction (
    id                          INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY
    , created_at                DATETIME NOT NULL DEFAULT NOW()
    , updated_at                DATETIME NOT NULL DEFAULT NOW() ON UPDATE NOW()
    , deleted_at                DATETIME NOT NULL DEFAULT '0000-00-00 00:00:00'
    , uuid                      CHAR(36) NOT NULL
    , account_id                INT UNSIGNED NOT NULL
    , status                    ENUM('pending', 'approved', 'denied') NOT NULL DEFAULT 'pending'
    , name                      VARCHAR(100) NOT NULL
    , credit_card_number        CHAR(16) NOT NULL
    , credit_card_security_code CHAR(3) NOT NULL
    , credit_card_expires       DATE NOT NULL
    , amount                    DECIMAL(12,2)

    , UNIQUE INDEX `uuid_uidx` (`uuid` ASC)
    
    , CONSTRAINT fk_transaction_account_id
        FOREIGN KEY (account_id)
        REFERENCES tb_account (id)
        ON DELETE NO ACTION
        ON UPDATE NO ACTION
);