/*
 Navicat MySQL Data Transfer

 Source Server         : lu
 Source Server Type    : MySQL
 Source Server Version : 50724
 Source Host           : localhost:3306
 Source Schema         : deercoder-chat

 Target Server Type    : MySQL
 Target Server Version : 50724
 File Encoding         : 65001

 Date: 18/01/2019 15:00:47
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for group_last_msg
-- ----------------------------
DROP TABLE IF EXISTS `group_last_msg`;
CREATE TABLE `group_last_msg`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `group_id` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '群聊id,获得其中成员',
  `uid` int(11) NULL DEFAULT NULL COMMENT '用户id,下次接收离线消息用',
  `last_group_msg_uuid` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '记录的最后消息id',
  `is_read` tinyint(1) NULL DEFAULT 0 COMMENT '0未读,1已读',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `group_id`(`last_group_msg_uuid`, `group_id`, `uid`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for group_msg
-- ----------------------------
DROP TABLE IF EXISTS `group_msg`;
CREATE TABLE `group_msg`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `uuid` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '唯一标识',
  `group_id` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '群聊id,获得其中成员',
  `content` blob NULL COMMENT '消息内容,可存储表情',
  `from_uid` int(11) NULL DEFAULT NULL COMMENT '由谁发送',
  `create_time` datetime(0) NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
  `content_type` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT 'text' COMMENT '内容类,默认文本类型',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uuid`(`uuid`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for group_users
-- ----------------------------
DROP TABLE IF EXISTS `group_users`;
CREATE TABLE `group_users`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `group_id` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '好友列表(包括群组)id',
  `uid` int(11) NULL DEFAULT NULL COMMENT '用户id',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `group_id`(`group_id`, `uid`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 100000002 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of group_users
-- ----------------------------
INSERT INTO `group_users` VALUES (100000000, '93f65451-efc4-11e8-918b-34e6d7558043', 1);
INSERT INTO `group_users` VALUES (100000001, '93f65451-efc4-11e8-918b-34e6d7558043', 2);

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '用户名',
  `headimg` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '头像',
  `password` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '密码',
  `createtime` datetime(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, '姓名666', 'xxx.img', '123456', NULL);
INSERT INTO `user` VALUES (2, '姓名1', 'xxx.img', '2c39f9867429', '2018-11-24 19:15:54');

SET FOREIGN_KEY_CHECKS = 1;
