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

 Date: 19/09/2022 13:23:43
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for product
-- ----------------------------
DROP TABLE IF EXISTS `product`;
CREATE TABLE `product` (
  `id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `product_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `price` int NOT NULL,
  `seller_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`id`),
  KEY `id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of product
-- ----------------------------
BEGIN;
INSERT INTO `product` VALUES ('1663507151719420000', 'Product I Seller 2', 'Ini Product I Seller 2', 1000, '1663424129812111000');
INSERT INTO `product` VALUES ('1663508598665130000', 'Product I Seller 1', 'ini Product I Seller 1', 5000, '1663424171520107000');
INSERT INTO `product` VALUES ('1663511606558036000', 'Test Produk', 'Test description', 1500, '1663424171520107000');
INSERT INTO `product` VALUES ('1663511679594088000', 'Test Produk', 'Test description', 1500, '1663424171520107000');
INSERT INTO `product` VALUES ('1663511872233008000', 'Test Produk 1', 'Test description 1', 1500, '1663424171520107000');
INSERT INTO `product` VALUES ('1663511962264435000', 'Test Produk 1', 'Test description 1', 1500, '1663424129812111000');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
