CREATE DATABASE IF NOT EXISTS `blog_service` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;

USE blog_service;

DROP TABLE IF EXISTS `example`;
CREATE TABLE `example` (
    `id` INT(11) UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(32) NOT NULL DEFAULT '' COMMENT '名称',
    `alias` VARCHAR(32) NOT NULL DEFAULT '' COMMENT '别称',
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Hexagonal示例表';

CREATE TABLE IF NOT EXISTS `blog_tag` (
    `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(100) DEFAULT ''NOT NULL COMMENT '标签名称',
    `state` TINYINT(3) UNSIGNED DEFAULT '1' NOT NULL COMMENT '状态 0: 禁用, 1: 启用',
    `created_on` INT(10) UNSIGNED DEFAULT '0' NOT NULL COMMENT '创建时间',
    `created_by` VARCHAR(100) DEFAULT '' NOT NULL COMMENT '创建人',
    `modified_on` INT(10) UNSIGNED DEFAULT '0' NOT NULL COMMENT '修改时间',
    `modified_by` VARCHAR(100) DEFAULT '' NOT NULL COMMENT '修改人',
    `deleted_on` INT(10) UNSIGNED DEFAULT '0' NOT NULL COMMENT '删除时间',
    `is_del` TINYINT(3) UNSIGNED DEFAULT '0' NOT NULL COMMENT '是否删除 0: 未删除, 1: 已删除',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'created at',
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'updated at',
    `deleted_at` DATETIME COMMENT 'deleted at',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARACTER SET = utf8mb4 COMMENT = '标签表';

CREATE TABLE IF NOT EXISTS `blog_article` (
    `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `title` VARCHAR(100) DEFAULT '' COMMENT '文章标题',
    `desc` VARCHAR(255) DEFAULT '' COMMENT '文章简述',
    `cover_image_url` VARCHAR(255) DEFAULT '' COMMENT '封面图片地址',
    `content` LONGTEXT COMMENT '文章内容',
    `state` TINYINT(3) UNSIGNED DEFAULT '1' COMMENT '状态 0: 禁用, 1: 启用',
    `created_on` INT(10) UNSIGNED DEFAULT '0' COMMENT '创建时间',
    `created_by` VARCHAR(100) DEFAULT '' COMMENT '创建人',
    `modified_on` INT(10) UNSIGNED DEFAULT '0' COMMENT '修改时间',
    `modified_by` VARCHAR(100) DEFAULT '' COMMENT '修改人',
    `deleted_on` INT(10) UNSIGNED DEFAULT '0' COMMENT '删除时间',
    `is_del` TINYINT(3) UNSIGNED DEFAULT '0' COMMENT '是否删除 0: 未删除, 1: 已删除',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'created at',
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'updated at',
    `deleted_at` DATETIME COMMENT 'deleted at',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARACTER SET = utf8mb4 COMMENT = '文章表';

CREATE TABLE IF NOT EXISTS `blog_article_tag` (
    `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `article_id` INT(11) NOT NULL DEFAULT '0' COMMENT '文章ID',
    `tag_id` INT(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '标签ID',
    `created_on` INT(10) UNSIGNED DEFAULT '0' COMMENT '创建时间',
    `created_by` VARCHAR(100) DEFAULT '' COMMENT '创建人',
    `modified_on` INT(10) UNSIGNED DEFAULT '0' COMMENT '修改时间',
    `modified_by` VARCHAR(100) DEFAULT '' COMMENT '修改人',
    `deleted_on` INT(10) UNSIGNED DEFAULT '0' COMMENT '删除时间',
    `is_del` TINYINT(3) UNSIGNED DEFAULT '0' COMMENT '是否删除 0: 未删除, 1: 已删除',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARACTER SET = utf8mb4 COMMENT = '文章标签关联表';

CREATE TABLE IF NOT EXISTS `blog_auth` (
    `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `app_key` VARCHAR(20) NOT NULL DEFAULT '' COMMENT 'Key',
    `app_secret` VARCHAR(50) NOT NULL DEFAULT '' COMMENT 'Secret',
    `created_on` INT(10) UNSIGNED DEFAULT '0' COMMENT '创建时间',
    `created_by` VARCHAR(100) DEFAULT '' COMMENT '创建人',
    `modified_on` INT(10) UNSIGNED DEFAULT '0' COMMENT '修改时间',
    `modified_by` VARCHAR(100) DEFAULT '' COMMENT '修改人',
    `deleted_on` INT(10) UNSIGNED DEFAULT '0' COMMENT '删除时间',
    `is_del` TINYINT(3) UNSIGNED DEFAULT '0' COMMENT '是否删除 0: 未删除, 1: 已删除',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARACTER SET = utf8mb4 COMMENT = '认证表';
