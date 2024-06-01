SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for category
-- ----------------------------
CREATE TABLE `category`  (
                             `id` int NOT NULL AUTO_INCREMENT,
                             `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                             `parent_id` int NULL DEFAULT NULL,
                             PRIMARY KEY (`id`) USING BTREE,
                             INDEX `CATEGORY_parent_id_fkey`(`parent_id`) USING BTREE,
                             CONSTRAINT `CATEGORY_parent_id_fkey` FOREIGN KEY (`parent_id`) REFERENCES `category` (`id`) ON DELETE SET NULL ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 30 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for group
-- ----------------------------
CREATE TABLE `group`  (
                          `id` int NOT NULL AUTO_INCREMENT,
                          `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
                          `desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
                          PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for todo
-- ----------------------------
CREATE TABLE `todo`  (
                         `id` int NOT NULL AUTO_INCREMENT,
                         `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
                         `score` int NOT NULL DEFAULT 0,
                         `amount` float NOT NULL DEFAULT 0,
                         `status` enum('PENDING','PAUSED','COMPLETED','PROCESSING') CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'PENDING',
                         `createdAt` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         `updatedAt` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         `deadline` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         `priority` enum('HIGH','MEDIUM','LOW') CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'LOW',
                         `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
                         `category_id` int NULL DEFAULT 1,
                         PRIMARY KEY (`id`) USING BTREE,
                         INDEX `Todo_category_id_fkey`(`category_id`) USING BTREE,
                         CONSTRAINT `Todo_category_id_fkey` FOREIGN KEY (`category_id`) REFERENCES `category` (`id`) ON DELETE SET NULL ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for todo_detail
-- ----------------------------
CREATE TABLE `todo_detail`  (
                                `id` int NOT NULL AUTO_INCREMENT,
                                `desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
                                `img_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
                                `todo_id` int NOT NULL,
                                PRIMARY KEY (`id`) USING BTREE,
                                UNIQUE INDEX `Detail_todo_id_key`(`todo_id`) USING BTREE,
                                CONSTRAINT `Detail_todo_id_fkey` FOREIGN KEY (`todo_id`) REFERENCES `todo` (`id`) ON DELETE CASCADE ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for todo_tag
-- ----------------------------
CREATE TABLE `todo_tag`  (
                             `id` int NOT NULL AUTO_INCREMENT,
                             `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                             `todo_id` int NOT NULL,
                             `createdAt` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                             PRIMARY KEY (`id`) USING BTREE,
                             UNIQUE INDEX `tag_todo_id_key`(`name`, `todo_id`) USING BTREE,
                             INDEX `Todo_tag_id_fkey`(`todo_id`) USING BTREE,
                             CONSTRAINT `Todo_tag_id_fkey` FOREIGN KEY (`todo_id`) REFERENCES `todo` (`id`) ON DELETE CASCADE ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 50 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for todosongroups
-- ----------------------------
CREATE TABLE `todosongroups`  (
                                  `todo_id` int NOT NULL,
                                  `group_id` int NOT NULL,
                                  `assignedAt` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                  PRIMARY KEY (`todo_id`, `group_id`) USING BTREE,
                                  INDEX `TodosOnGroups_group_id_fkey`(`group_id`) USING BTREE,
                                  CONSTRAINT `TodosOnGroups_group_id_fkey` FOREIGN KEY (`group_id`) REFERENCES `group` (`id`) ON DELETE CASCADE ON UPDATE RESTRICT,
                                  CONSTRAINT `TodosOnGroups_todo_id_fkey` FOREIGN KEY (`todo_id`) REFERENCES `todo` (`id`) ON DELETE CASCADE ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

