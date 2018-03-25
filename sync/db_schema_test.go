package sync

import (
	"fmt"
	"testing"
)

var schemaData = "-- Adminer 4.3.1 MySQL dump\n\nSET NAMES utf8;\nSET time_zone = '+00:00';\nSET foreign_key_checks = 0;\nSET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';\n\nDROP TABLE IF EXISTS `code_invitation`;\nCREATE TABLE `code_invitation` (\n  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',\n  `code` varchar(40) NOT NULL COMMENT '邀请码',\n  `created` int(11) unsigned NOT NULL COMMENT '创建时间',\n  `uid` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建者',\n  `recv_uid` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '使用者',\n  `used` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '使用时间',\n  `start` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '有效时间',\n  `end` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '失效时间',\n  `disabled` enum('Y','N') NOT NULL DEFAULT 'N' COMMENT '是否禁用',\n  `role_ids` text NOT NULL COMMENT '注册为角色(多个用“,”分隔开)',\n  PRIMARY KEY (`id`),\n  UNIQUE KEY `code` (`code`)\n) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='邀请码';\n\n\nDROP TABLE IF EXISTS `code_verification`;\nCREATE TABLE `code_verification` (\n  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',\n  `code` varchar(40) NOT NULL COMMENT '验证码',\n  `created` int(11) unsigned NOT NULL COMMENT '创建时间',\n  `uid` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建者',\n  `used` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '使用时间',\n  `purpose` varchar(40) NOT NULL COMMENT '目的',\n  `start` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '有效时间',\n  `end` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '失效时间',\n  `disabled` enum('Y','N') NOT NULL DEFAULT 'N' COMMENT '是否禁用',\n  PRIMARY KEY (`id`),\n  UNIQUE KEY `code` (`code`)\n) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='验证码';\n"

func TestDBSchema(t *testing.T) {
	schema := NewMySchemaData(schemaData, `mysql`)
	fmt.Printf(`tables: %#v`+"\n", schema.GetTableNames())
	fmt.Printf(`schema: %s`+"\n", schema.GetTableSchema(`code_invitation`))
}
