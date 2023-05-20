CREATE TABLE `users` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `first_name` varchar(255),
  `last_name` varchar(255),
  `email` varchar(255) UNIQUE,
  `password` varchar(255),
  `created_at` timestamp,
  `updated_at` timestamp
);

CREATE TABLE `products` (
  `id` varchar(255) PRIMARY KEY,
  `name` varchar(255),
  `description` varchar(255),
  `price` decimal,
  `quantity` int,
  `category_id` json,
  `created_at` timestamp,
  `updated_at` timestamp
);

CREATE TABLE `categories` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255),
  `description` varchar(255),
  `created_at` timestamp,
  `updated_at` timestamp
);

CREATE TABLE `orders` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `user_id` int,
  `coupon_id` int,
  `total` decimal,
  `status` varchar(255),
  `created_at` timestamp,
  `updated_at` timestamp
);

CREATE TABLE `order_items` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `order_id` int,
  `product_id` varchar(255),
  `quantity` int,
  `price` decimal,
  `created_at` timestamp,
  `updated_at` timestamp
);

CREATE TABLE `carts` (
  `id` varchar(255) PRIMARY KEY,
  `user_id` int,
  `items` json,
  `created_at` timestamp,
  `updated_at` timestamp
);

CREATE TABLE `cart_items` (
  `id` varchar(255) PRIMARY KEY,
  `cart_id` varchar(255),
  `product_id` varchar(255),
  `quantity` int,
  `price` decimal,
  `created_at` timestamp,
  `updated_at` timestamp
);

CREATE TABLE `reviews` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `product_id` varchar(255),
  `user_id` int,
  `title` varchar(255),
  `body` text,
  `rating` int,
  `created_at` timestamp,
  `updated_at` timestamp
);

CREATE TABLE `coupons` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `code` varchar(255) UNIQUE,
  `type` varchar(255),
  `value` decimal,
  `created_at` timestamp,
  `updated_at` timestamp
);

CREATE TABLE `addresses` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `user_id` int,
  `line1` varchar(255),
  `line2` varchar(255),
  `city` varchar(255),
  `state` varchar(255),
  `country` varchar(255),
  `zip_code` varchar(255),
  `created_at` timestamp,
  `updated_at` timestamp
);
