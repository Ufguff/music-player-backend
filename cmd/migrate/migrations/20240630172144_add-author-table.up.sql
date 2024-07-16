CREATE TABLE IF NOT EXISTS authors (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` varchar(255) NOT NULL,
    `image` BLOB,

    PRIMARY KEY (id)
);
