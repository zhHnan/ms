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

CREATE TABLE `ms_project_template`  (
                                        `id` int(0) NOT NULL AUTO_INCREMENT,
                                        `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '类型名称',
                                        `description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '备注',
                                        `sort` tinyint(0) NULL DEFAULT 0,
                                        `create_time` bigint(20)  NULL DEFAULT 0,
                                        `organization_code` bigint(0) NULL DEFAULT NULL COMMENT '组织id',
                                        `cover` varchar(511) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '封面',
                                        `member_code` bigint(0) NULL DEFAULT NULL COMMENT '创建人',
                                        `is_system` tinyint(1) NULL DEFAULT 0 COMMENT '系统默认',
                                        PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 20 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '项目类型表' ROW_FORMAT = COMPACT;


CREATE TABLE `ms_task_stages_template`  (
                                            `id` int(0) NOT NULL AUTO_INCREMENT,
                                            `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '类型名称',
                                            `project_template_code` int(0) NULL DEFAULT 0 COMMENT '项目id',
                                            `create_time` bigint(0) NULL DEFAULT NULL,
                                            `sort` int(0) NULL DEFAULT 0,
                                            PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 84 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '任务列表模板表' ROW_FORMAT = COMPACT;


CREATE TABLE `ms_task_stages`  (
                                   `id` int(0) NOT NULL AUTO_INCREMENT,
                                   `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '类型名称',
                                   `project_code` bigint(0) NULL DEFAULT NULL COMMENT '项目id',
                                   `sort` int(0) NULL DEFAULT 0 COMMENT '排序',
                                   `description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '备注',
                                   `create_time` bigint(0) NULL DEFAULT NULL COMMENT '创建时间',
                                   `deleted` tinyint(1) NULL DEFAULT 0 COMMENT '删除标记',
                                   PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 77 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '任务列表表' ROW_FORMAT = COMPACT;


CREATE TABLE `ms_task`  (
                            `id` bigint(0) UNSIGNED NOT NULL AUTO_INCREMENT,
                            `project_code` bigint(0) NOT NULL DEFAULT 0 COMMENT '项目编号',
                            `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
                            `pri` tinyint(0) UNSIGNED NULL DEFAULT 0 COMMENT '紧急程度',
                            `execute_status` tinyint(0) NULL DEFAULT NULL COMMENT '执行状态',
                            `description` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '详情',
                            `create_by` bigint(0) NULL DEFAULT NULL COMMENT '创建人',
                            `done_by` bigint(0) NULL DEFAULT NULL COMMENT '完成人',
                            `done_time` bigint(0) NULL DEFAULT NULL COMMENT '完成时间',
                            `create_time` bigint(0) NULL DEFAULT NULL COMMENT '创建日期',
                            `assign_to` bigint(0) NULL DEFAULT NULL COMMENT '指派给谁',
                            `deleted` tinyint(1) NULL DEFAULT 0 COMMENT '回收站',
                            `stage_code` int(0) NULL DEFAULT NULL COMMENT '任务列表',
                            `task_tag` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '任务标签',
                            `done` tinyint(0) NULL DEFAULT 0 COMMENT '是否完成',
                            `begin_time` bigint(0) NULL DEFAULT NULL COMMENT '开始时间',
                            `end_time` bigint(0) NULL DEFAULT NULL COMMENT '截止时间',
                            `remind_time` bigint(0) NULL DEFAULT NULL COMMENT '提醒时间',
                            `pcode` bigint(0) NULL DEFAULT NULL COMMENT '父任务id',
                            `sort` int(0) NULL DEFAULT 0 COMMENT '排序',
                            `like` int(0) NULL DEFAULT 0 COMMENT '点赞数',
                            `star` int(0) NULL DEFAULT 0 COMMENT '收藏数',
                            `deleted_time` bigint(0) NULL DEFAULT NULL COMMENT '删除时间',
                            `private` tinyint(1) NULL DEFAULT 0 COMMENT '是否隐私模式',
                            `id_num` int(0) NULL DEFAULT 1 COMMENT '任务id编号',
                            `path` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '上级任务路径',
                            `schedule` int(0) NULL DEFAULT 0 COMMENT '进度百分比',
                            `version_code` bigint(0) NULL DEFAULT 0 COMMENT '版本id',
                            `features_code` bigint(0) NULL DEFAULT 0 COMMENT '版本库id',
                            `work_time` int(0) NULL DEFAULT 0 COMMENT '预估工时',
                            `status` tinyint(0) NULL DEFAULT 0 COMMENT '执行状态。0：未开始，1：已完成，2：进行中，3：挂起，4：测试中',
                            PRIMARY KEY (`id`, `project_code`) USING BTREE,
                            INDEX `stage_code`(`stage_code`) USING BTREE,
                            INDEX `project_code`(`project_code`) USING BTREE,
                            INDEX `pcode`(`pcode`) USING BTREE,
                            INDEX `sort`(`sort`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 12363 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '任务表' ROW_FORMAT = COMPACT;

CREATE TABLE `ms_task_member`  (
                                   `id` bigint(0) NOT NULL AUTO_INCREMENT,
                                   `task_code` bigint(0) NULL DEFAULT 0 COMMENT '任务ID',
                                   `is_executor` tinyint(1) NULL DEFAULT 0 COMMENT '执行者',
                                   `member_code` bigint(0) NULL DEFAULT NULL COMMENT '成员id',
                                   `join_time` bigint(0) NULL DEFAULT NULL,
                                   `is_owner` tinyint(1) NULL DEFAULT 0 COMMENT '是否创建人',
                                   PRIMARY KEY (`id`) USING BTREE,
                                   UNIQUE INDEX `id`(`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 273 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '任务-成员表' ROW_FORMAT = COMPACT;

CREATE TABLE `ms_project_log`  (
                                   `id` bigint(0) NOT NULL AUTO_INCREMENT,
                                   `member_code` bigint(0) NULL DEFAULT 0 COMMENT '操作人id',
                                   `content` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '操作内容',
                                   `remark` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
                                   `type` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT 'create' COMMENT '操作类型',
                                   `create_time` bigint(0) NULL DEFAULT NULL COMMENT '添加时间',
                                   `source_code` bigint(0) NULL DEFAULT 0 COMMENT '任务id',
                                   `action_type` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '场景类型',
                                   `to_member_code` bigint(0) NULL DEFAULT 0,
                                   `is_comment` tinyint(1) NULL DEFAULT 0 COMMENT '是否评论，0：否',
                                   `project_code` bigint(0) NULL DEFAULT NULL,
                                   `icon` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
                                   `is_robot` tinyint(1) NULL DEFAULT 0 COMMENT '是否机器人',
                                   PRIMARY KEY (`id`) USING BTREE,
                                   INDEX `member_code`(`member_code`) USING BTREE,
                                   INDEX `source_code`(`source_code`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5086 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '项目日志表' ROW_FORMAT = DYNAMIC;
CREATE TABLE `ms_task_work_time`  (
                                      `id` bigint(0) NOT NULL AUTO_INCREMENT,
                                      `task_code` bigint(0) NULL DEFAULT 0 COMMENT '任务ID',
                                      `member_code` bigint(0) NULL DEFAULT NULL COMMENT '成员id',
                                      `create_time` bigint(0) NULL DEFAULT NULL,
                                      `content` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '描述',
                                      `begin_time` bigint(0) NULL DEFAULT NULL COMMENT '开始时间',
                                      `num` int(0) NULL DEFAULT 0 COMMENT '工时',
                                      PRIMARY KEY (`id`) USING BTREE,
                                      UNIQUE INDEX `id`(`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '任务工时表' ROW_FORMAT = COMPACT;

DROP TABLE IF EXISTS `ms_project_menu`;
CREATE TABLE `ms_project_menu`  (
                                    `id` bigint(0) UNSIGNED NOT NULL AUTO_INCREMENT,
                                    `pid` bigint(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '父id',
                                    `title` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '名称',
                                    `icon` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '菜单图标',
                                    `url` varchar(400) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '链接',
                                    `file_path` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '文件路径',
                                    `params` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '链接参数',
                                    `node` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '#' COMMENT '权限节点',
                                    `sort` int(0) UNSIGNED NULL DEFAULT 0 COMMENT '菜单排序',
                                    `status` tinyint(0) UNSIGNED NULL DEFAULT 1 COMMENT '状态(0:禁用,1:启用)',
                                    `create_by` bigint(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建人',
                                    `is_inner` tinyint(1) NULL DEFAULT 0 COMMENT '是否内页',
                                    `values` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '参数默认值',
                                    `show_slider` tinyint(1) NULL DEFAULT 1 COMMENT '是否显示侧栏',
                                    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 176 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '项目菜单表' ROW_FORMAT = DYNAMIC;



CREATE TABLE `ms_file`  (
                            `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT,
                            `path_name` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '相对路径',
                            `title` char(90) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '名称',
                            `extension` char(30) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '扩展名',
                            `size` int(0) UNSIGNED NULL DEFAULT 0 COMMENT '文件大小',
                            `object_type` char(30) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '对象类型',
                            `organization_code` bigint(0) NULL DEFAULT NULL COMMENT '组织编码',
                            `task_code` bigint(0) NULL DEFAULT NULL COMMENT '任务编码',
                            `project_code` bigint(0) NULL DEFAULT NULL COMMENT '项目编码',
                            `create_by` bigint(0) NULL DEFAULT NULL COMMENT '上传人',
                            `create_time` bigint(0) NULL DEFAULT NULL COMMENT '创建时间',
                            `downloads` mediumint(0) UNSIGNED NULL DEFAULT 0 COMMENT '下载次数',
                            `extra` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '额外信息',
                            `deleted` tinyint(1) NULL DEFAULT 0 COMMENT '删除标记',
                            `file_url` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '完整地址',
                            `file_type` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '文件类型',
                            `deleted_time` bigint(0) NULL DEFAULT NULL COMMENT '删除时间',
                            PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 44 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '文件表' ROW_FORMAT = COMPACT;

CREATE TABLE `ms_source_link`  (
                                   `id` int(0) NOT NULL AUTO_INCREMENT,
                                   `source_type` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '资源类型',
                                   `source_code` bigint(0) NULL DEFAULT NULL COMMENT '资源编号',
                                   `link_type` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '关联类型',
                                   `link_code` bigint(0) NULL DEFAULT NULL COMMENT '关联编号',
                                   `organization_code` bigint(0) NULL DEFAULT NULL COMMENT '组织编码',
                                   `create_by` bigint(0) NULL DEFAULT NULL COMMENT '创建人',
                                   `create_time` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '创建时间',
                                   `sort` int(0) NULL DEFAULT 0 COMMENT '排序',
                                   PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '资源关联表' ROW_FORMAT = COMPACT;

CREATE TABLE `ms_member_account`  (
                                      `id` bigint(0) NOT NULL AUTO_INCREMENT,
                                      `member_code` bigint(0) NULL DEFAULT NULL COMMENT '所属账号id',
                                      `organization_code` bigint(0) NULL DEFAULT NULL COMMENT '所属组织',
                                      `department_code` bigint(0) NULL DEFAULT NULL COMMENT '部门编号',
                                      `authorize` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '角色',
                                      `is_owner` tinyint(1) NULL DEFAULT 0 COMMENT '是否主账号',
                                      `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '姓名',
                                      `mobile` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '手机号码',
                                      `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '邮件',
                                      `create_time` bigint(0) NULL DEFAULT NULL COMMENT '创建时间',
                                      `last_login_time` bigint(0) NULL DEFAULT NULL COMMENT '上次登录时间',
                                      `status` tinyint(1) NULL DEFAULT 0 COMMENT '状态0禁用 1使用中',
                                      `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '描述',
                                      `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '头像',
                                      `position` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '职位',
                                      `department` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '部门',
                                      PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 35 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '组织账号表' ROW_FORMAT = COMPACT;

CREATE TABLE `ms_department`  (
                                  `id` bigint(0) NOT NULL AUTO_INCREMENT,
                                  `organization_code` bigint(0) NULL DEFAULT NULL COMMENT '组织编号',
                                  `name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名称',
                                  `sort` int(0) NULL DEFAULT 0 COMMENT '排序',
                                  `pcode` bigint(0) NULL DEFAULT NULL COMMENT '上级编号',
                                  `icon` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '图标',
                                  `create_time` bigint(0) NULL DEFAULT NULL COMMENT '创建时间',
                                  `path` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '上级路径',
                                  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '部门表' ROW_FORMAT = COMPACT;

CREATE TABLE `ms_department_member`  (
                                         `id` bigint(0) NOT NULL AUTO_INCREMENT,
                                         `department_code` bigint(0) NULL DEFAULT NULL COMMENT '部门id',
                                         `organization_code` bigint(0) NULL DEFAULT NULL COMMENT '组织id',
                                         `account_code` bigint(0) NULL DEFAULT NULL COMMENT '成员id',
                                         `join_time` bigint(0) NULL DEFAULT NULL COMMENT '加入时间',
                                         `is_principal` tinyint(1) NULL DEFAULT NULL COMMENT '是否负责人',
                                         `is_owner` tinyint(1) NULL DEFAULT 0 COMMENT '拥有者',
                                         `authorize` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '角色',
                                         PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 38 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '部门-成员表' ROW_FORMAT = COMPACT;

CREATE TABLE `ms_project_auth_node`  (
                                         `id` bigint(0) UNSIGNED NOT NULL AUTO_INCREMENT,
                                         `auth` bigint(0) UNSIGNED NULL DEFAULT NULL COMMENT '角色ID',
                                         `node` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '节点路径',
                                         PRIMARY KEY (`id`) USING BTREE,
                                         INDEX `index_system_auth_auth`(`auth`) USING BTREE,
                                         INDEX `index_system_auth_node`(`node`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5280 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '项目角色与节点绑定' ROW_FORMAT = COMPACT;

CREATE TABLE `ms_project_node`  (
                                    `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT,
                                    `node` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '节点代码',
                                    `title` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '节点标题',
                                    `is_menu` tinyint(0) UNSIGNED NULL DEFAULT 0 COMMENT '是否可设置为菜单',
                                    `is_auth` tinyint(0) UNSIGNED NULL DEFAULT 1 COMMENT '是否启动RBAC权限控制',
                                    `is_login` tinyint(0) UNSIGNED NULL DEFAULT 1 COMMENT '是否启动登录控制',
                                    `create_at` bigint(0) NULL DEFAULT NULL COMMENT '创建时间',
                                    PRIMARY KEY (`id`) USING BTREE,
                                    INDEX `index_system_node_node`(`node`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 641 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '项目端节点表' ROW_FORMAT = COMPACT;

CREATE TABLE `ms_project_auth`  (
                                    `id` bigint(0) UNSIGNED NOT NULL AUTO_INCREMENT,
                                    `title` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '权限名称',
                                    `status` tinyint(0) UNSIGNED NULL DEFAULT 1 COMMENT '状态(0:禁用,1:启用)',
                                    `sort` smallint(0) UNSIGNED NULL DEFAULT 0 COMMENT '排序权重',
                                    `desc` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '备注说明',
                                    `create_by` bigint(0) UNSIGNED NULL DEFAULT 0 COMMENT '创建人',
                                    `create_at` bigint(0) NULL DEFAULT NULL COMMENT '创建时间',
                                    `organization_code` bigint(0) NULL DEFAULT NULL COMMENT '所属组织',
                                    `is_default` tinyint(1) NULL DEFAULT 0 COMMENT '是否默认',
                                    `type` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '权限类型',
                                    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '项目权限表' ROW_FORMAT = COMPACT;
