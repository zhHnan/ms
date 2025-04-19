CREATE TABLE `ms_member`  (
                              `id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT '系统前台用户表',
                              `account` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '用户登陆账号',
                              `password` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '登陆密码',
                              `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '用户昵称',
                              `mobile` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '手机',
                              `realname` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '真实姓名',
                              `create_time` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '创建时间',
                              `status` tinyint(1) NULL DEFAULT 0 COMMENT '状态',
                              `last_login_time` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '上次登录时间',
                              `sex` tinyint(0) NULL DEFAULT 0 COMMENT '性别',
                              `avatar` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '头像',
                              `idcard` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '身份证',
                              `province` int(0) NULL DEFAULT 0 COMMENT '省',
                              `city` int(0) NULL DEFAULT 0 COMMENT '市',
                              `area` int(0) NULL DEFAULT 0 COMMENT '区',
                              `address` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '所在地址',
                              `description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '备注',
                              `email` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '邮箱',
                              `dingtalk_openid` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '钉钉openid',
                              `dingtalk_unionid` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '钉钉unionid',
                              `dingtalk_userid` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '钉钉用户id',
                              PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1000 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '用户表' ROW_FORMAT = COMPACT;

CREATE TABLE `ms_organization`  (
                                    `id` bigint(0) NOT NULL AUTO_INCREMENT,
                                    `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名称',
                                    `avatar` varchar(511) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '头像',
                                    `description` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '描述',
                                    `member_id` bigint(0) NULL DEFAULT NULL COMMENT '拥有者',
                                    `create_time` bigint(0) NULL DEFAULT NULL COMMENT '创建时间',
                                    `personal` tinyint(1) NULL DEFAULT 0 COMMENT '是否个人项目',
                                    `address` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '地址',
                                    `province` int(0) NULL DEFAULT 0 COMMENT '省',
                                    `city` int(0) NULL DEFAULT 0 COMMENT '市',
                                    `area` int(0) NULL DEFAULT 0 COMMENT '区',
                                    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '组织表' ROW_FORMAT = COMPACT;

CREATE TABLE `ms_project`  (
                               `id` bigint(0) UNSIGNED NOT NULL AUTO_INCREMENT,
                               `cover` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '封面',
                               `name` varchar(90) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '名称',
                               `description` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '描述',
                               `access_control_type` tinyint(0) NULL DEFAULT 0 COMMENT '访问控制l类型',
                               `white_list` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '可以访问项目的权限组（白名单）',
                               `order` int(0) UNSIGNED NULL DEFAULT 0 COMMENT '排序',
                               `deleted` tinyint(1) NULL DEFAULT 0 COMMENT '删除标记',
                               `template_code` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '项目类型',
                               `schedule` double(5, 2) NULL DEFAULT 0.00 COMMENT '进度',
                               `create_time` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '创建时间',
                               `organization_code` bigint(0) NULL DEFAULT NULL COMMENT '组织id',
                               `deleted_time` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '删除时间',
                               `private` tinyint(1) NULL DEFAULT 1 COMMENT '是否私有',
                               `prefix` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '项目前缀',
                               `open_prefix` tinyint(1) NULL DEFAULT 0 COMMENT '是否开启项目前缀',
                               `archive` tinyint(1) NULL DEFAULT 0 COMMENT '是否归档',
                               `archive_time` bigint(0) NULL DEFAULT NULL COMMENT '归档时间',
                               `open_begin_time` tinyint(1) NULL DEFAULT 0 COMMENT '是否开启任务开始时间',
                               `open_task_private` tinyint(1) NULL DEFAULT 0 COMMENT '是否开启新任务默认开启隐私模式',
                               `task_board_theme` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT 'default' COMMENT '看板风格',
                               `begin_time` bigint(0) NULL DEFAULT NULL COMMENT '项目开始日期',
                               `end_time` bigint(0) NULL DEFAULT NULL COMMENT '项目截止日期',
                               `auto_update_schedule` tinyint(1) NULL DEFAULT 0 COMMENT '自动更新项目进度',
                               PRIMARY KEY (`id`) USING BTREE,
                               INDEX `project`(`order`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 13043 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '项目表' ROW_FORMAT = COMPACT;

CREATE TABLE `ms_project_member`  (
                                      `id` bigint(0) NOT NULL AUTO_INCREMENT,
                                      `project_code` bigint(0) NULL DEFAULT NULL COMMENT '项目id',
                                      `member_code` bigint(0) NULL DEFAULT NULL COMMENT '成员id',
                                      `join_time` bigint(0) NULL DEFAULT NULL COMMENT '加入时间',
                                      `is_owner` bigint(0) NULL DEFAULT 0 COMMENT '拥有者',
                                      `authorize` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '角色',
                                      PRIMARY KEY (`id`) USING BTREE,
                                      UNIQUE INDEX `unique`(`project_code`, `member_code`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 37 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '项目-成员表' ROW_FORMAT = COMPACT;