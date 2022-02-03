# OceanCT和数据库底层操作

## Preparations

项目是基于gorm、gin框架的，使用以下命令来导入它们
```
go get -u github.com/jinzhu/gorm
go get -u -v github.com/gin-gonic/gin
```
接下来，为了方便使用，OceanCT建议您使用git pull命令在我的代码的基础上进一步完成剩下的内容。

当然，为了将我的代码与您需要完成的代码分隔开，您也可以使用如下代码
```
go get -u github.com/ByteDanceBackendProject18/Project
```

## 如何使用我提供的接口来完成接下来的部分？

### 整体项目概况

在我的预想中，项目的代码是分在Dao(数据库底层操作层),Service(业务逻辑层),APIController(API层)，Types(通用数据结构定义层)，在一期中，我完成了Dao层，其他组员应该完成部分Service层和API层，都需要单元测试

### Dao层实现思路

在我提供的Dao层实现接口里，非常明显的可以分为两类，Dao类接口和非Dao类接口，如下的代码可以简单地说明他们的区别
```
var course Types.TCourse
var newCourse Types.Tourse
var courseID string
// 使用Dao类接口完成CRUD
dao := TCourseDao.MakeTCourseDao(course)
TCourseDao.InsertCourseByDao(dao)
TCourseDao.UpdateCourseByDao(dao, newCourse)
TCourseDao.DeleteCourseByDao(dao)
// 可以看到使用Dao类接口的关键在于产生*Dao变量
// （由MakeTCourseDao函数完成）
// 然后您对它的CRUD就可以以它为句柄
// 无需考虑数据库中的其他数据对CRUD的影响

// 使用非Dao类接口完成CRUD
TCourseDao.InsertCourse(course)
TCourseDao.FindCourseByID(courseID)
TCourseDao.UpdateCourseByID(courseID,newCourse)
TCourseDao.DeleteCourseByName(course.Name)
// 这些接口的适用性更广
// 连续对同一对象的操作的效率或许会比上面的接口稍差
// 使用时需要考虑数据库中其他数据的影响：
// 如：在一种我们都不想见到的情况下，CourseID重复了
// 这时所有的CRUD就不可避免的互相影响
```
需要注意的是，根据组长的告知，所有的ID不是数据库层自己生成的，而是Service层调用时传进来的，因此数据库层不会对ID是否重复进行检查

例如：在插入用户时，不允许有重复的用户名，那么Service层使用时应该先使用Find类函数进行查找，根据返回结果判断是否重复，如果不重复再使用Insert类函数插入数据库,而非直接使用Insert类函数根据其返回状态判断

## 尾言

使用时建议参考函数上的注解（由于篇幅限制，这里不会对每个函数进行列出讲解，如果实在需要或许会写API文档，个人认为看函数名就可以正常使用）

尽管进行了单元测试，或许仍有不足之处，还望谅解
