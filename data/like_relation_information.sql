CREATE TABLE `like_relation_information` (
  `id` int(11) NOT NULL,
  `user_id` char(255) NOT NULL,
  `target_user_id` char(255) NOT NULL,
  `is_matched` tinyint(1) NOT NULL,
PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
