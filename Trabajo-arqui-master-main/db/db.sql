-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Servidor: localhost
-- Tiempo de generación: 19-07-2022 a las 04:07:57
-- Versión del servidor: 10.4.21-MariaDB
-- Versión de PHP: 7.4.29

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Base de datos: `Pruebas`
--


-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `orders`
--

CREATE TABLE `orders` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `address_id` int(11) NOT NULL,
  `date` date NOT NULL,
  `total` decimal(10,0) NOT NULL,
  `currency_id` varchar(10) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Volcado de datos para la tabla `orders`
--

INSERT INTO `orders` (`id`, `user_id`, `address_id`, `date`, `total`, `currency_id`) VALUES
(1, 1, 0, '2022-07-14', '19978', 'ARS'),
(19, 1, 0, '2022-07-14', '173499', 'ARS'),
(20, 1, 0, '2022-07-14', '3229', 'ARS');


-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `order_details`
--

CREATE TABLE `order_details` (
  `order_detail_id` int(11) DEFAULT NULL,
  `order_id` int(11) DEFAULT NULL,
  `product_id` int(11) DEFAULT NULL,
  `quantity` int(11) NOT NULL,
  `price` decimal(10,0) NOT NULL,
  `currency_id` varchar(10) NOT NULL,
  `name` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Volcado de datos para la tabla `order_details`
--

INSERT INTO `order_details` (`order_detail_id`, `order_id`, `product_id`, `quantity`, `price`, `currency_id`, `name`) VALUES
(NULL, 1, 261, 2, '9989', 'ARS', 'Boxer Hombre Briganti Algod?n Combinado Pack X 6 Hxim09524'),
(NULL, 19, 55, 1, '173499', 'ARS', 'Apple iPhone 11 (128 GB) - Blanco'),
(NULL, 20, 47, 1, '3229', 'ARS', 'Alcatel 3H 64 GB elegant black 3 GB RAM');


--
-- Estructura de tabla para la tabla `users`
--

CREATE TABLE `users` (
  `user_id` int(11) DEFAULT NULL,
  `username` varchar(40) NOT NULL,
  `password` varchar(255) NOT NULL,
  `first_name` varchar(100) NOT NULL,
  `last_name` varchar(100) NOT NULL,
  `email` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Volcado de datos para la tabla `users`
--

INSERT INTO `users` (`user_id`, `username`, `password`, `first_name`, `last_name`, `email`) VALUES
(1, 'juli912', 'river123', 'Julian', 'Gergolet', '2012598@ucc.edu.ar'),
(2, 'franll', 'fran1', 'Fransisco', 'Llables', '2012791@ucc.edu.ar'),
(3, 'sant1', 'santiagoaguero', 'Santiago', 'Aguero', '2212471@ucc.edu.ar');

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `products`
--

CREATE TABLE `products` (
  `product_id` int(11) DEFAULT NULL,
  `name` varchar(100) NOT NULL,
  `price` decimal(10,0) NOT NULL,
  `picture` varchar(100) NOT NULL,
  `category_id` int(11) DEFAULT NULL,
  `stock` int(11) NOT NULL,
  `description` varchar(255) NOT NULL,
  `currency_id` varchar(10) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Volcado de datos para la tabla `products`
--

INSERT INTO `products` (`product_id`, `name`, `price`, `picture`, `category_id`, `stock`, `description`, `currency_id`) VALUES
(174, 'Adaptador Enchufe Viajero Ingl?s Uk Europeo Subte A', '164', 'https://http2.mlstatic.com/D_NQ_NP_721733-MLA47919939567_102021-V.webp', 3, 34, '', ''),
(47, 'Alcatel 3H 64 GB elegant black 3 GB RAM', '3229', 'https://http2.mlstatic.com/D_NQ_NP_705727-MLA47762780436_102021-V.webp', 2, 30, '', ''),
(196, 'Alfajor Deluxe Bla Bla Arroz Choco Bl Y Dulce De Leche 28gr', '63', 'https://http2.mlstatic.com/D_NQ_NP_886131-MLA47263301400_082021-W.webp', 7, 19, '', ''),
(226, 'Alfajor Simple Milka Mousse X 42 Gr', '81', 'https://http2.mlstatic.com/D_NQ_NP_827777-MLA47778215710_102021-W.webp', 7, 11, '', ''),
(238, 'Almendras Peladas X 1kg - Envios A Todo El Pais', '1842', 'https://http2.mlstatic.com/D_NQ_NP_885131-MLA49874625347_052022-W.webp', 7, 16, '', ''),
(121, 'APC Back-UPS BR900G-AR 900VA entrada y salida de 230V negro', '44599', 'https://http2.mlstatic.com/D_NQ_NP_607316-MLA44132992420_112020-V.webp', 8, 13, '', ''),
(55, 'Apple iPhone 11 (128 GB) - Blanco', '173499', 'https://http2.mlstatic.com/D_NQ_NP_796892-MLA46114829828_052021-V.webp', 2, 36, '', ''),
(60, 'Apple iPhone 12 (128 GB) - Azul', '223499', 'https://http2.mlstatic.com/D_NQ_NP_839818-MLA45719894955_042021-V.webp', 2, 17, '', ''),
(292, 'Auto A Control Remoto Gadnic Twist 15 Naranja Camioneta 4 Direcciones Monster Truck 4x4 Radiocontrol', '11623', 'https://http2.mlstatic.com/D_NQ_NP_932170-MLA49738807505_042022-V.webp', 6, 26, '', ''),
(327, 'Auto A Control Remoto Gadnic Twist 19 4x4 360 Radiocontrol', '12999', 'https://http2.mlstatic.com/D_NQ_NP_680931-MLA49738808455_042022-V.webp', 6, 30, '', ''),
(289, 'Auto A Radio Control Remoto 4x4 Camioneta Juguete Luz Led Rc', '6999', 'https://http2.mlstatic.com/D_NQ_NP_645542-MLA45170628476_032021-V.webp', 6, 13, '', ''),
(301, 'Auto Barbie Original Tv Con Accesorios Y Stickers Lelab', '22', 'https://http2.mlstatic.com/D_NQ_NP_930397-MLA31009126521_062019-V.webp', 6, 29, '', ''),
(291, 'Banco De Herramientas 27 Piezas Mi Primer Taller Lionel`s', '416', 'https://http2.mlstatic.com/D_NQ_NP_622595-MLA31621082686_072019-V.webp', 6, 19, '', ''),
(147, 'Banco Multifuncion 1000, Combo De Barra, Discos Y Bco Scott', '35371', 'https://http2.mlstatic.com/D_NQ_NP_995639-MLA41305048828_042020-V.webp', 4, 29, '', ''),
(135, 'Banco Pecho Plegable Tabla Gimnasia Regulable Plano Camilla', '15499', 'https://http2.mlstatic.com/D_NQ_NP_955966-MLA29738613227_032019-V.webp', 4, 22, '', ''),
(162, 'Banco Pecho Regulable Multigimnasio Df + Envio Gratis Caba!', '349', 'https://http2.mlstatic.com/D_NQ_NP_799216-MLA48318294058_112021-V.webp', 4, 36, '', ''),
(239, 'Barbacoa La Campagnola en doy pack 250?g', '134', 'https://http2.mlstatic.com/D_NQ_NP_681205-MLA44923488233_022021-W.webp', 7, 31, '', ''),
(132, 'Barra + 2 Mancuernas + 30 Kg En Discos Con Manija Kit Cuotas', '95', 'https://http2.mlstatic.com/D_NQ_NP_662216-MLA47709384098_092021-V.webp', 4, 36, '', ''),
(136, 'Barra + 2 Mancuernas + 30 Kg En Discos Con Manija Kit Set', '12961', 'https://http2.mlstatic.com/D_NQ_NP_673113-MLA45879745991_052021-V.webp', 4, 27, '', ''),
(125, 'Barra De Dominadas Exahome Para Pared Con Triple Agarre', '4619', 'https://http2.mlstatic.com/D_NQ_NP_618463-MLA43558393443_092020-V.webp', 4, 35, '', ''),
(134, 'Barra Dominadas Extensible Puerta Fitness Ejercicio Gym', '1599', 'https://http2.mlstatic.com/D_NQ_NP_876256-MLA46037553855_052021-V.webp', 4, 12, '', ''),
(133, 'Barras Paralelas Push Up Importadas Flexiones Calistenia', '899', 'https://http2.mlstatic.com/D_NQ_NP_761654-MLA49117699510_022022-V.webp', 4, 16, '', ''),
(323, 'Bebote Bebe Real Con Ropa Y Chupete 50cm', '459', 'https://http2.mlstatic.com/D_NQ_NP_725474-MLA49591803312_042022-V.webp', 6, 24, '', ''),
(139, 'Bicicleta Fija Gadnic F-3050 Resistencia Ajustable Piernas Brazos Rehabilitaci?n Cardio', '5499', 'https://http2.mlstatic.com/D_NQ_NP_929111-MLA48531825926_122021-V.webp', 4, 30, '', ''),
(148, 'Bicicleta fija Helitec Indoor Profesional 18KG para spinning negra', '40985', 'https://http2.mlstatic.com/D_NQ_NP_928429-MLA46738304662_072021-V.webp', 4, 29, '', ''),
(165, 'Bicicleta Fija Pedalera Rehabilita Piernas Y Brazos Portatil', '3329', 'https://http2.mlstatic.com/D_NQ_NP_814814-MLA47949000706_102021-V.webp', 4, 24, '', ''),
(157, 'Bicicleta fija Randers ARG - 101 tradicional azul y blanca', '16999', 'https://http2.mlstatic.com/D_NQ_NP_742704-MLA48914290619_012022-V.webp', 4, 39, '', ''),
(138, 'Bicicleta fija Randers ARG-840SP para spinning roja y negra', '39999', 'https://http2.mlstatic.com/D_NQ_NP_688257-MLA46950962197_082021-V.webp', 4, 12, '', ''),
(146, 'Bicicleta fija Unnic U-Fitness Spin Bike 105 para spinning negra', '39999', 'https://http2.mlstatic.com/D_NQ_NP_959821-MLA46834566046_072021-V.webp', 4, 29, '', ''),
(129, 'Bicicleta Spinning Fija Spin4100 Gadnic Profesional Est?tica Display Led Potencia Regulable Hasta 10', '44999', 'https://http2.mlstatic.com/D_NQ_NP_856338-MLA49740302636_042022-V.webp', 4, 18, '', ''),
(309, 'Bloques para armar Lego Marvel Avengers Thor mech armor 139 piezas en caja', '4627', 'https://http2.mlstatic.com/D_NQ_NP_808134-MLA48252261870_112021-V.webp', 6, 35, '', ''),
(337, 'Bloques para armar Lego Spider-Man and Sandman showdown 45 piezas en caja', '5499', 'https://http2.mlstatic.com/D_NQ_NP_814699-MLA48242663593_112021-V.webp', 6, 21, '', ''),
(320, 'Bloques para armar Lego Spider-Man Miles Morales mech armor 125 piezas en caja', '5999', 'https://http2.mlstatic.com/D_NQ_NP_646910-MLA48242502553_112021-V.webp', 6, 26, '', ''),
(298, 'Bloques para armar magn?ticos Magnific KBM-56ARU en caja', '1375', 'https://http2.mlstatic.com/D_NQ_NP_787590-MLA48250552858_112021-V.webp', 6, 24, '', ''),
(302, 'Bloques para armar Rasti Blokoco 120 piezas en balde', '2608', 'https://http2.mlstatic.com/D_NQ_NP_689649-MLA48732811884_012022-V.webp', 6, 33, '', ''),
(295, 'Bloques para armar Rasti Kit accesorios Kit de piezas b?sicas 400 en caja', '569', 'https://http2.mlstatic.com/D_NQ_NP_920965-MLA48250552789_112021-V.webp', 6, 33, '', ''),
(297, 'Bloques para armar Rasti Kit accesorios Kit de piezas b?sicas 600 en bolsa', '5249', 'https://http2.mlstatic.com/D_NQ_NP_733727-MLA48680431199_122021-V.webp', 6, 20, '', ''),
(300, 'Bloques para armar Rasti Mix 300 en caja', '459', 'https://http2.mlstatic.com/D_NQ_NP_825040-MLA48242784927_112021-V.webp', 6, 17, '', ''),
(260, 'Borcego Bota Cuero Briganti Hombre Zapato Goma - Hcbo00917', '21977', 'https://http2.mlstatic.com/D_NQ_NP_809557-MLA42435164408_072020-W.webp', 5, 18, '', ''),
(246, 'Bota Hombre Cuero Borcego Botita Briganti Zapato - Hcbo00902', '1898', 'https://http2.mlstatic.com/D_NQ_NP_905821-MLA31114113460_062019-W.webp', 5, 35, '', ''),
(261, 'Boxer Hombre Briganti Algod?n Combinado Pack X 6 Hxim09524', '9989', 'https://http2.mlstatic.com/D_NQ_NP_774437-MLA48062997448_102021-W.webp', 5, 12, '', ''),
(280, 'Buzo Billabong Con Capucha Fundamental Hombre', '12499', 'https://http2.mlstatic.com/D_NQ_NP_882147-MLA49312642773_032022-W.webp', 5, 28, '', ''),
(274, 'Buzo Polar Bordado Medio Cierre Hombre Moda Mistral 80021', '125', 'https://http2.mlstatic.com/D_NQ_NP_782831-MLA49660570254_042022-W.webp', 5, 11, '', ''),
(204, 'C?psulas Caf? Cabrales Colombia Nacionales X 10 U X 5,5 Gr', '979', 'https://http2.mlstatic.com/D_NQ_NP_985431-MLA47481211132_092021-W.webp', 7, 15, '', ''),
(224, 'Caf? au lait en c?psula Nescaf? Dolce Gusto 16?u', '949', 'https://http2.mlstatic.com/D_NQ_NP_852928-MLA43773298639_102020-W.webp', 7, 35, '', ''),
(240, 'Caf? avvolgente en c?psula Lavazza 10?u', '1161', 'https://http2.mlstatic.com/D_NQ_NP_973315-MLA43773297360_102020-W.webp', 7, 35, '', ''),
(218, 'Caf? cappuccino en c?psula Nescaf? Dolce Gusto 16?u', '949', 'https://http2.mlstatic.com/D_NQ_NP_666656-MLA43773282709_102020-W.webp', 7, 30, '', ''),
(232, 'Caf? colombia en c?psula La Morenita 10?u', '718', 'https://http2.mlstatic.com/D_NQ_NP_728273-MLA45259644695_032021-W.webp', 7, 26, '', ''),
(228, 'Caf? equilibrado n?7 en c?psula La Virginia 10?u', '689', 'https://http2.mlstatic.com/D_NQ_NP_782611-MLA45259617774_032021-W.webp', 7, 10, '', ''),
(203, 'Caf? Instant?neo Cabrales La Planta Suave Frasco 100 Gr', '689', 'https://http2.mlstatic.com/D_NQ_NP_760600-MLA47417281924_092021-W.webp', 7, 14, '', ''),
(227, 'Caf? instant?neo suave Nescaf? Dolca doypack 170?g', '649', 'https://http2.mlstatic.com/D_NQ_NP_877349-MLA43567415505_092020-W.webp', 7, 29, '', ''),
(230, 'Caf? intenso n?9 en c?psula La Virginia 10?u', '660', 'https://http2.mlstatic.com/D_NQ_NP_704728-MLA45259664433_032021-W.webp', 7, 23, '', ''),
(223, 'Caf? late macchiato en c?psula Nescaf? Dolce Gusto 16?u', '949', 'https://http2.mlstatic.com/D_NQ_NP_612744-MLA43773300940_102020-W.webp', 7, 20, '', ''),
(237, 'Caf? lungo en c?psula Nescaf? Dolce Gusto sin TACC 16?u', '949', 'https://http2.mlstatic.com/D_NQ_NP_641393-MLA43773298339_102020-W.webp', 7, 16, '', ''),
(330, 'Caja Registradora Luz Y Sonidos 16 Cm Con Carro De Super', '499', 'https://http2.mlstatic.com/D_NQ_NP_755240-MLA43601717601_092020-V.webp', 6, 22, '', ''),
(269, 'Campera Abrigo Hombre Brooksfield Moda Sport Importada 7017b', '22495', 'https://http2.mlstatic.com/D_NQ_NP_672876-MLA50061022938_052022-W.webp', 5, 27, '', ''),
(285, 'Campera Cuello Alto Abrigo Moda Hombre Mistral 70046', '23', 'https://http2.mlstatic.com/D_NQ_NP_728698-MLA49576209308_042022-W.webp', 5, 26, '', ''),
(264, 'Campera Hombre Farenheite Importada Roux', '23655', 'https://http2.mlstatic.com/D_NQ_NP_720103-MLA49946621775_052022-W.webp', 5, 25, '', ''),
(253, 'Campera Polar Cuello Alto Bordada Hombre Mistral 70058', '149', 'https://http2.mlstatic.com/D_NQ_NP_883401-MLA49941404328_052022-W.webp', 5, 28, '', ''),
(281, 'Campera Polar Cuello Alto Bordada Hombre Mistral 86019', '125', 'https://http2.mlstatic.com/D_NQ_NP_610857-MLA49658490032_042022-W.webp', 5, 22, '', ''),
(266, 'Campera Tofino O\'neill', '1881', 'https://http2.mlstatic.com/D_NQ_NP_904895-MLA50020871768_052022-W.webp', 5, 16, '', ''),
(250, 'Campera Topper Gd Men Negro Hombre', '9999', 'https://http2.mlstatic.com/D_NQ_NP_609282-MLA50158861918_062022-W.webp', 5, 11, '', ''),
(259, 'Canguro Buzo Capucha Algod?n Moda Hombre Mistral 83002f', '1195', 'https://http2.mlstatic.com/D_NQ_NP_999372-MLA45657493404_042021-W.webp', 5, 16, '', ''),
(177, 'Canilla Doble Lavarropas Esferica 2 Salidas 3/4', '1183', 'https://http2.mlstatic.com/D_NQ_NP_622879-MLA44029062644_112020-V.webp', 3, 13, '', ''),
(221, 'Caramelo Masticable Billiken X 200u - Oferta En Sweet Market', '379', 'https://http2.mlstatic.com/D_NQ_NP_763174-MLA49971231004_052022-W.webp', 7, 27, '', ''),
(231, 'Caramelos Masticables Flynn Paff Tutifruti 8 Gr X 70 U', '492', 'https://http2.mlstatic.com/D_NQ_NP_629701-MLA49760927314_042022-W.webp', 7, 40, '', ''),
(186, 'Cepillo Alambre Acero Bronceado Curvo P/ Limpieza - Subte A', '350', 'https://http2.mlstatic.com/D_NQ_NP_921037-MLA31351316304_072019-V.webp', 3, 25, '', ''),
(178, 'Cepillo Matricero Acero Bronceado Plastico P/ Limpieza X3pzs', '333', 'https://http2.mlstatic.com/D_NQ_NP_824461-MLA31351256225_072019-V.webp', 3, 36, '', ''),
(42, 'Chevrolet Cruze Ii 1,4 Lt 153cv', '3395000', 'https://http2.mlstatic.com/D_NQ_NP_946154-MLA50108764896_052022-W.webp', 1, 26, '', ''),
(26, 'Chevrolet Onix 1,4 Ltz 98cv', '2590000', 'https://http2.mlstatic.com/D_NQ_NP_823488-MLA50008180504_052022-W.webp', 1, 31, '', ''),
(44, 'Chevrolet S10 2,8 Ls Cd Tdci 200cv', '4600000', 'https://http2.mlstatic.com/D_NQ_NP_731311-MLA50020140886_052022-W.webp', 1, 17, '', ''),
(8, 'Chevrolet S10 2,8 Ls Cd Tdci 200cv 4x2', '4290000', 'https://http2.mlstatic.com/D_NQ_NP_912204-MLA49094650869_022022-W.webp', 1, 30, '', ''),
(23, 'Chevrolet Spin 1,8 Ltz 7as 105cv', '2750000', 'https://http2.mlstatic.com/D_NQ_NP_689750-MLA50173081551_062022-W.webp', 1, 10, '', ''),
(193, 'Chimichurri Recetas De Entonces X 170 Gr', '497', 'https://http2.mlstatic.com/D_NQ_NP_790501-MLA47246846494_082021-W.webp', 7, 13, '', ''),
(233, 'Chocolate chococcino en c?psula Nescaf? Dolce Gusto 16?u', '1099', 'https://http2.mlstatic.com/D_NQ_NP_953430-MLA43773285906_102020-W.webp', 7, 30, '', ''),
(185, 'Cinta Bifaz - Doble Faz Espumosa 18mm - Subte A Carabobo', '234', 'https://http2.mlstatic.com/D_NQ_NP_704920-MLA49367978426_032022-V.webp', 3, 26, '', ''),
(155, 'Cinta de correr el?ctrica Femmto Runner 14K 220V negra', '51312', 'https://http2.mlstatic.com/D_NQ_NP_646711-MLA50233676892_062022-V.webp', 4, 15, '', ''),
(1, 'Citro?n C3 1,5 Feel 90cv', '2290000', 'https://http2.mlstatic.com/D_NQ_NP_966219-MLA50111078092_052022-W.webp', 1, 39, '', ''),
(39, 'Citro?n C3 1,6 I Exclusive Am74', '1390000', 'https://http2.mlstatic.com/D_NQ_NP_778461-MLA50168765967_062022-W.webp', 1, 13, '', ''),
(130, 'Colchoneta + 2 Mancuernas + 2 Tobilleras 2k Fitness Gimnasia', '1899', 'https://http2.mlstatic.com/D_NQ_NP_705523-MLA45719717515_042021-V.webp', 4, 20, '', ''),
(159, 'Colchoneta Gimnasia Plegable Yoga Abdominales Exahome Oferta', '1139', 'https://http2.mlstatic.com/D_NQ_NP_662625-MLA44152319544_112020-V.webp', 4, 19, '', ''),
(154, 'Colchoneta Plegable Yoga Mat 10mm Pilates Caucho Nbr + Bolso', '2999', 'https://http2.mlstatic.com/D_NQ_NP_952169-MLA49804613220_042022-V.webp', 4, 34, '', ''),
(98, 'Computadora Completa Amd Intel Dual Core 8gb Led 19 Lol', '67999', 'https://http2.mlstatic.com/D_NQ_NP_751786-MLA31060099570_062019-V.webp', 8, 27, '', ''),
(76, 'Computadora Completa Amd Intel Dual Core 8gb Led 22 Lol', '69999', 'https://http2.mlstatic.com/D_NQ_NP_751786-MLA31060099570_062019-V.webp', 8, 17, '', ''),
(70, 'Computadora Completa Amd Intel Dual Core Monitor Led 19 Lol', '59999', 'https://http2.mlstatic.com/D_NQ_NP_638519-MLA31031803216_062019-V.webp', 8, 26, '', ''),
(74, 'Computadora Pc Cpu Gamerpro Rgb Intel Core I7 32gb 480g Ssd', '180998', 'https://http2.mlstatic.com/D_NQ_NP_880324-MLA49795873014_042022-V.webp', 8, 35, '', ''),
(118, 'Computadora Pc Cpu Solarmax Intel Core I3 10ma 8gb 240g Ssd', '6019', 'https://http2.mlstatic.com/D_NQ_NP_787165-MLA49850844464_052022-V.webp', 8, 12, '', ''),
(113, 'Computadora Pc Cpu Solarmax Intel Core I5 10ma 16gb 480g Ssd', '8089', 'https://http2.mlstatic.com/D_NQ_NP_996257-MLA49850389277_052022-V.webp', 8, 40, '', ''),
(293, 'Cry Babies Coney IMC Toys 10598IM', '13599', 'https://http2.mlstatic.com/D_NQ_NP_823936-MLA32147564166_092019-V.webp', 6, 11, '', ''),
(176, 'Cuenta Ganado Contador Manual Teatro - Subte Carabobo', '505', 'https://http2.mlstatic.com/D_NQ_NP_961163-MLA45840883850_052021-V.webp', 3, 21, '', ''),
(275, 'Customs Ba Camperas Hombre Parka Con Piel Campera Invierno B', '1499', 'https://http2.mlstatic.com/D_NQ_NP_926160-MLA50222101504_062022-W.webp', 5, 10, '', ''),
(244, 'Customs Ba Camperas Hombre Parka Corderito Campera Full', '1899', 'https://http2.mlstatic.com/D_NQ_NP_963813-MLA49850341919_052022-W.webp', 5, 26, '', ''),
(248, 'Customs Ba Camperas Hombre Parka Corderito Campera Negra Us', '1899', 'https://http2.mlstatic.com/D_NQ_NP_810392-MLA50181952141_062022-W.webp', 5, 33, '', ''),
(341, 'Dinosaurio T. Rex Super Gigante Juguete De Goma 70x60cm', '2755', 'https://http2.mlstatic.com/D_NQ_NP_751809-MLA45698936302_042021-V.webp', 6, 20, '', ''),
(311, 'Dinosaurios De Goma 5 Cm Juguete Souvenir Pi?ata X 20', '135', 'https://http2.mlstatic.com/D_NQ_NP_889125-MLA44213883583_122020-V.webp', 6, 14, '', ''),
(166, 'Disco Pesa Fundicion Fitness 5 Kg Mancuerna Barra', '229', 'https://http2.mlstatic.com/D_NQ_NP_944022-MLA44814565736_022021-V.webp', 4, 11, '', ''),
(187, 'Discos De Corte Minitorno Dremel Sierra Circulares X6 Pcs', '1512', 'https://http2.mlstatic.com/D_NQ_NP_657225-MLA46519509858_062021-V.webp', 3, 23, '', ''),
(150, 'Ejercitador Abdominal Abt - Importado!! - World Fitness', '1023', 'https://http2.mlstatic.com/D_NQ_NP_748143-MLA47886803301_102021-V.webp', 4, 26, '', ''),
(182, 'Espatulas X 3 Con Mango Madera Subte A Carabobo', '292', 'https://http2.mlstatic.com/D_NQ_NP_925811-MLA20633290178_032016-V.webp', 3, 17, '', ''),
(94, 'Estabilizador de tensi?n Atomlux Linea H H1000@ 1000VA entrada y salida de 220V blanco', '3035', 'https://http2.mlstatic.com/D_NQ_NP_653460-MLA44456543268_122020-V.webp', 8, 11, '', ''),
(111, 'Estabilizador de tensi?n Forza FVR 2200VA Series FVR-2202A 2200VA entrada y salida de 220V CA negro', '5999', 'https://http2.mlstatic.com/D_NQ_NP_858185-MLA45743826886_042021-V.webp', 8, 38, '', ''),
(104, 'Estabilizador de tensi?n Lyonn TCA 1200NV 1200VA entrada y salida de 220V CA negro', '3519', 'https://http2.mlstatic.com/D_NQ_NP_919319-MLA44111139770_112020-V.webp', 8, 35, '', ''),
(84, 'Estabilizador de tensi?n Lyonn TCA TCA-2000NV 2000VA entrada y salida de 220V CA negro', '5412', 'https://http2.mlstatic.com/D_NQ_NP_954479-MLA44446945155_122020-V.webp', 8, 17, '', ''),
(77, 'Estabilizador de tensi?n SurElectric Future ER 1000 1000VA entrada y salida de 220V blanco', '2399', 'https://http2.mlstatic.com/D_NQ_NP_969593-MLA44112672036_112020-V.webp', 8, 27, '', ''),
(82, 'Estabilizador de tensi?n TRV Electronics Concept USB 2000VA entrada y salida de 220V negro', '5488', 'https://http2.mlstatic.com/D_NQ_NP_852897-MLA44443256213_122020-V.webp', 8, 16, '', ''),
(75, 'Estabilizador de tensi?n UPS Forza NT Series NT-1002A 1000VA entrada y salida de 220V CA negro', '1599', 'https://http2.mlstatic.com/D_NQ_NP_648717-MLA45743512493_042021-V.webp', 8, 10, '', ''),
(24, 'Fiat 500x 1,4 T 16v Pop', '4450000', 'https://http2.mlstatic.com/D_NQ_NP_734290-MLA50170808894_062022-W.webp', 1, 12, '', ''),
(29, 'Fiat Argo 1,3 Drive Gse Manual', '1900000', 'https://http2.mlstatic.com/D_NQ_NP_945738-MLA49997944818_052022-W.webp', 1, 11, '', ''),
(2, 'Fiat Ducato 2,3 Maxicargo Mjet Te C/abs+aa', '6250500', 'https://http2.mlstatic.com/D_NQ_NP_957009-MLA49880859642_052022-W.webp', 1, 26, '', ''),
(6, 'Fiat Toro 2,0 Volcano 4x4 At', '4415500', 'https://http2.mlstatic.com/D_NQ_NP_830294-MLA49535246400_032022-W.webp', 1, 37, '', ''),
(9, 'Fiat Toro 2,0 Volcano 4x4 At Pack Premium', '4290000', 'https://http2.mlstatic.com/D_NQ_NP_696969-MLA49094487641_022022-W.webp', 1, 14, '', ''),
(206, 'Flan Light Royal Vainilla X 16 Gr', '90', 'https://http2.mlstatic.com/D_NQ_NP_697623-MLA47529144971_092021-W.webp', 7, 24, '', ''),
(213, 'Flan Royal Vainilla X 60 Gr', '48', 'https://http2.mlstatic.com/D_NQ_NP_978541-MLA47684566993_092021-W.webp', 7, 33, '', ''),
(184, 'Flexible De Ducha Manguera Acero 2 Mt Subte A', '611', 'https://http2.mlstatic.com/D_NQ_NP_987432-MLA31575947875_072019-V.webp', 3, 37, '', ''),
(32, 'Ford Fiesta 1,6 Max Ambiente Plus Mp3', '1050000', 'https://http2.mlstatic.com/D_NQ_NP_926383-MLA50172384733_062022-W.webp', 1, 12, '', ''),
(27, 'Ford Fiesta Kinetic Design 1,6 Sedan Titanium Powershift 120cv', '2350000', 'https://http2.mlstatic.com/D_NQ_NP_875296-MLA50008219846_052022-W.webp', 1, 35, '', ''),
(40, 'Ford Focus Iii 1,6 S', '2395000', 'https://http2.mlstatic.com/D_NQ_NP_695521-MLA50157239092_062022-W.webp', 1, 37, '', ''),
(35, 'Ford Ka 1,5 Freestyle Se', '2890000', 'https://http2.mlstatic.com/D_NQ_NP_954752-MLA50171779462_062022-W.webp', 1, 36, '', ''),
(30, 'Ford Ka 1,5 S 5 P', '2350000', 'https://http2.mlstatic.com/D_NQ_NP_868158-MLA50174535386_062022-W.webp', 1, 33, '', ''),
(17, 'Ford Ranger 3,2 Cd 4x2 Xlt Tdci 200cv', '4580000', 'https://http2.mlstatic.com/D_NQ_NP_606030-MLA50057476080_052022-W.webp', 1, 21, '', ''),
(171, 'Fresa Lima Rotativa Metal Duro Para Taladro Doble Corte X 10', '4279', 'https://http2.mlstatic.com/D_NQ_NP_824169-MLA47577340406_092021-V.webp', 3, 19, '', ''),
(181, 'Fresas Mini Torno Multiproposito Madera Minitorno X6 Pcs', '853', 'https://http2.mlstatic.com/D_NQ_NP_949042-MLA47246869756_082021-V.webp', 3, 39, '', ''),
(339, 'Frutas Y Verduras Con Abrojo Para Cortar Cuchillo Y Tabla', '1643', 'https://http2.mlstatic.com/D_NQ_NP_882660-MLA47124085686_082021-V.webp', 6, 25, '', ''),
(189, 'Funda P/ Tabla De Planchar Metalizada No Se Quema !!', '973', 'https://http2.mlstatic.com/D_NQ_NP_845512-MLA31351168144_072019-V.webp', 3, 10, '', ''),
(195, 'Galleta De Arroz Deluxe Bla Bla Sin Sal X 98 Gramos', '98', 'https://http2.mlstatic.com/D_NQ_NP_679184-MLA47260784805_082021-W.webp', 7, 26, '', ''),
(194, 'Galletas Deluxe Bla Bla Arroz Ba?adas Frutilla 115gr', '162', 'https://http2.mlstatic.com/D_NQ_NP_600180-MLA47260785286_082021-W.webp', 7, 35, '', ''),
(225, 'Galletita Chocolinas 250?g', '175', 'https://http2.mlstatic.com/D_NQ_NP_673690-MLA43815375417_102020-W.webp', 7, 26, '', ''),
(215, 'Galletitas Oreo Cl?sicas Tripack X 354 Gr', '426', 'https://http2.mlstatic.com/D_NQ_NP_957898-MLA47588758747_092021-W.webp', 7, 17, '', ''),
(216, 'Galletitas Pepitos Regular X 119 Gr', '156', 'https://http2.mlstatic.com/D_NQ_NP_890823-MLA47588759343_092021-W.webp', 7, 15, '', ''),
(14, 'Geely Emgrand X7 Sport 2,4 Gl At', '6340000', 'https://http2.mlstatic.com/D_NQ_NP_773621-MLA48420027763_122021-W.webp', 1, 12, '', ''),
(207, 'Gelatina Royal Light Cereza Sobre X 25 Gr', '133', 'https://http2.mlstatic.com/D_NQ_NP_988103-MLA47527768978_092021-W.webp', 7, 21, '', ''),
(210, 'Gelatina Royal Light Frutilla 25 Gramos', '133', 'https://http2.mlstatic.com/D_NQ_NP_655805-MLA47520091052_092021-W.webp', 7, 15, '', ''),
(211, 'Gelatina Royal Light Frutos Rojos X 25 Gr', '133', 'https://http2.mlstatic.com/D_NQ_NP_903949-MLA47498406652_092021-W.webp', 7, 36, '', ''),
(205, 'Gelatina Royal Regular Cereza X 40 Gr', '115', 'https://http2.mlstatic.com/D_NQ_NP_688396-MLA47498407052_092021-W.webp', 7, 17, '', ''),
(212, 'Gelatina Royal Regular Frambuesa X 40 Gr', '115', 'https://http2.mlstatic.com/D_NQ_NP_875863-MLA48335736180_112021-W.webp', 7, 36, '', ''),
(13, 'Great Wall Wingle 6 2,0 Tdi Dc 4wd Dignity', '3890000', 'https://http2.mlstatic.com/D_NQ_NP_661064-MLA49647933961_042022-W.webp', 1, 37, '', ''),
(167, 'Gymform Multigym Equipo De Entrenamiento F?sico- Tevecompras', '127875', 'https://http2.mlstatic.com/D_NQ_NP_674576-MLA49851764370_052022-V.webp', 4, 17, '', ''),
(312, 'Helic?ptero Dron Inteligente Sensorial Vuela Sube Y Baja', '285', 'https://http2.mlstatic.com/D_NQ_NP_669932-MLA48454376572_122021-V.webp', 6, 36, '', ''),
(267, 'Jean Tascani Skinny Trecker Plus Negro', '12325', 'https://http2.mlstatic.com/D_NQ_NP_884061-MLA43460523963_092020-W.webp', 5, 16, '', ''),
(20, 'Jeep Compass 2,4 Sport Manual', '8020000', 'https://http2.mlstatic.com/D_NQ_NP_764172-MLA49739017965_042022-W.webp', 1, 21, '', ''),
(188, 'Juego De Abrazaderas 12 Pcs 8 Mm - 27 Mm', '418', 'https://http2.mlstatic.com/D_NQ_NP_900366-MLA31352281474_072019-V.webp', 3, 38, '', ''),
(331, 'Juego de mesa Bad dog Ditoys 1991', '5778', 'https://http2.mlstatic.com/D_NQ_NP_829339-MLA44735959555_012021-V.webp', 6, 15, '', ''),
(340, 'Juego de mesa Pictionary Jr. Ruibal 7901', '4', 'https://http2.mlstatic.com/D_NQ_NP_628483-MLA47660926712_092021-V.webp', 6, 14, '', ''),
(318, 'Juego de mesa T.E.G. Tradicional Yetem', '5495', 'https://http2.mlstatic.com/D_NQ_NP_659090-MLA44676313817_012021-V.webp', 6, 31, '', ''),
(180, 'Juego Llaves Allen Hexagonales X10 Piezas - Subte Carabobo', '292', 'https://http2.mlstatic.com/D_NQ_NP_769205-MLA31351132147_072019-V.webp', 3, 36, '', ''),
(173, 'Juego Mechas Para Madera Espiral X 5 Pcs 4, 5, 6, 8, 10 Mm', '396', 'https://http2.mlstatic.com/D_NQ_NP_897717-MLA47499630396_092021-V.webp', 3, 23, '', ''),
(220, 'K?tchup picante La Campagnola en doy pack 250?g', '130', 'https://http2.mlstatic.com/D_NQ_NP_906524-MLA44923420146_022021-W.webp', 7, 22, '', ''),
(36, 'Kia Sportage 2,0 Ex At 154cv 4x2', '6680000', 'https://http2.mlstatic.com/D_NQ_NP_673773-MLA50171648848_062022-W.webp', 1, 30, '', ''),
(71, 'Kit de teclado y mouse gamer Noga NKB-300 Espa?ol Espa?a de color negro y rojo', '936', 'https://http2.mlstatic.com/D_NQ_NP_720123-MLA46445221737_062021-V.webp', 8, 29, '', ''),
(127, 'Kit Fitness X5 Bandas El?sticas De Resistencia Set Ejercicio', '1259', 'https://http2.mlstatic.com/D_NQ_NP_925468-MLA48885803832_012022-V.webp', 4, 23, '', ''),
(172, 'Kit Herramientas 3 Pinzas Punta Alicate Universal Subte A', '1241', 'https://http2.mlstatic.com/D_NQ_NP_642021-MLA31352575191_072019-V.webp', 3, 30, '', ''),
(151, 'Kit Mancuernas Y Barra Con Tope A Rosca + 20 Kg Discos Pvc', '7892', 'https://http2.mlstatic.com/D_NQ_NP_717857-MLA46897614292_072021-V.webp', 4, 35, '', ''),
(168, 'Kit Tobilleras 2 Kg + Mancuernas 2 Kg + Soga+ Banda Circular', '4369', 'https://http2.mlstatic.com/D_NQ_NP_686812-MLA50040739324_052022-V.webp', 4, 25, '', ''),
(156, 'Kit Yoga Pilates Mat Colchoneta 6mm + 2 Ladrillos + Cintur?n', '3499', 'https://http2.mlstatic.com/D_NQ_NP_786706-MLA49612699306_042022-V.webp', 4, 30, '', ''),
(5, 'Lifan Foison Cargo 1,3 Cargo 92cv', '2749000', 'https://http2.mlstatic.com/D_NQ_NP_831966-MLA49094417022_022022-W.webp', 1, 25, '', ''),
(149, 'Mancuerna 2kg. Pesa Recubierta En Goma Importadas', '999', 'https://http2.mlstatic.com/D_NQ_NP_660935-MLA49021780266_022022-V.webp', 4, 10, '', ''),
(145, 'Mancuerna Hexagonal Engomada 6 Kg Goma Pesa Maciza Gym', '5409', 'https://http2.mlstatic.com/D_NQ_NP_727451-MLA42722763090_072020-V.webp', 4, 18, '', ''),
(201, 'Mantequilla Man? Untable Dame Mani Crocante 510g', '436', 'https://http2.mlstatic.com/D_NQ_NP_966807-MLA47683608324_092021-W.webp', 7, 16, '', ''),
(202, 'Mantequilla Man? Untable Dame Mani Natural 510g', '436', 'https://http2.mlstatic.com/D_NQ_NP_845399-MLA47338311856_092021-W.webp', 7, 38, '', ''),
(191, 'Martillo De Goma - Plastico Chapista Profesional', '582', 'https://http2.mlstatic.com/D_NQ_NP_948530-MLA49863358899_052022-V.webp', 3, 33, '', ''),
(170, 'Martillo Maza De Goma 500 Gr Mango Madera', '509', 'https://http2.mlstatic.com/D_NQ_NP_817177-MLA46043526008_052021-V.webp', 3, 37, '', ''),
(229, 'Mayonesa Hellmann\'s Cl?sica en doy pack 475?g', '224', 'https://http2.mlstatic.com/D_NQ_NP_809505-MLA44923510034_022021-W.webp', 7, 26, '', ''),
(219, 'Mayonesa Hellmann\'s Cl?sica en doy pack 950?g', '474', 'https://http2.mlstatic.com/D_NQ_NP_914391-MLA44923458010_022021-W.webp', 7, 37, '', ''),
(120, 'Memoria RAM color verde 4GB 1 Markvision MVD34096MLD-16', '3046', 'https://http2.mlstatic.com/D_NQ_NP_647177-MLA49041660035_022022-V.webp', 8, 23, '', ''),
(73, 'Memoria RAM color verde 8GB 1 Crucial CT8G4DFS8266', '4999', 'https://http2.mlstatic.com/D_NQ_NP_998250-MLA41465180090_042020-V.webp', 8, 12, '', ''),
(69, 'Memoria RAM Fury Beast DDR4 gamer color negro 8GB 1 Kingston KF426C16BB/8', '6099', 'https://http2.mlstatic.com/D_NQ_NP_908770-MLA48636149604_122021-V.webp', 8, 21, '', ''),
(114, 'Memoria RAM Fury Beast DDR4 gamer color negro 8GB 1 Kingston KF432C16BB/8', '7149', 'https://http2.mlstatic.com/D_NQ_NP_612752-MLA48636149729_122021-V.webp', 8, 34, '', ''),
(100, 'Memoria RAM Fury Beast DDR4 RGB color negro 8GB 1 Kingston KF426C16BBA/8', '6999', 'https://http2.mlstatic.com/D_NQ_NP_675714-MLA48636416221_122021-V.webp', 8, 18, '', ''),
(106, 'Memoria RAM Fury Beast DDR4 RGB gamer color negro 16GB 1 Kingston KF432C16BB1A/16', '13499', 'https://http2.mlstatic.com/D_NQ_NP_613170-MLA48878114544_012022-V.webp', 8, 18, '', ''),
(88, 'Memoria RAM Value Select gamer color negro 16GB 1 Corsair CMV16GX4M1A2400C16', '10499', 'https://http2.mlstatic.com/D_NQ_NP_852659-MLA41522341298_042020-V.webp', 8, 33, '', ''),
(68, 'Memoria RAM Vengeance LPX gamer color black 16GB 1 Corsair CMK16GX4M1A2400C14', '11399', 'https://http2.mlstatic.com/D_NQ_NP_779401-MLA41497604625_042020-V.webp', 8, 27, '', ''),
(105, 'Memoria RAM Vengeance LPX gamer color negro 8GB 1 Corsair CMK8GX4M1A2400C16', '6151', 'https://http2.mlstatic.com/D_NQ_NP_676976-MLA48049766096_102021-V.webp', 8, 24, '', ''),
(67, 'Memoria RAM Vengeance LPX gamer color negro 8GB 1 Corsair CMK8GX4M1D3000C16', '7199', 'https://http2.mlstatic.com/D_NQ_NP_912545-MLA48832758167_012022-V.webp', 8, 38, '', ''),
(336, 'Mi Primera Licuadora Juliana Sisjul005', '405', 'https://http2.mlstatic.com/D_NQ_NP_652167-MLA44628749279_012021-V.webp', 6, 22, '', ''),
(83, 'Micr?fono HyperX QuadCast condensador multipatr?n negro', '18199', 'https://http2.mlstatic.com/D_NQ_NP_841463-MLA32553444734_102019-V.webp', 8, 29, '', ''),
(308, 'Monster Truck a control remoto Spin Master Monster Jam Monster Jam 1:24 Megalodon', '8275', 'https://http2.mlstatic.com/D_NQ_NP_655809-MLA48578839702_122021-V.webp', 6, 21, '', ''),
(343, 'Monster Truck a control remoto Spin Master Monster Jam Monster Jam 1:24 Monster Mutt Dalmatian', '7499', 'https://http2.mlstatic.com/D_NQ_NP_840977-MLA48578576637_122021-V.webp', 6, 29, '', ''),
(51, 'Moto G200 5G 128 GB azul oscuro 8 GB RAM', '89999', 'https://http2.mlstatic.com/D_NQ_NP_614394-MLA49460542461_032022-V.webp', 2, 16, '', ''),
(53, 'Moto G41 128 GB negro ?nix 4 GB RAM', '43999', 'https://http2.mlstatic.com/D_NQ_NP_711274-MLA49535290493_032022-V.webp', 2, 34, '', ''),
(45, 'Moto G60s 128 GB azul 6 GB RAM', '55999', 'https://http2.mlstatic.com/D_NQ_NP_788340-MLA48112659213_112021-V.webp', 2, 30, '', ''),
(54, 'Moto G71 128 GB arctic blue 6 GB RAM', '59999', 'https://http2.mlstatic.com/D_NQ_NP_916341-MLA49272367960_032022-V.webp', 2, 35, '', ''),
(46, 'Motorola Edge 20 Lite 128 GB gris 6 GB RAM', '72999', 'https://http2.mlstatic.com/D_NQ_NP_833619-MLA48233725917_112021-V.webp', 2, 22, '', ''),
(316, 'Motosierra Con Luz Y Sonido El Duende Azul', '219', 'https://http2.mlstatic.com/D_NQ_NP_880896-MLA49003512698_022022-V.webp', 6, 14, '', ''),
(85, 'Mouse Genius DX-110 USB blanco elegante', '490', 'https://http2.mlstatic.com/D_NQ_NP_654369-MLA44840092222_022021-V.webp', 8, 29, '', ''),
(108, 'Mouse Genius DX-120 spring green', '319', 'https://http2.mlstatic.com/D_NQ_NP_865723-MLA32845768629_112019-V.webp', 8, 30, '', ''),
(115, 'Mouse Pad GTC PAD-212 de neoprene 190mm x 230mm x 22mm negro', '648', 'https://http2.mlstatic.com/D_NQ_NP_754705-MLA40736232212_022020-V.webp', 8, 27, '', ''),
(209, 'Mousse Royal Light Chocolate X 40 Gr', '147', 'https://http2.mlstatic.com/D_NQ_NP_977595-MLA47498406510_092021-W.webp', 7, 22, '', ''),
(324, 'Mu?eco Juguete Sonic The Hedgehog Calidad Unica Y Mucho +!', '13', 'https://http2.mlstatic.com/D_NQ_NP_766712-MLA48183709820_112021-V.webp', 6, 22, '', ''),
(290, 'Mu?eco Soft Batman Liga De La Justicia Original Dc Comics', '395', 'https://http2.mlstatic.com/D_NQ_NP_808790-MLA41053646728_032020-V.webp', 6, 25, '', ''),
(314, 'Mu?ecos Five Nights At Freddys Fnaf Animatr?nicos Desarmable', '149', 'https://http2.mlstatic.com/D_NQ_NP_835705-MLA45601125650_042021-V.webp', 6, 38, '', ''),
(342, 'Mu?ecos Five Nights At Freddys Pizza Simulator Desarmables', '169', 'https://http2.mlstatic.com/D_NQ_NP_756611-MLA46530474269_062021-V.webp', 6, 28, '', ''),
(306, 'Mu?ecos Minecraft Juguetes. Set X 10.', '5214', 'https://http2.mlstatic.com/D_NQ_NP_742485-MLA40120647725_122019-V.webp', 6, 18, '', ''),
(310, 'Mu?ecos Naruto Shippuden Hokage Six Paths Mercado Juguete', '169', 'https://http2.mlstatic.com/D_NQ_NP_729507-MLA47740151982_102021-V.webp', 6, 14, '', ''),
(161, 'Multigimnasio 7116 Protector De Lingotera C/75 Kilos', '103699', 'https://http2.mlstatic.com/D_NQ_NP_923669-MLA40317046008_012020-V.webp', 4, 13, '', ''),
(123, 'Multigimnasio 75kg C/estruct. P/lingotera World Fitness 7116', '84999', 'https://http2.mlstatic.com/D_NQ_NP_624524-MLA31047194237_062019-V.webp', 4, 40, '', ''),
(153, 'Multigimnasio C/50kg Reforzado 30 Funcion World Fitness 7000', '51072', 'https://http2.mlstatic.com/D_NQ_NP_845478-MLA31113059763_062019-V.webp', 4, 19, '', ''),
(124, 'Multigimnasio Con Bco Scott+discos+barra 150cm World Fitness', '28993', 'https://http2.mlstatic.com/D_NQ_NP_675092-MLA31034902555_062019-V.webp', 4, 36, '', ''),
(144, 'Multigimnasio Randers Arg-63140 Incluye Pesas 50kg Gym', '66499', 'https://http2.mlstatic.com/D_NQ_NP_648441-MLA31008279823_062019-V.webp', 4, 16, '', ''),
(143, 'Multigimnasio Randers Multigym Arg-63167 Incluye Pesas 100kg', '107299', 'https://http2.mlstatic.com/D_NQ_NP_735795-MLA48197098132_112021-V.webp', 4, 24, '', ''),
(79, 'Notebook HP 15-dw3033dx natural silver 15.6\"\", Intel Core i3 1115G4 8GB de RAM 256GB SSD, Intel UHD ', '99999', 'https://http2.mlstatic.com/D_NQ_NP_999103-MLA48632024149_122021-V.webp', 8, 21, '', ''),
(179, 'O Ring X 50 Pcs Repuesto Para Canilla Vastagos', '404', 'https://http2.mlstatic.com/D_NQ_NP_693580-MLA46568974562_062021-V.webp', 3, 35, '', ''),
(163, 'Par De Mancuernas 1kg C/u. Pesas Recubiertas Entrenamiento', '1899', 'https://http2.mlstatic.com/D_NQ_NP_920937-MLA49002628043_022022-V.webp', 4, 39, '', ''),
(142, 'Par De Mancuernas Recubiertas Goma 2 Kg C/u Pesas Importadas', '1999', 'https://http2.mlstatic.com/D_NQ_NP_663500-MLA49002510579_022022-V.webp', 4, 35, '', ''),
(241, 'Parka Campera Hombre Gabardina Capucha Piel Abrigrado Invier', '11899', 'https://http2.mlstatic.com/D_NQ_NP_732905-MLA45190788799_032021-W.webp', 5, 31, '', ''),
(87, 'Pc Armada Completa Cpu Computadora I5 8gb 1tb O Ssd Tareas Oficina', '621', 'https://http2.mlstatic.com/D_NQ_NP_933200-MLA45195347150_032021-V.webp', 8, 14, '', ''),
(103, 'Pc Armada Cpu Computadora I5 8gb 1tb O Ssd - Cuotas', '873', 'https://http2.mlstatic.com/D_NQ_NP_945983-MLA45195575304_032021-V.webp', 8, 19, '', ''),
(109, 'Pc Cpu Computadora Amd 3.5ghz Dual-core Nueva C/garantia', '32999', 'https://http2.mlstatic.com/D_NQ_NP_672029-MLA31725850773_082019-V.webp', 8, 15, '', ''),
(91, 'Pc Gamer Escritorio Completa Amd Computadora Cpu Nueva Ssd', '106859', 'https://http2.mlstatic.com/D_NQ_NP_664554-MLA48690667850_122021-V.webp', 8, 17, '', ''),
(137, 'Pelota Esferodinamia 75 Cm Fitness Pilates Yoga + Inflador', '1799', 'https://http2.mlstatic.com/D_NQ_NP_685932-MLA49116383724_022022-V.webp', 4, 16, '', ''),
(164, 'Pelota Esferodinamia Fitness Pilates Yoga 55 Cm Oferta', '1099', 'https://http2.mlstatic.com/D_NQ_NP_716280-MLA49116455567_022022-V.webp', 4, 14, '', ''),
(140, 'Pelota Esferodinamia Fitness Pilates Yoga 65 Cm Oferta!!', '1427', 'https://http2.mlstatic.com/D_NQ_NP_607772-MLA49116356903_022022-V.webp', 4, 27, '', ''),
(126, 'Pesa 10kg Con Discos Ajustables Estuche Mercado Libre Basics', '3899', 'https://http2.mlstatic.com/D_NQ_NP_762388-MLA43761990728_102020-V.webp', 4, 19, '', ''),
(21, 'Peugeot 208 1,2 Like', '3000101', 'https://http2.mlstatic.com/D_NQ_NP_739036-MLA48218778060_112021-W.webp', 1, 23, '', ''),
(169, 'Pinza Niveladora P/ Porcelanato Ceramica Colocadora Subte A', '1386', 'https://http2.mlstatic.com/D_NQ_NP_653295-MLA47645529072_092021-V.webp', 3, 12, '', ''),
(141, 'Piso Encastrable Pack X 4 Goma Eva 60x60cm Entrenamiento Gym', '2999', 'https://http2.mlstatic.com/D_NQ_NP_967182-MLA45740888976_042021-V.webp', 4, 27, '', ''),
(333, 'Playmobil Lancha Con Motor 3142', '4', 'https://http2.mlstatic.com/D_NQ_NP_916373-MLA46316690799_062021-V.webp', 6, 27, '', ''),
(255, 'Polar Columbia Cierre Hombre Fast Trek Ii Manchester United', '159', 'https://http2.mlstatic.com/D_NQ_NP_617496-MLA47339636567_092021-W.webp', 5, 19, '', ''),
(117, 'Porta Base Cruz De Madera Notebook Tablet Computadora', '1286', 'https://http2.mlstatic.com/D_NQ_NP_606527-MLA50233731869_062022-V.webp', 8, 32, '', ''),
(214, 'Postre Royal Chocolate X 65 Gr', '109', 'https://http2.mlstatic.com/D_NQ_NP_733258-MLA47498406813_092021-W.webp', 7, 17, '', ''),
(208, 'Postre Royal Light Vainilla X 43 Gr', '119', 'https://http2.mlstatic.com/D_NQ_NP_634378-MLA47498407180_092021-W.webp', 7, 25, '', ''),
(96, 'Procesador AMD Athlon 3000G YD3000C6FHBOX de 2 n?cleos y 3.5GHz de frecuencia con gr?fica integrada', '17299', 'https://http2.mlstatic.com/D_NQ_NP_620009-MLA41419343230_042020-V.webp', 8, 33, '', ''),
(112, 'Procesador gamer Intel Core i5-11400 BX8070811400 de 6 n?cleos y 4.4GHz de frecuencia con gr?fica in', '34199', 'https://http2.mlstatic.com/D_NQ_NP_766062-MLA47177160024_082021-V.webp', 8, 26, '', ''),
(80, 'Procesador Intel Celeron G5905 BX80701G5905 de 2 n?cleos y 3.5GHz de frecuencia con gr?fica integrad', '6715', 'https://http2.mlstatic.com/D_NQ_NP_677873-MLA44347499239_122020-V.webp', 8, 36, '', ''),
(95, 'Procesador Intel Core i3-10100 BX8070110100 de 4 n?cleos y 4.3GHz de frecuencia con gr?fica integrad', '20799', 'https://http2.mlstatic.com/D_NQ_NP_872687-MLA43117152530_082020-V.webp', 8, 37, '', ''),
(116, 'Procesador Intel Core i3-10105 BX8070110105 de 4 n?cleos y 4.4GHz de frecuencia con gr?fica integrad', '1999', 'https://http2.mlstatic.com/D_NQ_NP_681640-MLA47573219581_092021-V.webp', 8, 13, '', ''),
(86, 'Procesador Intel Core i5-10400 BX8070110400 de 6 n?cleos y 4.3GHz de frecuencia con gr?fica integrad', '27999', 'https://http2.mlstatic.com/D_NQ_NP_686262-MLA42885718700_072020-V.webp', 8, 11, '', ''),
(107, 'Procesador Intel Core i7-10700F BX8070110700F de 8 n?cleos y 4.8GHz de frecuencia', '47699', 'https://http2.mlstatic.com/D_NQ_NP_674338-MLA44771883129_022021-V.webp', 8, 11, '', ''),
(152, 'Protector Cervical Importado Cubre Barra Get Fit', '143', 'https://http2.mlstatic.com/D_NQ_NP_607046-MLA48790785980_012022-V.webp', 4, 10, '', ''),
(128, 'Rack Sentadilla Piernas Posiciones Regulable Ajustable Pecho', '33999', 'https://http2.mlstatic.com/D_NQ_NP_760502-MLA48282751663_112021-V.webp', 4, 12, '', ''),
(183, 'Regla Met?lica Acero 30 Cm / 12 Pulgadas', '230', 'https://http2.mlstatic.com/D_NQ_NP_729741-MLA40763605126_022020-V.webp', 3, 17, '', ''),
(282, 'Remera Lacoste 2042522 Mangas Cortas Blanco Para Hombre', '9', 'https://http2.mlstatic.com/D_NQ_NP_797413-MLA49435613793_032022-W.webp', 5, 11, '', ''),
(11, 'Renault Alaskan 2,3 Bit 16v Iconic At 4x4', '9115000', 'https://http2.mlstatic.com/D_NQ_NP_988915-MLA49737038580_042022-W.webp', 1, 24, '', ''),
(28, 'Renault Captur 2,0 Intens Manual', '3860000', 'https://http2.mlstatic.com/D_NQ_NP_722834-MLA49729384878_042022-W.webp', 1, 23, '', ''),
(31, 'Renault Clio 1,2 M?o Expression Pack I', '1550000', 'https://http2.mlstatic.com/D_NQ_NP_973548-MLA50172690631_062022-W.webp', 1, 16, '', ''),
(16, 'Renault Duster 2,0 Ph2 4x2 Privilege', '3980000', 'https://http2.mlstatic.com/D_NQ_NP_743200-MLA50055763595_052022-W.webp', 1, 28, '', ''),
(334, 'Rompecabezas Antex Argentina 3036 de 36 piezas', '125', 'https://http2.mlstatic.com/D_NQ_NP_775702-MLA46653151719_072021-V.webp', 6, 32, '', ''),
(328, 'Rompecabezas Antex Mapamundi 1243 de 36 piezas', '1153', 'https://http2.mlstatic.com/D_NQ_NP_788747-MLA42683887476_072020-V.webp', 6, 20, '', ''),
(307, 'Rompecabezas Diente de Le?n Provincias Argentinas de 77 piezas', '238', 'https://http2.mlstatic.com/D_NQ_NP_837155-MLA46650570716_072021-V.webp', 6, 39, '', ''),
(332, 'Rompecabezas Trefl Santorini, Greece 26119 de 1500 piezas', '5399', 'https://http2.mlstatic.com/D_NQ_NP_951575-MLA42685385677_072020-V.webp', 6, 35, '', ''),
(329, 'Rompecabezas Trefl Tiger\'s Nest, Bhutan 27092 de 2000 piezas', '5999', 'https://http2.mlstatic.com/D_NQ_NP_850331-MLA42685385253_072020-V.webp', 6, 32, '', ''),
(131, 'Rueda Abdominales Doble Entrenamiento Fitness Ejercicio Gym', '1099', 'https://http2.mlstatic.com/D_NQ_NP_774894-MLA48985017840_012022-V.webp', 4, 22, '', ''),
(175, 'Saca Pelusas Motas Pelotitas Para Ropa Subte A Carabobo', '563', 'https://http2.mlstatic.com/D_NQ_NP_975850-MLA31577737543_072019-V.webp', 3, 39, '', ''),
(279, 'Saco Lap?n D Cacique Campera Pluma Duvet - Impermeable', '42998', 'https://http2.mlstatic.com/D_NQ_NP_631715-MLA42790099292_072020-W.webp', 5, 29, '', ''),
(199, 'Sal Entrefina Salero Colosal X 500 Gr', '123', 'https://http2.mlstatic.com/D_NQ_NP_677563-MLA47291821232_082021-W.webp', 7, 40, '', ''),
(198, 'Sal Fina Colosal Salero Dosis Flip X 150 Gr', '73', 'https://http2.mlstatic.com/D_NQ_NP_974110-MLA47291821220_082021-W.webp', 7, 23, '', ''),
(200, 'Sal Fina Salero Colosal X 500 Gr', '129', 'https://http2.mlstatic.com/D_NQ_NP_706580-MLA47291818859_082021-W.webp', 7, 35, '', ''),
(234, 'Salsa golf Hellmann\'s en doy pack 250?g', '124', 'https://http2.mlstatic.com/D_NQ_NP_645368-MLA44923397179_022021-W.webp', 7, 22, '', ''),
(63, 'Samsung Galaxy A03 128 GB negro 4 GB RAM', '35999', 'https://http2.mlstatic.com/D_NQ_NP_737913-MLA49433851234_032022-V.webp', 2, 17, '', ''),
(52, 'Samsung Galaxy S20 FE 5G 128 GB cloud navy 6 GB RAM', '91999', 'https://http2.mlstatic.com/D_NQ_NP_716191-MLA44281674442_122020-V.webp', 2, 29, '', ''),
(58, 'Samsung Galaxy S21 Fe Gris Oscuro 5g', '138999', 'https://http2.mlstatic.com/D_NQ_NP_730981-MLA48800008115_012022-V.webp', 2, 17, '', ''),
(50, 'Samsung Galaxy S22 Ultra 12gb 256gb Green', '203999', 'https://http2.mlstatic.com/D_NQ_NP_668050-MLA49303776893_032022-V.webp', 2, 10, '', ''),
(59, 'Samsung Galaxy S22+ 8gb 256gb Phantom White', '167999', 'https://http2.mlstatic.com/D_NQ_NP_668120-MLA49303776830_032022-V.webp', 2, 19, '', ''),
(56, 'Samsung Galaxy Z Flip3 5G 128 GB phantom black 8 GB RAM', '126999', 'https://http2.mlstatic.com/D_NQ_NP_676914-MLA47887111268_102021-V.webp', 2, 18, '', ''),
(313, 'Set De Cocina Y Accesorios Piccola Rondi', '485', 'https://http2.mlstatic.com/D_NQ_NP_736489-MLA41120613580_032020-V.webp', 6, 28, '', ''),
(335, 'Set Lol Surprise Tweens Francy Gurl Incluye Accesorios', '966', 'https://http2.mlstatic.com/D_NQ_NP_960874-MLA48357162412_112021-V.webp', 6, 14, '', ''),
(299, 'Set X6 Mu?ecos Roblox Desarmables Articulados Con Accesorios', '293', 'https://http2.mlstatic.com/D_NQ_NP_906066-MLA48468329996_122021-V.webp', 6, 30, '', ''),
(10, 'Shineray T30 T30 Minitruck Cab Simple', '2739000', 'https://http2.mlstatic.com/D_NQ_NP_726761-MLA49094618739_022022-W.webp', 1, 34, '', ''),
(3, 'Shineray X30 X30 Furgon', '2749000', 'https://http2.mlstatic.com/D_NQ_NP_890007-MLA49092194539_022022-W.webp', 1, 15, '', ''),
(296, 'Silla mecedora para beb? Funny Smile FS 8620 marr?n', '9999', 'https://http2.mlstatic.com/D_NQ_NP_971217-MLA49910396873_052022-V.webp', 6, 40, '', ''),
(197, 'Snacks De Ma?z Deluxe Bla Bla Queso Cheddar X 80 Gr', '153', 'https://http2.mlstatic.com/D_NQ_NP_788821-MLA47264397115_082021-W.webp', 7, 20, '', ''),
(102, 'Soporte Bandeja Notebook Computadora Mesa Cama Home Office', '5599', 'https://http2.mlstatic.com/D_NQ_NP_649325-MLA46786579204_072021-V.webp', 8, 31, '', ''),
(93, 'Soporte Notebook Mesa Bandeja Computadora Cama Home Office', '5664', 'https://http2.mlstatic.com/D_NQ_NP_692409-MLA48405408467_122021-V.webp', 8, 40, '', ''),
(294, 'Supermercado Registrado Petit Gourmet Fast Food Casa Valente', '6699', 'https://http2.mlstatic.com/D_NQ_NP_783393-MLA44050984638_112020-V.webp', 6, 29, '', ''),
(263, 'Sweater Cuello Redondo Algod?n Hombre Mistral 14790n', '109', 'https://http2.mlstatic.com/D_NQ_NP_867323-MLA45792848519_052021-W.webp', 5, 10, '', ''),
(257, 'Sweater Pullover Escote V Algod?n Hombre Mistral 14791-2', '109', 'https://http2.mlstatic.com/D_NQ_NP_973357-MLA45754784907_042021-W.webp', 5, 18, '', ''),
(57, 'TCL 20 SE 256 GB nuit black 6 GB RAM', '41999', 'https://http2.mlstatic.com/D_NQ_NP_623923-MLA48764808772_012022-V.webp', 2, 23, '', ''),
(122, 'Teclado Gamer Retroiluminado Rgb Computadora Usb Led Efecto', '1299', 'https://http2.mlstatic.com/D_NQ_NP_973423-MLA49985099453_052022-V.webp', 8, 39, '', ''),
(101, 'Teclado gamer Soul XK700 QWERTY ingl?s color negro con luz RGB', '1283', 'https://http2.mlstatic.com/D_NQ_NP_613689-MLA43542876224_092020-V.webp', 8, 15, '', ''),
(78, 'Teclado Global Electronics K103BKUSB QWERTY espa?ol Espa?a color negro', '370', 'https://http2.mlstatic.com/D_NQ_NP_666059-MLA43541492818_092020-V.webp', 8, 13, '', ''),
(92, 'Teclado Noganet NKB-78011 QWERTY espa?ol color negro', '523', 'https://http2.mlstatic.com/D_NQ_NP_634788-MLA42519841414_072020-V.webp', 8, 26, '', ''),
(158, 'Tobillera Con Peso 2 Kg - El Par - Reforzadas Fitness', '799', 'https://http2.mlstatic.com/D_NQ_NP_723171-MLA49011823483_022022-V.webp', 4, 39, '', ''),
(321, 'Tocador Infantil De Juguete Little Star Con Luz Y Maletin', '799', 'https://http2.mlstatic.com/D_NQ_NP_804168-MLA42362087330_062020-V.webp', 6, 24, '', ''),
(192, 'Tonfa,baston Extensible,cachiporra + Manopla Acero Subte A', '1576', 'https://http2.mlstatic.com/D_NQ_NP_957666-MLA31351889349_072019-V.webp', 3, 16, '', ''),
(242, 'Topper Boro II Hombre Adultos', '9144', 'https://http2.mlstatic.com/D_NQ_NP_731302-MLA49200679366_022022-W.webp', 5, 38, '', ''),
(247, 'Topper Gondor II Hombre Adultos', '11899', 'https://http2.mlstatic.com/D_NQ_NP_875526-MLA49201325445_022022-W.webp', 5, 33, '', ''),
(243, 'Topper Squat Hombre Adultos', '11099', 'https://http2.mlstatic.com/D_NQ_NP_950723-MLA47397900331_092021-W.webp', 5, 29, '', ''),
(268, 'Topper Strong Pace III Hombre Adultos', '9099', 'https://http2.mlstatic.com/D_NQ_NP_898412-MLA48322250504_112021-W.webp', 5, 15, '', ''),
(272, 'Topper X-Forcer Hombre Adultos', '11599', 'https://http2.mlstatic.com/D_NQ_NP_614458-MLA49249945961_032022-W.webp', 5, 32, '', ''),
(66, 'Toyota Corolla 2014 tremendo maquinon', '4000000', 'https://http2.mlstatic.com/D_NQ_NP_909388-MLA48831366604_012022-W.webp', 1, 33, '', ''),
(19, 'Toyota Etios 1,5 Sedan Xls My19', '2800000', 'https://http2.mlstatic.com/D_NQ_NP_896100-MLA50013283773_052022-W.webp', 1, 12, '', ''),
(4, 'Toyota Etios 1,5 Xls At', '3300000', 'https://http2.mlstatic.com/D_NQ_NP_712027-MLA49294846519_032022-W.webp', 1, 12, '', ''),
(15, 'Toyota Sw4 2,8 Sr 177cv 4x4', '8900000', 'https://http2.mlstatic.com/D_NQ_NP_760592-MLA50007858053_052022-W.webp', 1, 11, '', ''),
(18, 'Toyota Yaris 1,5 107cv S', '4500000', 'https://http2.mlstatic.com/D_NQ_NP_655420-MLA50110044685_052022-W.webp', 1, 32, '', ''),
(190, 'Transformador 220 V Ac 110 V Ac 300w Reforzado', '3758', 'https://http2.mlstatic.com/D_NQ_NP_909339-MLA47249255302_082021-V.webp', 3, 13, '', ''),
(305, 'Triciclo Rondi 1er Triciclo c/Barra rosa y violeta', '475', 'https://http2.mlstatic.com/D_NQ_NP_839560-MLA48483830272_122021-V.webp', 6, 18, '', ''),
(317, 'Triciclo Rondi Glam rosa', '55', 'https://http2.mlstatic.com/D_NQ_NP_796775-MLA48571851015_122021-V.webp', 6, 32, '', ''),
(322, 'Triciclo Rondi Junior Rider azul, rojo y blanco', '319', 'https://http2.mlstatic.com/D_NQ_NP_621284-MLA48481279613_122021-V.webp', 6, 36, '', ''),
(235, 'Turr?n De Man? Nat X 25 G', '14', 'https://http2.mlstatic.com/D_NQ_NP_970719-MLA48809469366_012022-W.webp', 7, 37, '', ''),
(97, 'Ups Estabilizador Kaise 800va Display Lcd 4 Salidas Modem', '1195', 'https://http2.mlstatic.com/D_NQ_NP_855471-MLA41754794987_052020-V.webp', 8, 13, '', ''),
(99, 'UPS TRV Electronics NEO 650 650VA entrada y salida de 220V negro', '903', 'https://http2.mlstatic.com/D_NQ_NP_633060-MLA44453027235_122020-V.webp', 8, 15, '', ''),
(89, 'UPS TRV Electronics NEO 850 850VA entrada y salida de 220V negro', '11711', 'https://http2.mlstatic.com/D_NQ_NP_858870-MLA44453012150_122020-V.webp', 8, 10, '', ''),
(338, 'Valija Juliana Doctora Chica Con Accesorios D012', '1818', 'https://http2.mlstatic.com/D_NQ_NP_678224-MLA43035351430_082020-V.webp', 6, 38, '', ''),
(303, 'Valija Juliana Make Up Unicorn Maquillaje Unicornio', '2651', 'https://http2.mlstatic.com/D_NQ_NP_716852-MLA43035677103_082020-V.webp', 6, 37, '', ''),
(325, 'Valija Juliana Make Up Unicorn Pinturas+accesorios Fionatoys', '1994', 'https://http2.mlstatic.com/D_NQ_NP_966105-MLA44083168460_112020-V.webp', 6, 15, '', ''),
(217, 'Vauquita X18 Unidades Light Y Clasica - Oferta Sweet Market', '810', 'https://http2.mlstatic.com/D_NQ_NP_858867-MLA49182941719_022022-W.webp', 7, 23, '', ''),
(7, 'Volkswagen Amarok 2,0 Cd Tdi 140cv 4x2 Startline', '4290000', 'https://http2.mlstatic.com/D_NQ_NP_942782-MLA49094761364_022022-W.webp', 1, 22, '', ''),
(22, 'Volkswagen Amarok 2,0 Cd Tdi 140cv Trendline', '5791000', 'https://http2.mlstatic.com/D_NQ_NP_935220-MLA42592455505_072020-W.webp', 1, 17, '', ''),
(12, 'Volkswagen Amarok 2,0 Cd Tdi 180cv 4x4 Highline Pack At', '7450000', 'https://http2.mlstatic.com/D_NQ_NP_828700-MLA49616177165_042022-W.webp', 1, 29, '', ''),
(41, 'Volkswagen Fox 1,6 Trendline', '2200000', 'https://http2.mlstatic.com/D_NQ_NP_683607-MLA50147976901_052022-W.webp', 1, 15, '', ''),
(25, 'Volkswagen Scirocco 2,0 Tsi 211cv', '5200000', 'https://http2.mlstatic.com/D_NQ_NP_725416-MLA50010728855_052022-W.webp', 1, 12, '', ''),
(110, 'Woox - Atril Base De Dise?o Apoya Notebook Computadora', '1969', 'https://http2.mlstatic.com/D_NQ_NP_948147-MLA50212466670_062022-V.webp', 8, 17, '', ''),
(62, 'Xiaomi Mi 11 Lite 5G NE Dual SIM 128 GB negro trufa 8 GB RAM', '81999', 'https://http2.mlstatic.com/D_NQ_NP_768134-MLA48496137270_122021-V.webp', 2, 10, '', ''),
(49, 'Xiaomi Redmi Note 10 Pro (Global) Dual SIM 128 GB azul glaciar 6 GB RAM', '75659', 'https://http2.mlstatic.com/D_NQ_NP_769912-MLA46490092919_062021-V.webp', 2, 28, '', ''),
(64, 'Xiaomi Redmi Note 11 (Snapdragon) Dual SIM 128 GB gris grafito 4 GB RAM', '59999', 'https://http2.mlstatic.com/D_NQ_NP_664692-MLA49654660759_042022-V.webp', 2, 24, '', ''),
(48, 'Xiaomi Redmi Note 11s Dual SIM 128 GB gris grafito 6 GB RAM', '72499', 'https://http2.mlstatic.com/D_NQ_NP_734580-MLA49397636457_032022-V.webp', 2, 19, '', ''),
(271, 'Zapartillas Coreracer adidas Sport 78 Tienda Oficial', '11999', 'https://http2.mlstatic.com/D_NQ_NP_828977-MLA50128052624_052022-W.webp', 5, 15, '', ''),
(270, 'Zapatilla Hombre Topper T.700 Marino Ensign/ Azul Mood', '10439', 'https://http2.mlstatic.com/D_NQ_NP_904375-MLA45670837676_042021-W.webp', 5, 39, '', ''),
(273, 'Zapatilla Puma Hombre Urbana Cuero Negra Caven', '11999', 'https://http2.mlstatic.com/D_NQ_NP_761373-MLA49725845798_042022-W.webp', 5, 38, '', ''),
(249, 'Zapatilla Reaction Negro Atomik Hombre', '949', 'https://http2.mlstatic.com/D_NQ_NP_863117-MLA47357700099_092021-W.webp', 5, 20, '', ''),
(265, 'Zapatilla Topper Hombre Vr Speed Gris/verde', '9519', 'https://http2.mlstatic.com/D_NQ_NP_745772-MLA48537408394_122021-W.webp', 5, 24, '', ''),
(283, 'Zapatilla Topper Strong Pace Iii Azul Moro/gris Micro', '9099', 'https://http2.mlstatic.com/D_NQ_NP_684209-MLA47765996182_102021-W.webp', 5, 31, '', '');
INSERT INTO `products` (`product_id`, `name`, `price`, `picture`, `category_id`, `stock`, `description`, `currency_id`) VALUES
(256, 'Zapatilla Topper Tenis Tie Break Blanco Azul Ras', '9349', 'https://http2.mlstatic.com/D_NQ_NP_703474-MLA40436495686_012020-W.webp', 5, 29, '', ''),
(287, 'Zapatilla Trekking Cuero Gamuzado Hombre Ringo Gorc 02', '178', 'https://http2.mlstatic.com/D_NQ_NP_932103-MLA44306564607_122020-W.webp', 5, 26, '', ''),
(284, 'Zapatillas Casual Hombre Caterpillar Camber Gri Beige Caucho', '249', 'https://http2.mlstatic.com/D_NQ_NP_993708-MLA46872347936_072021-W.webp', 5, 17, '', ''),
(262, 'Zapatillas Deportivas Trekking Hombre Topper Gondor Ii', '10709', 'https://http2.mlstatic.com/D_NQ_NP_874254-MLA48700023207_122021-W.webp', 5, 19, '', ''),
(286, 'Zapatillas Fila Discovery Azul Gris Hombre', '9599', 'https://http2.mlstatic.com/D_NQ_NP_806256-MLA49866897269_052022-W.webp', 5, 36, '', ''),
(278, 'Zapatillas Montagne Bota Outdoor Hombre Ona Verde-neg Cli', '16499', 'https://http2.mlstatic.com/D_NQ_NP_818901-MLA50107128581_052022-W.webp', 5, 40, '', ''),
(254, 'Zapatillas Montagne Outdoor Hombre Daylite Gris-nja Cli', '13999', 'https://http2.mlstatic.com/D_NQ_NP_975684-MLA50100102466_052022-W.webp', 5, 25, '', ''),
(277, 'Zapatillas Olympikus Running Hombre G Iris Negro-dorado Cli', '11699', 'https://http2.mlstatic.com/D_NQ_NP_921069-MLA49335993463_032022-W.webp', 5, 28, '', ''),
(252, 'Zapatillas Puma Hombre Deportivas Negra Better Foam Emerge', '12149', 'https://http2.mlstatic.com/D_NQ_NP_707795-MLA49725845888_042022-W.webp', 5, 15, '', ''),
(245, 'Zapatillas Topper Dakota Negro', '10259', 'https://http2.mlstatic.com/D_NQ_NP_883023-MLA45814509264_052021-W.webp', 5, 32, '', ''),
(251, 'Zapatillas Topper Outdoor Hombre Gondor Il Negro-azul Cli', '11899', 'https://http2.mlstatic.com/D_NQ_NP_640213-MLA44945553140_022021-W.webp', 5, 11, '', ''),
(276, 'Zapatillas Topper Padel Hombre T-padel Gris-negro Cli', '11499', 'https://http2.mlstatic.com/D_NQ_NP_690151-MLA49851403744_052022-W.webp', 5, 29, '', ''),
(258, 'Zapatillas Vans Lifestyle Hombre Authentic Negro Cli', '11099', 'https://http2.mlstatic.com/D_NQ_NP_996814-MLA45840310269_052021-W.webp', 5, 37, '', ''),
(288, 'Zapato Cuero Hombre Briganti Acordonado Negro - Hcac00969', '20978', 'https://http2.mlstatic.com/D_NQ_NP_766356-MLA47133342258_082021-W.webp', 5, 28, '', '');

-- --------------------------------------------------------



--
-- Índices para tablas volcadas
--

--
-- Indices de la tabla `addresses`
--


--
-- Indices de la tabla `orders`
--
ALTER TABLE `orders`
  ADD PRIMARY KEY (`id`);

--
-- Indices de la tabla `order_details`
--
ALTER TABLE `order_details`
  ADD UNIQUE KEY `order_detail_id` (`order_detail_id`);

--
-- Indices de la tabla `products`
--
ALTER TABLE `products`
  ADD UNIQUE KEY `name` (`name`);

--
-- Indices de la tabla `users`
--
ALTER TABLE `users`
  ADD UNIQUE KEY `username` (`username`),
  ADD UNIQUE KEY `email` (`email`);

--
-- AUTO_INCREMENT de las tablas volcadas
--

--
-- AUTO_INCREMENT de la tabla `orders`
--
ALTER TABLE `orders`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=25;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
