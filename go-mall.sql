/*
 Navicat Premium Data Transfer

 Source Server         : localmysql
 Source Server Type    : MySQL
 Source Server Version : 80021
 Source Host           : localhost:3306
 Source Schema         : go-mall

 Target Server Type    : MySQL
 Target Server Version : 80021
 File Encoding         : 65001

 Date: 19/07/2021 22:08:15
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for address
-- ----------------------------
DROP TABLE IF EXISTS `address`;
CREATE TABLE `address`  (
  `id` int(0) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `user_id` int(0) NULL DEFAULT NULL COMMENT '用户id',
  `name` char(32) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL COMMENT '姓名',
  `mobile` char(32) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL COMMENT '手机号码',
  `province` char(30) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL COMMENT '省',
  `city` char(100) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL COMMENT '市',
  `county` char(30) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL COMMENT '县/区',
  `address` char(200) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL COMMENT '详细地址',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime(0) NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of address
-- ----------------------------

-- ----------------------------
-- Table structure for goods
-- ----------------------------
DROP TABLE IF EXISTS `goods`;
CREATE TABLE `goods`  (
  `id` int(0) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `price` float(32, 2) NULL DEFAULT NULL COMMENT '原价',
  `discount_price` float(32, 2) NULL DEFAULT NULL COMMENT '折后价',
  `name` char(200) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL COMMENT '商品名称',
  `desc` varchar(2000) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL COMMENT '商品描述',
  `cid` int(0) NULL DEFAULT NULL COMMENT '分类ID',
  `remark` varchar(2000) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL COMMENT '备注',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime(0) NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8 COLLATE = utf8_bin ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of goods
-- ----------------------------

-- ----------------------------
-- Table structure for order
-- ----------------------------
DROP TABLE IF EXISTS `order`;
CREATE TABLE `order`  (
  `id` int(0) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `amount` float(255, 2) NULL DEFAULT NULL COMMENT '订单金额',
  `status` int(0) NULL DEFAULT NULL COMMENT '订单状态(1.待支付、2.待发货、3.已发货，4.已签收、5.已关闭、6.已取消、7.已拒收、8.无效订单)',
  `pay_amount` float(255, 2) NULL DEFAULT NULL COMMENT '支付金额',
  `pay_way` int(0) NULL DEFAULT NULL COMMENT '支付方式',
  `type` int(0) NULL DEFAULT NULL COMMENT '订单类型(1.销售订单、2.退货订单、3.补发订单)',
  `pay_status` int(0) NULL DEFAULT NULL COMMENT '支付状态(1.已支付、0.未支付)',
  `source` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL COMMENT '订单来源',
  `goods_name` varchar(0) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL COMMENT '商品名称',
  `goods_id` int(0) NULL DEFAULT NULL COMMENT '商品ID',
  `goods_num` int(0) NULL DEFAULT NULL COMMENT '商品数量',
  `use_id` int(0) NULL DEFAULT NULL COMMENT '用户ID',
  `address_id` int(0) NULL DEFAULT NULL COMMENT '地址ID',
  `username` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL COMMENT '用户名',
  `pay_time` datetime(0) NULL DEFAULT NULL COMMENT '支付时间',
  `shipment_num` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL COMMENT '物流单号',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime(0) NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_bin ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of order
-- ----------------------------

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `username` char(32) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL,
  `password` char(255) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL,
  `email` char(32) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL,
  `phone` char(255) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL,
  `nickname` char(32) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL,
  `avatar` char(255) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8 COLLATE = utf8_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (10, '哈哈', '$2a$10$ta/g6Pn60dVO4T/gphzlqexW8i3XAWDplD.Tp59gL.xt4CiH/ZK3u', '963732141@qq.com', '18328584905', '呵呵', '', '2021-07-18 11:51:32', '2021-07-18 12:57:05', NULL);

SET FOREIGN_KEY_CHECKS = 1;
