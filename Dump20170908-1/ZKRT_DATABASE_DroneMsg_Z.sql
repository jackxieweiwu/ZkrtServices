CREATE DATABASE  IF NOT EXISTS `ZKRT_DATABASE` /*!40100 DEFAULT CHARACTER SET latin1 */;
USE `ZKRT_DATABASE`;
-- MySQL dump 10.13  Distrib 5.7.19, for Linux (x86_64)
--
-- Host: localhost    Database: ZKRT_DATABASE
-- ------------------------------------------------------
-- Server version	5.7.19-0ubuntu0.17.04.1

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
-- Table structure for table `DroneMsg_Z`
--

DROP TABLE IF EXISTS `DroneMsg_Z`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `DroneMsg_Z` (
  `DId` int(11) NOT NULL AUTO_INCREMENT,
  `DroneID` varchar(100) NOT NULL,
  `DroneAlt` float NOT NULL,
  `DroneYaw` float NOT NULL,
  `DronePitch` float NOT NULL,
  `DroneSpeed` float NOT NULL,
  `LevelName` varchar(255) DEFAULT NULL,
  `DroneBool` int(11) NOT NULL,
  `DroneDateTime` datetime NOT NULL,
  `DroneDate` time NOT NULL,
  PRIMARY KEY (`DId`),
  UNIQUE KEY `DId_UNIQUE` (`DId`),
  UNIQUE KEY `DroneID_UNIQUE` (`DroneID`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `DroneMsg_Z`
--

LOCK TABLES `DroneMsg_Z` WRITE;
/*!40000 ALTER TABLE `DroneMsg_Z` DISABLE KEYS */;
INSERT INTO `DroneMsg_Z` VALUES (2,'1113',12.3,2.4,2.1,13,NULL,1,'2017-09-08 06:28:04','06:28:04');
/*!40000 ALTER TABLE `DroneMsg_Z` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2017-09-08 23:03:40