-- MySQL dump 10.13  Distrib 5.7.25, for osx10.9 (x86_64)
--
-- Host: 127.0.0.1    Database: loket
-- ------------------------------------------------------
-- Server version	5.7.25

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `customer`
--

DROP TABLE IF EXISTS `customer`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `customer` (
  `id_customer` varchar(128) NOT NULL,
  `name` varchar(100) NOT NULL,
  `email` varchar(100) NOT NULL,
  `phone` varchar(100) NOT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id_customer`),
  UNIQUE KEY `customer_id_customer_uindex` (`id_customer`),
  UNIQUE KEY `customer_email_uindex` (`email`),
  UNIQUE KEY `customer_phone_uindex` (`phone`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `customer`
--

LOCK TABLES `customer` WRITE;
/*!40000 ALTER TABLE `customer` DISABLE KEYS */;
INSERT INTO `customer` (`id_customer`, `name`, `email`, `phone`, `created_at`, `updated_at`) VALUES ('6e0be810-1b80-11eb-97ee-2f8a3df60fc6','Q','ka@k.com','4231','2020-10-31 13:53:18',NULL),('855e09f0-1b7e-11eb-8bcb-1a151d514255','K','k@k.com','1212312312','2020-10-31 13:39:38',NULL);
/*!40000 ALTER TABLE `customer` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `event`
--

DROP TABLE IF EXISTS `event`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `event` (
  `id_event` varchar(128) NOT NULL,
  `id_location` varchar(128) NOT NULL,
  `name` varchar(100) NOT NULL,
  `desc` text,
  `start_date` datetime NOT NULL,
  `end_date` datetime NOT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id_event`),
  UNIQUE KEY `event_id_event_uindex` (`id_event`),
  KEY `event_location_id_location_fk` (`id_location`),
  CONSTRAINT `event_location_id_location_fk` FOREIGN KEY (`id_location`) REFERENCES `location` (`id_location`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `event`
--

LOCK TABLES `event` WRITE;
/*!40000 ALTER TABLE `event` DISABLE KEYS */;
INSERT INTO `event` (`id_event`, `id_location`, `name`, `desc`, `start_date`, `end_date`, `created_at`, `updated_at`) VALUES ('48d25390-1b71-11eb-9d1a-f94f590013b6','81865ce0-1b67-11eb-8dc2-8105b2bf0401','Over the cloud','up','2020-10-31 03:53:53','2021-10-31 03:53:54','2020-10-31 12:04:53',NULL),('6765a0a0-1b71-11eb-9d1a-f94f590013b6','81865ce0-1b67-11eb-8dc2-8105b2bf0401','Over the cloud 2','up','2020-10-31 03:53:53','2020-10-31 03:53:54','2020-10-31 12:05:44',NULL),('cfb3f480-1b2c-11eb-8bcb-1a151d514255','a3fbcf16-1b2c-11eb-8bcb-1a151d514255','Sapu Gerbang','Sapu Gerbang','2020-10-31 10:53:53','2021-10-31 10:53:54','2020-10-31 10:54:44',NULL);
/*!40000 ALTER TABLE `event` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `location`
--

DROP TABLE IF EXISTS `location`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `location` (
  `id_location` varchar(128) NOT NULL,
  `name` varchar(100) NOT NULL,
  `address` text NOT NULL,
  `street` text NOT NULL,
  `city` varchar(100) NOT NULL,
  `country` varchar(100) NOT NULL,
  `zip` varchar(10) NOT NULL,
  `latitude` varchar(20) NOT NULL,
  `longitude` varchar(20) NOT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id_location`),
  UNIQUE KEY `location_id_location_uindex` (`id_location`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `location`
--

LOCK TABLES `location` WRITE;
/*!40000 ALTER TABLE `location` DISABLE KEYS */;
INSERT INTO `location` (`id_location`, `name`, `address`, `street`, `city`, `country`, `zip`, `latitude`, `longitude`, `created_at`, `updated_at`) VALUES ('81865ce0-1b67-11eb-8dc2-8105b2bf0401','Over the cloud','up','12','up','cloud','24123','-232.3412','62.2313','2020-10-31 10:54:53',NULL),('a3fbcf16-1b2c-11eb-8bcb-1a151d514255','Lapangan Sukarame','Sukarame','Sukarame','Sukarame','Sukarame','23124','123','123','2020-10-31 10:53:30',NULL);
/*!40000 ALTER TABLE `location` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `order`
--

DROP TABLE IF EXISTS `order`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `order` (
  `id_order` varchar(128) NOT NULL,
  `id_transaction` varchar(128) NOT NULL,
  `id_ticket` varchar(128) DEFAULT NULL,
  `order_qty` int(11) DEFAULT NULL,
  `order_amount` int(11) DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id_order`),
  UNIQUE KEY `order_id_order_uindex` (`id_order`),
  KEY `order_ticket_id_ticket_fk` (`id_ticket`),
  KEY `order_transaction_id_transaction_fk` (`id_transaction`),
  CONSTRAINT `order_ticket_id_ticket_fk` FOREIGN KEY (`id_ticket`) REFERENCES `ticket` (`id_ticket`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `order_transaction_id_transaction_fk` FOREIGN KEY (`id_transaction`) REFERENCES `transaction` (`id_transaction`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `order`
--

LOCK TABLES `order` WRITE;
/*!40000 ALTER TABLE `order` DISABLE KEYS */;
INSERT INTO `order` (`id_order`, `id_transaction`, `id_ticket`, `order_qty`, `order_amount`, `created_at`, `updated_at`) VALUES ('04bc0e50-1ba1-11eb-be51-775568593ebf','04b9c460-1ba1-11eb-be51-775568593ebf','14b47100-1b77-11eb-b7bc-b5b5561267b1',2,24,'2020-10-31 17:46:34',NULL),('90c24ef0-1ba1-11eb-bdbb-f3a563d2b916','90bde220-1ba1-11eb-bdbb-f3a563d2b916','14b47100-1b77-11eb-b7bc-b5b5561267b1',2,24,'2020-10-31 17:50:29',NULL),('90c24ef1-1ba1-11eb-bdbb-f3a563d2b916','90bde220-1ba1-11eb-bdbb-f3a563d2b916','7081e2f6-1b40-11eb-8bcb-1a151d514255',2,20000,'2020-10-31 17:50:29',NULL),('edb5bd00-1ba0-11eb-858f-7bc0ed743898','edb23a90-1ba0-11eb-858f-7bc0ed743898','14b47100-1b77-11eb-b7bc-b5b5561267b1',1,12,'2020-10-31 17:45:56',NULL),('fabc64e0-1ba0-11eb-be51-775568593ebf','fab90980-1ba0-11eb-be51-775568593ebf','14b47100-1b77-11eb-b7bc-b5b5561267b1',1,12,'2020-10-31 17:46:18',NULL);
/*!40000 ALTER TABLE `order` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ticket`
--

DROP TABLE IF EXISTS `ticket`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ticket` (
  `id_ticket` varchar(128) NOT NULL,
  `id_event` varchar(128) NOT NULL,
  `name` varchar(100) NOT NULL,
  `desc` text NOT NULL,
  `stock` int(11) NOT NULL,
  `price` int(11) NOT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id_ticket`),
  UNIQUE KEY `ticket_id_ticket_uindex` (`id_ticket`),
  KEY `ticket_event_id_event_fk` (`id_event`),
  CONSTRAINT `ticket_event_id_event_fk` FOREIGN KEY (`id_event`) REFERENCES `event` (`id_event`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ticket`
--

LOCK TABLES `ticket` WRITE;
/*!40000 ALTER TABLE `ticket` DISABLE KEYS */;
INSERT INTO `ticket` (`id_ticket`, `id_event`, `name`, `desc`, `stock`, `price`, `created_at`, `updated_at`) VALUES ('14b47100-1b77-11eb-b7bc-b5b5561267b1','6765a0a0-1b71-11eb-9d1a-f94f590013b6','Over the cloud','up',6,12,'2020-10-31 12:46:22','2020-10-31 17:50:29'),('7081e2f6-1b40-11eb-8bcb-1a151d514255','cfb3f480-1b2c-11eb-8bcb-1a151d514255','Event Sapu','Event Sapu',8,10000,'2020-10-31 13:15:14','2020-10-31 17:50:29');
/*!40000 ALTER TABLE `ticket` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `transaction`
--

DROP TABLE IF EXISTS `transaction`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `transaction` (
  `id_transaction` varchar(128) NOT NULL,
  `id_customer` varchar(128) NOT NULL,
  `invoice_number` varchar(128) NOT NULL,
  `invoice_amount` int(20) DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id_transaction`),
  UNIQUE KEY `transaction_id_transaction_uindex` (`id_transaction`),
  KEY `transaction_customer_id_customer_fk` (`id_customer`),
  CONSTRAINT `transaction_customer_id_customer_fk` FOREIGN KEY (`id_customer`) REFERENCES `customer` (`id_customer`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `transaction`
--

LOCK TABLES `transaction` WRITE;
/*!40000 ALTER TABLE `transaction` DISABLE KEYS */;
INSERT INTO `transaction` (`id_transaction`, `id_customer`, `invoice_number`, `invoice_amount`, `created_at`, `updated_at`) VALUES ('04b9c460-1ba1-11eb-be51-775568593ebf','6e0be810-1b80-11eb-97ee-2f8a3df60fc6','LK-MTYwNDE2NjM5NDc4Ng==',24,'2020-10-31 17:46:34',NULL),('90bde220-1ba1-11eb-bdbb-f3a563d2b916','6e0be810-1b80-11eb-97ee-2f8a3df60fc6','LK-MTYwNDE2NjYyOTY5NA==',20024,'2020-10-31 17:50:29',NULL),('edb23a90-1ba0-11eb-858f-7bc0ed743898','6e0be810-1b80-11eb-97ee-2f8a3df60fc6','LK-MTYwNDE2NjM1NjE0Nw==',12,'2020-10-31 17:45:56',NULL),('fab90980-1ba0-11eb-be51-775568593ebf','6e0be810-1b80-11eb-97ee-2f8a3df60fc6','LK-MTYwNDE2NjM3Nzk5Ng==',12,'2020-10-31 17:46:18',NULL);
/*!40000 ALTER TABLE `transaction` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-11-01  1:22:18
