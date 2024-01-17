CREATE TABLE `tb_user` (
       `id` int unsigned NOT NULL AUTO_INCREMENT,
       `name` varchar(50) DEFAULT NULL,
       `age` int DEFAULT NULL,
       `createtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
       PRIMARY KEY (`id`)
) ENGINE=InnoDB