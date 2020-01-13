-- phpMyAdmin SQL Dump
-- version 4.9.2
-- https://www.phpmyadmin.net/
--
-- Máy chủ: 127.0.0.1
-- Thời gian đã tạo: Th1 12, 2020 lúc 04:29 AM
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
-- Cấu trúc bảng cho bảng `notification`
--

CREATE TABLE `notification` (
  `id` int(11) NOT NULL,
  `title` varchar(255) NOT NULL,
  `content` text NOT NULL,
  `image_path` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Đang đổ dữ liệu cho bảng `notification`
--

INSERT INTO `notification` (`id`, `title`, `content`, `image_path`) VALUES
(1, 'Khuyến mãi nạp thẻ', 'Từ ngày 13/01/2020 đến 25/01/2020 nạp thẻ điện thoại bằng ví điện tử Vinid có cơ hội nhận khuyến mãi lên đến 505. Nhanh tay nạp tiền cho dế yêu của bạn đi nào!', 'https://vinid.net/wp-content/uploads/2019/12/VinID-Topup-KV_thang-1_Banner-web-1.jpg'),
(2, 'Chợ tết sập giá', 'Chào xuân Canh Tí, không canh là phí. Vinid mở chương trình khuyến mãi chào xuân với các mặt hàng đồng giá chỉ còn 3 nghìn đồng, Nhanh tay săn hàng cùng Vinid và Scan&Go nào', 'https://vinid.net/wp-content/uploads/2020/01/2019-01-03-cho-tet-dong-gia-3k_BannerWeb_1920X1080px-2.jpg'),
(3, 'Thẻ quà tặng', 'Chế độ chính sách đặc biệt với những khác hàng sử dụng thẻ quà tặng. Chiết khấu cực khủng với các hóa đơn trên 20 triệu, xuất hóa đơn vat 10%, tiêu dùng tại hơn 2000 điểm kinh doanh của các cửa hàng.', 'https://vinid.net/wp-content/uploads/2019/12/20191204_Giftcard-TET_Banner-Web_1920x1080.jpg');

--
-- Chỉ mục cho các bảng đã đổ
--

--
-- Chỉ mục cho bảng `notification`
--
ALTER TABLE `notification`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT cho các bảng đã đổ
--

--
-- AUTO_INCREMENT cho bảng `notification`
--
ALTER TABLE `notification`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
