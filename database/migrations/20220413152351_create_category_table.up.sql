CREATE TABLE categories (
                            id         INT AUTO_INCREMENT NOT NULL,
                            name      VARCHAR(128) NOT NULL,
                            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                            PRIMARY KEY (`id`)
)