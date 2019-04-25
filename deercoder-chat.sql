-- MySQL dump 10.13  Distrib 5.7.25, for Linux (x86_64)
--
-- Host: localhost    Database: deercoder-chat
-- ------------------------------------------------------
-- Server version	5.7.25-0ubuntu0.18.04.2

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
-- Table structure for table `group_last_msg`
--

DROP TABLE IF EXISTS `group_last_msg`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `group_last_msg` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `group_id` varchar(50) DEFAULT NULL COMMENT '群聊id,获得其中成员',
  `uid` int(11) DEFAULT NULL COMMENT '用户id,下次接收离线消息用',
  `last_group_msg_uuid` varchar(50) DEFAULT NULL COMMENT '记录的最后消息id',
  `is_read` tinyint(1) DEFAULT '0' COMMENT '0未读,1已读',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `group_id` (`last_group_msg_uuid`,`group_id`,`uid`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `group_last_msg`
--

LOCK TABLES `group_last_msg` WRITE;
/*!40000 ALTER TABLE `group_last_msg` DISABLE KEYS */;
/*!40000 ALTER TABLE `group_last_msg` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `group_msg`
--

DROP TABLE IF EXISTS `group_msg`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `group_msg` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `uuid` varchar(50) DEFAULT NULL COMMENT '唯一标识',
  `group_id` varchar(50) DEFAULT '' COMMENT '群聊id,获得其中成员',
  `content` blob COMMENT '消息内容,可存储表情',
  `from_uid` int(11) DEFAULT NULL COMMENT '由谁发送',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `content_type` varchar(10) DEFAULT 'text' COMMENT '内容类,默认文本类型',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uuid` (`uuid`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `group_msg`
--

LOCK TABLES `group_msg` WRITE;
/*!40000 ALTER TABLE `group_msg` DISABLE KEYS */;
/*!40000 ALTER TABLE `group_msg` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `group_users`
--

DROP TABLE IF EXISTS `group_users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `group_users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `group_id` varchar(50) DEFAULT '' COMMENT '好友列表(包括群组)id',
  `uid` int(11) DEFAULT NULL COMMENT '用户id',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `group_id` (`group_id`,`uid`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=100000002 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `group_users`
--

LOCK TABLES `group_users` WRITE;
/*!40000 ALTER TABLE `group_users` DISABLE KEYS */;
INSERT INTO `group_users` VALUES (100000000,'93f65451-efc4-11e8-918b-34e6d7558043',1,NULL),(100000001,'93f65451-efc4-11e8-918b-34e6d7558043',2,NULL);
/*!40000 ALTER TABLE `group_users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(30) DEFAULT '' COMMENT '用户名',
  `headimg` varchar(255) DEFAULT '' COMMENT '头像',
  `password` varchar(50) DEFAULT '' COMMENT '密码',
  `createtime` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (1,'admin','xxx.img','2c39f9867429',NULL),(2,'test','xxx.img','2c39f9867429','2018-11-24 19:15:54');
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping routines for database 'deercoder-chat'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2019-04-24 14:18:17
