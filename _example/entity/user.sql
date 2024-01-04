
CREATE TABLE `user1` (
`id` int NOT NULL AUTO_INCREMENT,
`user_name` varchar(50) DEFAULT NULL,
`age` int DEFAULT NULL,
`CITY_CODE` varchar(50) comment 'the city code of user',
ts timestamp not null default current_timestamp comment 'the time of writing time',
PRIMARY KEY (`id`)
) ENGINE=InnoDB;

insert into user1(user_name,age,CITY_CODE) values('user1',1,'10000'),(null,null,null),('user3',18,'20000');



