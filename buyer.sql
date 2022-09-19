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

 Date: 19/09/2022 13:23:18
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for buyer
-- ----------------------------
DROP TABLE IF EXISTS `buyer`;
CREATE TABLE `buyer` (
  `id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `email` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`id`),
  KEY `id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of buyer
-- ----------------------------
BEGIN;
INSERT INTO `buyer` VALUES ('1663424122746090000', 'buyer1@gmail.com', 'Buyer 1', '5cbd9d629096842872fdc665d2d03ba3', 'Bandung');
INSERT INTO `buyer` VALUES ('1663424168281696000', 'buyer2@gmail.com', 'Buyer 2', 'ba71d29d4efdd8753c516db594fab6d8', 'Jakarta');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
