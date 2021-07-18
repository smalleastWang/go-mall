/*
 Navicat Premium Data Transfer

 Source Server         : local-sql
 Source Server Type    : MySQL
 Source Server Version : 50721
 Source Host           : localhost:3306
 Source Schema         : go-mall

 Target Server Type    : MySQL
 Target Server Version : 50721
 File Encoding         : 65001

 Date: 18/07/2021 09:04:37
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for address
-- ----------------------------
DROP TABLE IF EXISTS `address`;
CREATE TABLE `address` (
  `id` int(32) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `user_id` int(32) DEFAULT NULL COMMENT '用户id',
  `name` char(32) COLLATE utf8_bin DEFAULT NULL COMMENT '姓名',
  `mobile` char(32) COLLATE utf8_bin DEFAULT NULL COMMENT '手机号码',
  `province` char(30) COLLATE utf8_bin DEFAULT NULL COMMENT '省',
  `city` char(100) COLLATE utf8_bin DEFAULT NULL COMMENT '市',
  `county` char(30) COLLATE utf8_bin DEFAULT NULL COMMENT '县/区',
  `address` char(200) COLLATE utf8_bin DEFAULT NULL COMMENT '详细地址',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
-- Table structure for goods
-- ----------------------------
DROP TABLE IF EXISTS `goods`;
CREATE TABLE `goods` (
  `id` int(32) NOT NULL COMMENT '主键',
  `price` float(32,2) DEFAULT NULL COMMENT '原价',
  `discount_price` float(32,2) DEFAULT NULL COMMENT '折后价',
  `name` char(200) COLLATE utf8_bin DEFAULT NULL COMMENT '商品名称',
  `desc` varchar(2000) COLLATE utf8_bin DEFAULT NULL COMMENT '商品描述',
  `remark` varchar(2000) COLLATE utf8_bin DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int(32) NOT NULL AUTO_INCREMENT,
  `username` char(32) COLLATE utf8_bin DEFAULT NULL,
  `password` char(32) COLLATE utf8_bin DEFAULT NULL,
  `email` char(32) COLLATE utf8_bin DEFAULT NULL,
  `phone` char(255) COLLATE utf8_bin DEFAULT NULL,
  `nickname` char(32) COLLATE utf8_bin DEFAULT NULL,
  `avatar` char(255) COLLATE utf8_bin DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
-- Records of users
-- ----------------------------
BEGIN;
INSERT INTO `users` VALUES (6, '哈哈12', '123', '963732141@qq.com', '18328584905', 'haha', '123', NULL, NULL, '2021-07-16 15:16:49');
INSERT INTO `users` VALUES (7, '哈哈112', '123', '963732141@qq.com', '18328584905', 'haha', '123', '2021-07-16 14:56:21', '2021-07-16 15:07:59', NULL);
INSERT INTO `users` VALUES (8, '哈哈12', '123', '963732141@qq.com', '18328584905', '', '', '2021-07-16 15:30:24', '2021-07-16 15:30:24', NULL);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
