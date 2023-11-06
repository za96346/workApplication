-- MySQL dump 10.13  Distrib 8.0.33, for Win64 (x86_64)
--
-- Host: localhost    Database: new_workapp
-- ------------------------------------------------------
-- Server version	8.0.33

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `company`
--

DROP TABLE IF EXISTS `company`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `company` (
  `companyId` int NOT NULL COMMENT '公司id',
  `companyCode` varchar(50) NOT NULL COMMENT '公司代碼',
  `companyName` varchar(200) DEFAULT NULL COMMENT '公司名稱',
  `companyLocation` varchar(200) DEFAULT NULL COMMENT '公司位置',
  `companyPhoneNumber` varchar(20) DEFAULT NULL COMMENT '公司電話',
  `bossId` int NOT NULL DEFAULT '-1' COMMENT '負責人userId',
  `createTime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '創建時間',
  `lastModify` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '最後更新時間',
  PRIMARY KEY (`companyId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='公司';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `company`
--

LOCK TABLES `company` WRITE;
/*!40000 ALTER TABLE `company` DISABLE KEYS */;
INSERT INTO `company` VALUES (0,'ddd','ddd','wjiefoijwi','0906930873',5,'2023-11-06 15:11:34','2023-11-06 15:23:14');
/*!40000 ALTER TABLE `company` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `company_banch`
--

DROP TABLE IF EXISTS `company_banch`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `company_banch` (
  `companyId` int NOT NULL COMMENT '公司Id',
  `banchId` int NOT NULL COMMENT '部門id',
  `banchName` varchar(50) NOT NULL COMMENT '部門名稱',
  `deleteFlag` char(1) DEFAULT 'N' COMMENT '刪除旗標 ( N, Y )',
  `deleteTime` datetime DEFAULT NULL COMMENT '刪除時間',
  `createTime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '創建時間',
  `lastModify` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '最後更新時間',
  PRIMARY KEY (`companyId`,`banchId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='公司部門';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `company_banch`
--

LOCK TABLES `company_banch` WRITE;
/*!40000 ALTER TABLE `company_banch` DISABLE KEYS */;
INSERT INTO `company_banch` VALUES (0,1,'資訊組','N',NULL,'2023-11-04 17:14:03','2023-11-04 17:14:03'),(0,2,'輔導組','N',NULL,'2023-11-04 17:14:03','2023-11-04 17:14:03'),(0,3,'研發部部','Y','2023-11-06 16:21:50','2023-11-06 16:14:49','2023-11-06 16:21:50');
/*!40000 ALTER TABLE `company_banch` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `function_item`
--

DROP TABLE IF EXISTS `function_item`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `function_item` (
  `funcCode` varchar(255) NOT NULL COMMENT '功能代碼(banchManager)',
  `funcName` varchar(30) DEFAULT NULL COMMENT '功能名稱',
  `createTime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '創建時間',
  `lastModify` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '最近修改',
  `scopeRoleEnable` char(1) DEFAULT 'Y' COMMENT '可編輯角色範圍',
  `scopeBanchEnable` char(1) DEFAULT 'Y' COMMENT '可編輯部門範圍',
  PRIMARY KEY (`funcCode`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='功能項目;';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `function_item`
--

LOCK TABLES `function_item` WRITE;
/*!40000 ALTER TABLE `function_item` DISABLE KEYS */;
INSERT INTO `function_item` VALUES ('banchManage','部門管理','2023-11-02 14:48:33','2023-11-02 14:48:33','N','Y'),('companyData','公司基本資料','2023-11-02 14:47:56','2023-11-02 14:47:56','N','N'),('employeeManage','員工管理','2023-11-02 14:49:06','2023-11-02 14:49:06','Y','Y'),('performance','績效管理','2023-11-02 14:50:03','2023-11-02 14:50:03','Y','Y'),('roleManage','角色管理','2023-11-04 14:55:50','2023-11-04 14:55:50','N','N'),('selfData','基本資料','2023-11-02 14:47:03','2023-11-02 14:47:03','N','N'),('shift','排班管理','2023-11-02 14:51:32','2023-11-02 14:51:32','Y','Y'),('shiftSetting','班表設定','2023-11-02 14:51:01','2023-11-02 14:51:01','N','Y'),('yearPerformance','年度績效管理','2023-11-02 14:50:03','2023-11-02 14:50:03','Y','Y');
/*!40000 ALTER TABLE `function_item` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `operation_item`
--

DROP TABLE IF EXISTS `operation_item`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `operation_item` (
  `operationCode` varchar(255) NOT NULL COMMENT '操作代碼',
  `operationName` varchar(255) DEFAULT NULL COMMENT '操作名稱',
  `createTime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '創建時間',
  `lastModify` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '最近修改',
  PRIMARY KEY (`operationCode`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='操作代碼;';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `operation_item`
--

LOCK TABLES `operation_item` WRITE;
/*!40000 ALTER TABLE `operation_item` DISABLE KEYS */;
INSERT INTO `operation_item` VALUES ('add','新增','2023-11-02 14:56:43','2023-11-02 14:56:43'),('copy','複製','2023-11-02 14:56:43','2023-11-02 14:56:43'),('delete','刪除','2023-11-02 14:56:43','2023-11-02 14:56:43'),('edit','編輯','2023-11-02 14:56:43','2023-11-02 14:56:43'),('inquire','查詢','2023-11-02 14:56:43','2023-11-02 14:56:43'),('print','列印','2023-11-02 14:57:21','2023-11-02 14:57:21');
/*!40000 ALTER TABLE `operation_item` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `role`
--

DROP TABLE IF EXISTS `role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `role` (
  `companyId` int NOT NULL COMMENT '公司id',
  `roleId` int NOT NULL COMMENT '角色id',
  `roleName` varchar(30) DEFAULT NULL COMMENT '角色名稱',
  `stopFlag` char(1) DEFAULT 'N' COMMENT '停用旗標 ( N, Y )',
  `deleteFlag` varchar(255) DEFAULT 'N' COMMENT '刪除旗標 ( N, Y )',
  `deleteTime` datetime DEFAULT NULL COMMENT '刪除時間',
  `createTime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '創建時間',
  `lastModify` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '最近修改',
  PRIMARY KEY (`companyId`,`roleId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='角色表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `role`
--

LOCK TABLES `role` WRITE;
/*!40000 ALTER TABLE `role` DISABLE KEYS */;
INSERT INTO `role` VALUES (0,1,'admin','N','N',NULL,'2023-11-02 19:53:02','2023-11-06 16:10:49'),(0,2,'admin','N','N',NULL,'2023-11-02 19:53:16','2023-11-04 13:42:08');
/*!40000 ALTER TABLE `role` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `role_struct`
--

DROP TABLE IF EXISTS `role_struct`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `role_struct` (
  `companyId` int NOT NULL COMMENT '公司id',
  `roleId` int NOT NULL COMMENT '角色id',
  `funcCode` varchar(255) NOT NULL COMMENT '功能代碼( banchManage, shiftedit )',
  `itemCode` varchar(255) NOT NULL COMMENT '操作代碼(edit, delete...)',
  `scopeBanch` text COMMENT '可操作部門範圍 ( 部門ID[] )',
  `scopeRole` text COMMENT '可操作角色範圍 ( 角色ID[] )',
  `createTime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '創建時間',
  `lastModify` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '最近修改',
  PRIMARY KEY (`companyId`,`roleId`,`funcCode`,`itemCode`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='角色結構表;';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `role_struct`
--

LOCK TABLES `role_struct` WRITE;
/*!40000 ALTER TABLE `role_struct` DISABLE KEYS */;
INSERT INTO `role_struct` VALUES (0,1,'banchManage','add','self','[1,2]','2023-11-06 16:10:49','2023-11-06 16:10:49'),(0,1,'banchManage','delete','all','[1,2,3]','2023-11-06 16:10:49','2023-11-06 16:10:49'),(0,1,'banchManage','edit','all','[1,2,3]','2023-11-06 16:10:49','2023-11-06 16:10:49'),(0,1,'banchManage','inquire','self','[1,2]','2023-11-06 16:10:49','2023-11-06 16:10:49'),(0,1,'companyData','add','self','self','2023-11-06 16:10:49','2023-11-06 16:10:49'),(0,1,'companyData','delete','all','[1,2,3]','2023-11-06 16:10:49','2023-11-06 16:10:49'),(0,1,'companyData','edit','all','self','2023-11-06 16:10:49','2023-11-06 16:10:49'),(0,1,'companyData','inquire','self','self','2023-11-06 16:10:49','2023-11-06 16:10:49'),(0,1,'employeeManage','add','self','self','2023-11-06 16:10:49','2023-11-06 16:10:49'),(0,1,'employeeManage','delete','all','[1,2,3]','2023-11-06 16:10:49','2023-11-06 16:10:49'),(0,1,'employeeManage','edit','all','self','2023-11-06 16:10:49','2023-11-06 16:10:49'),(0,1,'employeeManage','inquire','self','self','2023-11-06 16:10:49','2023-11-06 16:10:49'),(0,1,'selfData','delete','all','all','2023-11-06 16:10:49','2023-11-06 16:10:49'),(0,1,'selfData','edit','all','[1,2,3]','2023-11-06 16:10:49','2023-11-06 16:10:49'),(0,2,'banchManage','delete','All','[1,2,3]','2023-11-04 13:42:08','2023-11-04 13:42:08'),(0,2,'banchManage','edit','All','[1,2,3]','2023-11-04 13:42:08','2023-11-04 13:42:08'),(0,2,'selfData','delete','All','All','2023-11-04 13:42:08','2023-11-04 13:42:08'),(0,2,'selfData','edit','All','[1,2,3]','2023-11-04 13:42:08','2023-11-04 13:42:08');
/*!40000 ALTER TABLE `role_struct` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user` (
  `companyId` varchar(50) NOT NULL COMMENT '公司id',
  `userId` int NOT NULL COMMENT '使用者id',
  `roleId` int DEFAULT '2' COMMENT '使用者套用的角色id',
  `banchId` int DEFAULT '-1' COMMENT '部門id',
  `userName` varchar(20) DEFAULT NULL COMMENT '使用者名稱',
  `employeeNumber` varchar(30) DEFAULT NULL COMMENT '使用者員工編號',
  `account` varchar(50) NOT NULL COMMENT '使用者帳號',
  `password` varchar(50) NOT NULL COMMENT '使用者密碼',
  `onWorkDay` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '開始工作日',
  `deleteFlag` char(1) DEFAULT 'N' COMMENT '刪除旗標 ( N, Y )',
  `deleteTime` datetime DEFAULT NULL COMMENT '刪除時間',
  `createTime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '創建時間',
  `lastModify` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '最後更新時間',
  PRIMARY KEY (`companyId`,`userId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='使用者資料';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES ('0',1,1,1,'SIOU','0906930873','za96346','aa20010722','2023-11-04 16:31:18','N',NULL,'2023-11-04 16:31:18','2023-11-04 16:31:18'),('0',2,2,1,'a00001','a00001','a00001','a00001','2023-11-04 16:31:56','N',NULL,'2023-11-04 16:31:56','2023-11-04 16:31:56'),('0',3,1,1,'黃明修00','za999','za999','aa20010722','2006-01-02 23:04:05','Y','2023-11-06 14:46:30','2023-11-04 16:31:56','2023-11-06 14:46:30');
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-11-06 16:27:37
