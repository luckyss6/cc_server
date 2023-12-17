SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

DROP TABLE IF EXISTS `user_info`;
CREATE TABLE `user_info`
(
    `id`         bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `created_at` datetime(3)  DEFAULT NULL,
    `updated_at` datetime(3)  DEFAULT NULL,
    `deleted_at` datetime(3)  DEFAULT NULL,
    `user_id`    varchar(32)         NOT NULL COMMENT '''用户id''',
    `password`   varchar(100)        NOT NULL COMMENT '''密码''',
    `nickname`   varchar(20)  DEFAULT NULL COMMENT '''昵称''',
    `phone`      varchar(20)  DEFAULT NULL COMMENT '''手机号''',
    `wechat`     varchar(20)  DEFAULT NULL COMMENT '''微信''',
    `email`      varchar(100) DEFAULT NULL COMMENT '''邮箱''',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

DROP TABLE IF EXISTS `user_role`;
CREATE TABLE `user_role`
(
    `id`         bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `created_at` datetime(3) DEFAULT NULL,
    `updated_at` datetime(3) DEFAULT NULL,
    `deleted_at` datetime(3) DEFAULT NULL,
    `user_id`    bigint(20) unsigned NOT NULL COMMENT '''用户id''',
    `role_id`    bigint(20) unsigned NOT NULL COMMENT '''角色id''',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

DROP TABLE IF EXISTS `role`;
CREATE TABLE `role`
(
    `id`          bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `created_at`  datetime(3)  DEFAULT NULL,
    `updated_at`  datetime(3)  DEFAULT NULL,
    `deleted_at`  datetime(3)  DEFAULT NULL,
    `name`        varchar(20)         NOT NULL COMMENT '''角色名称''',
    `description` varchar(100) DEFAULT NULL COMMENT '''角色描述''',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

DROP TABLE IF EXISTS `role_permission`;
CREATE TABLE `role_permission`
(
    `id`            bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `created_at`    datetime(3) DEFAULT NULL,
    `updated_at`    datetime(3) DEFAULT NULL,
    `deleted_at`    datetime(3) DEFAULT NULL,
    `role_id`       bigint(20) unsigned NOT NULL COMMENT '''角色id''',
    `permission_id` bigint(20) unsigned NOT NULL COMMENT '''权限id''',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

DROP TABLE IF EXISTS `permission`;
CREATE TABLE `permission`
(
    `id`          bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `created_at`  datetime(3)  DEFAULT NULL,
    `updated_at`  datetime(3)  DEFAULT NULL,
    `deleted_at`  datetime(3)  DEFAULT NULL,
    `name`        varchar(20)         NOT NULL COMMENT '''权限名称''',
    `description` varchar(100) DEFAULT NULL COMMENT '''权限描述''',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

DROP TABLE IF EXISTS `operation_log`;
CREATE TABLE `operation_log`
(
    `id`         bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `created_at` datetime(3)  DEFAULT NULL,
    `updated_at` datetime(3)  DEFAULT NULL,
    `deleted_at` datetime(3)  DEFAULT NULL,
    `user_id`    bigint(20) unsigned NOT NULL COMMENT '''用户id''',
    `path`       varchar(100) DEFAULT NULL COMMENT '''请求路径''',
    `method`     varchar(10)  DEFAULT NULL COMMENT '''请求方法''',
    `ip`         varchar(20)  DEFAULT NULL COMMENT '''请求ip''',
    `input`      text         DEFAULT NULL COMMENT '''请求参数''',
    `output`     text         DEFAULT NULL COMMENT '''响应参数''',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;



