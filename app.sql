/*
 Navicat Premium Data Transfer

 Source Server         : 127.0.0.1
 Source Server Type    : MySQL
 Source Server Version : 50744
 Source Host           : localhost:3306
 Source Schema         : app

 Target Server Type    : MySQL
 Target Server Version : 50744
 File Encoding         : 65001

 Date: 22/07/2024 17:27:42
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for app_admin
-- ----------------------------
DROP TABLE IF EXISTS `app_admin`;
CREATE TABLE `app_admin`  (
                              `id` tinyint(3) UNSIGNED NOT NULL AUTO_INCREMENT,
                              `user_id` int(30) NOT NULL DEFAULT 0 COMMENT '用户id',
                              `is_use` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '是否启用 1:启用 0:禁用',
                              `created_at` datetime NULL DEFAULT NULL,
                              `updated_at` datetime NULL DEFAULT NULL,
                              PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Fixed;

-- ----------------------------
-- Records of app_admin
-- ----------------------------
INSERT INTO `app_admin` VALUES (1, 1, 1, NULL, NULL);

-- ----------------------------
-- Table structure for app_admin_role
-- ----------------------------
DROP TABLE IF EXISTS `app_admin_role`;
CREATE TABLE `app_admin_role`  (
                                   `id` smallint(5) UNSIGNED NOT NULL AUTO_INCREMENT,
                                   `admin_id` smallint(5) UNSIGNED NOT NULL COMMENT '管理员id',
                                   `role_id` smallint(5) UNSIGNED NOT NULL COMMENT '角色id',
                                   `created_at` datetime NULL DEFAULT NULL,
                                   `updated_at` datetime NULL DEFAULT NULL,
                                   PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '会员角色表' ROW_FORMAT = Fixed;

-- ----------------------------
-- Records of app_admin_role
-- ----------------------------
INSERT INTO `app_admin_role` VALUES (1, 1, 1, '2024-07-09 15:01:52', '2024-07-09 15:01:55');

-- ----------------------------
-- Table structure for app_privilege
-- ----------------------------
DROP TABLE IF EXISTS `app_privilege`;
CREATE TABLE `app_privilege`  (
                                  `id` smallint(5) UNSIGNED NOT NULL AUTO_INCREMENT,
                                  `pri_name` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '权限名称',
                                  `action_name` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '方法名称',
                                  `parent_id` smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '上级权限的ID，0：代表顶级权限',
                                  `created_at` datetime NULL DEFAULT NULL,
                                  `updated_at` datetime NULL DEFAULT NULL,
                                  PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 4 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '权限表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of app_privilege
-- ----------------------------
INSERT INTO `app_privilege` VALUES (1, '用户管理', '', 0, NULL, NULL);
INSERT INTO `app_privilege` VALUES (2, '用户列表', '/admin/user/list', 1, NULL, NULL);
INSERT INTO `app_privilege` VALUES (3, '权限管理', '/admin/pri/list', 0, NULL, NULL);

-- ----------------------------
-- Table structure for app_role
-- ----------------------------
DROP TABLE IF EXISTS `app_role`;
CREATE TABLE `app_role`  (
                             `id` smallint(5) UNSIGNED NOT NULL AUTO_INCREMENT,
                             `role_name` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '角色名称',
                             `created_at` datetime NULL DEFAULT NULL,
                             `updated_at` datetime NULL DEFAULT NULL,
                             PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '角色表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of app_role
-- ----------------------------
INSERT INTO `app_role` VALUES (1, '运营', NULL, NULL);

-- ----------------------------
-- Table structure for app_role_privilege
-- ----------------------------
DROP TABLE IF EXISTS `app_role_privilege`;
CREATE TABLE `app_role_privilege`  (
                                       `id` smallint(5) UNSIGNED NOT NULL AUTO_INCREMENT,
                                       `pri_id` smallint(5) UNSIGNED NOT NULL COMMENT '权限的id',
                                       `role_id` smallint(5) UNSIGNED NOT NULL COMMENT '角色的id',
                                       `created_at` datetime NULL DEFAULT NULL,
                                       `updated_at` datetime NULL DEFAULT NULL,
                                       PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '角色权限表' ROW_FORMAT = Fixed;

-- ----------------------------
-- Records of app_role_privilege
-- ----------------------------
INSERT INTO `app_role_privilege` VALUES (1, 1, 1, '2024-07-10 14:11:23', '2024-07-10 14:11:25');
INSERT INTO `app_role_privilege` VALUES (2, 2, 1, '2024-07-10 14:11:49', '2024-07-10 14:11:52');

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
                         `id` int(11) NOT NULL AUTO_INCREMENT,
                         `user_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
                         `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
                         `phone` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
                         `ip` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
                         `status` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
                         `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '',
                         `token` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
                         `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
                         `created_at` datetime NULL DEFAULT NULL,
                         `updated_at` datetime NULL DEFAULT NULL,
                         PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, 'admin', '123456', '1000000', '', '', 'admin@123.com', '', '', '2024-06-27 13:34:15', '2024-06-27 13:34:15');

SET FOREIGN_KEY_CHECKS = 1;
