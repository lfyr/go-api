/*
Navicat MySQL Data Transfer

Source Server         : localhot
Source Server Version : 50744
Source Host           : localhost:3306
Source Database       : app

Target Server Type    : MYSQL
Target Server Version : 50744
File Encoding         : 65001

Date: 2024-07-10 15:20:04
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for app_admin
-- ----------------------------
DROP TABLE IF EXISTS `app_admin`;
CREATE TABLE `app_admin` (
  `id` tinyint(3) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(30) NOT NULL DEFAULT '0' COMMENT '用户id',
  `is_use` tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '是否启用 1:启用 0:禁用',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of app_admin
-- ----------------------------
INSERT INTO `app_admin` VALUES ('1', '1', '1', null, null);

-- ----------------------------
-- Table structure for app_admin_role
-- ----------------------------
DROP TABLE IF EXISTS `app_admin_role`;
CREATE TABLE `app_admin_role` (
  `id` smallint(5) unsigned NOT NULL AUTO_INCREMENT,
  `admin_id` smallint(5) unsigned NOT NULL COMMENT '管理员id',
  `role_id` smallint(5) unsigned NOT NULL COMMENT '角色id',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COMMENT='会员角色表';

-- ----------------------------
-- Records of app_admin_role
-- ----------------------------
INSERT INTO `app_admin_role` VALUES ('1', '1', '1', '2024-07-09 15:01:52', '2024-07-09 15:01:55');

-- ----------------------------
-- Table structure for app_privilege
-- ----------------------------
DROP TABLE IF EXISTS `app_privilege`;
CREATE TABLE `app_privilege` (
  `id` smallint(5) unsigned NOT NULL AUTO_INCREMENT,
  `pri_name` varchar(30) NOT NULL COMMENT '权限名称',
  `action_name` varchar(20) NOT NULL COMMENT '方法名称',
  `parent_id` smallint(5) unsigned NOT NULL DEFAULT '0' COMMENT '上级权限的ID，0：代表顶级权限',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 COMMENT='权限表';

-- ----------------------------
-- Records of app_privilege
-- ----------------------------
INSERT INTO `app_privilege` VALUES ('1', '用户管理', '', '0', null, null);
INSERT INTO `app_privilege` VALUES ('2', '用户列表', '/admin/user/list', '1', null, null);
INSERT INTO `app_privilege` VALUES ('3', '权限管理', '/admin/pri/list', '0', null, null);

-- ----------------------------
-- Table structure for app_role
-- ----------------------------
DROP TABLE IF EXISTS `app_role`;
CREATE TABLE `app_role` (
  `id` smallint(5) unsigned NOT NULL AUTO_INCREMENT,
  `role_name` varchar(30) NOT NULL COMMENT '角色名称',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COMMENT='角色表';

-- ----------------------------
-- Records of app_role
-- ----------------------------
INSERT INTO `app_role` VALUES ('1', '运营', null, null);

-- ----------------------------
-- Table structure for app_role_privilege
-- ----------------------------
DROP TABLE IF EXISTS `app_role_privilege`;
CREATE TABLE `app_role_privilege` (
  `id` smallint(5) unsigned NOT NULL AUTO_INCREMENT,
  `pri_id` smallint(5) unsigned NOT NULL COMMENT '权限的id',
  `role_id` smallint(5) unsigned NOT NULL COMMENT '角色的id',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COMMENT='角色权限表';

-- ----------------------------
-- Records of app_role_privilege
-- ----------------------------
INSERT INTO `app_role_privilege` VALUES ('1', '1', '1', '2024-07-10 14:11:23', '2024-07-10 14:11:25');
INSERT INTO `app_role_privilege` VALUES ('2', '2', '1', '2024-07-10 14:11:49', '2024-07-10 14:11:52');

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_name` varchar(255) NOT NULL DEFAULT '',
  `password` varchar(255) NOT NULL DEFAULT '',
  `phone` varchar(255) NOT NULL DEFAULT '',
  `ip` varchar(255) NOT NULL DEFAULT '',
  `status` varchar(255) NOT NULL DEFAULT '',
  `email` varchar(255) DEFAULT '',
  `token` varchar(255) NOT NULL DEFAULT '',
  `avatar` varchar(255) NOT NULL DEFAULT '',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES ('1', 'admin', '123456', '1000000', '', '', 'admin@123.com', '', '', '2024-06-27 13:34:15', '2024-06-27 13:34:15');
