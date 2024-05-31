/*
 Navicat Premium Data Transfer

 Source Server         : 本地mysql
 Source Server Type    : MySQL
 Source Server Version : 80035
 Source Host           : localhost:3306
 Source Schema         : goadmin

 Target Server Type    : MySQL
 Target Server Version : 80035
 File Encoding         : 65001

 Date: 31/05/2024 14:21:48
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for goadmin_menu
-- ----------------------------
CREATE TABLE `goadmin_menu`  (
                                 `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
                                 `parent_id` int UNSIGNED NOT NULL DEFAULT 0,
                                 `type` tinyint UNSIGNED NOT NULL DEFAULT 0,
                                 `order` int UNSIGNED NOT NULL DEFAULT 0,
                                 `title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
                                 `icon` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
                                 `uri` varchar(3000) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
                                 `header` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
                                 `plugin_name` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
                                 `uuid` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
                                 `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                                 `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                                 PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for goadmin_operation_log
-- ----------------------------
CREATE TABLE `goadmin_operation_log`  (
                                          `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
                                          `user_id` int UNSIGNED NOT NULL,
                                          `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
                                          `method` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
                                          `ip` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
                                          `input` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
                                          `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                                          `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                                          PRIMARY KEY (`id`) USING BTREE,
                                          INDEX `admin_operation_log_user_id_index`(`user_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for goadmin_permissions
-- ----------------------------
CREATE TABLE `goadmin_permissions`  (
                                        `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
                                        `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
                                        `slug` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
                                        `http_method` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
                                        `http_path` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
                                        `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                                        `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                                        PRIMARY KEY (`id`) USING BTREE,
                                        UNIQUE INDEX `admin_permissions_name_unique`(`name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for goadmin_role_menu
-- ----------------------------
CREATE TABLE `goadmin_role_menu`  (
                                      `role_id` int UNSIGNED NOT NULL,
                                      `menu_id` int UNSIGNED NOT NULL,
                                      `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                                      `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                                      INDEX `admin_role_menu_role_id_menu_id_index`(`role_id`, `menu_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for goadmin_role_permissions
-- ----------------------------
CREATE TABLE `goadmin_role_permissions`  (
                                             `role_id` int UNSIGNED NOT NULL,
                                             `permission_id` int UNSIGNED NOT NULL,
                                             `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                                             `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                                             UNIQUE INDEX `admin_role_permissions`(`role_id`, `permission_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for goadmin_role_users
-- ----------------------------
CREATE TABLE `goadmin_role_users`  (
                                       `role_id` int UNSIGNED NOT NULL,
                                       `user_id` int UNSIGNED NOT NULL,
                                       `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                                       `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                                       UNIQUE INDEX `admin_user_roles`(`role_id`, `user_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for goadmin_roles
-- ----------------------------
CREATE TABLE `goadmin_roles`  (
                                  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
                                  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
                                  `slug` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
                                  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                                  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                                  PRIMARY KEY (`id`) USING BTREE,
                                  UNIQUE INDEX `admin_roles_name_unique`(`name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for goadmin_session
-- ----------------------------
CREATE TABLE `goadmin_session`  (
                                    `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
                                    `sid` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
                                    `values` varchar(3000) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
                                    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                                    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                                    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for goadmin_site
-- ----------------------------
CREATE TABLE `goadmin_site`  (
                                 `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
                                 `key` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
                                 `value` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL,
                                 `description` varchar(3000) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
                                 `state` tinyint UNSIGNED NOT NULL DEFAULT 0,
                                 `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                                 `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                                 PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for goadmin_user_permissions
-- ----------------------------
CREATE TABLE `goadmin_user_permissions`  (
                                             `user_id` int UNSIGNED NOT NULL,
                                             `permission_id` int UNSIGNED NOT NULL,
                                             `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                                             `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                                             UNIQUE INDEX `admin_user_permissions`(`user_id`, `permission_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for goadmin_users
-- ----------------------------
CREATE TABLE `goadmin_users`  (
                                  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
                                  `username` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
                                  `password` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
                                  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
                                  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
                                  `remember_token` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
                                  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                                  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                                  PRIMARY KEY (`id`) USING BTREE,
                                  UNIQUE INDEX `admin_users_username_unique`(`username`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;


SET FOREIGN_KEY_CHECKS = 1;
