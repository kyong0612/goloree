CREATE TABLE `users` (
  `user_no` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '氏名',
  `address` varchar(254) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `token` text COLLATE utf8mb4_unicode_ci COMMENT '再認証トークン',
  `dept_code` char(4) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '部担コード',
  `scope` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '職制',
  `version` int(11) NOT NULL DEFAULT '0' COMMENT 'バージョン',
  `created_at` datetime NOT NULL COMMENT '登録日時',
  `updated_at` datetime NOT NULL COMMENT '更新日時',
  `deleted_at` datetime DEFAULT NULL COMMENT '削除日時',
  PRIMARY KEY (`user_no`),
  KEY `users_name_IDX` (`name`) USING BTREE,
  KEY `users_address_IDX` (`address`) USING BTREE,
  KEY `users_dept_code_IDX` (`dept_code`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;