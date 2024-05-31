-- ----------------------------
-- Records of goadmin_menu
-- ----------------------------
INSERT INTO `goadmin_menu` VALUES (1, 0, 1, 2, 'Admin', 'fa-tasks', '', NULL, '', NULL, '2019-09-10 00:00:00', '2019-09-10 00:00:00');
INSERT INTO `goadmin_menu` VALUES (2, 1, 1, 2, 'Users', 'fa-users', '/info/manager', NULL, '', NULL, '2019-09-10 00:00:00', '2019-09-10 00:00:00');
INSERT INTO `goadmin_menu` VALUES (3, 1, 1, 3, 'Roles', 'fa-user', '/info/roles', NULL, '', NULL, '2019-09-10 00:00:00', '2019-09-10 00:00:00');
INSERT INTO `goadmin_menu` VALUES (4, 1, 1, 4, 'Permission', 'fa-ban', '/info/permission', NULL, '', NULL, '2019-09-10 00:00:00', '2019-09-10 00:00:00');
INSERT INTO `goadmin_menu` VALUES (5, 1, 1, 5, 'Menu', 'fa-bars', '/menu', NULL, '', NULL, '2019-09-10 00:00:00', '2019-09-10 00:00:00');
INSERT INTO `goadmin_menu` VALUES (6, 1, 1, 6, 'Operation log', 'fa-history', '/info/op', NULL, '', NULL, '2019-09-10 00:00:00', '2019-09-10 00:00:00');
INSERT INTO `goadmin_menu` VALUES (7, 0, 1, 1, 'Dashboard', 'fa-bar-chart', '/', NULL, '', NULL, '2019-09-10 00:00:00', '2019-09-10 00:00:00');

-- ----------------------------
-- Records of goadmin_permissions
-- ----------------------------
INSERT INTO `goadmin_permissions` VALUES (1, 'All permission', '*', '', '*', '2019-09-10 00:00:00', '2019-09-10 00:00:00');
INSERT INTO `goadmin_permissions` VALUES (2, 'Dashboard', 'dashboard', 'GET,PUT,POST,DELETE', '/', '2019-09-10 00:00:00', '2019-09-10 00:00:00');
-- ----------------------------
-- Records of goadmin_role_menu
-- ----------------------------
INSERT INTO `goadmin_role_menu` VALUES (1, 1, '2019-09-10 00:00:00', '2019-09-10 00:00:00');
INSERT INTO `goadmin_role_menu` VALUES (1, 7, '2019-09-10 00:00:00', '2019-09-10 00:00:00');
INSERT INTO `goadmin_role_menu` VALUES (2, 7, '2019-09-10 00:00:00', '2019-09-10 00:00:00');
INSERT INTO `goadmin_role_menu` VALUES (1, 8, '2019-09-11 10:20:55', '2019-09-11 10:20:55');
INSERT INTO `goadmin_role_menu` VALUES (2, 8, '2019-09-11 10:20:55', '2019-09-11 10:20:55');

-- ----------------------------
-- Records of goadmin_role_permissions
-- ----------------------------
INSERT INTO `goadmin_role_permissions` VALUES (1, 1, '2019-09-10 00:00:00', '2019-09-10 00:00:00');
INSERT INTO `goadmin_role_permissions` VALUES (1, 2, '2019-09-10 00:00:00', '2019-09-10 00:00:00');
INSERT INTO `goadmin_role_permissions` VALUES (2, 2, '2019-09-10 00:00:00', '2019-09-10 00:00:00');
-- ----------------------------
-- Records of goadmin_role_users
-- ----------------------------
INSERT INTO `goadmin_role_users` VALUES (1, 1, '2019-09-10 00:00:00', '2019-09-10 00:00:00');
INSERT INTO `goadmin_role_users` VALUES (2, 2, '2019-09-10 00:00:00', '2019-09-10 00:00:00');
-- ----------------------------
-- Records of goadmin_roles
-- ----------------------------
INSERT INTO `goadmin_roles` VALUES (1, 'Administrator', 'administrator', '2019-09-10 00:00:00', '2019-09-10 00:00:00');
INSERT INTO `goadmin_roles` VALUES (2, 'Operator', 'operator', '2019-09-10 00:00:00', '2019-09-10 00:00:00');
-- ----------------------------
-- Records of goadmin_user_permissions
-- ----------------------------
INSERT INTO `goadmin_user_permissions` VALUES (1, 1, '2019-09-10 00:00:00', '2019-09-10 00:00:00');
INSERT INTO `goadmin_user_permissions` VALUES (2, 2, '2019-09-10 00:00:00', '2019-09-10 00:00:00');
-- ----------------------------
-- Records of goadmin_users
-- ----------------------------
INSERT INTO `goadmin_users` VALUES (1, 'admin', '$2a$10$U3F/NSaf2kaVbyXTBp7ppOn0jZFyRqXRnYXB.AMioCjXl3Ciaj4oy', 'admin', '', 'tlNcBVK9AvfYH7WEnwB1RKvocJu8FfRy4um3DJtwdHuJy0dwFsLOgAc0xUfh', '2019-09-10 00:00:00', '2019-09-10 00:00:00');
INSERT INTO `goadmin_users` VALUES (2, 'operator', '$2a$10$rVqkOzHjN2MdlEprRflb1eGP0oZXuSrbJLOmJagFsCd81YZm0bsh.', 'Operator', '', NULL, '2019-09-10 00:00:00', '2019-09-10 00:00:00');