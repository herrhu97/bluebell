drop table if exists `user`;
create table `user` (
                        `id` bigint(20) not null auto_increment,
                        `user_id` bigint(20) not null,
                        `username` varbinary(64) not null,
                        `password` varbinary(64) not null,
                        `email` varchar(64),
                        `gender` tinyint(4) not null default '0',
                        `create_time` timestamp null default current_timestamp,
                        `update_time` timestamp null default current_timestamp on update current_timestamp,
                        primary key (`id`),
                        unique key `idx_username` (`username`) using btree,
                        unique key `idx_user_id` (`user_id`) using btree
)engine=InnoDB default character set=utf8mb4 collate=utf8mb4_general_ci;