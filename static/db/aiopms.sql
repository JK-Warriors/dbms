/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 50505
Source Host           : localhost:3306
Source Database       : aiopms

Target Server Type    : MYSQL
Target Server Version : 50505
File Encoding         : 65001

Date: 2017-03-28 17:23:47
*/

SET FOREIGN_KEY_CHECKS=0;




-- ----------------------------
-- Table structure for pms_roles
-- ----------------------------
DROP TABLE IF EXISTS `pms_roles`;
CREATE TABLE `pms_roles` (
  `id` bigint(20) NOT NULL,
  `name` varchar(30) DEFAULT NULL COMMENT '角色名称',
  `summary` varchar(500) DEFAULT NULL COMMENT '角色描述',
  `created` int(10) DEFAULT NULL,
  `changed` int(10) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `INDEX_NCC` (`name`,`created`,`changed`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='角色表';

-- ----------------------------
-- Records of pms_roles
-- ----------------------------
INSERT INTO `pms_roles` VALUES ('1', '管理员', '系统管理员', UNIX_TIMESTAMP(now()), UNIX_TIMESTAMP(now()));
INSERT INTO `pms_roles` VALUES ('2', '业务员', '业务员', UNIX_TIMESTAMP(now()), UNIX_TIMESTAMP(now()));
INSERT INTO `pms_roles` VALUES ('3', '审计员', '审计员', UNIX_TIMESTAMP(now()), UNIX_TIMESTAMP(now()));


-- ----------------------------
-- Table structure for pms_permissions
-- ----------------------------
DROP TABLE IF EXISTS `pms_permissions`;
CREATE TABLE `pms_permissions` (
  `id` bigint(20) NOT NULL,
  `parent_id` bigint(20) DEFAULT NULL,
  `name` varchar(50) DEFAULT NULL COMMENT '中文名称',
  `ename` varchar(50) DEFAULT NULL COMMENT '英文名称',
  `url` varchar(255) DEFAULT '0' COMMENT 'URL地址',
  `icon` varchar(20) DEFAULT NULL,
  `is_nav` tinyint(1) DEFAULT '0' COMMENT '1是0否导航',
  `is_show` tinyint(1) DEFAULT '0' COMMENT '0不显示1显示',
  `sort` tinyint(4) DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `INDEX_PNETW` (`parent_id`,`name`,`ename`,`is_show`,`sort`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='权限表';

-- ----------------------------
-- Records of pms_permissions
-- ----------------------------
INSERT INTO `pms_permissions` VALUES ('0', '0', '根节点', 'root', '/root', 'fa-root', '0', '0', '0');
INSERT INTO `pms_permissions` VALUES ('1', '0', '业务配置', 'config', '/config', 'fa-config', '1', '1', '1');
INSERT INTO `pms_permissions` VALUES ('2', '0', '权限管理', 'auth', '/auth', '', '1', '1', '2');
INSERT INTO `pms_permissions` VALUES ('3', '0', '系统管理', 'system', '/system', '', '1', '1', '3');

INSERT INTO `pms_permissions` VALUES ('1010', '2', '用户管理', 'user-manage', '/user/manage', 'fa-user', '1', '1', '0');
INSERT INTO `pms_permissions` VALUES ('1011', '2', '添加用户', 'user-add', '/user/add', null, '0', '0', '0');
INSERT INTO `pms_permissions` VALUES ('1012', '2', '编辑用户', 'user-edit', '/user/edit', null, '0', '0', '0');
INSERT INTO `pms_permissions` VALUES ('1013', '2', '删除用户', 'user-delete', '/user/delete', '', '0', '0', '0');
INSERT INTO `pms_permissions` VALUES ('1020', '2', '角色管理', 'role-manage', '/role/manage', 'fa-user', '1', '1', '1');
INSERT INTO `pms_permissions` VALUES ('1021', '2', '添加角色', 'role-add', '/role/add', '', '0', '0', '0');
INSERT INTO `pms_permissions` VALUES ('1022', '2', '编辑角色', 'role-edit', '/role/edit', '', '0', '0', '0');
INSERT INTO `pms_permissions` VALUES ('1023', '2', '删除角色', 'role-delete', '/role/delete', '', '0', '0', '0');
INSERT INTO `pms_permissions` VALUES ('1024', '2', '角色权限', 'role-permission', '/role/permission', '', '0', '0', '0');
-- INSERT INTO `pms_permissions` VALUES ('1025', '2', '角色成员', 'role-user', '/role/user', '', '0', '0', '0');
-- INSERT INTO `pms_permissions` VALUES ('1026', '2', '添加角色', 'role-user-add', '/role/useradd', '', '0', '0', '0');
-- INSERT INTO `pms_permissions` VALUES ('1027', '2', '删除角色', 'role-user-delete', '/role/userdelete', '', '0', '0', '0');
INSERT INTO `pms_permissions` VALUES ('1030', '2', '权限管理', 'permission-manage', '/permission/manage', 'fa-user', '1', '1', '2');
INSERT INTO `pms_permissions` VALUES ('1031', '2', '添加权限', 'permission-add', '/permission/add', '', '0', '0', '0');
INSERT INTO `pms_permissions` VALUES ('1032', '2', '编辑权限', 'permission-edit', '/permission/edit', '', '0', '0', '0');
INSERT INTO `pms_permissions` VALUES ('1033', '2', '删除权限', 'permission-delete', '/permission/delete', '', '0', '0', '0');


INSERT INTO `pms_permissions` VALUES ('1040', '1', '数据库配置', 'dbconfig-manage', '/dbconfig/manage', 'fa-user', '1', '1', '1');
INSERT INTO `pms_permissions` VALUES ('1041', '1', '添加数据库', 'dbconfig-add', '/dbconfig/add', '', '0', '0', '0');
INSERT INTO `pms_permissions` VALUES ('1042', '1', '编辑数据库', 'dbconfig-edit', '/dbconfig/edit', '', '0', '0', '0');
INSERT INTO `pms_permissions` VALUES ('1043', '1', '删除数据库', 'dbconfig-delete', '/dbconfig/delete', '', '0', '0', '0');

INSERT INTO `pms_permissions` VALUES ('1050', '3', '日志管理', 'log-manage', '/log/manage', 'fa-user', '1', '1', '1');
INSERT INTO `pms_permissions` VALUES ('1051', '3', '日志删除', 'log-delete', '/log/delete', 'fa-user', '0', '0', '0');

INSERT INTO `pms_permissions` VALUES ('1060', '3', '消息管理', 'message-manage', '/message/manage', 'fa-user', '0', '0', '0');
INSERT INTO `pms_permissions` VALUES ('1061', '3', '消息删除', 'message-delete', '/message/delete', 'fa-user', '0', '0', '0');

INSERT INTO `pms_permissions` VALUES ('1070', '1', '数据库查看', 'dbmonitor-manage', '/dbmonitor/manage', 'fa-user', '1', '1', '2');


-- ----------------------------
-- Table structure for pms_role_permission
-- ----------------------------
DROP TABLE IF EXISTS `pms_role_permission`;
CREATE TABLE `pms_role_permission` (
  `id` bigint(20) NOT NULL,
  `role_id` bigint(20) DEFAULT NULL,
  `permission_id` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `INDEX_GP` (`role_id`,`permission_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='角色权限表';

-- ----------------------------
-- Records of pms_groups_permission
-- ----------------------------
INSERT INTO `pms_role_permission` VALUES ('1 ', '1', '1');
INSERT INTO `pms_role_permission` VALUES ('2 ', '1', '2');
INSERT INTO `pms_role_permission` VALUES ('3 ', '1', '3');
INSERT INTO `pms_role_permission` VALUES ('4 ', '1', '4');
INSERT INTO `pms_role_permission` VALUES ('5 ', '1', '1010');
INSERT INTO `pms_role_permission` VALUES ('6 ', '1', '1011');
INSERT INTO `pms_role_permission` VALUES ('7 ', '1', '1012');
INSERT INTO `pms_role_permission` VALUES ('8 ', '1', '1013');
INSERT INTO `pms_role_permission` VALUES ('9 ', '1', '1020');
INSERT INTO `pms_role_permission` VALUES ('10', '1', '1021');
INSERT INTO `pms_role_permission` VALUES ('11', '1', '1022');
INSERT INTO `pms_role_permission` VALUES ('12', '1', '1023');
INSERT INTO `pms_role_permission` VALUES ('13', '1', '1024');
INSERT INTO `pms_role_permission` VALUES ('14', '1', '1025');
INSERT INTO `pms_role_permission` VALUES ('15', '1', '1026');
INSERT INTO `pms_role_permission` VALUES ('16', '1', '1027');
INSERT INTO `pms_role_permission` VALUES ('17', '1', '1030');
INSERT INTO `pms_role_permission` VALUES ('18', '1', '1031');
INSERT INTO `pms_role_permission` VALUES ('19', '1', '1032');
INSERT INTO `pms_role_permission` VALUES ('20', '1', '1033');
INSERT INTO `pms_role_permission` VALUES ('21', '1', '1040');
INSERT INTO `pms_role_permission` VALUES ('22', '1', '1041');
INSERT INTO `pms_role_permission` VALUES ('23', '1', '1042');
INSERT INTO `pms_role_permission` VALUES ('24', '1', '1043');
INSERT INTO `pms_role_permission` VALUES ('25', '1', '1050');
INSERT INTO `pms_role_permission` VALUES ('26', '1', '1051');
INSERT INTO `pms_role_permission` VALUES ('27', '1', '1060');
INSERT INTO `pms_role_permission` VALUES ('28', '1', '1061');
INSERT INTO `pms_role_permission` VALUES ('29', '1', '1070');



-- ----------------------------
-- Table structure for pms_users
-- ----------------------------
DROP TABLE IF EXISTS `pms_users`;
CREATE TABLE `pms_users` (
  `userid` bigint(20) NOT NULL,
  `profile_id` bigint(20) DEFAULT NULL,
  `username` varchar(15) DEFAULT NULL COMMENT '用户名',
  `password` varchar(255) DEFAULT NULL COMMENT '密码',
  `avatar` varchar(100) DEFAULT NULL,
  `status` tinyint(1) DEFAULT '1' COMMENT '状态1正常，2禁用',
  PRIMARY KEY (`userid`),
  KEY `INDEX_US` (`username`,`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户主表';

-- ----------------------------
-- Records of pms_users
-- ----------------------------
INSERT INTO `pms_users` VALUES ('1', '1', 'admin', 'e10adc3949ba59abbe56e057f20f883e', '/static/img/avatar/1.jpg', '1');
INSERT INTO `pms_users` VALUES ('2', '2', 'audit', 'e10adc3949ba59abbe56e057f20f883e', '/static/uploadfile/2017-3/28/5b41faa955a4c1acdb6d7e6c116bce2f-cropper.jpg', '1');

-- ----------------------------
-- Table structure for pms_users_profile
-- ----------------------------
DROP TABLE IF EXISTS `pms_users_profile`;
CREATE TABLE `pms_users_profile` (
  `userid` bigint(20) NOT NULL auto_increment,
  `realname` varchar(15) DEFAULT NULL COMMENT '姓名',
  `sex` tinyint(1) DEFAULT '1' COMMENT '1男2女',
  `birth` varchar(15) DEFAULT NULL,
  `email` varchar(30) DEFAULT NULL COMMENT '邮箱',
  `webchat` varchar(15) DEFAULT NULL COMMENT '微信号',
  `qq` varchar(15) DEFAULT NULL COMMENT 'qq号',
  `phone` varchar(15) DEFAULT NULL COMMENT '手机',
  `tel` varchar(20) DEFAULT NULL COMMENT '电话',
  `address` varchar(100) DEFAULT NULL COMMENT '地址',
  `emercontact` varchar(15) DEFAULT NULL COMMENT '紧急联系人',
  `emerphone` varchar(15) DEFAULT NULL COMMENT '紧急电话',
  `lognum` int(10) DEFAULT '0' COMMENT '登录次数',
  `ip` varchar(15) DEFAULT NULL COMMENT '最近登录IP',
  `lasted` int(10) DEFAULT NULL COMMENT '最近登录时间',
  PRIMARY KEY (`userid`),
  KEY `INDEX_RSL` (`realname`,`sex`,`lasted`)
) ENGINE=InnoDB AUTO_INCREMENT=1001 DEFAULT CHARSET=utf8 COMMENT='用户详情表';

-- ----------------------------
-- Records of pms_users_profile
-- ----------------------------
INSERT INTO `pms_users_profile` VALUES ('1', 'admin', '1', '1993-03-06', 'admin@tom.com', '', '', '13282176663', '', '', '', '',  '0', '', '0');
INSERT INTO `pms_users_profile` VALUES ('2', 'audit', '1', '1985-12-12', 'audit@163.com', '', '', '', '', '', '', '',  '0', '', '0');

-- ----------------------------
-- Table structure for pms_role_user
-- ----------------------------
DROP TABLE IF EXISTS `pms_role_user`;
CREATE TABLE `pms_role_user` (
  `id` bigint(20) NOT NULL auto_increment,
  `role_id` bigint(20) DEFAULT NULL,
  `user_id` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `INDEX_GU` (`role_id`,`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=100 DEFAULT CHARSET=utf8 COMMENT='角色成员';

-- ----------------------------
-- Records of pms_groups_user
-- ----------------------------
INSERT INTO `pms_role_user` VALUES ('1', '1', '1');
INSERT INTO `pms_role_user` VALUES ('2', '2', '2');


-- ----------------------------
-- Table structure for pms_messages
-- ----------------------------
DROP TABLE IF EXISTS `pms_messages`;
CREATE TABLE `pms_messages` (
  `msgid` bigint(20) NOT NULL,
  `userid` bigint(20) DEFAULT NULL,
  `touserid` bigint(20) DEFAULT NULL,
  `type` tinyint(2) DEFAULT NULL COMMENT '类型1评论2赞3审批',
  `subtype` tinyint(3) DEFAULT NULL COMMENT '11知识评论12相册评论21知识赞22相册赞31请假审批32加班33报销34出差35外出36物品',
  `title` varchar(200) DEFAULT NULL,
  `url` varchar(200) DEFAULT NULL,
  `view` tinyint(1) DEFAULT '1' COMMENT '1未看，2已看',
  `created` int(10) DEFAULT NULL,
  PRIMARY KEY (`msgid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT=' 消息表';

-- ----------------------------
-- Records of pms_messages
-- ----------------------------
INSERT INTO `pms_messages` VALUES ('66618325785907200', '1461312703628858832', '1469024587469707428', '4', '31', '去审批处理', '/leave/approval/66618286464307200', '1', '1490685934');
INSERT INTO `pms_messages` VALUES ('66626417378463744', '1461312703628858832', '1461312703628858832', '1', '11', 'OPMS 1.2 版本更新发布', '/knowledge/66618679508340736', '1', '1490687863');
INSERT INTO `pms_messages` VALUES ('66639445431947264', '1461312703628858832', '1461312703628858832', '1', '12', '油菜花', '/album/66621262012616704', '1', '1490690969');


-- ----------------------------
-- Table structure for pms_admin_log
-- ----------------------------
DROP TABLE IF EXISTS `pms_admin_log`;
CREATE TABLE `pms_admin_log` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `user_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '用户ID',
  `username` varchar(30) NOT NULL DEFAULT '' COMMENT '用户名称',
  `url` varchar(1500) NOT NULL DEFAULT '' COMMENT '操作页面',
  `title` varchar(100) NOT NULL DEFAULT '' COMMENT '日志标题',
  `content` text NOT NULL COMMENT '内容',
  `ip` varchar(50) NOT NULL DEFAULT '' COMMENT 'IP',
  `useragent` varchar(255) NOT NULL DEFAULT '' COMMENT 'User-Agent',
  `created` int(10) DEFAULT NULL COMMENT '操作时间',
  PRIMARY KEY (`id`),
  KEY `name` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT='系统日志表';


-- -----------------------------------------------------------------------------
-- Table structure for pms_db_config
-- -----------------------------------------------------------------------------
DROP TABLE IF EXISTS `pms_db_config`;
CREATE TABLE `pms_db_config` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `db_type` tinyint(2) DEFAULT NULL,
  `host` varchar(30) NOT NULL DEFAULT '' COMMENT '数据库IP',
  `port` int(10) NOT NULL DEFAULT 0 COMMENT '数据库端口',
  `alias` varchar(255) DEFAULT '' COMMENT '别名',
  `instance_name` varchar(50) DEFAULT '' COMMENT '实例名',
  `db_name` varchar(50) DEFAULT '' COMMENT '数据库名',
  `username` varchar(30) DEFAULT '' COMMENT '用户名',
  `password` varchar(255) DEFAULT '' COMMENT '密码',
  `role` tinyint(2) DEFAULT 1 COMMENT '1：主；2: 备',
  `status` tinyint(2) DEFAULT 1 COMMENT '1: 激活；0：禁用',
  `is_delete` tinyint(2) DEFAULT 0 COMMENT '1: 删除；0：未删除',
  `retention` int(10) NOT NULL DEFAULT 0 COMMENT '保留时间，默认单位为天',
  `created` int(10) DEFAULT NULL COMMENT '操作时间',
  `updated` int(10) DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `host` (`host`),
  KEY `alias` (`alias`)
)ENGINE=InnoDB AUTO_INCREMENT=101 DEFAULT CHARSET=utf8 COMMENT='数据库配置表';


-- -----------------------------------------------------------------------------
-- Table structure for pms_template
-- -----------------------------------------------------------------------------
DROP TABLE IF EXISTS `pms_template`;
CREATE TABLE `pms_template` (
  `template_id` int(10) NOT NULL AUTO_INCREMENT,
  `db_type`     varchar(50) DEFAULT NULL,
  `scraper_name` varchar(255)  DEFAULT NULL,
  `subsystem` varchar(255)  DEFAULT NULL,
  `metrix_name` varchar(255)  DEFAULT NULL,
  `label` varchar(255)  DEFAULT NULL,
  `value_type` tinyint(2) DEFAULT 0 COMMENT '1: Counter；2: Gauge；3：Histogram；4：Summary；5：Untyped',
  PRIMARY KEY (`template_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='模板表';


-- -----------------------------------------------------------------------------
-- Table structure for pms_items
-- -----------------------------------------------------------------------------
DROP TABLE IF EXISTS `pms_items`;
CREATE TABLE `pms_items` (
	`item_id`                bigint unsigned                           NOT NULL         AUTO_INCREMENT,
	`type`                   integer         DEFAULT '0'               NOT NULL,
	`template_id`            bigint unsigned                           NULL,
	`obj_type`               varchar(50)                               NOT NULL,
	`obj_id`                 bigint unsigned                           NOT NULL,
	`name`                   varchar(255)    DEFAULT ''                NOT NULL,
	`key_`                   varchar(255)    DEFAULT ''                NOT NULL,
	`label`                  varchar(255)    DEFAULT ''                NOT NULL,
	`value_type`             integer         DEFAULT '1'               NOT NULL,
	`units`                  varchar(255)    DEFAULT ''                NOT NULL,
	`status`                 integer         DEFAULT '1'               NOT NULL,
  PRIMARY KEY (`item_id`),
  KEY `idx_items_1` (`obj_id`,`key_`,`label`)
) ENGINE=InnoDB AUTO_INCREMENT=101 DEFAULT CHARSET=utf8 COMMENT='items表';

-- -----------------------------------------------------------------------------
-- Table structure for pms_item_data
-- -----------------------------------------------------------------------------
DROP TABLE IF EXISTS `pms_item_data`;
CREATE TABLE `pms_item_data` (
	`itemid`                 bigint unsigned                           NOT NULL,
	`time`                   int(10)         DEFAULT '0'               NOT NULL,
	`value`                  double(16,4)    DEFAULT '0.0000'          NOT NULL,
	`ns`                     integer         DEFAULT '0'               NOT NULL,
  PRIMARY KEY (`itemid`,`time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='数据表';


-- ----------------------------
-- Table structure for pms_db_status
-- ----------------------------
DROP TABLE IF EXISTS `pms_db_status`;
CREATE TABLE `pms_db_status` (
  `id` int(10) unsigned NOT NULL COMMENT 'ID',
  `db_type` tinyint(2) DEFAULT NULL COMMENT '数据库类型',
  `alias`   varchar(255) DEFAULT NULL COMMENT '别名',
  `connect` tinyint(2) DEFAULT NULL COMMENT '连接',
  `role`    varchar(30) DEFAULT NULL COMMENT '角色',
  `version` varchar(30) DEFAULT NULL COMMENT '版本',
  `created` int(10) DEFAULT NULL COMMENT '操作时间',
  `updated` int(10) DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*
CREATE TABLE `history_uint` (
	`itemid`                 bigint unsigned                           NOT NULL,
	`time`                  integer         DEFAULT '0'               NOT NULL,
	`value`                  bigint unsigned DEFAULT '0'               NOT NULL,
	`ns`                     integer         DEFAULT '0'               NOT NULL
) ENGINE=InnoDB;
CREATE INDEX `history_uint_1` ON `history_uint` (`itemid`,`clock`);
CREATE TABLE `history_str` (
	`itemid`                 bigint unsigned                           NOT NULL,
	`time`                  integer         DEFAULT '0'               NOT NULL,
	`value`                  varchar(255)    DEFAULT ''                NOT NULL,
	`ns`                     integer         DEFAULT '0'               NOT NULL
) ENGINE=InnoDB;
CREATE INDEX `history_str_1` ON `history_str` (`itemid`,`clock`);
CREATE TABLE `history_log` (
	`itemid`                 bigint unsigned                           NOT NULL,
	`time`                   integer         DEFAULT '0'               NOT NULL,
	`timestamp`              integer         DEFAULT '0'               NOT NULL,
	`source`                 varchar(64)     DEFAULT ''                NOT NULL,
	`severity`               integer         DEFAULT '0'               NOT NULL,
	`value`                  text                                      NOT NULL,
	`logeventid`             integer         DEFAULT '0'               NOT NULL,
	`ns`                     integer         DEFAULT '0'               NOT NULL
) ENGINE=InnoDB;
CREATE INDEX `history_log_1` ON `history_log` (`itemid`,`clock`);
CREATE TABLE `history_text` (
	`itemid`                 bigint unsigned                           NOT NULL,
	`clock`                  integer         DEFAULT '0'               NOT NULL,
	`value`                  text                                      NOT NULL,
	`ns`                     integer         DEFAULT '0'               NOT NULL
) ENGINE=InnoDB;
CREATE INDEX `history_text_1` ON `history_text` (`itemid`,`clock`);
*/