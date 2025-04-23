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

INSERT INTO `ms_project_template`(`id`, `name`, `description`, `sort`, `create_time`, `organization_code`, `cover`, `member_code`, `is_system`) VALUES (11, '产品进展', '适用于互联网产品人员对产品计划、跟进及发布管理', 0, 1670904236057, 17, 'https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fbpic.51yuansu.com%2Fpic3%2Fcover%2F01%2F91%2F92%2F5982adf6c88ea_610.jpg&refer=http%3A%2F%2Fbpic.51yuansu.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=auto?sec=1673496114&t=956c5614481fedea97794e161deddb00', NULL, 1);
INSERT INTO `ms_project_template`(`id`, `name`, `description`, `sort`, `create_time`, `organization_code`, `cover`, `member_code`, `is_system`) VALUES (12, '需求管理', '适用于产品部门对需求的收集、评估及反馈管理', 0, 1670904236057, 17, 'https://img0.baidu.com/it/u=437485064,4277010738&fm=253&fmt=auto&app=138&f=JPEG?w=610&h=491', NULL, 1);
INSERT INTO `ms_project_template`(`id`, `name`, `description`, `sort`, `create_time`, `organization_code`, `cover`, `member_code`, `is_system`) VALUES (13, '机械制造', '适用于制造商对图纸设计及制造安装的工作流程管理', 0, 1670904236057, 17, 'https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fbpic.51yuansu.com%2Fpic2%2Fcover%2F00%2F38%2F93%2F5812ca7a24020_610.jpg&refer=http%3A%2F%2Fbpic.51yuansu.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=auto?sec=1673496114&t=6d03fb91b230058fc43f1b7ae00f73e3', NULL, 1);
INSERT INTO `ms_project_template`(`id`, `name`, `description`, `sort`, `create_time`, `organization_code`, `cover`, `member_code`, `is_system`) VALUES (19, 'OKR 管理', '适用于团队的 OKR 管理', 0, 1670904236057, 17, 'https://img2.baidu.com/it/u=2241642503,1613686234&fm=253&fmt=auto&app=138&f=JPEG?w=603&h=500', 1015, 0);

CREATE TABLE `ms_task_stages_template`  (
                                            `id` int(0) NOT NULL AUTO_INCREMENT,
                                            `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '类型名称',
                                            `project_template_code` int(0) NULL DEFAULT 0 COMMENT '项目id',
                                            `create_time` bigint(0) NULL DEFAULT NULL,
                                            `sort` int(0) NULL DEFAULT 0,
                                            PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 84 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '任务列表模板表' ROW_FORMAT = COMPACT;

INSERT INTO `ms_task_stages_template`(`id`, `name`, `project_template_code`, `create_time`, `sort`) VALUES (61, '待处理', 19, 1670904236057, 1);
INSERT INTO `ms_task_stages_template`(`id`, `name`, `project_template_code`, `create_time`, `sort`) VALUES (62, '进行中', 19, 1670904236057, 0);
INSERT INTO `ms_task_stages_template`(`id`, `name`, `project_template_code`, `create_time`, `sort`) VALUES (63, '已完成', 19, 1670904236057, 0);
INSERT INTO `ms_task_stages_template`(`id`, `name`, `project_template_code`, `create_time`, `sort`) VALUES (65, '协议签订', 13, 1670904236057, 0);
INSERT INTO `ms_task_stages_template`(`id`, `name`, `project_template_code`, `create_time`, `sort`) VALUES (66, '图纸设计', 13, 1670904236057, 0);
INSERT INTO `ms_task_stages_template`(`id`, `name`, `project_template_code`, `create_time`, `sort`) VALUES (67, '评审及打样', 13, 1670904236057, 0);
INSERT INTO `ms_task_stages_template`(`id`, `name`, `project_template_code`, `create_time`, `sort`) VALUES (68, '构件采购', 13, 1670904236057, 0);
INSERT INTO `ms_task_stages_template`(`id`, `name`, `project_template_code`, `create_time`, `sort`) VALUES (69, '制造安装', 13, 1670904236057, 0);
INSERT INTO `ms_task_stages_template`(`id`, `name`, `project_template_code`, `create_time`, `sort`) VALUES (70, '内部检验', 13, 1670904236057, 0);
INSERT INTO `ms_task_stages_template`(`id`, `name`, `project_template_code`, `create_time`, `sort`) VALUES (71, '验收', 13, 1670904236057, 0);
INSERT INTO `ms_task_stages_template`(`id`, `name`, `project_template_code`, `create_time`, `sort`) VALUES (72, '需求收集', 12, 1670904236057, 0);
INSERT INTO `ms_task_stages_template`(`id`, `name`, `project_template_code`, `create_time`, `sort`) VALUES (73, '评估确认', 12, 1670904236057, 0);
INSERT INTO `ms_task_stages_template`(`id`, `name`, `project_template_code`, `create_time`, `sort`) VALUES (74, '需求暂缓', 12, 1670904236057, 0);
INSERT INTO `ms_task_stages_template`(`id`, `name`, `project_template_code`, `create_time`, `sort`) VALUES (75, '研发中', 12, 1670904236057, 0);
INSERT INTO `ms_task_stages_template`(`id`, `name`, `project_template_code`, `create_time`, `sort`) VALUES (76, '内测中', 12, 1670904236057, 0);
INSERT INTO `ms_task_stages_template`(`id`, `name`, `project_template_code`, `create_time`, `sort`) VALUES (77, '通知用户', 12, 1670904236057, 0);
INSERT INTO `ms_task_stages_template`(`id`, `name`, `project_template_code`, `create_time`, `sort`) VALUES (78, '已完成&归档', 12, 1670904236057, 0);
INSERT INTO `ms_task_stages_template`(`id`, `name`, `project_template_code`, `create_time`, `sort`) VALUES (79, '产品计划', 11, 1670904236057, 0);
INSERT INTO `ms_task_stages_template`(`id`, `name`, `project_template_code`, `create_time`, `sort`) VALUES (80, '即将发布', 11, 1670904236057, 0);
INSERT INTO `ms_task_stages_template`(`id`, `name`, `project_template_code`, `create_time`, `sort`) VALUES (81, '测试', 11, 1670904236057, 0);
INSERT INTO `ms_task_stages_template`(`id`, `name`, `project_template_code`, `create_time`, `sort`) VALUES (82, '准备发布', 11, 1670904236057, 0);
INSERT INTO `ms_task_stages_template`(`id`, `name`, `project_template_code`, `create_time`, `sort`) VALUES (83, '发布成功', 11, 1670904236057, 0);

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


