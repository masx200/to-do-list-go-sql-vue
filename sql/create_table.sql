create database if not exists todolist;

use todolist;

CREATE TABLE if not exists `to_do_items`(
    `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
    `created_at` datetime(3) NULL DEFAULT NULL,
    `updated_at` datetime(3) NULL DEFAULT NULL,
    `deleted_at` datetime(3) NULL DEFAULT NULL,
    `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
    `completed` tinyint(1) NOT NULL,
    `author` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
    PRIMARY KEY (`id`) USING BTREE,
    INDEX `idx_to_do_items_completed`(`completed` ASC) USING BTREE,
    INDEX `idx_to_do_items_deleted_at`(`deleted_at` ASC) USING BTREE,
    INDEX `idx_to_do_items_author`(`author` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 20 DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;