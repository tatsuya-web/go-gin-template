-- SQL dump generated using DBML (dbml-lang.org)
-- Database: MySQL
-- Generated at: 2023-04-07T05:32:12.337Z

CREATE TABLE `users` (
  `id` bigint PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `email` varchar(256) NOT NULL COMMENT 'メールアドレス',
  `password` varchar(256) NOT NULL COMMENT 'パスワードハッシュ',
  `role` varchar(256) NOT NULL COMMENT 'ユーザー権限',
  `created_at` timestamp COMMENT '投稿時刻',
  `updated_at` timestamp COMMENT '更新時刻'
);

CREATE TABLE `posts` (
  `id` bigint PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `title` varchar(256) NOT NULL COMMENT '投稿のタイトル',
  `content` varchar(255) DEFAULT "" COMMENT '投稿内容',
  `user_id` bigint COMMENT '投稿ユーザー',
  `created_at` timestamp COMMENT '投稿時刻',
  `updated_at` timestamp COMMENT '更新時刻',
  PRIMARY KEY (`id`),
    CONSTRAINT `fk_user_id`
        FOREIGN KEY (`user_id`) REFERENCES `user` (`id`)
            ON DELETE CASCADE ON UPDATE CASCADE
);
