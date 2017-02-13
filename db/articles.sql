CREATE TABLE `articles` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `title` varchar(255) NOT NULL COMMENT 'title',
  `body` varchar(255) NOT NULL COMMENT 'body',
  `created` timestamp NOT NULL DEFAULT NOW() COMMENT 'when created',
  `updated` timestamp NOT NULL DEFAULT NOW() ON UPDATE CURRENT_TIMESTAMP COMMENT 'when last updated',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COMMENT='list of articles';
