# MS 项目

## 项目概述

这是一个前后端分离的项目，前端使用Vue.js框架，后端使用Go语言开发。项目包含多个微服务模块，使用Docker Compose进行容器化部署。

## 环境要求

- Docker 20.10+ 
- Docker Compose 1.29+
- Node.js 16+ (前端)
- Go 1.18+ (后端)

## 前端项目

### 安装依赖
```bash
cd front
npm install
```

### 开发模式
```bash
npm run serve
```

### 生产构建
```bash
npm run build
```

## 后端项目

### 运行服务
```bash
cd serve
run.bat
```

### Docker Compose 部署
```bash
docker-compose up -d
```

## 服务端口

- MySQL: 3309
- Redis: 6379
- Nacos: 8848
- Jaeger: 16686
- Kafka UI: 9000
- Elasticsearch: 9200
- Kibana: 5601

## 配置说明

环境变量配置在`.env`文件中，请根据实际情况修改。