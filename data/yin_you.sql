/*
 Navicat Premium Data Transfer

 Source Server         : local
 Source Server Type    : MySQL
 Source Server Version : 100108
 Source Host           : localhost:3306
 Source Schema         : yin_you

 Target Server Type    : MySQL
 Target Server Version : 100108
 File Encoding         : 65001

 Date: 28/07/2019 00:53:03
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for music_information
-- ----------------------------
DROP TABLE IF EXISTS `music_information`;
CREATE TABLE `music_information`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(512) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '歌名或者歌手名称或者标签',
  `type` int(11) NULL DEFAULT NULL COMMENT '类型，具体自己定义',
  `addtional_fields` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT 'json字符串，此处需要前端定义不同类型需要展现的字段以及爬虫能入库的信息',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of music_information
-- ----------------------------
INSERT INTO `music_information` VALUES (1, 'test', 0, '{\"f1\":\"hi\"}');
INSERT INTO `music_information` VALUES (2, 'this is test for song 2', 0, '{\"f2\":\"hi\"}');

SET FOREIGN_KEY_CHECKS = 1;
