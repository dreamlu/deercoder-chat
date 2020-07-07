-- MySQL dump 10.13  Distrib 8.0.16, for Linux (x86_64)
--
-- Host: localhost    Database: deercoder-chat
-- ------------------------------------------------------
-- Server version	8.0.16

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
 SET NAMES utf8mb4 ;
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
 SET character_set_client = utf8mb4 ;
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
 SET character_set_client = utf8mb4 ;
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
) ENGINE=InnoDB AUTO_INCREMENT=74 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `group_msg`
--

LOCK TABLES `group_msg` WRITE;
/*!40000 ALTER TABLE `group_msg` DISABLE KEYS */;
INSERT INTO `group_msg` VALUES (1,'asfddgdsfgshfdsffgdf','93f65451-efc4-11e8-918b-34e6d7558045','',1,'2019-04-27 18:00:55','text'),(2,'0beeab13-6a6b-11e9-940b-e86a6477cf1b','93f65451-efc4-11e8-918b-34e6d7558045',_binary '阿斯蒂芬',1,'2019-04-29 18:39:20','text'),(3,'492daf2e-6a6b-11e9-940b-e86a6477cf1b','93f65451-efc4-11e8-918b-34e6d7558045',_binary '测测撒旦法规的规范',1,'2019-04-29 18:41:03','text'),(4,'5ab6a484-6a6b-11e9-940b-e86a6477cf1b','93f65451-efc4-11e8-918b-34e6d7558044',_binary '的方式发给',3,'2019-04-29 18:41:32','text'),(5,'5c99ad4b-6a6b-11e9-940b-e86a6477cf1b','93f65451-efc4-11e8-918b-34e6d7558044',_binary '第三方公司',3,'2019-04-29 18:41:35','text'),(6,'5e841e4c-6a6b-11e9-940b-e86a6477cf1b','93f65451-efc4-11e8-918b-34e6d7558044',_binary '的好方法付付付付付付付付付付付付付付',3,'2019-04-29 18:41:39','text'),(7,'d799a394-6a6c-11e9-b006-e86a6477cf1b','93f65451-efc4-11e8-918b-34e6d7558044',_binary '阿斯蒂芬',1,'2019-04-29 18:52:11','text'),(8,'25097d39-6a72-11e9-b006-e86a6477cf1b','93f65451-efc4-11e8-918b-34e6d7558044',_binary '是打发',1,'2019-04-29 19:30:09','text'),(9,'3f563ace-6a73-11e9-b006-e86a6477cf1b','93f65451-efc4-11e8-918b-34e6d7558045',_binary '撒旦法师打',1,'2019-04-29 19:38:02','text'),(10,'40dc2078-6a73-11e9-b006-e86a6477cf1b','93f65451-efc4-11e8-918b-34e6d7558045',_binary '撒旦法师打撒旦法',1,'2019-04-29 19:38:05','text'),(11,'b05cf9be-6a73-11e9-b006-e86a6477cf1b','93f65451-efc4-11e8-918b-34e6d7558044',_binary '阿斯蒂芬',1,'2019-04-29 19:41:12','text'),(12,'b1bfd216-6a73-11e9-b006-e86a6477cf1b','93f65451-efc4-11e8-918b-34e6d7558044',_binary '阿斯蒂芬',1,'2019-04-29 19:41:14','text'),(13,'b3fe8787-6a73-11e9-b006-e86a6477cf1b','93f65451-efc4-11e8-918b-34e6d7558045',_binary '阿斯蒂芬',1,'2019-04-29 19:41:18','text'),(14,'dae21df7-6a73-11e9-b006-e86a6477cf1b','93f65451-efc4-11e8-918b-34e6d7558045',_binary '阿斯蒂芬',1,'2019-04-29 19:42:23','text'),(15,'f21cba3a-6a73-11e9-b006-e86a6477cf1b','93f65451-efc4-11e8-918b-34e6d7558044',_binary '萨芬的',1,'2019-04-29 19:43:02','text'),(16,'f6d4e27b-6a73-11e9-b006-e86a6477cf1b','93f65451-efc4-11e8-918b-34e6d7558044',_binary '阿斯蒂芬阿斯顿发那地方',1,'2019-04-29 19:43:10','text'),(17,'722c7b94-6a75-11e9-b006-e86a6477cf1b','93f65451-efc4-11e8-918b-34e6d7558044',_binary '撒旦法',3,'2019-04-29 19:53:47','text'),(18,'7772e437-6a75-11e9-b006-e86a6477cf1b','93f65451-efc4-11e8-918b-34e6d7558044',_binary '阿斯蒂芬',1,'2019-04-29 19:53:55','text'),(19,'7a3dddff-6a75-11e9-b006-e86a6477cf1b','93f65451-efc4-11e8-918b-34e6d7558044',_binary '阿斯蒂芬',3,'2019-04-29 19:54:00','text'),(20,'829a5b39-6a75-11e9-b006-e86a6477cf1b','93f65451-efc4-11e8-918b-34e6d7558045',_binary '是打发',1,'2019-04-29 19:54:14','text'),(21,'873ba523-6a75-11e9-b006-e86a6477cf1b','93f65451-efc4-11e8-918b-34e6d7558043',_binary '阿斯顿发多少',1,'2019-04-29 19:54:22','text'),(22,'cff1eb83-6a75-11e9-b006-e86a6477cf1b','93f65451-efc4-11e8-918b-34e6d7558044',_binary '长整形',3,'2019-04-29 19:56:24','text'),(23,'d1bc8c54-6a75-11e9-b006-e86a6477cf1b','93f65451-efc4-11e8-918b-34e6d7558044',_binary '撒旦法',3,'2019-04-29 19:56:27','text'),(24,'d4ab4c5c-6a75-11e9-b006-e86a6477cf1b','93f65451-efc4-11e8-918b-34e6d7558044',_binary '阿斯蒂芬',3,'2019-04-29 19:56:32','text'),(25,'c74e476e-6a76-11e9-b006-e86a6477cf1b','93f65451-efc4-11e8-918b-34e6d7558044',_binary 'sadf',3,'2019-04-29 20:03:19','text'),(26,'9462cc9d-6a7d-11e9-a9b4-e86a6477cf1b','93f65451-efc4-11e8-918b-34e6d7558044',_binary 'sdf ',1,'2019-04-29 20:52:00','text'),(27,'9ee01329-6a7f-11e9-a9b4-e86a6477cf1b','93f65451-efc4-11e8-918b-34e6d7558043',_binary 'sdf',1,'2019-04-29 21:06:37','text'),(28,'a0420f04-6a7f-11e9-a9b4-e86a6477cf1b','93f65451-efc4-11e8-918b-34e6d7558043',_binary '大师傅',1,'2019-04-29 21:06:39','text'),(29,'ade69115-7316-11e9-a78b-e86a6477cf1b','93f65451-efc4-11e8-918b-34e6d7558044',_binary 'df',1,'2019-05-10 19:28:05','text'),(34,'650413f9-7322-11e9-a78b-e86a6477cf1b','93f65451-efc4-11e8-918b-34e6d7558044',_binary 'static/file/20190510205157.png',1,'2019-05-10 20:51:57','file'),(35,'59b52d1e-7325-11e9-a78b-e86a6477cf1b','93f65451-efc4-11e8-918b-34e6d7558043',_binary 'static/file/20190510211306.png',1,'2019-05-10 21:13:06','file'),(36,'398e1974-7326-11e9-b800-e86a6477cf1b','68c4c74d-725d-11e9-9e58-e86a6477cf1b',_binary '好友测试',1,'2019-05-10 21:19:22','text'),(37,'3fa22ec4-7326-11e9-b800-e86a6477cf1b','68c4c74d-725d-11e9-9e58-e86a6477cf1b',_binary 'static/file/20190510211932.png',1,'2019-05-10 21:19:32','file'),(38,'9dd73e5a-7575-11e9-9432-e86a6477cf1b','93f65451-efc4-11e8-918b-34e6d7558044',_binary '回复测试',3,'2019-05-13 19:52:43','text'),(39,'a330bb82-7575-11e9-9432-e86a6477cf1b','93f65451-efc4-11e8-918b-34e6d7558044',_binary '在线测试',1,'2019-05-13 19:52:52','text'),(40,'acf36518-7575-11e9-9432-e86a6477cf1b','93f65451-efc4-11e8-918b-34e6d7558043',_binary '测试',1,'2019-05-13 19:53:08','text'),(41,'233cd3c9-7909-11e9-bd72-e86a6477cf1b','68c4c74d-725d-11e9-9e58-e86a6477cf1b',_binary '回复测试',1,'2019-05-18 09:06:16','text'),(71,'600051c8-7918-11e9-84b5-e86a6477cf1b','68c4c74d-725d-11e9-9e58-e86a6477cf1b',_binary '测试',1,'2019-05-18 10:55:20','text'),(72,'6307fb55-7918-11e9-84b5-e86a6477cf1b','68c4c74d-725d-11e9-9e58-e86a6477cf1b',_binary 'dsfaljl',4,'2019-05-18 10:55:25','text'),(73,'83032762-7918-11e9-84b5-e86a6477cf1b','68c4c74d-725d-11e9-9e58-e86a6477cf1b',_binary 'static/file/20190518105619.png',4,'2019-05-18 10:56:19','file');
/*!40000 ALTER TABLE `group_msg` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `group_users`
--

DROP TABLE IF EXISTS `group_users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `group_users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `group_id` varchar(50) DEFAULT '' COMMENT '好友列表(包括群组)id',
  `uid` int(11) DEFAULT NULL COMMENT '用户id',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `group_id` (`group_id`,`uid`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=100000010 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `group_users`
--

LOCK TABLES `group_users` WRITE;
/*!40000 ALTER TABLE `group_users` DISABLE KEYS */;
INSERT INTO `group_users` VALUES (100000000,'93f65451-efc4-11e8-918b-34e6d7558043',1,NULL),(100000001,'93f65451-efc4-11e8-918b-34e6d7558043',2,NULL),(100000002,'93f65451-efc4-11e8-918b-34e6d7558044',1,NULL),(100000003,'93f65451-efc4-11e8-918b-34e6d7558044',3,NULL),(100000004,'93f65451-efc4-11e8-918b-34e6d7558045',1,NULL),(100000006,'68c4c74d-725d-11e9-9e58-e86a6477cf1b',4,'2019-05-09 21:21:52'),(100000007,'68c4c74d-725d-11e9-9e58-e86a6477cf1b',1,'2019-05-09 21:21:52'),(100000008,'96ece7f6-7918-11e9-99d7-e86a6477cf1b',2,'2019-05-18 10:56:52'),(100000009,'96ece7f6-7918-11e9-99d7-e86a6477cf1b',4,'2019-05-18 10:56:52');
/*!40000 ALTER TABLE `group_users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(30) DEFAULT '' COMMENT '用户名',
  `headimg` varchar(255) DEFAULT '' COMMENT '头像',
  `password` varchar(50) DEFAULT '' COMMENT '密码',
  `introduce` varchar(100) DEFAULT NULL COMMENT '自我介绍',
  `createtime` datetime DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (1,'admin','static/file/head1.jpeg','H5mIc17Z93HWDPkU+1tLQg==','admin测试','2018-11-24 19:15:54'),(2,'test','static/file/head2.jpeg','H5mIc17Z93HWDPkU+1tLQg==','test测试,这是一份测试','2018-11-24 19:15:54'),(3,'user','static/file/head3.jpeg','H5mIc17Z93HWDPkU+1tLQg==','user测试','2018-11-24 19:15:54'),(4,'lululu','static/file/head1.jpeg','H5mIc17Z93HWDPkU+1tLQg==','lullulu的世界','2018-11-24 19:15:54'),(5,'从前有个鹿，lululu～','static/file/head2.jpeg','H5mIc17Z93HWDPkU+1tLQg==','各位看官好，我是鹿成','2018-11-24 19:15:54'),(6,'test1','static/file/head2.jpeg','H5mIc17Z93HWDPkU+1tLQg==','test1测试','2018-11-24 19:15:54'),(7,'test2','static/file/head2.jpeg','H5mIc17Z93HWDPkU+1tLQg==','test2测试','2018-11-24 19:15:54'),(8,'test3','static/file/head2.jpeg','H5mIc17Z93HWDPkU+1tLQg==','test3测试v2','2018-11-24 19:15:54');
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

-- Dump completed on 2019-06-28 14:43:14
