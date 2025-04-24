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




INSERT INTO `ms_project_auth`(`id`, `title`, `status`, `sort`, `desc`, `create_by`, `create_at`, `organization_code`, `is_default`, `type`) VALUES (1, '管理员', 1, 0, '管理员', 0, replace(unix_timestamp(now(3)),'.',''), null, 0, 'admin');
INSERT INTO `ms_project_auth`(`id`, `title`, `status`, `sort`, `desc`, `create_by`, `create_at`, `organization_code`, `is_default`, `type`) VALUES (2, '成员', 1, 0, '成员', 0, replace(unix_timestamp(now(3)),'.',''), null, 1, 'member');



INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3097, 1, 'project');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3098, 1, 'project/account');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3099, 1, 'project/account/index');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3100, 1, 'project/account/auth');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3101, 1, 'project/account/add');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3102, 1, 'project/account/edit');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3103, 1, 'project/account/del');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3104, 1, 'project/account/forbid');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3105, 1, 'project/account/resume');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3106, 1, 'project/auth');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3107, 1, 'project/auth/index');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3108, 1, 'project/auth/apply');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3109, 1, 'project/auth/add');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3110, 1, 'project/auth/edit');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3111, 1, 'project/auth/forbid');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3112, 1, 'project/auth/resume');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3113, 1, 'project/auth/setdefault');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3114, 1, 'project/auth/del');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3115, 1, 'project/department');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3116, 1, 'project/department/index');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3117, 1, 'project/department/read');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3118, 1, 'project/department/save');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3119, 1, 'project/department/edit');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3120, 1, 'project/department/delete');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3121, 1, 'project/department_member');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3122, 1, 'project/department_member/index');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3123, 1, 'project/department_member/searchinvitemember');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3124, 1, 'project/department_member/invitemember');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3125, 1, 'project/department_member/removemember');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3126, 1, 'project/index');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3127, 1, 'project/index/index');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3128, 1, 'project/index/changecurrentorganization');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3129, 1, 'project/index/systemconfig');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3130, 1, 'project/index/info');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3131, 1, 'project/index/editpersonal');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3132, 1, 'project/index/editpassword');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3133, 1, 'project/index/uploadimg');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3134, 1, 'project/index/uploadavatar');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3135, 1, 'project/menu');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3136, 1, 'project/menu/menu');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3137, 1, 'project/menu/menuadd');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3138, 1, 'project/menu/menuedit');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3139, 1, 'project/menu/menuforbid');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3140, 1, 'project/menu/menuresume');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3141, 1, 'project/menu/menudel');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3142, 1, 'project/node');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3143, 1, 'project/node/index');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3144, 1, 'project/node/alllist');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3145, 1, 'project/node/clear');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3146, 1, 'project/node/save');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3147, 1, 'project/notify');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3148, 1, 'project/notify/index');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3149, 1, 'project/notify/noreads');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3150, 1, 'project/notify/setreadied');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3151, 1, 'project/notify/batchdel');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3152, 1, 'project/notify/read');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3153, 1, 'project/notify/delete');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3154, 1, 'project/organization');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3155, 1, 'project/organization/index');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3156, 1, 'project/organization/save');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3157, 1, 'project/organization/read');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3158, 1, 'project/organization/edit');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3159, 1, 'project/organization/delete');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3160, 1, 'project/project');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3161, 1, 'project/project/index');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3162, 1, 'project/project/selflist');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3163, 1, 'project/project/save');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3164, 1, 'project/project/read');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3165, 1, 'project/project/edit');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3166, 1, 'project/project/uploadcover');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3167, 1, 'project/project/recycle');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3168, 1, 'project/project/recovery');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3169, 1, 'project/project/archive');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3170, 1, 'project/project/recoveryarchive');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3171, 1, 'project/project/quit');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3172, 1, 'project/project_collect');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3173, 1, 'project/project_collect/collect');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3174, 1, 'project/project_member');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3175, 1, 'project/project_member/index');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3176, 1, 'project/project_member/searchinvitemember');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3177, 1, 'project/project_member/invitemember');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3178, 1, 'project/project_template');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3179, 1, 'project/project_template/index');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3180, 1, 'project/project_template/save');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3181, 1, 'project/project_template/uploadcover');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3182, 1, 'project/project_template/edit');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3183, 1, 'project/project_template/delete');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3184, 1, 'project/task');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3185, 1, 'project/task/index');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3186, 1, 'project/task/selflist');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3187, 1, 'project/task/read');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3188, 1, 'project/task/save');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3189, 1, 'project/task/taskdone');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3190, 1, 'project/task/assigntask');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3191, 1, 'project/task/sort');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3192, 1, 'project/task/createcomment');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3193, 1, 'project/task/edit');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3194, 1, 'project/task/like');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3195, 1, 'project/task/star');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3196, 1, 'project/task/recycle');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3197, 1, 'project/task/recovery');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3198, 1, 'project/task/delete');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3199, 1, 'project/task_log');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3200, 1, 'project/task_log/index');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3201, 1, 'project/task_log/getlistbyselfproject');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3202, 1, 'project/task_member');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3203, 1, 'project/task_member/index');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3204, 1, 'project/task_member/searchinvitemember');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3205, 1, 'project/task_member/invitemember');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3206, 1, 'project/task_member/invitememberbatch');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3207, 1, 'project/task_stages');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3208, 1, 'project/task_stages/index');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3209, 1, 'project/task_stages/tasks');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3210, 1, 'project/task_stages/sort');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3211, 1, 'project/task_stages/save');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3212, 1, 'project/task_stages/edit');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3213, 1, 'project/task_stages/delete');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3214, 1, 'project/task_stages_template');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3215, 1, 'project/task_stages_template/index');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3216, 1, 'project/task_stages_template/save');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3217, 1, 'project/task_stages_template/edit');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3218, 1, 'project/task_stages_template/delete');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3219, 2, 'project/account/index');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3220, 2, 'project/auth/index');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3221, 2, 'project/index/index');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3222, 2, 'project/index');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3223, 2, 'project/index/changecurrentorganization');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3224, 2, 'project/index/systemconfig');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3225, 2, 'project/index/info');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3226, 2, 'project/index/editpersonal');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3227, 2, 'project/index/editpassword');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3228, 2, 'project/index/uploadimg');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3229, 2, 'project/index/uploadavatar');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3230, 2, 'project/menu/menu');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3231, 2, 'project/node/index');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3232, 2, 'project/node/alllist');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3233, 2, 'project/notify/index');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3234, 2, 'project/notify');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3235, 2, 'project/notify/noreads');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3236, 2, 'project/notify/setreadied');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3237, 2, 'project/notify/batchdel');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3238, 2, 'project/notify/read');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3239, 2, 'project/notify/delete');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3240, 2, 'project/organization/index');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3241, 2, 'project/organization');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3242, 2, 'project/organization/save');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3243, 2, 'project/organization/read');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3244, 2, 'project/organization/edit');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3245, 2, 'project/organization/delete');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3246, 2, 'project/project/index');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3247, 2, 'project/project/read');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3248, 2, 'project/project_collect/collect');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3249, 2, 'project/project_collect');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3250, 2, 'project/project_member/index');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3251, 2, 'project/project_template/index');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3252, 2, 'project/task/index');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3253, 2, 'project/task/read');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3254, 2, 'project/task/save');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3255, 2, 'project/task/taskdone');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3256, 2, 'project/task/assigntask');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3257, 2, 'project/task/sort');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3258, 2, 'project/task/createcomment');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3259, 2, 'project/task/like');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3260, 2, 'project/task/star');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3261, 2, 'project/task_log/index');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3262, 2, 'project/task_log');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3263, 2, 'project/task_log/getlistbyselfproject');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3264, 2, 'project/task_member/index');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3265, 2, 'project/task_member/searchinvitemember');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3266, 2, 'project/task_stages/index');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3267, 2, 'project/task_stages/tasks');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3268, 2, 'project/task_stages/sort');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3269, 2, 'project/task_stages_template/index');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3270, 2, 'project/department/index');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3271, 2, 'project/department/read');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3272, 2, 'project/department_member/index');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3273, 2, 'project/department_member/searchinvitemember');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3274, 2, 'project/project/selflist');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3275, 2, 'project/project/save');
INSERT INTO `ms_project_auth_node`(`id`, `auth`, `node`) VALUES (3276, 2, 'project/task/selflist');


INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (360, 'project', '项目管理模块', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (361, 'project/index/info', '详情', 0, 0, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (362, 'project/index', '基础版块', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (363, 'project/index/index', '框架布局', 0, 0, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (364, 'project/index/systemconfig', '系统信息', 0, 0, 0, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (365, 'project/index/editpersonal', '修改个人资料', 0, 0, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (366, 'project/index/uploadavatar', '上传头像', 0, 0, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (370, 'project/account', '账号管理', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (371, 'project/account/index', '账号列表', 0, 0, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (372, 'project/organization/index', '组织列表', 0, 0, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (373, 'project/organization/save', '创建组织', 0, 0, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (374, 'project/organization/read', '组织信息', 0, 0, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (375, 'project/organization/edit', '编辑组织', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (376, 'project/organization/delete', '删除组织', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (377, 'project/organization', '组织管理', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (388, 'project/auth/index', '权限列表', 0, 0, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (389, 'project/auth/add', '添加权限角色', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (390, 'project/auth/edit', '编辑权限', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (391, 'project/auth/forbid', '禁用权限', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (392, 'project/auth/resume', '启用权限', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (393, 'project/auth/del', '删除权限', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (394, 'project/auth', '访问授权', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (395, 'project/auth/apply', '应用权限', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (396, 'project/notify/index', '通知列表', 0, 0, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (397, 'project/notify/noreads', '未读通知', 0, 0, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (399, 'project/notify/read', '通知信息', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (401, 'project/notify/delete', '删除通知', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (402, 'project/notify', '通知管理', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (434, 'project/account/auth', '授权管理', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (435, 'project/account/add', '添加账号', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (436, 'project/account/edit', '编辑账号', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (437, 'project/account/del', '删除账号', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (438, 'project/account/forbid', '禁用账号', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (439, 'project/account/resume', '启用账号', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (498, 'project/notify/setreadied', '设置已读', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (499, 'project/notify/batchdel', '批量删除', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (500, 'project/auth/setdefault', '设置默认权限', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (501, 'project/department', '部门管理', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (502, 'project/department/index', '部门列表', 0, 0, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (503, 'project/department/read', '部门信息', 0, 0, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (504, 'project/department/save', '创建部门', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (505, 'project/department/edit', '编辑部门', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (506, 'project/department/delete', '删除部门', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (507, 'project/department_member', '部门成员管理', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (508, 'project/department_member/index', '部门成员列表', 0, 0, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (509, 'project/department_member/searchinvitemember', '搜索部门成员', 0, 0, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (510, 'project/department_member/invitemember', '添加部门成员', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (511, 'project/department_member/removemember', '移除部门成员', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (512, 'project/index/changecurrentorganization', '切换当前组织', 0, 0, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (513, 'project/index/editpassword', '修改密码', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (514, 'project/index/uploadimg', '上传图片', 0, 0, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (515, 'project/menu', '菜单管理', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (516, 'project/menu/menu', '菜单列表', 0, 0, 0, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (517, 'project/menu/menuadd', '添加菜单', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (518, 'project/menu/menuedit', '编辑菜单', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (519, 'project/menu/menuforbid', '禁用菜单', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (520, 'project/menu/menuresume', '启用菜单', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (521, 'project/menu/menudel', '删除菜单', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (522, 'project/node', '节点管理', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (523, 'project/node/index', '节点列表', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (524, 'project/node/alllist', '全部节点列表', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (525, 'project/node/clear', '清理节点', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (526, 'project/node/save', '编辑节点', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (527, 'project/project', '项目管理', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (528, 'project/project/index', '项目列表', 0, 0, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (529, 'project/project/selflist', '个人项目列表', 0, 0, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (530, 'project/project/save', '创建项目', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (531, 'project/project/read', '项目信息', 0, 0, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (532, 'project/project/edit', '编辑项目', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (533, 'project/project/uploadcover', '上传项目封面', 0, 0, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (534, 'project/project/recycle', '项目放入回收站', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (535, 'project/project/recovery', '恢复项目', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (536, 'project/project/archive', '归档项目', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (537, 'project/project/recoveryarchive', '取消归档项目', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (538, 'project/project/quit', '退出项目', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (539, 'project/project_collect', '项目收藏管理', 0, 0, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (540, 'project/project_collect/collect', '收藏项目', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (541, 'project/project_member', '项目成员管理', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (542, 'project/project_member/index', '项目成员列表', 0, 0, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (543, 'project/project_member/searchinvitemember', '搜索项目成员', 0, 0, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (544, 'project/project_member/invitemember', '邀请项目成员', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (545, 'project/project_template', '项目模板管理', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (546, 'project/project_template/index', '项目模板列表', 0, 0, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (547, 'project/project_template/save', '创建项目模板', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (548, 'project/project_template/uploadcover', '上传项目模板封面', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (549, 'project/project_template/edit', '编辑项目模板', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (550, 'project/project_template/delete', '删除项目模板', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (551, 'project/task/index', '任务列表', 0, 0, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (552, 'project/task/selflist', '个人任务列表', 0, 0, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (553, 'project/task/read', '任务信息', 0, 0, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (554, 'project/task/save', '创建任务', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (555, 'project/task/taskdone', '更改任务状态', 0, 0, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (556, 'project/task/assigntask', '指派任务执行者', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (557, 'project/task/sort', '任务排序', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (558, 'project/task/createcomment', '发表任务评论', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (559, 'project/task/edit', '编辑任务', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (560, 'project/task/like', '点赞任务', 0, 0, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (561, 'project/task/star', '收藏任务', 0, 0, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (562, 'project/task/recycle', '移动任务到回收站', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (563, 'project/task/recovery', '恢复任务', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (564, 'project/task/delete', '删除任务', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (565, 'project/task', '任务管理', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (569, 'project/task_member/index', '任务成员列表', 0, 0, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (570, 'project/task_member/searchinvitemember', '搜索任务成员', 0, 0, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (571, 'project/task_member/invitemember', '添加任务成员', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (572, 'project/task_member/invitememberbatch', '批量添加任务成员', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (573, 'project/task_member', '任务成员管理', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (574, 'project/task_stages', '任务分组管理', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (575, 'project/task_stages/index', '任务分组列表', 0, 0, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (576, 'project/task_stages/tasks', '任务分组任务列表', 0, 0, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (577, 'project/task_stages/sort', '任务分组排序', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (578, 'project/task_stages/save', '添加任务分组', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (579, 'project/task_stages/edit', '编辑任务分组', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (580, 'project/task_stages/delete', '删除任务分组', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (581, 'project/task_stages_template/index', '任务分组模板列表', 0, 0, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (582, 'project/task_stages_template/save', '创建任务分组模板', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (583, 'project/task_stages_template/edit', '编辑任务分组模板', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (584, 'project/task_stages_template/delete', '删除任务分组模板', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (585, 'project/task_stages_template', '任务分组模板管理', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (587, 'project/project_member/removemember', '移除项目成员', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (588, 'project/task/datetotalforproject', '任务统计', 0, 0, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (589, 'project/task/tasksources', '任务资源列表', 0, 0, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (590, 'project/file', '文件管理', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (591, 'project/file/index', '文件列表', 0, 0, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (592, 'project/file/read', '文件详情', 0, 0, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (593, 'project/file/uploadfiles', '上传文件', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (594, 'project/file/edit', '编辑文件', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (595, 'project/file/recycle', '文件移至回收站', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (596, 'project/file/recovery', '恢复文件', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (597, 'project/file/delete', '删除文件', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (598, 'project/project/getlogbyselfproject', '项目概况', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (599, 'project/source_link', '资源关联管理', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (600, 'project/source_link/delete', '取消关联', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (601, 'project/task/tasklog', '任务动态', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (602, 'project/task/recyclebatch', '批量移动任务到回收站', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (603, 'project/invite_link', '邀请链接管理', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (604, 'project/invite_link/save', '创建邀请链接', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (605, 'project/task/setprivate', '设置任务隐私模式', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (606, 'project/account/read', '账号信息', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (607, 'project/task/batchassigntask', '批量指派任务', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (608, 'project/task/tasktotags', '任务标签', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (609, 'project/task/settag', '设置任务标签', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (610, 'project/task_tag', '任务标签管理', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (611, 'project/task_tag/index', '任务标签列表', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (612, 'project/task_tag/save', '创建任务标签', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (613, 'project/task_tag/edit', '编辑任务标签', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (614, 'project/task_tag/delete', '删除任务标签', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (615, 'project/project_features', '项目版本库管理', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (616, 'project/project_features/index', '版本库列表', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (617, 'project/project_features/save', '添加版本库', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (618, 'project/project_features/edit', '编辑版本库', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (619, 'project/project_features/delete', '删除版本库', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (620, 'project/project_version', '项目版本管理', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (621, 'project/project_version/index', '项目版本列表', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (622, 'project/project_version/save', '添加项目版本', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (623, 'project/project_version/edit', '编辑项目版本', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (624, 'project/project_version/changestatus', '更改项目版本状态', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (625, 'project/project_version/read', '项目版本详情', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (626, 'project/project_version/addversiontask', '关联项目版本任务', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (627, 'project/project_version/removeversiontask', '移除项目版本任务', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (628, 'project/project_version/delete', '删除项目版本', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (629, 'project/task/getlistbytasktag', '标签任务列表', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (630, 'project/task_workflow', '任务流转管理', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (631, 'project/task_workflow/index', '任务流转列表', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (632, 'project/task_workflow/save', '添加任务流转', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (633, 'project/task_workflow/edit', '编辑任务流转', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (634, 'project/task_workflow/delete', '删除任务流转', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (635, 'project/department_member/detail', '部门成员详情', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (636, 'project/department_member/uploadfile', '上传头像', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (637, 'project/task/savetaskworktime', '保存任务流转', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (638, 'project/task/edittaskworktime', '编辑任务流转', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (639, 'project/task/deltaskworktime', '删除任务流转', 0, 1, 1, 1673277965322);
INSERT INTO  `ms_project_node`(`id`, `node`, `title`, `is_menu`, `is_auth`, `is_login`, `create_at`) VALUES (640, 'project/task/uploadfile', '上传文件', 0, 1, 1, 1673277965322);
