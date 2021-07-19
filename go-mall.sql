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

 Date: 19/07/2021 18:37:08
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
-- Table structure for banner
-- ----------------------------
DROP TABLE IF EXISTS `banner`;
CREATE TABLE `banner` (
  `id` int(11) NOT NULL COMMENT '主键',
  `url` varchar(0) COLLATE utf8_bin DEFAULT NULL COMMENT 'url地址',
  `goods_id` int(11) DEFAULT NULL COMMENT '商品ID',
  `title` varchar(0) COLLATE utf8_bin DEFAULT NULL COMMENT '标题',
  `desc` varchar(0) COLLATE utf8_bin DEFAULT NULL COMMENT '描述',
  `remark` varchar(0) COLLATE utf8_bin DEFAULT NULL COMMENT '备注(预留字段)',
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
  `id` int(32) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `price` float(32,2) DEFAULT NULL COMMENT '原价',
  `discount_price` float(32,2) DEFAULT NULL COMMENT '折后价',
  `name` char(200) COLLATE utf8_bin DEFAULT NULL COMMENT '商品名称',
  `desc` varchar(2000) COLLATE utf8_bin DEFAULT NULL COMMENT '商品描述',
  `cid` int(32) DEFAULT NULL COMMENT '分类ID',
  `remark` varchar(2000) COLLATE utf8_bin DEFAULT NULL COMMENT '备注',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 COLLATE=utf8_bin ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of goods
-- ----------------------------
BEGIN;
INSERT INTO `goods` VALUES (3, 12.00, 12.00, '苹果', '哈哈', 0, 'ok', '2021-07-19 18:24:42', '2021-07-19 18:24:42', NULL);
COMMIT;

-- ----------------------------
-- Table structure for order
-- ----------------------------
DROP TABLE IF EXISTS `order`;
CREATE TABLE `order` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `amount` float(255,2) DEFAULT NULL COMMENT '订单金额',
  `status` int(11) DEFAULT NULL COMMENT '订单状态(1.待支付、2.待发货、3.已发货，4.已签收、5.已关闭、6.已取消、7.已拒收、8.无效订单)',
  `pay_amount` float(255,2) DEFAULT NULL COMMENT '支付金额',
  `pay_way` int(11) DEFAULT NULL COMMENT '支付方式',
  `type` int(11) DEFAULT NULL COMMENT '订单类型(1.销售订单、2.退货订单、3.补发订单)',
  `pay_status` int(11) DEFAULT NULL COMMENT '支付状态(1.已支付、0.未支付)',
  `source` varchar(255) COLLATE utf8_bin DEFAULT NULL COMMENT '订单来源',
  `goods_name` varchar(0) COLLATE utf8_bin DEFAULT NULL COMMENT '商品名称',
  `goods_id` int(11) DEFAULT NULL COMMENT '商品ID',
  `goods_num` int(11) DEFAULT NULL COMMENT '商品数量',
  `pay_time` datetime DEFAULT NULL COMMENT '支付时间',
  `shipment_num` varchar(255) COLLATE utf8_bin DEFAULT NULL COMMENT '物流单号',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(32) NOT NULL AUTO_INCREMENT,
  `username` char(32) COLLATE utf8_bin DEFAULT NULL,
  `password` char(255) COLLATE utf8_bin DEFAULT NULL,
  `email` char(32) COLLATE utf8_bin DEFAULT NULL,
  `phone` char(255) COLLATE utf8_bin DEFAULT NULL,
  `nickname` char(32) COLLATE utf8_bin DEFAULT NULL,
  `avatar` char(255) COLLATE utf8_bin DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT INTO `user` VALUES (9, '哈哈23', '$2a$10$vQ7LY7SzphqGucjMJGxeyepOQaULwtVnacFbPvjlUfYB3zcJ4JZeO', '963732141@qq.com', '18328584905', '呵呵', '', '2021-07-19 11:12:29', '2021-07-19 11:12:29', NULL);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
