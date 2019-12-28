-- phpMyAdmin SQL Dump
-- version 4.9.2
-- https://www.phpmyadmin.net/
--
-- Máy chủ: 127.0.0.1
-- Thời gian đã tạo: Th12 28, 2019 lúc 04:09 PM
-- Phiên bản máy phục vụ: 10.4.10-MariaDB
-- Phiên bản PHP: 7.3.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Cơ sở dữ liệu: `vinid_project`
--

-- --------------------------------------------------------

--
-- Cấu trúc bảng cho bảng `store`
--

CREATE TABLE `store` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `image_path` varchar(50) NOT NULL,
  `address` varchar(255) NOT NULL,
  `latitude` decimal(20,10) NOT NULL,
  `longitude` decimal(20,10) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Đang đổ dữ liệu cho bảng `store`
--

INSERT INTO `store` (`id`, `name`, `image_path`, `address`, `latitude`, `longitude`, `created_at`) VALUES
(1, 'VinMart+', '/resources/store_images/1.png', 'Nhà No9 KĐT, Pháp Vân, Hoàng Ma,Hà Nội', '20.9588310000', '105.8467490000', '2019-12-27 12:24:21'),
(2, 'Vinmart+', '/resources/store_images/2.png', 'Ô 13, khu, Ngõ 2 - Hoàng Liệt, Hoàng Liệt, Hoàng Mai, Hà Nội', '20.9666340000', '105.8376770000', '2019-12-27 12:27:10'),
(3, 'Vinmart+', '/resources/store_images/3.png', 'Ô số 13A, lô Ơ2, bán đảo, Khu đô thị Linh Đàm, Hoàng Mai, Hà Nội', '20.9676220000', '105.8333100000', '2019-12-27 12:30:30'),
(4, 'Vinmart+', '/resources/store_images/4.png', 'Nguyễn Duy Trinh, Khu đô thị Linh Đàm, Hoàng Mai, Hà Nội', '20.9641100000', '105.8318960000', '2019-12-27 12:32:13'),
(5, 'Vinmart+', '/resources/store_images/5.png', 'Toàn Nhà Hud 3, Linh Đàm, Hoàng Liệt, Thanh Xuân, Hà Nội', '20.9645230000', '105.8250960000', '2019-12-27 12:36:55'),
(6, 'Vinmart+', '/resources/store_images/6.png', 'Số 10, ngõ 15 phố Bằng Liệt, Tây Nam Linh Đàm, Hoàng Mai, Hà Nội', '20.9646810000', '105.8208780000', '2019-12-27 12:45:49'),
(7, 'Vinmart+', '/resources/store_images/7.png', '182 Đại Từ, Đại Kim, Hoàng Mai, Hà Nội', '20.9713430000', '105.8335200000', '2019-12-27 12:46:38'),
(8, 'Vinmart+', '/resources/store_images/8.png', '18 Cầu Dậu, huyện Thanh Trì, Thanh Trì, Thanh Trì Hà Nội', '20.9693050000', '105.8236250000', '2019-12-27 12:47:45'),
(9, 'Vinmart+', '/resources/store_images/9.png', '532 Đ. Kim Giang, Thanh Liệt, Thanh Trì, Hà Nội', '20.9691560000', '105.8234840000', '2019-12-27 12:48:32'),
(10, 'Vinmart+', '/resources/store_images/10.png', 'Số 19 ngõ 42 phố, Thịnh Liệt, Hoàng Mai, Hà Nội', '20.9739070000', '105.8457300000', '2019-12-27 12:49:53'),
(11, 'Vinmart+', '/resources/store_images/11.png', 'chung cư, Tầng 1, đường Vũ Tông Phan, Đại Kim, Hoàng Mai, Hà Nội', '20.9755850000', '105.8255910000', '2019-12-27 12:51:02'),
(12, 'Vinmart+', '/resources/store_images/12.png', '100 Ngõ 168 Đường Kim Giang, Đại Kim, Thanh Xuân, Hà Nội', '20.9793400000', '105.8156630000', '2019-12-27 12:52:53'),
(13, 'Vinmart+', '/resources/store_images/13.png', 'Số 142 - Kim Giang - Hoàng Mai - Hà Nội, Kim Giang - Ngõ 142 Ngõ 168 Đường Kim Giang, Đại Kim, Hoàng Mai, Hà Nội', '20.9809850000', '105.8178850000', '2019-12-27 12:54:01'),
(14, 'Vinmart+', '/resources/store_images/14.png', '639 Vũ Tông Phan, Khương Đình, Thanh Xuân, Hà Nội', '20.9828780000', '105.8165920000', '2019-12-27 12:54:49'),
(15, 'Vinmart+', '/resources/store_images/15.png', 'Ngõ 245 Định Công, Định Công, Thanh Xuân, Hà Nội', '20.9811160000', '105.8334550000', '2019-12-27 12:55:51'),
(16, 'Vinmart+', '/resources/store_images/1.png', '261, Tân Mai, Quận Hoàng Mai, Thành Phố Hà Nội, Tân Mai, Hoàng Mai, Hà Nội', '20.9838990000', '105.8508880000', '2019-12-27 12:58:19'),
(17, 'Vinmart+', '/resources/store_images/2.png', '589 Trương Định, Giáp Bát, Hai Bà Trưng, Hà Nội', '20.9814760000', '105.8454340000', '2019-12-27 13:00:09'),
(18, 'Vinmart+', '/resources/store_images/3.png', '1 Kim Đồng, Giáp Bát, Hoàng Mai, Hà Nội', '20.9835140000', '105.8437820000', '2019-12-27 13:01:25'),
(19, 'Vinmart+', '/resources/store_images/4.png', 'Tòa nhà Tiến Phú, Hoàng Văn Thụ, Hoàng Mai, Hà Nội', '20.9837550000', '105.8665080000', '2019-12-27 13:02:24'),
(20, 'Vinmart+', '/resources/store_images/5.png', '59 Đền Lừ 2, Hoàng Văn Thụ, Hoàng Mai, Hà Nội', '20.9860900000', '105.8578210000', '2019-12-27 13:03:31'),
(21, 'Vinmart+', '/resources/store_images/6.png', '55A Ngõ 1 Mai Động, Mai Động, Hoàng Mai, Hà Nội', '20.9917770000', '105.8625340000', '2019-12-27 13:05:07'),
(22, 'Vinmart+', '/resources/store_images/7.png', '33 Lương Khánh Thiện, Tương Mai, Hoàng Mai, Hà Nội', '20.9862590000', '105.8510130000', '2019-12-27 13:06:08'),
(23, 'Vinmart+', '/resources/store_images/8.png', '62 Nguyễn Đức Cảnh, Tân Mai, Hoàng Mai, Hà Nội', '20.9873080000', '105.8498940000', '2019-12-27 13:07:07'),
(24, 'Vinmart+', '/resources/store_images/9.png', 'Toà nhà A1, Nguyễn Đức Cảnh, Ngõ 151 Hanoi Hà Nội, ngõ 151 Nguyễn Đức Cảnh, Tương Mai, Hoàng Mai, Hà Nội', '20.9880790000', '105.8517070000', '2019-12-27 13:08:11'),
(25, 'Vinmart+', '/resources/store_images/10.png', '110 Ngõ 553Giải Phóng, Đường Giáp Bát, Giáp Bát, Hai Bà Trưng, Hà Nội', '20.9901100000', '105.8426610000', '2019-12-27 13:09:35'),
(26, 'Vinmart+', '/resources/store_images/11.png', '164 Trương Định, Hoàng Mai, Hà Nội', '20.9917000000', '105.8489610000', '2019-12-27 13:10:45'),
(27, 'Vinmart+', '/resources/store_images/12.png', '194 Phố Minh Khai, Minh Khai, Hai Bà Trưng, Hà Nội', '20.9952770000', '105.8542700000', '2019-12-27 13:11:47'),
(28, 'Vinmart+', '/resources/store_images/13.png', '91 Vĩnh Hưng, Hoàng Mai, Hà Nội', '20.9859350000', '105.8759850000', '2019-12-27 13:14:02'),
(29, 'Vinmart+', '/resources/store_images/14.png', '283 Vĩnh Hưng, Hoàng Mai, Hà Nội', '20.9911000000', '105.8789580000', '2019-12-27 13:14:58'),
(31, 'Vinmart+', '/resources/store_images/1.png', 'Tầng 1 toà Park 6, Khu đô thị Times City, Hai Bà Trưng, Hà Nội', '20.9913360000', '105.8685800000', '2019-12-27 13:18:15'),
(32, 'Vinmart+', '/resources/store_images/2.png', ' Park 1 Times City, 12, Ngõ 454 Minh Khai, Mai Động, Hai Bà Trưng, Hà Nội', '20.9931760000', '105.8678730000', '2019-12-27 13:19:02'),
(33, 'Vinmart+', '/resources/store_images/3.png', 'T10, Khu đô thị Times City, Hai Bà Trưng, Hà Nội', '20.9942020000', '105.8685080000', '2019-12-27 13:19:44'),
(34, 'Vinmart+', '/resources/store_images/4.png', 'L1-03, 458 Phố Minh Khai, Khu đô thị Times City, Hai Bà Trưng, Hà Nội', '20.9961150000', '105.8694090000', '2019-12-27 13:20:29'),
(35, 'Cửa Hàng Tiện Ích VinMart', '/resources/store_images/5.png', ' 79 phố Dương Văn Bé, Vĩnh Phú, Hai Bà Trưng, Hà Nội', '20.9992190000', '105.8730850000', '2019-12-27 13:21:44'),
(36, 'Cửa Hàng Vinmart+', '/resources/store_images/6.png', '31 Mạc Thị Bưởi, Vĩnh Tuy, Hai Bà Trưng, Hà Nội', '21.0007330000', '105.8691160000', '2019-12-27 13:22:39'),
(37, 'Cửa Hàng Vinmart+', '/resources/store_images/7.png', ' số 49B Ngõ 651 - Phố Minh Khai, Thanh Lương, Hai Bà Trưng, Hà Nội', '21.0034800000', '105.8682740000', '2019-12-27 13:23:34'),
(38, 'Cửa Hàng Vinmart+', '/resources/store_images/8.png', 'Ngõ 230 Lạc Trung, Thanh Lương, Hai Bà Trưng, Hà Nội', '21.0033190000', '105.8667530000', '2019-12-27 13:24:25'),
(39, 'Vinmart+', '/resources/store_images/9.png', '183 Hồng Mai, Quỳnh Lôi, Hai Bà Trưng, Hà Nội', '20.9989710000', '105.8555200000', '2019-12-27 13:26:49'),
(40, 'Vinmart+', '/resources/store_images/10.png', '69 Hồng Mai, Trương Định, Hai Bà Trưng, Hà Nội', '21.0010090000', '105.8506770000', '2019-12-27 13:27:53'),
(41, 'Vinmart+', '/resources/store_images/11.png', '47, 52 ngõ 187 Hồng Mai, Bạch Mai, Hai Bà Trưng, Hà Nội', '21.0004740000', '105.8556570000', '2019-12-27 13:28:51'),
(42, 'Vinmart+', '/resources/store_images/12.png', '409 Bạch Mai, Hai Bà Trưng, Hà Nội', '21.0013190000', '105.8506990000', '2019-12-27 13:29:48'),
(43, 'Vinmart+', '/resources/store_images/13.png', '244 Lê Thanh Nghị, Đồng Tâm, Hai Bà Trưng, Hà Nội', '21.0019800000', '105.8424290000', '2019-12-27 13:30:33'),
(44, 'Vinmart+', '/resources/store_images/14.png', '171 Giải Phóng, Phương Mai, Đống Đa, Hà Nội', '21.0014600000', '105.8413090000', '2019-12-27 13:31:14');

--
-- Chỉ mục cho các bảng đã đổ
--

--
-- Chỉ mục cho bảng `store`
--
ALTER TABLE `store`
  ADD PRIMARY KEY (`id`),
  ADD KEY `index_address` (`address`);

--
-- AUTO_INCREMENT cho các bảng đã đổ
--

--
-- AUTO_INCREMENT cho bảng `store`
--
ALTER TABLE `store`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=45;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
