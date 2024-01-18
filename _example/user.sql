CREATE TABLE `tb_user`
(
    `id`         int unsigned NOT NULL AUTO_INCREMENT,
    `name`       varchar(50)        DEFAULT NULL,
    `age`        int                DEFAULT NULL,
    `createtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);
insert into tb_user (name,age) values("user1",1);
insert into tb_user (name,age) values("user10",10);
insert into tb_user (name,age) values("user2",20);
insert into tb_user (name,age) values("user3",30);
insert into tb_user (name,age) values("user4",40);
insert into tb_user (name,age) values("user5",50);