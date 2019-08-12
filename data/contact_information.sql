CREATE TABLE `contact_information` (
  `id` int(11) NOT NULL,
  `user_id` varchar(0) DEFAULT NULL,
  `contact_persion_id` varchar(0) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;