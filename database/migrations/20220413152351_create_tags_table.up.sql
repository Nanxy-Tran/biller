CREATE TABLE tags (
                            id         INT AUTO_INCREMENT NOT NULL,
                            name      VARCHAR(128) NOT NULL,
                            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                            INDEX (id)
) ENGINE=INNODB