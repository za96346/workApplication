-- MySQL dump 10.13  Distrib 8.0.29, for macos12 (x86_64)
--
-- Host: 127.0.0.1    Database: new_workapp
-- ------------------------------------------------------
-- Server version	8.2.0

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
INSERT INTO `company` VALUES (0,'dddd','test_one','wjiefoijwi','0906930873',1,'2023-11-06 15:11:34','2023-11-25 15:12:52');
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
INSERT INTO `company_banch` VALUES (0,1,'資訊組','N',NULL,'2023-11-04 17:14:03','2023-11-13 17:58:13'),(0,2,'輔導組','N',NULL,'2023-11-04 17:14:03','2023-11-25 15:02:47'),(0,3,'研發部部','N',NULL,'2023-11-06 16:14:49','2023-11-25 13:48:53'),(0,4,'公關組','N',NULL,'2023-11-13 15:17:29','2023-11-25 15:03:31'),(0,5,'jjj','Y','2023-11-16 18:47:15','2023-11-13 16:09:41','2023-11-16 18:47:15'),(0,6,'生輔不','N',NULL,'2023-11-13 18:13:02','2023-11-13 18:13:18'),(0,7,'人資部','N',NULL,'2023-11-13 18:13:53','2023-11-16 18:47:25'),(0,8,'會計部','N',NULL,'2023-11-15 16:04:54','2023-11-15 16:04:54'),(0,9,'保育組','N',NULL,'2023-11-22 17:30:31','2023-11-22 17:30:31'),(0,10,'efwfe','N',NULL,'2023-11-24 20:31:22','2023-11-24 20:31:22');
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
-- Table structure for table `log`
--

DROP TABLE IF EXISTS `log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `log` (
  `logId` int NOT NULL,
  `userId` int DEFAULT '-1',
  `companyId` int DEFAULT '-1',
  `routes` varchar(100) DEFAULT NULL,
  `ip` varchar(100) DEFAULT NULL,
  `params` text,
  `msg` text,
  `createTime` datetime DEFAULT CURRENT_TIMESTAMP,
  `lastModify` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`logId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='Log紀錄';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `log`
--

