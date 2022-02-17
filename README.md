# 选排课系统

## 项目模块

系统包括以下模块
- 登录模块
  - 考察登录的设计与实现，对 HTTP 协议的理解。
    - 账密登录
    - Cookie Session
- 成员模块
  - 考察工程实现能力。
    - CURD 及对数据库的操作
    - 参数校验
      - 参数长度
      - 弱密码校验
    - 权限判断
- 排课模块
  - 主要考察算法（二分图匹配）的实现。
- 抢课模块
  - 主要考察简单秒杀场景的设计。

## 项目结构
```
.
├── Controllers
│   ├── authController.go
│   ├── courseController.go
│   ├── scheduleCourseController.go
│   ├── secKillController.go
│   └── userControllers.go
├── Dao
│   ├── CapDao
│   │   └── capDao.go
│   ├── DBAccessor
│   │   └── DBAccessor.go
│   ├── RedisAccessor
│   │   └── RedisAccessor.go
│   ├── TCourseDao
│   │   ├── TCourseDao.go
│   │   └── TCourseDaoTest
│   │       └── test.go
│   ├── TMemberDao
│   │   ├── TmemberDao.go
│   │   └── TMemberDaoTest
│   │       └── test.go
│   ├── UserCourseDao
│   │   └── userCourseDao.go
│   └── UserDao
│       ├── userDao.go
│       └── UserDaoTest
│           └── test.go
├── go.mod
├── go.sum
├── main.go
├── README.md
├── Routers
│   └── router.go
├── Service
│   ├── ScheduleCourse
│   │   └── scheduleCourse.go
│   ├── SecKill
│   │   └── residueCheck.go
│   └── UserService
│       └── checkFunc.go
└── Types
    └── types.go
```

## 项目及接口文档
[https://bytedancecampus1.feishu.cn/docs/doccnens5GECVM1l9nGaRLUEIn3](文档)
