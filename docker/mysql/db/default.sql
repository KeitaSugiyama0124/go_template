CREATE DATABASE `default` DEFAULT CHARACTER SET utf8mb4;

USE `default`;


DROP TABLE IF EXISTS `contents_info`;

CREATE TABLE `contents_info` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `content_key` text CHARACTER SET utf8mb4 COLLATE utf8mb4_bin,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

LOCK TABLES `contents_info` WRITE;
/*!40000 ALTER TABLE `contents_info` DISABLE KEYS */;

INSERT INTO `contents_info` (`id`, `content_key`)
VALUES
	(1,X'414E4E');

/*!40000 ALTER TABLE `contents_info` ENABLE KEYS */;
UNLOCK TABLES;