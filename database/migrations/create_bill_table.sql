DROP TABLE IF EXISTS bills;
CREATE TABLE bills (
                       id         INT AUTO_INCREMENT NOT NULL,
                       name      VARCHAR(128) NOT NULL,
                       amount     INTEGER NOT NULL,
                       description   VARCHAR(255) DEFAULT 'description...',
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       PRIMARY KEY (`id`)
);

INSERT INTO bills
(name, amount, description)
VALUES
    ('Cafe Laviet', 45000, ''),
    ('Cafe G', 35000, 'Buoi chieu toi')
;