LOCK TABLES `log` WRITE;
/*!40000 ALTER TABLE `log` DISABLE KEYS */;
INSERT INTO `log` VALUES (1,1,0,'PUT : /workApp/performance/','::1','','Request Data 格式不正確 Key: \'Performance.UserId\' Error:Field validation for \'UserId\' failed on the \'required\' tag\nKey: \'Performance.Month\' Error:Field validation for \'Month\' failed on the \'required\' tag\nKey: \'Performance.Goal\' Error:Field validation for \'Goal\' failed on the \'required\' tag','2023-11-24 14:15:48','2023-11-24 14:15:48'),(2,1,0,'[PUT]/workApp/performance/','::1','','Request Data 格式不正確 Key: \'Performance.UserId\' Error:Field validation for \'UserId\' failed on the \'required\' tag\nKey: \'Performance.Month\' Error:Field validation for \'Month\' failed on the \'required\' tag\nKey: \'Performance.Goal\' Error:Field validation for \'Goal\' failed on the \'required\' tag','2023-11-24 14:31:27','2023-11-24 14:31:27'),(3,1,0,'[PUT]/workApp/performance/','::1','{}','Request Data 格式不正確 Key: \'Performance.UserId\' Error:Field validation for \'UserId\' failed on the \'required\' tag\nKey: \'Performance.Month\' Error:Field validation for \'Month\' failed on the \'required\' tag\nKey: \'Performance.Goal\' Error:Field validation for \'Goal\' failed on the \'required\' tag','2023-11-24 14:35:44','2023-11-24 14:35:44'),(4,1,0,'[PUT]/workApp/user/','::1',NULL,'新增失敗-帳號重複','2023-11-24 15:48:07','2023-11-24 15:48:07'),(5,1,0,'[POST]/workApp/user/password','::1',NULL,'無法插入此部門，尚無權限','2023-11-24 15:54:30','2023-11-24 15:54:30'),(6,1,0,'[POST]/workApp/user/password','::1',NULL,'更新失敗，password 與new password 不相等','2023-11-24 15:56:38','2023-11-24 15:56:38'),(7,1,0,'[POST]/workApp/user/password','::1',NULL,'更新失敗，舊密碼不同','2023-11-24 15:56:56','2023-11-24 15:56:56'),(8,1,0,'[POST]/workApp/user/password','::1',NULL,'更新失敗，舊密碼不同','2023-11-24 16:23:50','2023-11-24 16:23:50'),(9,1,0,'[POST]/workApp/user/password','::1',NULL,'更新失敗，password 與new password 不相等','2023-11-24 16:24:11','2023-11-24 16:24:11'),(10,1,0,'[POST]/workApp/entry/login','::1',NULL,'帳號或密碼錯誤','2023-11-24 16:24:27','2023-11-24 16:24:27'),(11,1,0,'[POST]/workApp/user/password','::1',NULL,'更新失敗，password 與new password 不相等','2023-11-24 16:49:13','2023-11-24 16:49:13'),(12,1,0,'[POST]/workApp/user/password','::1',NULL,'更新失敗，舊密碼不同','2023-11-24 16:49:19','2023-11-24 16:49:19'),(13,1,0,'[POST]/workApp/user/password','::1',NULL,'更新失敗，舊密碼不同','2023-11-24 16:50:36','2023-11-24 16:50:36'),(14,1,0,'[POST]/workApp/entry/login','::1',NULL,'帳號或密碼錯誤','2023-11-24 19:58:47','2023-11-24 19:58:47'),(15,1,0,'[POST]/workApp/entry/login','::1',NULL,'帳號或密碼錯誤','2023-11-24 20:00:38','2023-11-24 20:00:38'),(16,11,0,'[POST]/workApp/user/','::1',NULL,'權限驗證失敗--[funcCode: \'employeeManage\'][itemCode: \'edit\']','2023-11-24 20:01:39','2023-11-24 20:01:39'),(17,11,0,'[POST]/workApp/user/','::1',NULL,'權限驗證失敗--[funcCode: \'employeeManage\'][itemCode: \'edit\']','2023-11-24 20:01:55','2023-11-24 20:01:55'),(18,11,0,'[POST]/workApp/user/','::1',NULL,'權限驗證失敗--[funcCode: \'employeeManage\'][itemCode: \'edit\']','2023-11-24 20:02:00','2023-11-24 20:02:00'),(19,11,0,'[POST]/workApp/user/','::1',NULL,'權限驗證失敗--[funcCode: \'employeeManage\'][itemCode: \'edit\']','2023-11-24 20:02:13','2023-11-24 20:02:13'),(20,1,0,'[POST]/workApp/user/my','::1',NULL,'Request Data 格式不正確 EOF','2023-11-24 20:13:27','2023-11-24 20:13:27'),(21,1,0,'[PUT]/workApp/banch/','::1',NULL,'部門名稱重複','2023-11-24 20:23:08','2023-11-24 20:23:08'),(22,11,0,'[GET]/workApp/banch/','::1',NULL,'權限驗證失敗--[funcCode: \'banchManage\'][itemCode: \'inquire\']','2023-11-24 20:31:11','2023-11-24 20:31:11'),(23,11,0,'[GET]/workApp/banch/','::1',NULL,'權限驗證失敗--[funcCode: \'banchManage\'][itemCode: \'inquire\']','2023-11-24 20:31:14','2023-11-24 20:31:14'),(24,11,0,'[GET]/workApp/banch/','::1',NULL,'權限驗證失敗--[funcCode: \'banchManage\'][itemCode: \'inquire\']','2023-11-24 20:31:22','2023-11-24 20:31:22'),(25,11,0,'[GET]/workApp/banch/','::1',NULL,'權限驗證失敗--[funcCode: \'banchManage\'][itemCode: \'inquire\']','2023-11-24 20:31:38','2023-11-24 20:31:38'),(26,11,0,'[DELETE]/workApp/banch/','::1',NULL,'無法插入此部門，尚無權限','2023-11-24 20:35:22','2023-11-24 20:35:22'),(27,11,0,'[DELETE]/workApp/banch/','::1',NULL,'無法插入此部門，尚無權限','2023-11-24 20:35:50','2023-11-24 20:35:50'),(28,11,0,'[DELETE]/workApp/banch/','::1',NULL,'無法插入此部門，尚無權限','2023-11-24 20:36:48','2023-11-24 20:36:48'),(29,11,0,'[GET]/workApp/banch/','::1',NULL,'權限驗證失敗--[funcCode: \'banchManage\'][itemCode: \'inquire\']','2023-11-25 13:48:25','2023-11-25 13:48:25'),(30,11,0,'[GET]/workApp/banch/','::1',NULL,'權限驗證失敗--[funcCode: \'banchManage\'][itemCode: \'inquire\']','2023-11-25 13:48:28','2023-11-25 13:48:28'),(31,11,0,'[GET]/workApp/banch/','::1',NULL,'權限驗證失敗--[funcCode: \'banchManage\'][itemCode: \'inquire\']','2023-11-25 13:48:28','2023-11-25 13:48:28'),(32,11,0,'[GET]/workApp/banch/','::1',NULL,'權限驗證失敗--[funcCode: \'banchManage\'][itemCode: \'inquire\']','2023-11-25 13:48:44','2023-11-25 13:48:44'),(33,11,0,'[GET]/workApp/banch/','::1',NULL,'權限驗證失敗--[funcCode: \'banchManage\'][itemCode: \'inquire\']','2023-11-25 13:48:53','2023-11-25 13:48:53'),(34,11,0,'[GET]/workApp/banch/','::1',NULL,'權限驗證失敗--[funcCode: \'banchManage\'][itemCode: \'inquire\']','2023-11-25 13:49:16','2023-11-25 13:49:16'),(35,11,0,'[POST]/workApp/entry/login','::1',NULL,'帳號或密碼錯誤','2023-11-25 13:53:41','2023-11-25 13:53:41'),(36,11,0,'[POST]/workApp/banch/','::1',NULL,'無法插入此部門，尚無權限','2023-11-25 13:56:00','2023-11-25 13:56:00'),(37,11,0,'[POST]/workApp/company/','::1',NULL,'權限驗證失敗--[funcCode: \'companyData\'][itemCode: \'edit\']','2023-11-25 15:05:57','2023-11-25 15:05:57'),(38,11,0,'[GET]/workApp/company/','::1',NULL,'權限驗證失敗--[funcCode: \'companyData\'][itemCode: \'inquire\']','2023-11-25 15:13:20','2023-11-25 15:13:20'),(39,11,0,'[GET]/workApp/company/','::1',NULL,'權限驗證失敗--[funcCode: \'companyData\'][itemCode: \'inquire\']','2023-11-25 15:14:50','2023-11-25 15:14:50'),(40,1,0,'[PUT]/workApp/performance/','::1',NULL,'Request Data 格式不正確 Key: \'Performance.Goal\' Error:Field validation for \'Goal\' failed on the \'required\' tag','2023-11-25 15:28:51','2023-11-25 15:28:51'),(41,11,0,'[PUT]/workApp/performance/','::1',NULL,'Request Data 格式不正確 Key: \'Performance.Goal\' Error:Field validation for \'Goal\' failed on the \'required\' tag','2023-11-25 15:33:17','2023-11-25 15:33:17'),(42,11,0,'[PUT]/workApp/performance/','::1',NULL,'無法插入此部門，尚無權限','2023-11-25 16:00:21','2023-11-25 16:00:21'),(43,11,0,'[PUT]/workApp/performance/','::1',NULL,'新增失敗-檢查到重複資料','2023-11-25 16:09:36','2023-11-25 16:09:36'),(44,11,0,'[PUT]/workApp/performance/copy','::1',NULL,'新增失敗-檢查到重複資料','2023-11-25 16:33:30','2023-11-25 16:33:30');
/*!40000 ALTER TABLE `log` ENABLE KEYS */;
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
INSERT INTO `operation_item` VALUES ('add','新增','2023-11-02 14:56:43','2023-11-02 14:56:43'),('changeBanch','更換部門','2023-11-02 14:57:21','2023-11-25 16:44:36'),('copy','複製','2023-11-02 14:56:43','2023-11-02 14:56:43'),('delete','刪除','2023-11-02 14:56:43','2023-11-02 14:56:43'),('edit','編輯','2023-11-02 14:56:43','2023-11-02 14:56:43'),('inquire','查詢','2023-11-02 14:56:43','2023-11-02 14:56:43'),('print','列印','2023-11-02 14:57:21','2023-11-02 14:57:21');
/*!40000 ALTER TABLE `operation_item` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `performance`
--

DROP TABLE IF EXISTS `performance`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `performance` (
  `companyId` int NOT NULL COMMENT '公司id',
  `userId` int NOT NULL COMMENT '使用者id',
  `performanceId` int NOT NULL COMMENT '績效id',
  `year` int NOT NULL COMMENT '年分',
  `month` int NOT NULL COMMENT '月份',
  `banchId` int DEFAULT NULL COMMENT '部門id',
  `goal` varchar(1000) DEFAULT NULL COMMENT '績效目標',
  `attitude` int DEFAULT '0' COMMENT '態度分數',
  `efficiency` int DEFAULT '0' COMMENT '效率分數',
  `professional` int DEFAULT '0' COMMENT '專業分數',
  `directions` varchar(1000) DEFAULT NULL,
  `beLate` int DEFAULT '0' COMMENT '遲到',
  `dayOffNotOnRule` int DEFAULT '0' COMMENT '未依規定請假',
  `deleteFlag` char(1) DEFAULT 'N' COMMENT '刪除旗標 ( N, Y )',
  `deleteTime` datetime DEFAULT NULL COMMENT '刪除時間',
  `createTime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '創建時間',
  `lastModify` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '最後更新時間',
  PRIMARY KEY (`companyId`,`userId`,`performanceId`,`year`,`month`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='績效評核';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `performance`
--

LOCK TABLES `performance` WRITE;
/*!40000 ALTER TABLE `performance` DISABLE KEYS */;
INSERT INTO `performance` VALUES (0,1,1,111,1,1,NULL,0,0,0,NULL,0,0,'Y','2023-11-23 15:42:59','2023-11-23 12:31:15','2023-11-23 15:42:59'),(0,1,8,112,3,2,'efw',0,0,0,'',0,0,'N',NULL,'2023-11-25 15:30:41','2023-11-25 15:34:01'),(0,2,2,112,3,9,'kk',4,4,0,' 績效ewijfwiofj',0,4,'Y','2023-11-23 14:21:46','2023-11-23 14:02:34','2023-11-23 14:21:46'),(0,2,3,112,2,9,'jjj',11,5,0,'',0,0,'N',NULL,'2023-11-23 14:21:59','2023-11-24 12:15:36'),(0,2,4,112,5,9,'1、落實外賓、車輛、貨物管控。 2、出車零事故、提醒車輛保養定檢。 3、交付事項確時完成並積極幫助家園事務。',4,0,0,'願意配合家園事務，需建立自信心以應對夥伴的引導。',0,0,'N',NULL,'2023-11-23 16:29:35','2023-11-24 12:15:47'),(0,8,6,112,3,7,'sq2s',0,0,0,'',0,0,'N',NULL,'2023-11-24 11:48:22','2023-11-24 11:48:22'),(0,9,5,112,5,9,'dd',0,3,0,'',2,2,'N',NULL,'2023-11-24 11:48:01','2023-11-24 11:48:01'),(0,9,9,111,8,9,'goal',100,100,100,'',0,0,'N',NULL,'2023-11-25 15:33:25','2023-11-25 15:33:25'),(0,9,10,111,9,9,'sd',0,0,0,'',0,0,'N',NULL,'2023-11-25 16:09:43','2023-11-25 16:09:43'),(0,11,7,112,4,4,'wefwqf',0,0,0,'',0,0,'N',NULL,'2023-11-25 15:28:54','2023-11-25 15:28:54'),(0,11,11,112,5,4,'wefwqf',0,0,0,'',0,0,'N',NULL,'2023-11-25 16:33:38','2023-11-25 16:33:38');
/*!40000 ALTER TABLE `performance` ENABLE KEYS */;
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
INSERT INTO `role` VALUES (0,1,'最高管理員','N','N',NULL,'2023-11-02 19:53:02','2023-11-23 15:49:16'),(0,2,'test','N','N',NULL,'2023-11-15 13:39:55','2023-11-25 16:45:55'),(0,3,'保育組 [ 一般職員 ]','N','N',NULL,'2023-11-15 15:27:17','2023-11-22 17:55:18'),(0,4,'測試者','N','Y','2023-11-16 14:48:41','2023-11-16 14:46:03','2023-11-16 14:48:41'),(0,5,'保育組 [ 組長 ]','N','N',NULL,'2023-11-22 17:27:28','2023-11-22 18:57:19'),(0,6,'人資部 [ 一般職員 ]','N','N',NULL,'2023-11-22 17:56:24','2023-11-22 17:56:24'),(0,7,'人資部 [ 組長 ]','N','N',NULL,'2023-11-22 18:03:04','2023-11-22 18:03:04');
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
INSERT INTO `role_struct` VALUES (0,0,'banchManage','add','all','[1,2,3]','2023-11-15 12:52:57','2023-11-15 12:52:57'),(0,0,'banchManage','delete','all','all','2023-11-15 12:52:57','2023-11-15 12:52:57'),(0,0,'banchManage','edit','all','[1,2,3]','2023-11-15 12:52:57','2023-11-15 12:52:57'),(0,0,'banchManage','inquire','all','[1,2,3]','2023-11-15 12:52:57','2023-11-15 12:52:57'),(0,0,'companyData','add','all','[1,2,3]','2023-11-15 12:52:57','2023-11-15 12:52:57'),(0,0,'companyData','delete','all','all','2023-11-15 12:52:57','2023-11-15 12:52:57'),(0,0,'companyData','edit','all','[1,2,3]','2023-11-15 12:52:57','2023-11-15 12:52:57'),(0,0,'companyData','inquire','all','[1,2,3]','2023-11-15 12:52:57','2023-11-15 12:52:57'),(0,0,'employeeManage','add','all','[1,2,3]','2023-11-15 12:52:57','2023-11-15 12:52:57'),(0,0,'employeeManage','delete','all','all','2023-11-15 12:52:57','2023-11-15 12:52:57'),(0,0,'employeeManage','edit','all','[1,2,3]','2023-11-15 12:52:57','2023-11-15 12:52:57'),(0,0,'employeeManage','inquire','all','[1,2,3]','2023-11-15 12:52:57','2023-11-15 12:52:57'),(0,0,'performance','add','all','[1,2,3]','2023-11-15 12:52:57','2023-11-15 12:52:57'),(0,0,'performance','delete','all','all','2023-11-15 12:52:57','2023-11-15 12:52:57'),(0,0,'performance','edit','all','[1,2,3]','2023-11-15 12:52:57','2023-11-15 12:52:57'),(0,0,'performance','inquire','all','[1,2,3]','2023-11-15 12:52:57','2023-11-15 12:52:57'),(0,0,'roleManage','add','all','[1,2,3]','2023-11-15 12:52:57','2023-11-15 12:52:57'),(0,0,'roleManage','delete','all','all','2023-11-15 12:52:57','2023-11-15 12:52:57'),(0,0,'roleManage','edit','all','[1,2,3]','2023-11-15 12:52:57','2023-11-15 12:52:57'),(0,0,'roleManage','inquire','all','[1,2,3]','2023-11-15 12:52:57','2023-11-15 12:52:57'),(0,0,'selfData','add','all','[1,2,3]','2023-11-15 12:52:57','2023-11-15 12:52:57'),(0,0,'selfData','delete','all','all','2023-11-15 12:52:57','2023-11-15 12:52:57'),(0,0,'selfData','edit','all','[1,2,3]','2023-11-15 12:52:57','2023-11-15 12:52:57'),(0,0,'selfData','inquire','all','[1,2,3]','2023-11-15 12:52:57','2023-11-15 12:52:57'),(0,0,'shift','add','all','[1,2,3]','2023-11-15 12:52:57','2023-11-15 12:52:57'),(0,0,'shift','delete','all','all','2023-11-15 12:52:57','2023-11-15 12:52:57'),(0,0,'shift','edit','all','[1,2,3]','2023-11-15 12:52:57','2023-11-15 12:52:57'),(0,0,'shift','inquire','all','[1,2,3]','2023-11-15 12:52:57','2023-11-15 12:52:57'),(0,0,'shiftSetting','add','all','[1,2,3]','2023-11-15 12:52:57','2023-11-15 12:52:57'),(0,0,'shiftSetting','delete','all','all','2023-11-15 12:52:57','2023-11-15 12:52:57'),(0,0,'shiftSetting','edit','all','[1,2,3]','2023-11-15 12:52:57','2023-11-15 12:52:57'),(0,0,'shiftSetting','inquire','all','[1,2,3]','2023-11-15 12:52:57','2023-11-15 12:52:57'),(0,0,'yearPerformance','add','all','[1,2,3]','2023-11-15 12:52:57','2023-11-15 12:52:57'),(0,0,'yearPerformance','delete','all','all','2023-11-15 12:52:57','2023-11-15 12:52:57'),(0,0,'yearPerformance','edit','all','[1,2,3]','2023-11-15 12:52:57','2023-11-15 12:52:57'),(0,0,'yearPerformance','inquire','all','[1,2,3]','2023-11-15 12:52:57','2023-11-15 12:52:57'),(0,1,'banchManage','add','all','all','2023-11-23 15:49:16','2023-11-23 15:49:16'),(0,1,'banchManage','delete','all','all','2023-11-23 15:49:16','2023-11-23 15:49:16'),(0,1,'banchManage','edit','all','all','2023-11-23 15:49:16','2023-11-23 15:49:16'),(0,1,'banchManage','inquire','all','all','2023-11-23 15:49:16','2023-11-23 15:49:16'),(0,1,'companyData','add','all','all','2023-11-23 15:49:16','2023-11-23 15:49:16'),(0,1,'companyData','delete','all','all','2023-11-23 15:49:16','2023-11-23 15:49:16'),(0,1,'companyData','edit','all','all','2023-11-23 15:49:16','2023-11-23 15:49:16'),(0,1,'companyData','inquire','all','all','2023-11-23 15:49:16','2023-11-23 15:49:16'),(0,1,'employeeManage','add','all','all','2023-11-23 15:49:16','2023-11-23 15:49:16'),(0,1,'employeeManage','delete','all','all','2023-11-23 15:49:16','2023-11-23 15:49:16'),(0,1,'employeeManage','edit','all','all','2023-11-23 15:49:16','2023-11-23 15:49:16'),(0,1,'employeeManage','inquire','all','all','2023-11-23 15:49:16','2023-11-23 15:49:16'),(0,1,'performance','add','all','all','2023-11-23 15:49:16','2023-11-23 15:49:16'),(0,1,'performance','delete','all','all','2023-11-23 15:49:16','2023-11-23 15:49:16'),(0,1,'performance','edit','all','all','2023-11-23 15:49:16','2023-11-23 15:49:16'),(0,1,'performance','inquire','all','all','2023-11-23 15:49:16','2023-11-23 15:49:16'),(0,1,'performance','print','all','all','2023-11-23 15:49:16','2023-11-23 15:49:16'),(0,1,'roleManage','add','all','all','2023-11-23 15:49:16','2023-11-23 15:49:16'),(0,1,'roleManage','delete','all','all','2023-11-23 15:49:16','2023-11-23 15:49:16'),(0,1,'roleManage','edit','all','all','2023-11-23 15:49:16','2023-11-23 15:49:16'),(0,1,'roleManage','inquire','all','all','2023-11-23 15:49:16','2023-11-23 15:49:16'),(0,1,'selfData','add','all','all','2023-11-23 15:49:16','2023-11-23 15:49:16'),(0,1,'selfData','delete','all','all','2023-11-23 15:49:16','2023-11-23 15:49:16'),(0,1,'selfData','edit','all','all','2023-11-23 15:49:16','2023-11-23 15:49:16'),(0,1,'selfData','inquire','all','all','2023-11-23 15:49:16','2023-11-23 15:49:16'),(0,1,'shift','add','all','all','2023-11-23 15:49:16','2023-11-23 15:49:16'),(0,1,'shift','delete','all','all','2023-11-23 15:49:16','2023-11-23 15:49:16'),(0,1,'shift','edit','all','all','2023-11-23 15:49:16','2023-11-23 15:49:16'),(0,1,'shift','inquire','all','all','2023-11-23 15:49:16','2023-11-23 15:49:16'),(0,1,'shiftSetting','add','all','all','2023-11-23 15:49:16','2023-11-23 15:49:16'),(0,1,'shiftSetting','delete','all','all','2023-11-23 15:49:16','2023-11-23 15:49:16'),(0,1,'shiftSetting','edit','all','all','2023-11-23 15:49:16','2023-11-23 15:49:16'),(0,1,'shiftSetting','inquire','all','all','2023-11-23 15:49:16','2023-11-23 15:49:16'),(0,1,'yearPerformance','add','all','all','2023-11-23 15:49:16','2023-11-23 15:49:16'),(0,1,'yearPerformance','delete','all','all','2023-11-23 15:49:16','2023-11-23 15:49:16'),(0,1,'yearPerformance','edit','all','all','2023-11-23 15:49:16','2023-11-23 15:49:16'),(0,1,'yearPerformance','inquire','all','all','2023-11-23 15:49:16','2023-11-23 15:49:16'),(0,2,'banchManage','add','all','all','2023-11-25 16:45:55','2023-11-25 16:45:55'),(0,2,'banchManage','delete','self','all','2023-11-25 16:45:55','2023-11-25 16:45:55'),(0,2,'banchManage','edit','[2]','all','2023-11-25 16:45:55','2023-11-25 16:45:55'),(0,2,'banchManage','inquire','all','all','2023-11-25 16:45:55','2023-11-25 16:45:55'),(0,2,'companyData','edit','all','all','2023-11-25 16:45:55','2023-11-25 16:45:55'),(0,2,'companyData','inquire','all','all','2023-11-25 16:45:55','2023-11-25 16:45:55'),(0,2,'employeeManage','add','[1,4]','[1,2,3,4]','2023-11-25 16:45:55','2023-11-25 16:45:55'),(0,2,'employeeManage','delete','self','[6]','2023-11-25 16:45:55','2023-11-25 16:45:55'),(0,2,'employeeManage','edit','self','self','2023-11-25 16:45:55','2023-11-25 16:45:55'),(0,2,'employeeManage','inquire','[1]','all','2023-11-25 16:45:55','2023-11-25 16:45:55'),(0,2,'performance','add','[9]','all','2023-11-25 16:45:55','2023-11-25 16:45:55'),(0,2,'performance','changeBanch','[2]','all','2023-11-25 16:45:55','2023-11-25 16:45:55'),(0,2,'performance','copy','self','all','2023-11-25 16:45:55','2023-11-25 16:45:55'),(0,2,'performance','delete','self','all','2023-11-25 16:45:55','2023-11-25 16:45:55'),(0,2,'performance','edit','[9]','all','2023-11-25 16:45:55','2023-11-25 16:45:55'),(0,2,'performance','inquire','all','all','2023-11-25 16:45:55','2023-11-25 16:45:55'),(0,2,'selfData','edit','all','all','2023-11-25 16:45:55','2023-11-25 16:45:55'),(0,2,'selfData','inquire','all','all','2023-11-25 16:45:55','2023-11-25 16:45:55'),(0,2,'yearPerformance','inquire','[1]','self','2023-11-25 16:45:55','2023-11-25 16:45:55'),(0,3,'performance','inquire','all','[1]','2023-11-22 17:55:18','2023-11-22 17:55:18'),(0,4,'banchManage','copy','all','all','2023-11-16 14:46:03','2023-11-16 14:46:03'),(0,5,'employeeManage','add','all','all','2023-11-22 18:57:19','2023-11-22 18:57:19'),(0,5,'employeeManage','copy','all','all','2023-11-22 18:57:19','2023-11-22 18:57:19'),(0,5,'employeeManage','delete','all','all','2023-11-22 18:57:19','2023-11-22 18:57:19'),(0,5,'employeeManage','edit','all','all','2023-11-22 18:57:19','2023-11-22 18:57:19'),(0,5,'employeeManage','inquire','[7,9]','[3,6]','2023-11-22 18:57:19','2023-11-22 18:57:19'),(0,5,'employeeManage','print','all','all','2023-11-22 18:57:19','2023-11-22 18:57:19'),(0,5,'selfData','add','all','all','2023-11-22 18:57:19','2023-11-22 18:57:19'),(0,5,'selfData','copy','all','all','2023-11-22 18:57:19','2023-11-22 18:57:19'),(0,5,'selfData','delete','all','all','2023-11-22 18:57:19','2023-11-22 18:57:19'),(0,5,'selfData','edit','all','all','2023-11-22 18:57:19','2023-11-22 18:57:19'),(0,5,'selfData','inquire','all','all','2023-11-22 18:57:19','2023-11-22 18:57:19'),(0,5,'selfData','print','all','all','2023-11-22 18:57:19','2023-11-22 18:57:19');
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
INSERT INTO `user` VALUES ('0',1,1,2,'SIOU','0906930873','za96346','123456','2023-11-04 16:31:18','N',NULL,'2023-11-04 16:31:18','2023-11-24 20:14:02'),('0',2,5,9,'張小姐','a00001','a00001','a00001','2023-11-04 16:31:56','N',NULL,'2023-11-04 16:31:56','2023-11-24 19:23:34'),('0',3,1,1,'黃明修00','za999','za999','aa20010722','2006-01-02 23:04:05','Y','2023-11-06 14:46:30','2023-11-04 16:31:56','2023-11-06 14:46:30'),('0',4,2,4,'json','b00002','json','123456','2023-11-16 12:07:44','Y','2023-11-16 14:25:49','2023-11-16 12:08:26','2023-11-16 14:25:49'),('0',5,6,7,'黃小姐','b00002','b00002','b00002','2023-11-22 17:55:34','N',NULL,'2023-11-22 17:55:59','2023-11-24 19:23:11'),('0',6,6,7,'陳小姐','b00003','b00003','b00003','2023-11-22 17:58:58','N',NULL,'2023-11-22 17:59:19','2023-11-22 17:59:19'),('0',7,3,9,'阮小姐','a00002','a00002','a00002','2023-11-22 18:00:16','N',NULL,'2023-11-22 18:00:40','2023-11-22 18:00:40'),('0',8,7,7,'賴小姐','b00001','b00001','b00001','2023-11-22 18:02:18','N',NULL,'2023-11-22 18:02:40','2023-11-22 18:03:26'),('0',9,6,9,'蕭小姐','c00001','c00001','c00001','2023-11-22 18:06:00','N',NULL,'2023-11-22 18:06:31','2023-11-22 18:06:45'),('0',10,6,4,'d00001','d00001','d00001','d00001','2023-11-24 15:47:43','N',NULL,'2023-11-24 15:48:12','2023-11-24 15:48:12'),('0',11,2,4,'testgg','testef','test','123456','2023-11-24 19:59:54','N',NULL,'2023-11-24 20:00:11','2023-11-25 15:20:59'),('0',12,2,4,'test-2','test-2','test-2','test-2','2023-10-31 15:22:41','N',NULL,'2023-11-25 15:23:21','2023-11-25 15:23:21'),('0',13,2,1,'test-200','test-200','test-200','test-200','2023-11-25 15:24:17','N',NULL,'2023-11-25 15:24:35','2023-11-25 15:24:35');
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

-- Dump completed on 2023-11-27 15:38:59
