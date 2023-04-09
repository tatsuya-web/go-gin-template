CREATE TABLE `posts`
(
    `id`       BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'postsの識別子',
    `title`    VARCHAR(128) NOT NULL COMMENT 'タスクのタイトル',
    `status`   VARCHAR(20)  NOT NULL COMMENT 'タスクの状態',
    `created`  DATETIME(6) NOT NULL COMMENT 'レコード作成日時',
    `modified` DATETIME(6) NOT NULL COMMENT 'レコード修正日時',
    PRIMARY KEY (`id`),
    CONSTRAINT `fk_user_id`
        FOREIGN KEY (`user_id`) REFERENCES `user` (`id`)
            ON DELETE RESTRICT ON UPDATE RESTRICT
) Engine=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='タスク';