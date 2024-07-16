CREATE TABLE IF NOT EXISTS albums (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` varchar(255) NOT NULL,
    `authorID` INT UNSIGNED NOT NULL,
    `image` BLOB,

    PRIMARY KEY(`id`),
    FOREIGN KEY(`authorID`) REFERENCES authors(`id`)
);
