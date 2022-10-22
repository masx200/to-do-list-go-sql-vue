-- Active: 1666419568297@@127.0.0.1@3306@todolist
-- 创建库
create database if not exists todolist;

-- 切换库
use todolist;

CREATE TABLE  if not exists `to_do_items`(
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `created_at` datetime(3) DEFAULT NULL,
    `updated_at` datetime(3) DEFAULT NULL,
    `deleted_at` datetime(3) DEFAULT NULL,
    `content` longtext NOT NULL,
    `finished` tinyint(1) DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `idx_to_do_items_deleted_at` (`deleted_at`)
) ENGINE = InnoDB AUTO_INCREMENT = 20 DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;