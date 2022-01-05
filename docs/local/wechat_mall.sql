/*
 Navicat Premium Data Transfer

 Source Server         : 127.0.0.1
 Source Server Type    : MySQL
 Source Server Version : 50726
 Source Host           : 127.0.0.1:3306
 Source Schema         : wechat_mall

 Target Server Type    : MySQL
 Target Server Version : 50726
 File Encoding         : 65001

 Date: 12/06/2021 23:56:02
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for wechat_mall_banner
-- ----------------------------
DROP TABLE IF EXISTS `wechat_mall_banner`;
CREATE TABLE `wechat_mall_banner` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `picture` varchar(200) NOT NULL DEFAULT '' COMMENT '图片地址',
  `name` varchar(30) NOT NULL DEFAULT '' COMMENT '名称',
  `business_type` tinyint(2) NOT NULL DEFAULT '0' COMMENT '业务类型：1-商品',
  `business_id` int(1) NOT NULL DEFAULT '0' COMMENT '业务主键',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否显示：0-否 1-是',
  `is_del` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除：0-否 1-是',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COMMENT='小程序Banner表';

-- ----------------------------
-- Records of wechat_mall_banner
-- ----------------------------
BEGIN;
INSERT INTO `wechat_mall_banner` VALUES (1, '/assets/images/banner/c3213d88-250a-4dab-b0de-049e37d8d4bd.jpg', '首页', 1, 7, 1, 0, '2020-03-29 02:03:08', '2020-04-06 17:47:55');
INSERT INTO `wechat_mall_banner` VALUES (2, '/assets/images/banner/fcea8768-7803-43e5-814b-603e112e6628.jpg', '首页', 1, 0, 1, 0, '2020-03-29 02:04:32', '2020-04-01 06:05:34');
INSERT INTO `wechat_mall_banner` VALUES (3, '/assets/images/banner/2f7fe2a0-02cf-4247-a611-d5ed62c9be00.jpg', '首页', 1, 0, 1, 0, '2020-03-29 03:00:13', '2020-04-01 06:00:11');
COMMIT;

-- ----------------------------
-- Table structure for wechat_mall_category
-- ----------------------------
DROP TABLE IF EXISTS `wechat_mall_category`;
CREATE TABLE `wechat_mall_category` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `parent_id` int(11) NOT NULL DEFAULT '0' COMMENT '父级分类ID',
  `name` varchar(30) NOT NULL DEFAULT '' COMMENT '分类名称',
  `sort` int(11) NOT NULL DEFAULT '0' COMMENT '排序',
  `online` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否上线：0-否 1-是',
  `picture` varchar(200) NOT NULL DEFAULT '' COMMENT '图片地址',
  `description` varchar(50) NOT NULL DEFAULT '' COMMENT '分类描述',
  `is_del` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除：0-否 1-是',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_parent_id` (`parent_id`),
  KEY `idx_name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=utf8mb4 COMMENT='商城-分类表';

-- ----------------------------
-- Records of wechat_mall_category
-- ----------------------------
BEGIN;
INSERT INTO `wechat_mall_category` VALUES (1, 0, '数码', 1, 1, '/assets/images/category/a3d7ce35-acb1-4842-b423-80f605a10b49.png', 'Apple全系列产品', 0, '2020-03-12 12:29:46', '2020-04-06 19:18:22');
INSERT INTO `wechat_mall_category` VALUES (2, 0, '裤装', 0, 1, '/assets/images/category/1564cd21-7b1d-43eb-b782-dd33e4e0d2f0.png', '裤装', 1, '2020-03-12 12:30:02', '2020-03-29 04:42:22');
INSERT INTO `wechat_mall_category` VALUES (3, 0, 'B-3', 0, 0, '/assets/images/category/1cb8c60a-4ac0-46ff-827f-f035d895f5a1.jpg', '', 1, '2020-03-15 07:08:15', '2020-03-15 07:08:48');
INSERT INTO `wechat_mall_category` VALUES (4, 0, '特价区', 1, 1, '/assets/images/category/1b937e33-1086-4963-a475-144dbe3f5ade.png', '特价区', 1, '2020-03-15 07:09:05', '2020-03-29 04:42:08');
INSERT INTO `wechat_mall_category` VALUES (5, 0, '女装', 0, 0, '/assets/images/category/c5d79d66-48cf-407c-8122-77605c79a36b.png', '', 1, '2020-03-15 07:27:27', '2020-03-15 07:28:37');
INSERT INTO `wechat_mall_category` VALUES (6, 0, '女装', 0, 1, '/assets/images/category/6492cd75-7f8f-4d21-880d-ab186e2215fc.png', '', 1, '2020-03-15 07:29:01', '2020-03-15 07:29:35');
INSERT INTO `wechat_mall_category` VALUES (7, 1, '女装', 0, 1, '/assets/images/category/d225b217-fc0d-43e0-82fd-df24d5f92b4f.png', '', 1, '2020-03-15 07:30:44', '2020-03-15 07:34:28');
INSERT INTO `wechat_mall_category` VALUES (8, 1, 'iPhone', 0, 1, '/assets/images/category/d385f313-f518-4135-850c-44c6d4837ae6.jpg', 'iPhone系列', 0, '2020-03-15 07:34:54', '2020-04-06 19:18:22');
INSERT INTO `wechat_mall_category` VALUES (9, 2, '圆帽', 0, 1, '/assets/images/category/1f7cff02-f446-44db-806f-728c7608476b.png', '圆帽', 1, '2020-03-15 07:35:56', '2020-03-29 04:42:14');
INSERT INTO `wechat_mall_category` VALUES (10, 1, 'iPad', 0, 1, '/assets/images/category/9cc7a90c-3018-422a-8b87-062e0602fb8c.jpg', 'iPad系列', 0, '2020-03-16 04:16:49', '2020-04-06 19:18:22');
INSERT INTO `wechat_mall_category` VALUES (11, 2, '鸭舌帽', 0, 1, '/assets/images/category/62f10af5-e965-4b17-80fd-cfd69b00fac7.png', '鸭舌帽', 1, '2020-03-16 04:17:21', '2020-03-29 04:42:16');
INSERT INTO `wechat_mall_category` VALUES (12, 0, '裙装', 0, 1, '/assets/images/category/5d53f422-3cbf-4e64-8b98-314839bba40e.png', '裙装', 1, '2020-03-29 04:33:52', '2020-03-29 04:42:06');
INSERT INTO `wechat_mall_category` VALUES (13, 0, '套装', 0, 1, '/assets/images/category/50297fd9-3e1f-4def-8cb5-b4c242514129.png', '套装', 1, '2020-03-29 04:34:07', '2020-03-29 04:42:04');
INSERT INTO `wechat_mall_category` VALUES (14, 0, '包包', 0, 1, '/assets/images/category/0aa2a736-c8be-40d7-8436-67e8b5afc722.png', '包包', 1, '2020-03-29 04:35:47', '2020-03-29 04:36:04');
INSERT INTO `wechat_mall_category` VALUES (15, 1, 'MacBook', 0, 1, '/assets/images/category/84da1af5-81e0-4bae-8d9c-e3efe9c2bf75.jpg', 'MacBook笔记本系列', 0, '2020-03-29 04:36:50', '2020-04-06 19:18:22');
INSERT INTO `wechat_mall_category` VALUES (16, 4, '手机', 0, 1, '/assets/images/category/09bed3ef-62be-4ea8-8a34-e24b840a38c5.png', '手机', 1, '2020-03-29 04:38:06', '2020-03-29 04:41:54');
INSERT INTO `wechat_mall_category` VALUES (17, 4, '电脑', 0, 1, '/assets/images/category/519a196f-ca9f-438b-87b4-cd8a7d4b7f71.png', '电脑', 1, '2020-03-29 04:38:19', '2020-03-29 04:41:56');
INSERT INTO `wechat_mall_category` VALUES (18, 1, 'Mac Pro', 0, 1, '/assets/images/category/e6810fbe-88c0-4ef7-92d8-b1e0f2768ac9.jpg', 'Mac Pro主机', 0, '2020-03-29 06:34:15', '2020-04-06 19:18:22');
INSERT INTO `wechat_mall_category` VALUES (19, 1, 'iMac', 0, 1, '/assets/images/category/6bdf15a4-7d23-4044-80eb-f7847eaabc2c.jpg', 'iMac系列', 0, '2020-03-29 06:34:41', '2020-04-06 19:18:22');
INSERT INTO `wechat_mall_category` VALUES (20, 1, 'Watch', 0, 1, '/assets/images/category/6104efa5-ea7f-42a2-a69e-a3d5398c0524.jpg', 'Watch手表系列', 0, '2020-03-29 06:35:08', '2020-04-06 19:18:22');
INSERT INTO `wechat_mall_category` VALUES (21, 1, 'AirPods', 0, 1, '/assets/images/category/2dff8f98-462e-404b-bcd5-125fc5764e58.jpg', 'AirPods耳机系列', 0, '2020-03-29 06:35:30', '2020-04-06 19:18:22');
INSERT INTO `wechat_mall_category` VALUES (22, 1, '配件', 0, 1, '/assets/images/category/00ab3380-3280-4178-8f52-00a390d0d64d.jpg', '配件系列', 0, '2020-03-29 06:35:52', '2020-04-06 19:18:22');
INSERT INTO `wechat_mall_category` VALUES (23, 0, '1234', 0, 1, '/assets/images/category/c577b057-85b4-483f-997d-0109b30c7a6e.jpg', '321', 1, '2020-03-30 05:01:52', '2020-03-30 05:02:14');
INSERT INTO `wechat_mall_category` VALUES (24, 0, '1111', 0, 0, '/assets/images/category/c9f35aad-12ad-4eee-8dae-0fa500f749bc.jpg', '', 1, '2020-03-31 07:34:37', '2020-03-31 07:34:52');
COMMIT;

-- ----------------------------
-- Table structure for wechat_mall_coupon
-- ----------------------------
DROP TABLE IF EXISTS `wechat_mall_coupon`;
CREATE TABLE `wechat_mall_coupon` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `title` varchar(30) NOT NULL DEFAULT '' COMMENT '标题',
  `full_money` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '满减额',
  `minus` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '优惠额',
  `rate` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '折扣',
  `type` tinyint(1) NOT NULL DEFAULT '0' COMMENT '券类型：1-满减券 2-折扣券 3-代金券 4-满金额折扣券',
  `grant_num` int(11) NOT NULL DEFAULT '0' COMMENT '发券数量',
  `limit_num` int(11) NOT NULL DEFAULT '0' COMMENT '单人限领',
  `start_time` datetime NOT NULL COMMENT '开始时间',
  `end_time` datetime NOT NULL COMMENT '结束时间',
  `description` varchar(30) NOT NULL DEFAULT '' COMMENT '描述',
  `online` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否上架: 0-下架 1-上架',
  `is_del` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除：0-否 1-是',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COMMENT='商城-优惠券表';

-- ----------------------------
-- Records of wechat_mall_coupon
-- ----------------------------
BEGIN;
INSERT INTO `wechat_mall_coupon` VALUES (1, '满100减10', 100.00, 10.00, 0.00, 1, 100, 2, '2021-05-01 00:00:00', '2021-12-31 00:00:00', '', 1, 0, '2020-05-01 00:00:00', '2020-05-01 00:00:00');
INSERT INTO `wechat_mall_coupon` VALUES (2, '折扣券', 0.01, 0.01, 0.88, 2, 100, 2, '2021-05-01 00:00:00', '2021-12-31 00:00:00', '', 1, 0, '2020-05-01 00:00:00', '2020-05-01 00:00:00');
INSERT INTO `wechat_mall_coupon` VALUES (3, '100代金券', 0.00, 0.01, 0.00, 3, 100, 10, '2021-05-01 00:00:00', '2021-12-31 00:00:00', '', 1, 0, '2020-05-01 00:00:00', '2020-05-01 00:00:00');
INSERT INTO `wechat_mall_coupon` VALUES (4, '满100享5.5折', 100.00, 0.00, 0.55, 4, 100, 2, '2021-05-01 00:00:00', '2021-12-31 00:00:00', '', 1, 0, '2020-05-01 00:00:00', '2020-05-01 00:00:00');
COMMIT;

-- ----------------------------
-- Table structure for wechat_mall_coupon_log
-- ----------------------------
DROP TABLE IF EXISTS `wechat_mall_coupon_log`;
CREATE TABLE `wechat_mall_coupon_log` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `coupon_id` int(11) NOT NULL DEFAULT '0' COMMENT '优惠券ID',
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `use_time` datetime DEFAULT NULL COMMENT '使用时间',
  `expire_time` datetime NOT NULL COMMENT '过期时间',
  `status` int(11) NOT NULL DEFAULT '0' COMMENT '状态：0-未使用 1-已使用 2-已过期',
  `code` varchar(30) NOT NULL DEFAULT '' COMMENT '兑换码',
  `order_no` varchar(30) NOT NULL DEFAULT '' COMMENT '核销的订单号',
  `is_del` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除：0-否 1-是',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_coupon_id` (`coupon_id`),
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商城-优惠券领取记录表';

-- ----------------------------
-- Table structure for wechat_mall_goods
-- ----------------------------
DROP TABLE IF EXISTS `wechat_mall_goods`;
CREATE TABLE `wechat_mall_goods` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `brand_name` varchar(30) NOT NULL DEFAULT '' COMMENT '品牌名称',
  `title` varchar(80) NOT NULL DEFAULT '' COMMENT '标题',
  `price` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '价格',
  `discount_price` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '折扣',
  `category_id` int(11) NOT NULL DEFAULT '0' COMMENT '分类ID',
  `online` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否上架：0-下架 1-上架',
  `picture` varchar(200) NOT NULL DEFAULT '' COMMENT '主图',
  `banner_picture` text COMMENT '轮播图',
  `detail_picture` text COMMENT '详情图',
  `tags` varchar(100) NOT NULL DEFAULT '' COMMENT '标签，示例：包邮$热门',
  `sale_num` int(11) NOT NULL DEFAULT '0' COMMENT '商品销量',
  `is_del` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除：0-否 1-是',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_title` (`title`),
  KEY `idx_category_id` (`category_id`)
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8mb4 COMMENT='商城-商品表';

-- ----------------------------
-- Records of wechat_mall_goods
-- ----------------------------
BEGIN;
INSERT INTO `wechat_mall_goods` VALUES (1, 'Aj', '春季-爆款-001', 10.98, 0.10, 8, 1, '/assets/images/goods/4fa48b65-cd94-4c2e-8b37-7e54482745d7.jpg', '[\"/assets/images/goods/979c0865-0217-4397-8b30-67786558f094.jpg\",\"/assets/images/goods/c336fff2-987c-49db-a0d8-5c059fd30478.gif\"]', '[\"/assets/images/goods/dad3cda2-c400-456a-95cd-8c9c8dc8a4a0.png\",\"/assets/images/goods/af8a8492-e3e9-45e4-8e26-19f1eb621da4.png\"]', '', 0, 1, '2020-03-12 12:44:27', '2020-03-29 06:29:35');
INSERT INTO `wechat_mall_goods` VALUES (2, 'Aj2', '春季-新品-002', 10.98, 0.10, 8, 1, '/assets/images/goods/acae461d-ca22-4efd-ba01-967094139a48.gif', '[\"/assets/images/goods/7b4090c8-7920-412f-8a4b-ca2e375f53eb.jpg\"]', '[\"/assets/images/goods/88e8ec5d-0d25-4e96-9884-81c7e829d040.jpg\"]', '', 0, 1, '2020-03-12 12:45:21', '2020-03-29 06:29:37');
INSERT INTO `wechat_mall_goods` VALUES (3, '', '春季-新品-002', 0.01, 0.00, 10, 1, '/assets/images/goods/4ab255ac-ea62-4bd9-8161-20fc7692a114.jpg', '[\"/assets/images/goods/5bee2c17-e941-404e-b483-cebc6468cd78.jpg\"]', '[\"/assets/images/goods/602ac66e-00fb-4404-b403-bbf32ee28b54.jpg\"]', '', 0, 1, '2020-03-16 04:42:53', '2020-03-16 04:45:57');
INSERT INTO `wechat_mall_goods` VALUES (4, '', '春季-新品-001', 4.01, 0.00, 10, 1, '/assets/images/goods/0edbca5b-8612-4ac0-81ea-4f9e2885f6de.jpg', '[\"/assets/images/goods/a427d307-ca70-4d13-8406-bf8c9f7a6e4a.jpg\"]', '[\"/assets/images/goods/9ec40814-168a-4fa8-88d7-d5ae30ce537b.jpg\"]', '', 0, 1, '2020-03-16 04:49:15', '2020-03-29 06:29:39');
INSERT INTO `wechat_mall_goods` VALUES (5, '', '春季-新品-005', 0.01, 0.00, 9, 1, '/assets/images/goods/81b928ae-e031-42ca-864e-e12a5f36dc67.jpg', '[\"/assets/images/goods/e6732201-98a0-4195-9258-5bc413f63acf.jpg\"]', '[\"/assets/images/goods/dded049e-61d7-4473-bb65-3abf212803e7.jpg\"]', '', 0, 1, '2020-03-18 02:55:44', '2020-03-18 02:57:34');
INSERT INTO `wechat_mall_goods` VALUES (6, '', '新品-005', 0.01, 0.00, 9, 1, '/assets/images/goods/8c8c0318-1925-4177-8160-652772635c85.jpg', '[\"/assets/images/goods/b0cc60a3-8fbe-4dad-b764-9a50f291b1f5.jpg\"]', '[\"/assets/images/goods/61ceca30-702e-42a7-8160-d4eec10c1a0d.jpg\"]', '', 0, 1, '2020-03-18 03:03:10', '2020-03-18 03:03:15');
INSERT INTO `wechat_mall_goods` VALUES (7, '', 'Apple iPhone 11 Pro (A2217) 64GB 暗夜绿色 移动联通电信4G手机 双卡双待', 8699.00, 0.00, 8, 1, '/assets/images/goods/706c5b82-42c7-4e6d-8f59-1eb3d6a03fb8.jpg', '[\"/assets/images/goods/93af8814-c6b6-4eb2-8b89-a26435d75677.jpg\",\"/assets/images/goods/3464f419-8041-4a33-ab96-0c2f1380e74d.jpg\",\"/assets/images/goods/aa21bfde-24b5-4558-9b5a-06fd0325fb58.jpg\"]', '[\"/assets/images/goods/9ddaa408-ce8c-4323-94f6-b426e094d086.png\",\"/assets/images/goods/4a13b0b7-ccc3-4340-83b6-09971da635d9.png\"]', '', 2, 0, '2020-03-29 07:22:18', '2020-04-06 19:18:22');
INSERT INTO `wechat_mall_goods` VALUES (8, '', 'Apple iPhone XS Max (A2104) 64GB 金色 移动联通电信4G手机 双卡双待', 6299.00, 0.00, 8, 1, '/assets/images/goods/5288be5a-eeec-43f7-84b3-5cc4cde7907d.jpg', '[\"/assets/images/goods/167806d7-3e92-4eb1-aa39-891e1e9658b3.jpg\"]', '[\"/assets/images/goods/2cef6f02-4285-4386-8b37-463ce5c8c286.jpg\",\"/assets/images/goods/5886d1da-d10b-410a-8d24-b16cdc5044f9.jpg\"]', '', 3, 0, '2020-03-29 07:33:09', '2020-04-06 19:18:22');
INSERT INTO `wechat_mall_goods` VALUES (9, '', 'Apple iPhone XR (A2108) 64GB 白色 移动联通电信4G手机 双卡双待', 4599.00, 0.00, 8, 1, '/assets/images/goods/c4c4a6d8-7bba-4362-8af1-b57b53321fb5.jpg', '[\"/assets/images/goods/f7067d52-f3f1-4d15-8f53-3e9c8ce2125f.jpg\",\"/assets/images/goods/5f72ecc9-6247-4d06-9cac-2a126dab7b72.jpg\"]', '[\"/assets/images/goods/40fddf2d-d23c-494d-8635-d2a2892003be.jpg\",\"/assets/images/goods/a8247bf0-4afe-4863-b0ee-449a4566ab1c.jpg\"]', '', 3, 0, '2020-03-29 07:40:10', '2020-04-06 19:18:22');
INSERT INTO `wechat_mall_goods` VALUES (10, '', 'Apple iPad Pro 12.9英寸平板电脑 2020年新款(128G WLAN版/全面屏/A12Z）', 7899.00, 0.00, 10, 1, '/assets/images/goods/5b15ea71-7576-4eff-8f60-dae6d4e9bb3d.jpg', '[\"/assets/images/goods/2972b059-9a3d-4015-83df-14893aebfcb4.jpg\",\"/assets/images/goods/118fc2f0-ad30-4632-8c41-aa1e5e45e686.jpg\",\"/assets/images/goods/78a03b57-a29a-4500-8e4a-1c86a9ae537a.jpg\"]', '[\"/assets/images/goods/530b712d-b87e-4322-a061-eab85122e4ff.jpg\",\"/assets/images/goods/92b1feed-8d9a-4ba9-b092-84e02423010b.jpg\"]', '', 1, 0, '2020-03-29 07:50:32', '2020-04-06 19:18:22');
INSERT INTO `wechat_mall_goods` VALUES (11, '', 'Apple iPad Pro 11英寸平板电脑2018年新款(1TB WLAN+Cellular版/全面屏/A12X）', 12499.00, 0.00, 10, 1, '/assets/images/goods/06da347c-a85d-4650-aec6-41c3807e248a.jpg', '[\"/assets/images/goods/a74fda9d-e13e-40cf-81e0-ac80af1cc748.jpg\",\"/assets/images/goods/c07ef42a-820c-4c65-8263-26bd1535fff3.jpg\"]', '[\"/assets/images/goods/ed634243-8b43-4b75-a951-5403afbb1db8.jpg\",\"/assets/images/goods/66d46a7a-ede1-408c-bb0c-be9530d9f4e9.jpg\"]', '', 0, 0, '2020-03-29 08:01:03', '2020-04-06 19:18:22');
INSERT INTO `wechat_mall_goods` VALUES (12, '', 'Apple 2019新品 MacBook Pro 16【带触控栏】九代六核i7 16G 512G 深空灰', 18999.00, 0.00, 15, 1, '/assets/images/goods/3030cd7d-17e0-4f31-859f-6dc7e7f690bb.jpg', '[\"/assets/images/goods/1c411487-13d4-42e2-8fe5-792be9c371ed.jpg\",\"/assets/images/goods/f7331f7c-ed3a-4fae-8af1-a69faa3def04.jpg\",\"/assets/images/goods/b6bbb8f8-43dc-4c35-9a7a-36bc217e2209.jpg\"]', '[\"/assets/images/goods/023129ea-1859-4ec4-80b7-051db1736c5d.jpg\",\"/assets/images/goods/d6eb6ad2-0c1f-49ff-a5f9-b0a09856411b.jpg\"]', '', 0, 0, '2020-03-29 08:11:39', '2020-04-06 19:18:22');
INSERT INTO `wechat_mall_goods` VALUES (13, '', '（新款）Mac Pro 3.5GHz 8 核 Intel Xeon W 处理器，32GB (4x8GB) DDR4 内存', 51999.00, 0.00, 18, 1, '/assets/images/goods/916098f3-f7bb-45ad-99f3-e67535abd0af.png', '[\"/assets/images/goods/42228842-7adf-4468-8bef-2ea22af9dda4.png\"]', '[\"/assets/images/goods/bb4ddbcf-a2fb-4bb5-87c9-76e03e5905e4.jpg\",\"/assets/images/goods/b0b75d8b-0de0-418a-8fcd-eb6e80e35222.png\"]', '', 0, 0, '2020-03-29 08:52:43', '2020-04-06 19:18:22');
INSERT INTO `wechat_mall_goods` VALUES (14, '', 'Apple iMac【2019新品】27英寸一体机5K屏 Core i5 8G 1TB融合 台式电脑主机', 13388.00, 0.00, 19, 1, '/assets/images/goods/e01594e9-f919-4550-9154-35452c0320fc.jpg', '[\"/assets/images/goods/5b9ca0d8-e63f-41b6-8bb3-d5ce93cf3ac6.jpg\"]', '[\"/assets/images/goods/069b6b21-0ed2-4aac-a65a-ce8af3111461.jpg\",\"/assets/images/goods/79764745-f8bb-4b3f-abf5-8c0fb23f69fb.jpg\",\"/assets/images/goods/17b10f42-c20b-4d1d-8edd-2a37d0070e85.jpg\"]', '', 0, 0, '2020-03-29 08:58:24', '2020-04-06 19:18:22');
INSERT INTO `wechat_mall_goods` VALUES (15, '', 'Apple Watch Series 5智能手表（GPS款 40毫米深空灰色铝金属表壳 MWV82CH/A)', 3199.00, 0.00, 20, 1, '/assets/images/goods/b45119c7-a959-4db6-ac97-c5cc57565dc7.jpg', '[\"/assets/images/goods/4be67dab-5e83-4108-8fcb-07b0ebe992e4.jpg\",\"/assets/images/goods/8591bd04-7181-40a6-b60f-56f59b61bb06.jpg\",\"/assets/images/goods/6611d0cf-7b9c-4d4d-84aa-0374e9e6d5c2.jpg\"]', '[\"/assets/images/goods/66d9cf89-f8df-4c31-a524-49232f1d5b24.jpg\",\"/assets/images/goods/f3cd8879-32b6-4391-af1b-ac05ba3958f5.jpg\"]', '', 0, 0, '2020-03-29 09:02:28', '2020-04-06 19:18:22');
INSERT INTO `wechat_mall_goods` VALUES (16, '', 'Apple AirPods Pro 主动降噪无线蓝牙耳机 适用iPhone/iPad/Apple Watch', 1999.00, 0.00, 21, 1, '/assets/images/goods/fee1e936-1625-4227-b79e-9c43d453a2fe.jpg', '[\"/assets/images/goods/480313c7-80fd-4844-8bbd-5f82cb095898.jpg\",\"/assets/images/goods/2246c12d-a5f0-4ca2-a13a-e4f6c9c1e144.jpg\",\"/assets/images/goods/4e28d2ef-6b4c-4213-8c02-8ac5e1adbded.jpg\"]', '[\"/assets/images/goods/e4664e44-6778-4b50-90f1-e823f6b40840.png\",\"/assets/images/goods/bf990ecc-758e-454c-9e00-5d7e3acbf38e.png\"]', '', 0, 0, '2020-03-29 09:05:32', '2020-04-06 19:18:22');
INSERT INTO `wechat_mall_goods` VALUES (17, '', 'Apple Magic Mouse/妙控鼠标 2代 - 银色 适用MacBook 无线鼠标', 564.00, 0.00, 22, 1, '/assets/images/goods/c54287b6-1a1e-4ffe-8e9c-a180018825b7.jpg', '[\"/assets/images/goods/3de2de74-b028-4828-890b-e9c8531a986f.jpg\",\"/assets/images/goods/3cfaab74-7c47-40f7-81f2-033b5ee59f11.jpg\",\"/assets/images/goods/e760ca78-7709-4cd0-8819-48c576d7ce7b.jpg\"]', '[\"/assets/images/goods/f1c943c9-a761-469a-857a-597148e4c17f.jpg\"]', '', 0, 0, '2020-03-29 09:14:46', '2020-04-06 19:18:22');
INSERT INTO `wechat_mall_goods` VALUES (18, '', 'Apple Lightning/闪电转 USB 连接线 (1 米) iPhone iPad 手机 平板 数据线 充电线', 135.00, 0.00, 22, 1, '/assets/images/goods/1ef52e99-fdca-47a2-8d5e-a76b1dc5bd49.jpg', '[\"/assets/images/goods/3386ad41-a3f9-45ae-84d3-81615d6f4ab6.jpg\",\"/assets/images/goods/300abd1b-3d83-4d31-807d-26783b94cdb1.jpg\"]', '[\"/assets/images/goods/6357a780-27bd-48df-b4ef-061db04f07e1.jpg\"]', '', 0, 0, '2020-03-29 09:16:59', '2020-04-06 19:18:22');
INSERT INTO `wechat_mall_goods` VALUES (19, '', '1234', 100.00, 0.00, 10, 0, '/assets/images/goods/0ba497e7-ef5d-475d-847b-e243ced84e96.jpg', '[\"/assets/images/goods/ea324478-4c7b-4ac7-ab1a-1b2abc99fcd6.jpg\"]', '[\"/assets/images/goods/bd101021-3043-436c-817c-49d670191c8f.jpg\"]', '', 0, 1, '2020-03-30 05:19:58', '2020-03-30 05:20:52');
INSERT INTO `wechat_mall_goods` VALUES (20, '', '1111', 0.01, 0.00, 20, 1, '/assets/images/goods/5a3e141a-7595-436d-86f2-1adf74ef5fcc.jpg', '[]', '[]', '', 0, 1, '2020-04-07 11:15:48', '2020-04-07 11:15:59');
COMMIT;

-- ----------------------------
-- Table structure for wechat_mall_goods_browse_record
-- ----------------------------
DROP TABLE IF EXISTS `wechat_mall_goods_browse_record`;
CREATE TABLE `wechat_mall_goods_browse_record` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `goods_id` int(11) NOT NULL DEFAULT '0' COMMENT '商品ID',
  `picture` varchar(200) NOT NULL DEFAULT '' COMMENT '商品图片',
  `title` varchar(80) NOT NULL DEFAULT '' COMMENT '商品名称',
  `price` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '商品价格',
  `is_del` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除：0-否 1-是',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_goods_id` (`goods_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商城-商品浏览记录';

-- ----------------------------
-- Table structure for wechat_mall_goods_spec
-- ----------------------------
DROP TABLE IF EXISTS `wechat_mall_goods_spec`;
CREATE TABLE `wechat_mall_goods_spec` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `goods_id` int(11) NOT NULL DEFAULT '0' COMMENT '商品ID',
  `spec_id` int(11) NOT NULL DEFAULT '0' COMMENT '规格ID',
  `is_del` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除：0-否 1-是',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_goods_id` (`goods_id`)
) ENGINE=InnoDB AUTO_INCREMENT=126 DEFAULT CHARSET=utf8mb4 COMMENT='商城-商品规格表';

-- ----------------------------
-- Records of wechat_mall_goods_spec
-- ----------------------------
BEGIN;
INSERT INTO `wechat_mall_goods_spec` VALUES (1, 1, 1, 1, '2020-03-12 12:44:27', '2020-03-12 12:44:27');
INSERT INTO `wechat_mall_goods_spec` VALUES (2, 1, 2, 1, '2020-03-12 12:44:27', '2020-03-12 12:44:27');
INSERT INTO `wechat_mall_goods_spec` VALUES (3, 2, 2, 1, '2020-03-12 12:45:21', '2020-03-12 12:45:21');
INSERT INTO `wechat_mall_goods_spec` VALUES (4, 2, 3, 1, '2020-03-12 12:45:21', '2020-03-12 12:45:21');
INSERT INTO `wechat_mall_goods_spec` VALUES (5, 2, 2, 1, '2020-03-12 12:45:40', '2020-03-12 12:45:40');
INSERT INTO `wechat_mall_goods_spec` VALUES (6, 2, 3, 1, '2020-03-12 12:45:40', '2020-03-12 12:45:40');
INSERT INTO `wechat_mall_goods_spec` VALUES (7, 2, 2, 1, '2020-03-15 12:55:08', '2020-03-15 12:55:08');
INSERT INTO `wechat_mall_goods_spec` VALUES (8, 2, 3, 1, '2020-03-15 12:55:08', '2020-03-15 12:55:08');
INSERT INTO `wechat_mall_goods_spec` VALUES (9, 1, 1, 1, '2020-03-16 03:27:18', '2020-03-16 03:27:18');
INSERT INTO `wechat_mall_goods_spec` VALUES (10, 1, 1, 1, '2020-03-16 03:27:21', '2020-03-16 03:27:21');
INSERT INTO `wechat_mall_goods_spec` VALUES (11, 1, 4, 1, '2020-03-16 03:27:21', '2020-03-16 03:27:21');
INSERT INTO `wechat_mall_goods_spec` VALUES (12, 1, 1, 1, '2020-03-16 03:27:55', '2020-03-16 03:27:55');
INSERT INTO `wechat_mall_goods_spec` VALUES (13, 1, 4, 1, '2020-03-16 03:27:55', '2020-03-16 03:27:55');
INSERT INTO `wechat_mall_goods_spec` VALUES (14, 1, 1, 1, '2020-03-16 03:28:38', '2020-03-16 03:28:38');
INSERT INTO `wechat_mall_goods_spec` VALUES (15, 1, 4, 1, '2020-03-16 03:28:38', '2020-03-16 03:28:38');
INSERT INTO `wechat_mall_goods_spec` VALUES (16, 1, 1, 1, '2020-03-16 04:15:17', '2020-03-16 04:15:17');
INSERT INTO `wechat_mall_goods_spec` VALUES (17, 1, 4, 1, '2020-03-16 04:15:17', '2020-03-16 04:15:17');
INSERT INTO `wechat_mall_goods_spec` VALUES (18, 1, 1, 1, '2020-03-16 04:18:09', '2020-03-16 04:18:09');
INSERT INTO `wechat_mall_goods_spec` VALUES (19, 1, 4, 1, '2020-03-16 04:18:09', '2020-03-16 04:18:09');
INSERT INTO `wechat_mall_goods_spec` VALUES (20, 1, 1, 1, '2020-03-16 04:19:42', '2020-03-16 04:19:42');
INSERT INTO `wechat_mall_goods_spec` VALUES (21, 1, 4, 1, '2020-03-16 04:19:42', '2020-03-16 04:19:42');
INSERT INTO `wechat_mall_goods_spec` VALUES (22, 1, 1, 1, '2020-03-16 04:20:54', '2020-03-16 04:20:54');
INSERT INTO `wechat_mall_goods_spec` VALUES (23, 1, 4, 1, '2020-03-16 04:20:54', '2020-03-16 04:20:54');
INSERT INTO `wechat_mall_goods_spec` VALUES (24, 1, 1, 1, '2020-03-16 04:21:10', '2020-03-16 04:21:10');
INSERT INTO `wechat_mall_goods_spec` VALUES (25, 1, 4, 1, '2020-03-16 04:21:10', '2020-03-16 04:21:10');
INSERT INTO `wechat_mall_goods_spec` VALUES (26, 1, 1, 1, '2020-03-16 04:37:19', '2020-03-16 04:37:19');
INSERT INTO `wechat_mall_goods_spec` VALUES (27, 1, 4, 1, '2020-03-16 04:37:19', '2020-03-16 04:37:19');
INSERT INTO `wechat_mall_goods_spec` VALUES (28, 2, 2, 1, '2020-03-16 04:41:49', '2020-03-16 04:41:49');
INSERT INTO `wechat_mall_goods_spec` VALUES (29, 2, 3, 1, '2020-03-16 04:41:49', '2020-03-16 04:41:49');
INSERT INTO `wechat_mall_goods_spec` VALUES (30, 4, 1, 1, '2020-03-16 04:49:16', '2020-03-16 04:49:16');
INSERT INTO `wechat_mall_goods_spec` VALUES (31, 4, 1, 1, '2020-03-16 04:49:43', '2020-03-16 04:49:43');
INSERT INTO `wechat_mall_goods_spec` VALUES (32, 4, 4, 1, '2020-03-16 04:49:43', '2020-03-16 04:49:43');
INSERT INTO `wechat_mall_goods_spec` VALUES (33, 4, 1, 1, '2020-03-16 04:49:48', '2020-03-16 04:49:48');
INSERT INTO `wechat_mall_goods_spec` VALUES (34, 4, 4, 1, '2020-03-16 04:49:48', '2020-03-16 04:49:48');
INSERT INTO `wechat_mall_goods_spec` VALUES (35, 4, 1, 1, '2020-03-16 04:49:55', '2020-03-16 04:49:55');
INSERT INTO `wechat_mall_goods_spec` VALUES (36, 4, 4, 1, '2020-03-16 04:49:55', '2020-03-16 04:49:55');
INSERT INTO `wechat_mall_goods_spec` VALUES (37, 1, 1, 1, '2020-03-16 07:10:01', '2020-03-16 07:10:01');
INSERT INTO `wechat_mall_goods_spec` VALUES (38, 1, 2, 1, '2020-03-16 07:10:01', '2020-03-16 07:10:01');
INSERT INTO `wechat_mall_goods_spec` VALUES (39, 4, 1, 1, '2020-03-18 02:21:13', '2020-03-18 02:21:13');
INSERT INTO `wechat_mall_goods_spec` VALUES (40, 4, 4, 1, '2020-03-18 02:21:13', '2020-03-18 02:21:13');
INSERT INTO `wechat_mall_goods_spec` VALUES (41, 4, 1, 1, '2020-03-18 02:54:56', '2020-03-18 02:54:56');
INSERT INTO `wechat_mall_goods_spec` VALUES (42, 4, 4, 1, '2020-03-18 02:54:56', '2020-03-18 02:54:56');
INSERT INTO `wechat_mall_goods_spec` VALUES (43, 4, 1, 1, '2020-03-18 02:55:02', '2020-03-18 02:55:02');
INSERT INTO `wechat_mall_goods_spec` VALUES (44, 4, 4, 1, '2020-03-18 02:55:02', '2020-03-18 02:55:02');
INSERT INTO `wechat_mall_goods_spec` VALUES (45, 5, 1, 1, '2020-03-18 02:55:44', '2020-03-18 02:55:44');
INSERT INTO `wechat_mall_goods_spec` VALUES (46, 5, 2, 1, '2020-03-18 02:55:44', '2020-03-18 02:55:44');
INSERT INTO `wechat_mall_goods_spec` VALUES (47, 6, 1, 1, '2020-03-18 03:03:10', '2020-03-18 03:03:10');
INSERT INTO `wechat_mall_goods_spec` VALUES (48, 7, 1, 1, '2020-03-29 07:29:03', '2020-03-29 07:29:03');
INSERT INTO `wechat_mall_goods_spec` VALUES (49, 7, 5, 1, '2020-03-29 07:29:03', '2020-03-29 07:29:03');
INSERT INTO `wechat_mall_goods_spec` VALUES (50, 9, 1, 0, '2020-03-29 07:40:10', '2020-03-29 07:40:10');
INSERT INTO `wechat_mall_goods_spec` VALUES (51, 9, 5, 0, '2020-03-29 07:40:10', '2020-03-29 07:40:10');
INSERT INTO `wechat_mall_goods_spec` VALUES (52, 10, 1, 0, '2020-03-29 07:50:32', '2020-03-29 07:50:32');
INSERT INTO `wechat_mall_goods_spec` VALUES (53, 10, 5, 0, '2020-03-29 07:50:32', '2020-03-29 07:50:32');
INSERT INTO `wechat_mall_goods_spec` VALUES (54, 11, 1, 0, '2020-03-29 08:01:03', '2020-03-29 08:01:03');
INSERT INTO `wechat_mall_goods_spec` VALUES (55, 11, 5, 0, '2020-03-29 08:01:03', '2020-03-29 08:01:03');
INSERT INTO `wechat_mall_goods_spec` VALUES (56, 12, 1, 1, '2020-03-29 08:11:39', '2020-03-29 08:11:39');
INSERT INTO `wechat_mall_goods_spec` VALUES (57, 12, 5, 1, '2020-03-29 08:11:39', '2020-03-29 08:11:39');
INSERT INTO `wechat_mall_goods_spec` VALUES (58, 13, 1, 0, '2020-03-29 08:52:43', '2020-03-29 08:52:43');
INSERT INTO `wechat_mall_goods_spec` VALUES (59, 14, 1, 0, '2020-03-29 08:58:24', '2020-03-29 08:58:24');
INSERT INTO `wechat_mall_goods_spec` VALUES (60, 15, 1, 0, '2020-03-29 09:02:28', '2020-03-29 09:02:28');
INSERT INTO `wechat_mall_goods_spec` VALUES (61, 16, 1, 1, '2020-03-29 09:05:32', '2020-03-29 09:05:32');
INSERT INTO `wechat_mall_goods_spec` VALUES (62, 17, 1, 1, '2020-03-29 09:14:46', '2020-03-29 09:14:46');
INSERT INTO `wechat_mall_goods_spec` VALUES (63, 17, 5, 1, '2020-03-29 09:14:46', '2020-03-29 09:14:46');
INSERT INTO `wechat_mall_goods_spec` VALUES (64, 18, 9, 1, '2020-03-29 09:16:59', '2020-03-29 09:16:59');
INSERT INTO `wechat_mall_goods_spec` VALUES (65, 18, 10, 1, '2020-03-29 09:16:59', '2020-03-29 09:16:59');
INSERT INTO `wechat_mall_goods_spec` VALUES (70, 8, 1, 0, '2020-03-29 11:01:59', '2020-03-29 11:01:59');
INSERT INTO `wechat_mall_goods_spec` VALUES (71, 8, 5, 0, '2020-03-29 11:01:59', '2020-03-29 11:01:59');
INSERT INTO `wechat_mall_goods_spec` VALUES (72, 19, 1, 1, '2020-03-30 05:20:27', '2020-03-30 05:20:27');
INSERT INTO `wechat_mall_goods_spec` VALUES (73, 19, 1, 0, '2020-03-30 05:20:40', '2020-03-30 05:20:40');
INSERT INTO `wechat_mall_goods_spec` VALUES (74, 19, 9, 0, '2020-03-30 05:20:40', '2020-03-30 05:20:40');
INSERT INTO `wechat_mall_goods_spec` VALUES (75, 7, 1, 1, '2020-03-30 14:35:57', '2020-03-30 14:35:57');
INSERT INTO `wechat_mall_goods_spec` VALUES (76, 7, 5, 1, '2020-03-30 14:35:57', '2020-03-30 14:35:57');
INSERT INTO `wechat_mall_goods_spec` VALUES (77, 7, 9, 1, '2020-03-30 14:35:57', '2020-03-30 14:35:57');
INSERT INTO `wechat_mall_goods_spec` VALUES (78, 12, 1, 1, '2020-03-31 10:45:47', '2020-03-31 10:45:47');
INSERT INTO `wechat_mall_goods_spec` VALUES (79, 12, 5, 1, '2020-03-31 10:45:47', '2020-03-31 10:45:47');
INSERT INTO `wechat_mall_goods_spec` VALUES (80, 12, 1, 0, '2020-03-31 10:46:45', '2020-03-31 10:46:45');
INSERT INTO `wechat_mall_goods_spec` VALUES (81, 12, 5, 0, '2020-03-31 10:46:45', '2020-03-31 10:46:45');
INSERT INTO `wechat_mall_goods_spec` VALUES (82, 7, 1, 1, '2020-03-31 12:32:27', '2020-03-31 12:32:27');
INSERT INTO `wechat_mall_goods_spec` VALUES (83, 7, 5, 1, '2020-03-31 12:32:27', '2020-03-31 12:32:27');
INSERT INTO `wechat_mall_goods_spec` VALUES (84, 7, 9, 1, '2020-03-31 12:32:27', '2020-03-31 12:32:27');
INSERT INTO `wechat_mall_goods_spec` VALUES (85, 7, 1, 1, '2020-03-31 12:32:41', '2020-03-31 12:32:41');
INSERT INTO `wechat_mall_goods_spec` VALUES (86, 7, 5, 1, '2020-03-31 12:32:41', '2020-03-31 12:32:41');
INSERT INTO `wechat_mall_goods_spec` VALUES (87, 7, 9, 1, '2020-03-31 12:32:41', '2020-03-31 12:32:41');
INSERT INTO `wechat_mall_goods_spec` VALUES (88, 7, 1, 1, '2020-03-31 12:32:55', '2020-03-31 12:32:55');
INSERT INTO `wechat_mall_goods_spec` VALUES (89, 7, 5, 1, '2020-03-31 12:32:55', '2020-03-31 12:32:55');
INSERT INTO `wechat_mall_goods_spec` VALUES (90, 7, 9, 1, '2020-03-31 12:32:55', '2020-03-31 12:32:55');
INSERT INTO `wechat_mall_goods_spec` VALUES (91, 7, 1, 1, '2020-03-31 13:09:13', '2020-03-31 13:09:13');
INSERT INTO `wechat_mall_goods_spec` VALUES (92, 7, 5, 1, '2020-03-31 13:09:13', '2020-03-31 13:09:13');
INSERT INTO `wechat_mall_goods_spec` VALUES (93, 7, 9, 1, '2020-03-31 13:09:13', '2020-03-31 13:09:13');
INSERT INTO `wechat_mall_goods_spec` VALUES (94, 7, 10, 1, '2020-03-31 13:09:13', '2020-03-31 13:09:13');
INSERT INTO `wechat_mall_goods_spec` VALUES (95, 7, 1, 1, '2020-04-01 05:33:08', '2020-04-01 05:33:08');
INSERT INTO `wechat_mall_goods_spec` VALUES (96, 7, 5, 1, '2020-04-01 05:33:08', '2020-04-01 05:33:08');
INSERT INTO `wechat_mall_goods_spec` VALUES (97, 7, 9, 1, '2020-04-01 05:33:08', '2020-04-01 05:33:08');
INSERT INTO `wechat_mall_goods_spec` VALUES (98, 7, 10, 1, '2020-04-01 05:33:08', '2020-04-01 05:33:08');
INSERT INTO `wechat_mall_goods_spec` VALUES (99, 7, 1, 1, '2020-04-01 05:33:40', '2020-04-01 05:33:40');
INSERT INTO `wechat_mall_goods_spec` VALUES (100, 7, 5, 1, '2020-04-01 05:33:40', '2020-04-01 05:33:40');
INSERT INTO `wechat_mall_goods_spec` VALUES (101, 7, 9, 1, '2020-04-01 05:33:40', '2020-04-01 05:33:40');
INSERT INTO `wechat_mall_goods_spec` VALUES (102, 7, 10, 1, '2020-04-01 05:33:40', '2020-04-01 05:33:40');
INSERT INTO `wechat_mall_goods_spec` VALUES (103, 7, 1, 1, '2020-04-01 05:34:24', '2020-04-01 05:34:24');
INSERT INTO `wechat_mall_goods_spec` VALUES (104, 7, 5, 1, '2020-04-01 05:34:24', '2020-04-01 05:34:24');
INSERT INTO `wechat_mall_goods_spec` VALUES (105, 7, 9, 1, '2020-04-01 05:34:24', '2020-04-01 05:34:24');
INSERT INTO `wechat_mall_goods_spec` VALUES (106, 18, 10, 0, '2020-04-01 05:57:49', '2020-04-01 05:57:49');
INSERT INTO `wechat_mall_goods_spec` VALUES (107, 18, 9, 0, '2020-04-01 05:57:49', '2020-04-01 05:57:49');
INSERT INTO `wechat_mall_goods_spec` VALUES (108, 17, 1, 0, '2020-04-02 09:25:39', '2020-04-02 09:25:39');
INSERT INTO `wechat_mall_goods_spec` VALUES (109, 7, 1, 1, '2020-04-02 09:33:51', '2020-04-02 09:33:51');
INSERT INTO `wechat_mall_goods_spec` VALUES (110, 7, 5, 1, '2020-04-02 09:33:51', '2020-04-02 09:33:51');
INSERT INTO `wechat_mall_goods_spec` VALUES (111, 7, 9, 1, '2020-04-02 09:33:51', '2020-04-02 09:33:51');
INSERT INTO `wechat_mall_goods_spec` VALUES (112, 7, 1, 1, '2020-04-02 09:34:05', '2020-04-02 09:34:05');
INSERT INTO `wechat_mall_goods_spec` VALUES (113, 7, 1, 1, '2020-04-02 09:36:27', '2020-04-02 09:36:27');
INSERT INTO `wechat_mall_goods_spec` VALUES (114, 7, 1, 1, '2020-04-02 09:36:37', '2020-04-02 09:36:37');
INSERT INTO `wechat_mall_goods_spec` VALUES (115, 7, 5, 1, '2020-04-02 09:36:37', '2020-04-02 09:36:37');
INSERT INTO `wechat_mall_goods_spec` VALUES (116, 7, 9, 1, '2020-04-02 09:36:37', '2020-04-02 09:36:37');
INSERT INTO `wechat_mall_goods_spec` VALUES (117, 7, 1, 1, '2020-04-02 09:37:25', '2020-04-02 09:37:25');
INSERT INTO `wechat_mall_goods_spec` VALUES (118, 7, 9, 1, '2020-04-02 09:37:25', '2020-04-02 09:37:25');
INSERT INTO `wechat_mall_goods_spec` VALUES (119, 7, 10, 1, '2020-04-02 09:37:25', '2020-04-02 09:37:25');
INSERT INTO `wechat_mall_goods_spec` VALUES (120, 7, 1, 0, '2020-04-02 09:37:30', '2020-04-02 09:37:30');
INSERT INTO `wechat_mall_goods_spec` VALUES (121, 7, 9, 0, '2020-04-02 09:37:30', '2020-04-02 09:37:30');
INSERT INTO `wechat_mall_goods_spec` VALUES (122, 7, 5, 0, '2020-04-02 09:37:30', '2020-04-02 09:37:30');
INSERT INTO `wechat_mall_goods_spec` VALUES (123, 16, 1, 1, '2020-04-02 10:05:39', '2020-04-02 10:05:39');
INSERT INTO `wechat_mall_goods_spec` VALUES (124, 16, 1, 0, '2020-04-02 10:05:46', '2020-04-02 10:05:46');
INSERT INTO `wechat_mall_goods_spec` VALUES (125, 20, 1, 1, '2020-04-07 11:15:48', '2020-04-07 11:15:48');
COMMIT;

-- ----------------------------
-- Table structure for wechat_mall_grid_category
-- ----------------------------
DROP TABLE IF EXISTS `wechat_mall_grid_category`;
CREATE TABLE `wechat_mall_grid_category` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `title` varchar(30) NOT NULL DEFAULT '' COMMENT '宫格标题',
  `name` varchar(30) NOT NULL DEFAULT '' COMMENT '宫格名',
  `category_id` int(11) NOT NULL DEFAULT '0' COMMENT '顶级分类ID',
  `picture` varchar(200) NOT NULL DEFAULT '' COMMENT '图片地址',
  `is_del` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除：0-否 1-是',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_category_id` (`category_id`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4 COMMENT='小程序-首页宫格表';

-- ----------------------------
-- Records of wechat_mall_grid_category
-- ----------------------------
BEGIN;
INSERT INTO `wechat_mall_grid_category` VALUES (9, '', 'iPhone', 8, '/assets/images/grid_category/90cb1d64-a3ea-48e6-8379-8c43d5865362.jpg', 0, '2020-03-29 06:39:28', '2020-04-06 17:56:13');
INSERT INTO `wechat_mall_grid_category` VALUES (10, '', 'iPad', 10, '/assets/images/grid_category/d1cb4eea-76fe-4e4b-a5c4-90ab66046a23.jpg', 0, '2020-03-29 06:39:47', '2020-03-29 06:39:47');
INSERT INTO `wechat_mall_grid_category` VALUES (11, '', 'MacBook', 15, '/assets/images/grid_category/612acbe6-d010-4a21-8085-19001020044a.jpg', 0, '2020-03-29 06:40:43', '2020-03-29 06:40:43');
INSERT INTO `wechat_mall_grid_category` VALUES (12, '', 'Mac Pro', 18, '/assets/images/grid_category/8f1b6283-6694-4dc1-8b79-929532019fe0.jpg', 0, '2020-03-29 06:41:03', '2020-03-29 06:41:03');
INSERT INTO `wechat_mall_grid_category` VALUES (13, '', 'iMac', 19, '/assets/images/grid_category/b2e3f6d2-b054-4c6a-9110-ad5b8b3b40e7.jpg', 0, '2020-03-29 06:41:17', '2020-03-29 06:41:17');
INSERT INTO `wechat_mall_grid_category` VALUES (14, '', 'Watch', 20, '/assets/images/grid_category/ef2693cf-17be-4a6c-a84d-9d75c2849836.jpg', 0, '2020-03-29 06:41:32', '2020-03-29 06:41:32');
INSERT INTO `wechat_mall_grid_category` VALUES (15, '', 'AirPods', 21, '/assets/images/grid_category/1cfbeecc-ba44-4953-8a27-449c0d8757a8.jpg', 0, '2020-03-29 06:41:47', '2020-03-29 06:41:47');
INSERT INTO `wechat_mall_grid_category` VALUES (16, '', '配件', 22, '/assets/images/grid_category/6bdc02a9-a087-49f4-9e9f-1d7a32ef3c55.jpg', 0, '2020-03-29 06:42:01', '2020-04-01 06:10:08');
COMMIT;

-- ----------------------------
-- Table structure for wechat_mall_order
-- ----------------------------
DROP TABLE IF EXISTS `wechat_mall_order`;
CREATE TABLE `wechat_mall_order` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `order_no` varchar(32) NOT NULL DEFAULT '' COMMENT '订单号',
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `pay_amount` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '订单金额（商品金额 + 运费 - 优惠金额）',
  `goods_amount` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '商品小计金额',
  `discount_amount` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '优惠金额',
  `dispatch_amount` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '运费',
  `pay_time` datetime NOT NULL DEFAULT '2006-01-02 15:04:05' COMMENT '支付时间',
  `deliver_time` datetime NOT NULL DEFAULT '2006-01-02 15:04:05' COMMENT '发货时间',
  `finish_time` datetime NOT NULL DEFAULT '2006-01-02 15:04:05' COMMENT '成交时间',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态 -1 已取消 0-待付款 1-待发货 2-待收货 3-已完成 4-（待发货）退款申请 5-已退款\n',
  `address_id` int(11) NOT NULL DEFAULT '0' COMMENT '收货地址ID',
  `address_snapshot` varchar(2000) NOT NULL DEFAULT '' COMMENT '收货地址快照',
  `wxapp_prepay_id` varchar(50) NOT NULL DEFAULT '' COMMENT '微信预支付ID',
  `transaction_id` varchar(50) NOT NULL DEFAULT '' COMMENT '微信支付单号',
  `remark` varchar(100) NOT NULL DEFAULT '' COMMENT '订单备注',
  `is_del` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除：0-否 1-是',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `idx_order_no` (`order_no`),
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商城订单表';

-- ----------------------------
-- Table structure for wechat_mall_order_goods
-- ----------------------------
DROP TABLE IF EXISTS `wechat_mall_order_goods`;
CREATE TABLE `wechat_mall_order_goods` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `order_no` varchar(30) NOT NULL DEFAULT '' COMMENT '订单号',
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `goods_id` int(11) NOT NULL DEFAULT '0' COMMENT '商品ID',
  `sku_id` int(11) NOT NULL DEFAULT '0' COMMENT 'sku ID',
  `picture` varchar(200) NOT NULL DEFAULT '' COMMENT '商品图片',
  `title` varchar(80) NOT NULL DEFAULT '' COMMENT '商品标题',
  `price` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '价格',
  `specs` varchar(500) NOT NULL DEFAULT '' COMMENT 'sku规格属性',
  `num` int(11) NOT NULL DEFAULT '0' COMMENT '数量',
  `lock_status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '锁定状态：0-锁定 1-解锁',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `idx_order_no` (`order_no`),
  KEY `idx_goods_id` (`goods_id`),
  KEY `idx_sku_id` (`sku_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商城订单-商品表';

-- ----------------------------
-- Table structure for wechat_mall_order_refund
-- ----------------------------
DROP TABLE IF EXISTS `wechat_mall_order_refund`;
CREATE TABLE `wechat_mall_order_refund` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `refund_no` varchar(30) NOT NULL DEFAULT '' COMMENT '退款编号',
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '平台用户ID',
  `order_no` varchar(30) NOT NULL DEFAULT '' COMMENT '订单号',
  `reason` varchar(30) NOT NULL DEFAULT '' COMMENT '退款原因',
  `refund_amount` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '退款金额',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态：0-退款申请 1-退款完成 2-撤销申请',
  `is_del` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除：0-否 1-是',
  `refund_time` datetime NOT NULL DEFAULT '2006-01-02 15:04:05' COMMENT '退款时间',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商城-订单退款申请表';

-- ----------------------------
-- Table structure for wechat_mall_sku
-- ----------------------------
DROP TABLE IF EXISTS `wechat_mall_sku`;
CREATE TABLE `wechat_mall_sku` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `title` varchar(20) NOT NULL DEFAULT '' COMMENT '标题',
  `price` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '价格',
  `code` varchar(30) NOT NULL DEFAULT '' COMMENT '编码',
  `stock` int(11) NOT NULL DEFAULT '0' COMMENT '库存量',
  `goods_id` int(11) NOT NULL DEFAULT '0' COMMENT '所属商品',
  `online` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否上架: 0-下架 1-上架',
  `picture` varchar(200) NOT NULL DEFAULT '' COMMENT '图片',
  `specs` varchar(500) NOT NULL DEFAULT '' COMMENT '规格属性',
  `is_del` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除：0-否 1-是',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_title` (`title`),
  KEY `idx_code` (`code`),
  KEY `idx_goods_id` (`goods_id`)
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8mb4 COMMENT='商城-SKU表';

-- ----------------------------
-- Records of wechat_mall_sku
-- ----------------------------
BEGIN;
INSERT INTO `wechat_mall_sku` VALUES (1, '爆款-3.5英寸', 1.20, '123', 7, 1, 1, '/assets/images/sku/b87f331c-d799-41a5-99ea-34fb8b33a178.jpg', '[{\"keyId\":1,\"key\":\"颜色\",\"valueId\":4,\"value\":\"蓝色\"},{\"keyId\":2,\"key\":\"尺寸\",\"valueId\":6,\"value\":\"3.5英寸\"}]', 1, '2020-03-12 13:15:04', '2020-03-29 06:29:27');
INSERT INTO `wechat_mall_sku` VALUES (2, '新品-3.2寸', 0.20, '1234', 10, 2, 1, '/assets/images/sku/364c0387-b64b-4552-8969-3dd500625788.jpg', '[{\"keyId\":2,\"key\":\"尺寸\",\"valueId\":2,\"value\":\"7英寸\"}]', 1, '2020-03-12 13:15:35', '2020-03-29 06:29:29');
INSERT INTO `wechat_mall_sku` VALUES (3, '新品-3.5寸', 0.01, '321#123', 3, 2, 1, '/assets/images/sku/6e1d43ed-cf1a-43d9-b107-ec9559dd5033.png', '[{\"keyId\":2,\"key\":\"尺寸\",\"valueId\":2,\"value\":\"7英寸\"}]', 1, '2020-03-16 10:03:32', '2020-03-16 13:05:16');
INSERT INTO `wechat_mall_sku` VALUES (4, '新品-ABC', 0.01, '#123', 9, 4, 1, '/assets/images/sku/c26d5f6b-4189-4d31-8cf5-417be791b1c7.jpg', '[{\"keyId\":1,\"key\":\"颜色\",\"valueId\":3,\"value\":\"红色\"},{\"keyId\":4,\"key\":\"型号\",\"valueId\":7,\"value\":\"ABC\"}]', 1, '2020-03-16 13:06:55', '2020-03-29 06:29:30');
INSERT INTO `wechat_mall_sku` VALUES (5, '暗夜绿色·64G', 8699.00, '', 3, 7, 1, '/assets/images/sku/ecac6375-15a8-45a4-bf0f-72168e35e576.jpg', '[{\"keyId\":1,\"key\":\"颜色\",\"valueId\":3,\"value\":\"暗夜绿色\"},{\"keyId\":5,\"key\":\"版本\",\"valueId\":10,\"value\":\"64GB\"},{\"keyId\":9,\"key\":\"长度\",\"valueId\":26,\"value\":\"2米\"}]', 0, '2020-03-30 02:47:51', '2020-05-03 20:11:24');
INSERT INTO `wechat_mall_sku` VALUES (6, '金色·512G', 8699.00, '', 94, 7, 1, '/assets/images/sku/f07512b0-2e83-4dca-bd5b-0178a424f195.jpg', '[{\"keyId\":1,\"key\":\"颜色\",\"valueId\":4,\"value\":\"金色\"},{\"keyId\":5,\"key\":\"版本\",\"valueId\":12,\"value\":\"512GB\"},{\"keyId\":9,\"key\":\"长度\",\"valueId\":26,\"value\":\"2米\"}]', 0, '2020-03-30 02:50:16', '2020-04-03 18:05:11');
INSERT INTO `wechat_mall_sku` VALUES (7, '深空灰色·64G', 7699.00, '', 96, 7, 1, '/assets/images/sku/3f16c65f-1e00-4017-8720-c672ed57f21a.jpg', '[{\"keyId\":1,\"key\":\"颜色\",\"valueId\":8,\"value\":\"深空灰色\"},{\"keyId\":5,\"key\":\"版本\",\"valueId\":10,\"value\":\"64GB\"},{\"keyId\":9,\"key\":\"长度\",\"valueId\":26,\"value\":\"2米\"}]', 0, '2020-03-30 04:35:44', '2020-05-03 14:48:28');
INSERT INTO `wechat_mall_sku` VALUES (8, '银色·256G', 8699.00, '', 100, 7, 1, '/assets/images/sku/581c71a1-086d-4393-ab1c-9ef59775034d.jpg', '[{\"keyId\":1,\"key\":\"颜色\",\"valueId\":9,\"value\":\"银色\"},{\"keyId\":5,\"key\":\"版本\",\"valueId\":11,\"value\":\"256GB\"},{\"keyId\":9,\"key\":\"长度\",\"valueId\":26,\"value\":\"2米\"},{\"keyId\":10,\"key\":\"接头\",\"valueId\":28,\"value\":\"USB-C充电线\"}]', 1, '2020-03-30 04:36:45', '2020-03-31 13:17:42');
INSERT INTO `wechat_mall_sku` VALUES (9, '1111', 0.01, '', 10, 7, 1, '/assets/images/sku/bf1fff51-98af-49f9-873e-5106263a69f9.jpg', '[{\"keyId\":1,\"key\":\"颜色\",\"valueId\":3,\"value\":\"暗夜绿色\"},{\"keyId\":5,\"key\":\"版本\",\"valueId\":11,\"value\":\"256GB\"},{\"keyId\":9,\"key\":\"长度\",\"valueId\":25,\"value\":\"1米\"}]', 1, '2020-04-01 05:37:08', '2020-04-01 05:39:47');
INSERT INTO `wechat_mall_sku` VALUES (10, '11111', 0.01, '', 1, 7, 0, '/assets/images/sku/48950f70-fb3e-4267-9c06-02a26e2d174d.gif', '[{\"keyId\":1,\"key\":\"颜色\",\"valueId\":4,\"value\":\"金色\"},{\"keyId\":5,\"key\":\"版本\",\"valueId\":16,\"value\":\"WLAN版 128G\"},{\"keyId\":9,\"key\":\"长度\",\"valueId\":25,\"value\":\"1米\"}]', 1, '2020-04-01 05:40:32', '2020-04-01 05:41:25');
INSERT INTO `wechat_mall_sku` VALUES (11, '111', 0.01, '', 1, 7, 0, '/assets/images/sku/c1abacaf-9b6b-4684-8df0-974a378ff56d.jpg', '[{\"keyId\":1,\"key\":\"颜色\",\"valueId\":4,\"value\":\"金色\"},{\"keyId\":5,\"key\":\"版本\",\"valueId\":10,\"value\":\"64GB\"},{\"keyId\":9,\"key\":\"长度\",\"valueId\":25,\"value\":\"1米\"}]', 1, '2020-04-01 05:44:11', '2020-04-01 05:44:20');
INSERT INTO `wechat_mall_sku` VALUES (12, '银色·64G', 8699.00, '', 100, 7, 1, '/assets/images/sku/417cd05d-09a3-4bd7-8181-73281bd56645.jpg', '[{\"keyId\":1,\"key\":\"颜色\",\"valueId\":9,\"value\":\"银色\"},{\"keyId\":9,\"key\":\"长度\",\"valueId\":25,\"value\":\"1米\"},{\"keyId\":5,\"key\":\"版本\",\"valueId\":10,\"value\":\"64GB\"}]', 0, '2020-04-02 08:30:54', '2020-04-02 16:55:02');
INSERT INTO `wechat_mall_sku` VALUES (13, '黑色·64G', 4599.00, '', 99, 9, 1, '/assets/images/sku/7df59d9a-3c82-403a-8f95-2b1789db3086.jpg', '[{\"keyId\":1,\"key\":\"颜色\",\"valueId\":14,\"value\":\"黑色\"},{\"keyId\":5,\"key\":\"版本\",\"valueId\":10,\"value\":\"64GB\"}]', 0, '2020-04-02 08:32:08', '2020-04-03 19:34:16');
INSERT INTO `wechat_mall_sku` VALUES (14, '红色·256G', 4599.00, '', 97, 9, 1, '/assets/images/sku/178c9daa-2365-469d-b4b4-aebf486adcc6.jpg', '[{\"keyId\":1,\"key\":\"颜色\",\"valueId\":15,\"value\":\"红色\"},{\"keyId\":5,\"key\":\"版本\",\"valueId\":11,\"value\":\"256GB\"}]', 0, '2020-04-02 08:32:53', '2020-05-03 14:48:28');
INSERT INTO `wechat_mall_sku` VALUES (15, '蓝色·64G', 4599.00, '', 100, 9, 1, '/assets/images/sku/8512adb7-38a9-42df-8083-7fc6ad374478.jpg', '[{\"keyId\":1,\"key\":\"颜色\",\"valueId\":13,\"value\":\"蓝色\"},{\"keyId\":5,\"key\":\"版本\",\"valueId\":10,\"value\":\"64GB\"}]', 0, '2020-04-02 08:33:26', '2020-04-02 08:33:26');
INSERT INTO `wechat_mall_sku` VALUES (16, '金色·64G', 6299.00, '', 97, 8, 1, '/assets/images/sku/21ff5e82-14e2-449e-a749-20da7b1d3ba1.jpg', '[{\"keyId\":1,\"key\":\"颜色\",\"valueId\":4,\"value\":\"金色\"},{\"keyId\":5,\"key\":\"版本\",\"valueId\":10,\"value\":\"64GB\"}]', 0, '2020-04-02 08:34:14', '2020-04-03 19:35:09');
INSERT INTO `wechat_mall_sku` VALUES (17, '深空灰色·256G', 6299.00, '', 97, 8, 1, '/assets/images/sku/f67963d7-1a10-448d-884f-470d4a3752fa.jpg', '[{\"keyId\":1,\"key\":\"颜色\",\"valueId\":8,\"value\":\"深空灰色\"},{\"keyId\":5,\"key\":\"版本\",\"valueId\":11,\"value\":\"256GB\"}]', 0, '2020-04-02 08:35:04', '2020-04-03 19:33:54');
INSERT INTO `wechat_mall_sku` VALUES (18, '银色·64G', 6299.00, '', 96, 8, 1, '/assets/images/sku/0dc3c950-3b23-415f-ad67-9524f0247796.jpg', '[{\"keyId\":1,\"key\":\"颜色\",\"valueId\":9,\"value\":\"银色\"},{\"keyId\":5,\"key\":\"版本\",\"valueId\":10,\"value\":\"64GB\"}]', 0, '2020-04-02 08:35:32', '2020-04-03 16:56:39');
INSERT INTO `wechat_mall_sku` VALUES (19, '深空灰色·WLAN版128G', 7899.00, '', 99, 10, 1, '/assets/images/sku/4556e258-c8dd-48aa-8ead-70f82f589d2c.jpg', '[{\"keyId\":1,\"key\":\"颜色\",\"valueId\":8,\"value\":\"深空灰色\"},{\"keyId\":5,\"key\":\"版本\",\"valueId\":16,\"value\":\"WLAN版 128G\"}]', 0, '2020-04-02 09:08:49', '2020-05-03 14:48:48');
INSERT INTO `wechat_mall_sku` VALUES (20, '银色·WLAN版 256G', 0.01, '', 0, 10, 1, '/assets/images/sku/48b66d13-4875-413b-aa79-2a31a08b6ce9.jpg', '[{\"keyId\":1,\"key\":\"颜色\",\"valueId\":9,\"value\":\"银色\"},{\"keyId\":5,\"key\":\"版本\",\"valueId\":17,\"value\":\"WLAN版 256G\"}]', 0, '2020-04-02 09:09:58', '2020-04-02 09:09:58');
INSERT INTO `wechat_mall_sku` VALUES (21, '深空灰色·WLAN版128G', 12499.00, '', 0, 11, 1, '/assets/images/sku/b959c3ad-769d-4992-819d-171f144737ca.jpg', '[{\"keyId\":1,\"key\":\"颜色\",\"valueId\":8,\"value\":\"深空灰色\"},{\"keyId\":5,\"key\":\"版本\",\"valueId\":16,\"value\":\"WLAN版 128G\"}]', 0, '2020-04-02 09:11:14', '2020-04-02 09:11:14');
INSERT INTO `wechat_mall_sku` VALUES (22, '银色·WLAN版256G', 12499.00, '', 100, 11, 1, '/assets/images/sku/c839f4b9-b151-435f-8543-188f33b1fdae.jpg', '[{\"keyId\":1,\"key\":\"颜色\",\"valueId\":9,\"value\":\"银色\"},{\"keyId\":5,\"key\":\"版本\",\"valueId\":17,\"value\":\"WLAN版 256G\"}]', 0, '2020-04-02 09:12:01', '2020-04-02 09:12:01');
INSERT INTO `wechat_mall_sku` VALUES (23, '深空灰色·512G', 18999.00, '', 100, 12, 1, '/assets/images/sku/84df2cff-cd00-498c-88f0-d63591e48985.jpg', '[{\"keyId\":1,\"key\":\"颜色\",\"valueId\":8,\"value\":\"深空灰色\"},{\"keyId\":5,\"key\":\"版本\",\"valueId\":12,\"value\":\"512GB\"}]', 0, '2020-04-02 09:13:51', '2020-04-02 09:13:51');
INSERT INTO `wechat_mall_sku` VALUES (24, '深空灰色', 51999.00, '', 1, 13, 0, '/assets/images/sku/61a16111-3905-46f1-837d-28c8af5cde17.png', '[{\"keyId\":1,\"key\":\"颜色\",\"valueId\":8,\"value\":\"深空灰色\"}]', 0, '2020-04-02 09:14:47', '2020-04-02 09:17:32');
INSERT INTO `wechat_mall_sku` VALUES (25, '新款8代i3', 13388.00, '', 97, 14, 1, '/assets/images/sku/c5ca0623-ff7f-4dc4-98d3-c28826102521.jpg', '[{\"keyId\":1,\"key\":\"颜色\",\"valueId\":8,\"value\":\"深空灰色\"}]', 0, '2020-04-02 09:18:53', '2020-04-02 16:11:31');
INSERT INTO `wechat_mall_sku` VALUES (26, '第五代·白色', 3199.00, '', 100, 15, 1, '/assets/images/sku/6458254e-2e6a-4272-8812-eaa0570c0d69.png', '[{\"keyId\":1,\"key\":\"颜色\",\"valueId\":30,\"value\":\"白色\"}]', 0, '2020-04-02 09:20:52', '2020-04-02 09:20:52');
INSERT INTO `wechat_mall_sku` VALUES (27, '第五代·黑色', 3199.00, '', 100, 15, 1, '/assets/images/sku/6fc15127-e850-4fc4-8ff1-eedc20d491c8.png', '[{\"keyId\":1,\"key\":\"颜色\",\"valueId\":14,\"value\":\"黑色\"}]', 0, '2020-04-02 09:21:22', '2020-04-02 09:21:22');
INSERT INTO `wechat_mall_sku` VALUES (28, '新款·公开版', 1999.00, '', 100, 16, 1, '/assets/images/sku/35c29829-b61f-4b19-8d21-5a13f05b3e61.jpg', '[{\"keyId\":1,\"key\":\"颜色\",\"valueId\":30,\"value\":\"白色\"}]', 0, '2020-04-02 09:23:01', '2020-04-02 09:23:01');
INSERT INTO `wechat_mall_sku` VALUES (29, '银色·鼠标', 564.00, '', 100, 17, 1, '/assets/images/sku/cadb8900-b4c0-4b5a-8d39-cd311c44d13b.jpg', '[{\"keyId\":1,\"key\":\"颜色\",\"valueId\":9,\"value\":\"银色\"}]', 0, '2020-04-02 09:26:06', '2020-04-02 09:26:06');
INSERT INTO `wechat_mall_sku` VALUES (30, '深空灰色·鼠标', 564.00, '', 100, 17, 1, '/assets/images/sku/74830341-e0d5-42f7-8649-8f42ea54691c.jpg', '[{\"keyId\":1,\"key\":\"颜色\",\"valueId\":8,\"value\":\"深空灰色\"}]', 0, '2020-04-02 09:26:38', '2020-04-02 09:26:38');
INSERT INTO `wechat_mall_sku` VALUES (31, '数据线·1米', 135.00, '', 100, 18, 1, '/assets/images/sku/71db4da9-2e58-40e8-808c-0fc8e53f0603.jpg', '[{\"keyId\":10,\"key\":\"接头\",\"valueId\":27,\"value\":\"闪电转USB\"},{\"keyId\":9,\"key\":\"长度\",\"valueId\":25,\"value\":\"1米\"}]', 0, '2020-04-02 09:29:53', '2020-04-02 09:29:53');
INSERT INTO `wechat_mall_sku` VALUES (32, '数据线·2米', 135.00, '', 10, 18, 1, '/assets/images/sku/46ffd860-ff64-4dd8-8f11-40dfd149a84f.jpg', '[{\"keyId\":10,\"key\":\"接头\",\"valueId\":27,\"value\":\"闪电转USB\"},{\"keyId\":9,\"key\":\"长度\",\"valueId\":26,\"value\":\"2米\"}]', 0, '2020-04-02 09:30:36', '2020-04-02 09:30:36');
INSERT INTO `wechat_mall_sku` VALUES (33, '1111', 0.01, '', 0, 7, 1, '/assets/images/sku/52aed024-6287-4102-b604-f3fddfd8856b.jpg', '[{\"keyId\":1,\"key\":\"颜色\",\"valueId\":3,\"value\":\"暗夜绿色\"},{\"keyId\":9,\"key\":\"长度\",\"valueId\":25,\"value\":\"1米\"},{\"keyId\":5,\"key\":\"版本\",\"valueId\":12,\"value\":\"512GB\"}]', 1, '2020-04-07 11:15:16', '2020-04-07 11:15:30');
COMMIT;

-- ----------------------------
-- Table structure for wechat_mall_sku_spec_attr
-- ----------------------------
DROP TABLE IF EXISTS `wechat_mall_sku_spec_attr`;
CREATE TABLE `wechat_mall_sku_spec_attr` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `sku_id` int(11) NOT NULL DEFAULT '0' COMMENT 'sku表主键',
  `spec_id` int(11) NOT NULL DEFAULT '0' COMMENT '规格ID',
  `attr_id` int(11) NOT NULL DEFAULT '0' COMMENT '规格-属性ID',
  `is_del` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除：0-否 1-是',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_sku_id` (`sku_id`),
  KEY `idx_spec_id` (`spec_id`),
  KEY `idx_attr_id` (`attr_id`)
) ENGINE=InnoDB AUTO_INCREMENT=88 DEFAULT CHARSET=utf8mb4 COMMENT='商城-SKU关联的规格属性';

-- ----------------------------
-- Records of wechat_mall_sku_spec_attr
-- ----------------------------
BEGIN;
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (1, 5, 1, 3, 1, '2020-04-01 05:34:55', '2020-04-01 05:34:55');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (2, 5, 5, 10, 1, '2020-04-01 05:34:55', '2020-04-01 05:34:55');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (3, 5, 9, 25, 1, '2020-04-01 05:34:55', '2020-04-01 05:34:55');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (4, 5, 1, 3, 1, '2020-04-01 05:35:51', '2020-04-01 05:35:51');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (5, 5, 5, 10, 1, '2020-04-01 05:35:52', '2020-04-01 05:35:52');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (6, 5, 9, 26, 1, '2020-04-01 05:35:52', '2020-04-01 05:35:52');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (7, 6, 1, 4, 1, '2020-04-01 05:36:01', '2020-04-01 05:36:01');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (8, 6, 5, 12, 1, '2020-04-01 05:36:01', '2020-04-01 05:36:01');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (9, 6, 9, 25, 1, '2020-04-01 05:36:01', '2020-04-01 05:36:01');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (10, 7, 1, 8, 1, '2020-04-01 05:36:04', '2020-04-01 05:36:04');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (11, 7, 5, 10, 1, '2020-04-01 05:36:04', '2020-04-01 05:36:04');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (12, 7, 9, 26, 1, '2020-04-01 05:36:04', '2020-04-01 05:36:04');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (13, 6, 1, 4, 0, '2020-04-01 05:36:16', '2020-04-01 05:36:16');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (14, 6, 5, 12, 0, '2020-04-01 05:36:16', '2020-04-01 05:36:16');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (15, 6, 9, 26, 0, '2020-04-01 05:36:16', '2020-04-01 05:36:16');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (16, 7, 1, 8, 1, '2020-04-01 05:36:26', '2020-04-01 05:36:26');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (17, 7, 5, 10, 1, '2020-04-01 05:36:26', '2020-04-01 05:36:26');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (18, 7, 9, 25, 1, '2020-04-01 05:36:26', '2020-04-01 05:36:26');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (19, 9, 1, 3, 1, '2020-04-01 05:37:08', '2020-04-01 05:37:08');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (20, 9, 5, 11, 1, '2020-04-01 05:37:08', '2020-04-01 05:37:08');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (21, 9, 9, 25, 1, '2020-04-01 05:37:08', '2020-04-01 05:37:08');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (22, 9, 1, 3, 1, '2020-04-01 05:39:47', '2020-04-01 05:39:47');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (23, 9, 5, 11, 1, '2020-04-01 05:39:47', '2020-04-01 05:39:47');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (24, 9, 9, 25, 1, '2020-04-01 05:39:47', '2020-04-01 05:39:47');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (25, 10, 1, 4, 1, '2020-04-01 05:40:32', '2020-04-01 05:40:32');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (26, 10, 5, 16, 1, '2020-04-01 05:40:32', '2020-04-01 05:40:32');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (27, 10, 9, 25, 1, '2020-04-01 05:40:32', '2020-04-01 05:40:32');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (28, 10, 1, 4, 0, '2020-04-01 05:41:25', '2020-04-01 05:41:25');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (29, 10, 5, 16, 0, '2020-04-01 05:41:25', '2020-04-01 05:41:25');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (30, 10, 9, 25, 0, '2020-04-01 05:41:25', '2020-04-01 05:41:25');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (31, 7, 1, 8, 0, '2020-04-01 05:43:13', '2020-04-01 05:43:13');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (32, 7, 5, 10, 0, '2020-04-01 05:43:13', '2020-04-01 05:43:13');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (33, 7, 9, 26, 0, '2020-04-01 05:43:13', '2020-04-01 05:43:13');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (34, 5, 1, 3, 1, '2020-04-01 05:43:26', '2020-04-01 05:43:26');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (35, 5, 5, 10, 1, '2020-04-01 05:43:26', '2020-04-01 05:43:26');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (36, 5, 9, 25, 1, '2020-04-01 05:43:26', '2020-04-01 05:43:26');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (37, 5, 1, 3, 0, '2020-04-01 05:43:40', '2020-04-01 05:43:40');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (38, 5, 5, 10, 0, '2020-04-01 05:43:40', '2020-04-01 05:43:40');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (39, 5, 9, 26, 0, '2020-04-01 05:43:40', '2020-04-01 05:43:40');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (40, 11, 1, 4, 1, '2020-04-01 05:44:11', '2020-04-01 05:44:11');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (41, 11, 5, 10, 1, '2020-04-01 05:44:11', '2020-04-01 05:44:11');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (42, 11, 9, 25, 1, '2020-04-01 05:44:11', '2020-04-01 05:44:11');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (43, 12, 1, 4, 1, '2020-04-02 08:30:54', '2020-04-02 16:55:02');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (44, 12, 5, 10, 1, '2020-04-02 08:30:54', '2020-04-02 16:55:02');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (45, 12, 9, 25, 1, '2020-04-02 08:30:54', '2020-04-02 16:55:02');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (46, 13, 1, 14, 0, '2020-04-02 08:32:08', '2020-04-02 08:32:08');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (47, 13, 5, 10, 0, '2020-04-02 08:32:08', '2020-04-02 08:32:08');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (48, 14, 1, 15, 0, '2020-04-02 08:32:53', '2020-04-02 08:32:53');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (49, 14, 5, 11, 0, '2020-04-02 08:32:53', '2020-04-02 08:32:53');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (50, 15, 1, 13, 0, '2020-04-02 08:33:26', '2020-04-02 08:33:26');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (51, 15, 5, 10, 0, '2020-04-02 08:33:26', '2020-04-02 08:33:26');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (52, 16, 1, 4, 0, '2020-04-02 08:34:14', '2020-04-02 08:34:14');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (53, 16, 5, 10, 0, '2020-04-02 08:34:14', '2020-04-02 08:34:14');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (54, 17, 1, 8, 0, '2020-04-02 08:35:04', '2020-04-02 08:35:04');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (55, 17, 5, 11, 0, '2020-04-02 08:35:04', '2020-04-02 08:35:04');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (56, 18, 1, 9, 0, '2020-04-02 08:35:32', '2020-04-02 08:35:32');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (57, 18, 5, 10, 0, '2020-04-02 08:35:32', '2020-04-02 08:35:32');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (58, 19, 1, 8, 1, '2020-04-02 09:08:49', '2020-04-02 09:09:20');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (59, 19, 5, 16, 1, '2020-04-02 09:08:49', '2020-04-02 09:09:20');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (60, 19, 1, 8, 0, '2020-04-02 09:09:20', '2020-04-02 09:09:20');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (61, 19, 5, 16, 0, '2020-04-02 09:09:20', '2020-04-02 09:09:20');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (62, 20, 1, 9, 0, '2020-04-02 09:09:58', '2020-04-02 09:09:58');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (63, 20, 5, 17, 0, '2020-04-02 09:09:58', '2020-04-02 09:09:58');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (64, 21, 1, 8, 0, '2020-04-02 09:11:14', '2020-04-02 09:11:14');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (65, 21, 5, 16, 0, '2020-04-02 09:11:14', '2020-04-02 09:11:14');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (66, 22, 1, 9, 0, '2020-04-02 09:12:01', '2020-04-02 09:12:01');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (67, 22, 5, 17, 0, '2020-04-02 09:12:01', '2020-04-02 09:12:01');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (68, 23, 1, 8, 0, '2020-04-02 09:13:51', '2020-04-02 09:13:51');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (69, 23, 5, 12, 0, '2020-04-02 09:13:51', '2020-04-02 09:13:51');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (70, 24, 1, 8, 1, '2020-04-02 09:14:47', '2020-04-02 09:17:31');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (71, 24, 1, 8, 0, '2020-04-02 09:17:32', '2020-04-02 09:17:32');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (72, 25, 1, 8, 0, '2020-04-02 09:18:53', '2020-04-02 09:18:53');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (73, 26, 1, 30, 0, '2020-04-02 09:20:52', '2020-04-02 09:20:52');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (74, 27, 1, 14, 0, '2020-04-02 09:21:22', '2020-04-02 09:21:22');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (75, 28, 1, 30, 0, '2020-04-02 09:23:01', '2020-04-02 09:23:01');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (76, 29, 1, 9, 0, '2020-04-02 09:26:06', '2020-04-02 09:26:06');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (77, 30, 1, 8, 0, '2020-04-02 09:26:38', '2020-04-02 09:26:38');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (78, 31, 10, 27, 0, '2020-04-02 09:29:53', '2020-04-02 09:29:53');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (79, 31, 9, 25, 0, '2020-04-02 09:29:53', '2020-04-02 09:29:53');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (80, 32, 10, 27, 0, '2020-04-02 09:30:36', '2020-04-02 09:30:36');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (81, 32, 9, 26, 0, '2020-04-02 09:30:36', '2020-04-02 09:30:36');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (82, 12, 1, 9, 0, '2020-04-02 16:55:02', '2020-04-02 16:55:02');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (83, 12, 9, 25, 0, '2020-04-02 16:55:02', '2020-04-02 16:55:02');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (84, 12, 5, 10, 0, '2020-04-02 16:55:02', '2020-04-02 16:55:02');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (85, 33, 1, 3, 1, '2020-04-07 11:15:16', '2020-04-07 11:15:30');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (86, 33, 9, 25, 1, '2020-04-07 11:15:16', '2020-04-07 11:15:30');
INSERT INTO `wechat_mall_sku_spec_attr` VALUES (87, 33, 5, 12, 1, '2020-04-07 11:15:16', '2020-04-07 11:15:30');
COMMIT;

-- ----------------------------
-- Table structure for wechat_mall_specification
-- ----------------------------
DROP TABLE IF EXISTS `wechat_mall_specification`;
CREATE TABLE `wechat_mall_specification` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(30) NOT NULL DEFAULT '' COMMENT '规格名名称',
  `description` varchar(30) NOT NULL DEFAULT '' COMMENT '规格名描述',
  `unit` varchar(10) NOT NULL DEFAULT '' COMMENT '单位',
  `standard` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否标准: 0-非标准 1-标准',
  `is_del` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除：0-否 1-是',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COMMENT='商城-规格表';

-- ----------------------------
-- Records of wechat_mall_specification
-- ----------------------------
BEGIN;
INSERT INTO `wechat_mall_specification` VALUES (1, '颜色', '颜色属性（通用）', '无', 0, 0, '2020-03-12 12:36:35', '2020-03-29 07:40:58');
INSERT INTO `wechat_mall_specification` VALUES (2, '尺寸', '衣服尺寸', '箱', 1, 1, '2020-03-12 12:36:48', '2020-03-29 07:28:02');
INSERT INTO `wechat_mall_specification` VALUES (3, '型号', '', '个', 1, 1, '2020-03-15 10:15:38', '2020-03-15 10:16:27');
INSERT INTO `wechat_mall_specification` VALUES (4, '型号', '规格-型号', '个', 1, 1, '2020-03-15 10:16:40', '2020-03-29 07:28:05');
INSERT INTO `wechat_mall_specification` VALUES (5, '版本', '存储容量版本', 'GB', 0, 0, '2020-03-15 10:17:06', '2020-03-29 07:26:37');
INSERT INTO `wechat_mall_specification` VALUES (6, '123', '', 'dfa', 0, 1, '2020-03-18 03:31:33', '2020-03-18 03:32:11');
INSERT INTO `wechat_mall_specification` VALUES (7, '12334', '431', 'dfa', 1, 1, '2020-03-18 03:32:19', '2020-03-18 03:32:45');
INSERT INTO `wechat_mall_specification` VALUES (8, '123', '', '333', 0, 1, '2020-03-18 03:33:44', '2020-03-18 03:33:46');
INSERT INTO `wechat_mall_specification` VALUES (9, '长度', '标准长度', '米', 1, 0, '2020-03-29 09:09:09', '2020-03-29 09:09:16');
INSERT INTO `wechat_mall_specification` VALUES (10, '接头', '数据线接头类型', '无', 0, 0, '2020-03-29 09:10:16', '2020-04-01 05:56:37');
COMMIT;

-- ----------------------------
-- Table structure for wechat_mall_specification_attr
-- ----------------------------
DROP TABLE IF EXISTS `wechat_mall_specification_attr`;
CREATE TABLE `wechat_mall_specification_attr` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `spec_id` int(11) NOT NULL DEFAULT '0' COMMENT '规格ID',
  `value` varchar(20) NOT NULL DEFAULT '' COMMENT '属性值',
  `extend` varchar(30) NOT NULL DEFAULT '' COMMENT '扩展',
  `is_del` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除：0-否 1-是',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_spec_id` (`spec_id`),
  KEY `idx_value` (`value`)
) ENGINE=InnoDB AUTO_INCREMENT=31 DEFAULT CHARSET=utf8mb4 COMMENT='商城-规格属性表';

-- ----------------------------
-- Records of wechat_mall_specification_attr
-- ----------------------------
BEGIN;
INSERT INTO `wechat_mall_specification_attr` VALUES (1, 1, '青芒色', '', 1, '2020-03-12 12:39:41', '2020-03-15 11:45:05');
INSERT INTO `wechat_mall_specification_attr` VALUES (2, 2, '7英寸', '', 1, '2020-03-12 12:40:16', '2020-03-29 07:27:42');
INSERT INTO `wechat_mall_specification_attr` VALUES (3, 1, '暗夜绿色', '', 0, '2020-03-15 12:00:46', '2020-03-29 07:24:00');
INSERT INTO `wechat_mall_specification_attr` VALUES (4, 1, '金色', '', 0, '2020-03-15 12:01:40', '2020-03-29 07:24:14');
INSERT INTO `wechat_mall_specification_attr` VALUES (5, 1, '紫色', '123', 1, '2020-03-15 12:02:12', '2020-03-15 12:02:19');
INSERT INTO `wechat_mall_specification_attr` VALUES (6, 2, '3.5英寸', '', 1, '2020-03-15 12:03:25', '2020-03-29 07:27:40');
INSERT INTO `wechat_mall_specification_attr` VALUES (7, 4, 'ABC', '', 1, '2020-03-16 13:06:02', '2020-03-29 07:27:54');
INSERT INTO `wechat_mall_specification_attr` VALUES (8, 1, '深空灰色', '', 0, '2020-03-29 07:24:23', '2020-03-29 07:24:23');
INSERT INTO `wechat_mall_specification_attr` VALUES (9, 1, '银色', '', 0, '2020-03-29 07:24:29', '2020-03-29 07:24:29');
INSERT INTO `wechat_mall_specification_attr` VALUES (10, 5, '64GB', '', 0, '2020-03-29 07:26:50', '2020-03-29 07:26:50');
INSERT INTO `wechat_mall_specification_attr` VALUES (11, 5, '256GB', '', 0, '2020-03-29 07:26:58', '2020-03-29 07:26:58');
INSERT INTO `wechat_mall_specification_attr` VALUES (12, 5, '512GB', '', 0, '2020-03-29 07:27:05', '2020-03-29 07:27:05');
INSERT INTO `wechat_mall_specification_attr` VALUES (13, 1, '蓝色', '', 0, '2020-03-29 07:40:36', '2020-03-29 07:40:36');
INSERT INTO `wechat_mall_specification_attr` VALUES (14, 1, '黑色', '', 0, '2020-03-29 07:40:41', '2020-03-29 07:40:41');
INSERT INTO `wechat_mall_specification_attr` VALUES (15, 1, '红色', '', 0, '2020-03-29 07:40:45', '2020-03-29 07:40:45');
INSERT INTO `wechat_mall_specification_attr` VALUES (16, 5, 'WLAN版 128G', '', 0, '2020-03-29 07:51:53', '2020-03-29 07:51:53');
INSERT INTO `wechat_mall_specification_attr` VALUES (17, 5, 'WLAN版 256G', '', 0, '2020-03-29 07:52:03', '2020-03-29 07:52:03');
INSERT INTO `wechat_mall_specification_attr` VALUES (18, 5, 'WLAN版 512G', '', 0, '2020-03-29 07:52:13', '2020-03-29 07:52:13');
INSERT INTO `wechat_mall_specification_attr` VALUES (19, 5, 'Cellular版 128G', '', 0, '2020-03-29 07:53:56', '2020-03-29 07:53:56');
INSERT INTO `wechat_mall_specification_attr` VALUES (20, 5, 'Cellular版 256G', '', 0, '2020-03-29 07:54:15', '2020-03-29 07:54:15');
INSERT INTO `wechat_mall_specification_attr` VALUES (21, 5, 'Cellular版 512G', '', 0, '2020-03-29 07:54:21', '2020-03-29 07:54:21');
INSERT INTO `wechat_mall_specification_attr` VALUES (22, 1, '16英寸 九代i7+512灰', '', 0, '2020-03-29 08:08:02', '2020-03-29 08:08:02');
INSERT INTO `wechat_mall_specification_attr` VALUES (23, 1, '16英寸 九代i7+512银', '', 0, '2020-03-29 08:08:53', '2020-03-29 08:08:53');
INSERT INTO `wechat_mall_specification_attr` VALUES (24, 1, '21.5英寸4K 8代i3', '', 0, '2020-03-29 08:56:07', '2020-03-29 08:56:07');
INSERT INTO `wechat_mall_specification_attr` VALUES (25, 9, '1米', '', 0, '2020-03-29 09:09:30', '2020-03-29 09:09:30');
INSERT INTO `wechat_mall_specification_attr` VALUES (26, 9, '2米', '', 0, '2020-03-29 09:09:34', '2020-04-01 04:54:06');
INSERT INTO `wechat_mall_specification_attr` VALUES (27, 10, '闪电转USB', '', 0, '2020-03-29 09:10:39', '2020-04-01 05:56:28');
INSERT INTO `wechat_mall_specification_attr` VALUES (28, 10, 'USB-C充电线', '', 0, '2020-03-29 09:10:47', '2020-04-01 05:56:30');
INSERT INTO `wechat_mall_specification_attr` VALUES (29, 10, '磁力充电器', '', 0, '2020-03-29 09:11:16', '2020-04-01 05:56:32');
INSERT INTO `wechat_mall_specification_attr` VALUES (30, 1, '白色', '', 0, '2020-04-02 09:19:56', '2020-04-02 09:19:56');
COMMIT;

-- ----------------------------
-- Table structure for wechat_mall_user
-- ----------------------------
DROP TABLE IF EXISTS `wechat_mall_user`;
CREATE TABLE `wechat_mall_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `openid` varchar(50) NOT NULL DEFAULT '' COMMENT '微信openid',
  `nickname` varchar(30) NOT NULL DEFAULT '' COMMENT '昵称',
  `avatar` varchar(200) NOT NULL DEFAULT '' COMMENT '微信头像',
  `mobile` varchar(11) NOT NULL DEFAULT '' COMMENT '手机号',
  `city` varchar(30) NOT NULL DEFAULT '' COMMENT '城市',
  `province` varchar(30) NOT NULL DEFAULT '' COMMENT '省份',
  `country` varchar(30) NOT NULL DEFAULT '' COMMENT '国家',
  `gender` tinyint(1) NOT NULL DEFAULT '0' COMMENT '性别 0：未知、1：男、2：女',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_openid` (`openid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='小程序-用户表';

-- ----------------------------
-- Table structure for wechat_mall_user_address
-- ----------------------------
DROP TABLE IF EXISTS `wechat_mall_user_address`;
CREATE TABLE `wechat_mall_user_address` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `contacts` varchar(15) NOT NULL DEFAULT '' COMMENT '联系人',
  `mobile` varchar(11) NOT NULL DEFAULT '' COMMENT '手机号',
  `province_id` varchar(10) NOT NULL DEFAULT '' COMMENT '省份编码',
  `city_id` varchar(10) NOT NULL DEFAULT '' COMMENT '城市编码',
  `area_id` varchar(10) NOT NULL DEFAULT '' COMMENT '地区编码',
  `province_str` varchar(10) NOT NULL DEFAULT '' COMMENT '省份',
  `city_str` varchar(10) NOT NULL DEFAULT '' COMMENT '城市',
  `area_str` varchar(10) NOT NULL DEFAULT '' COMMENT '地区',
  `address` varchar(30) NOT NULL DEFAULT '' COMMENT '详细地址',
  `is_default` tinyint(1) NOT NULL DEFAULT '0' COMMENT '默认收货地址：0-否 1-是',
  `is_del` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除：0-否 1-是',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商城-用户收货地址表';

-- ----------------------------
-- Table structure for wechat_mall_user_cart
-- ----------------------------
DROP TABLE IF EXISTS `wechat_mall_user_cart`;
CREATE TABLE `wechat_mall_user_cart` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `goods_id` int(11) NOT NULL DEFAULT '0' COMMENT '商品ID',
  `sku_id` int(11) NOT NULL DEFAULT '0' COMMENT 'sku ID',
  `num` int(11) NOT NULL DEFAULT '0' COMMENT '数量',
  `is_del` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除：0-否 1-是',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_goods_id` (`goods_id`),
  KEY `idx_sku_id` (`sku_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商城-购物车表';

-- ----------------------------
-- Table structure for wechat_mall_visitor_record
-- ----------------------------
DROP TABLE IF EXISTS `wechat_mall_visitor_record`;
CREATE TABLE `wechat_mall_visitor_record` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '平台用户ID',
  `ip` varchar(20) NOT NULL DEFAULT '' COMMENT '独立IP',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商城-访客记录表';

SET FOREIGN_KEY_CHECKS = 1;
