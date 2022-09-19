/*
 Navicat Premium Data Transfer

 Source Server         : Local SQL
 Source Server Type    : MySQL
 Source Server Version : 80029
 Source Host           : localhost:3306
 Source Schema         : gokomodo_test

 Target Server Type    : MySQL
 Target Server Version : 80029
 File Encoding         : 65001

 Date: 19/09/2022 13:23:33
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for order
-- ----------------------------
DROP TABLE IF EXISTS `order`;
CREATE TABLE `order` (
  `id` varchar(20) NOT NULL,
  `buyer_id` varchar(20) NOT NULL,
  `seller_id` varchar(20) NOT NULL,
  `delivery_source_address` varchar(255) NOT NULL,
  `delivery_destination_address` varchar(255) NOT NULL,
  `item_id` varchar(20) NOT NULL,
  `quantity` int NOT NULL,
  `price` int NOT NULL,
  `total_price` int NOT NULL,
  `status` varchar(20) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`),
  KEY `seller` (`seller_id`),
  KEY `buyer` (`buyer_id`),
  KEY `product` (`item_id`),
  CONSTRAINT `buyer` FOREIGN KEY (`buyer_id`) REFERENCES `buyer` (`id`),
  CONSTRAINT `product` FOREIGN KEY (`item_id`) REFERENCES `product` (`id`),
  CONSTRAINT `seller` FOREIGN KEY (`seller_id`) REFERENCES `seller` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of order
-- ----------------------------
BEGIN;
INSERT INTO `order` VALUES ('1663561375202319000', '1663424122746090000', '1663424129812111000', 'Jakarta', 'Bandung', '1663507151719420000', 5, 1000, 5000, 'accepted');
INSERT INTO `order` VALUES ('1663561540439144000', '1663424122746090000', '1663424129812111000', 'Jakarta', 'Bandung', '1663507151719420000', 2, 1000, 2000, 'pending');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
