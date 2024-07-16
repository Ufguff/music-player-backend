CREATE TABLE IF NOT EXISTS tracks (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` varchar(255) NOT NULL,
    `authorID` INT UNSIGNED NOT NULL,
    `recordPath` varchar(255) NOT NULL,
    `imagePath` varchar(255) NOT NULL,


    PRIMARY KEY (`id`),
    FOREIGN KEY(`authorID`) REFERENCES authors(`id`)
);
