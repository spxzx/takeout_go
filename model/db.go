package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"strings"
	"time"
)

var db *gorm.DB
var dbErr error
var dsn = "root:123456@tcp(127.0.0.1:3306)/takeout_go?charset=utf8mb4&parseTime=True&loc=Local"

func InitDb() {
	db, dbErr = gorm.Open(mysql.New(mysql.Config{
		DSN:               dsn,
		DefaultStringSize: 256,
	}), &gorm.Config{
		SkipDefaultTransaction: false,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   logger.Default.LogMode(logger.Info),
	})
	if dbErr != nil {
		log.Fatal("数据库连接失败，请检查参数 ", dbErr)
	}
	if err := db.AutoMigrate(
		&Employee{}, &Category{}, &Dish{}, &DishFlavor{},
	); err != nil {
		log.Fatal("迁移失败！", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("数据池错误，请检查", err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetConnMaxLifetime(8 * time.Hour)
}

type Universal struct {
	CreateTime time.Time `gorm:"autoCreateTime" json:"createTime"`
	UpdateTime time.Time `gorm:"autoUpdateTime" json:"updateTime"`
	CreateUser int       `json:"createUser"`
	UpdateUser int       `json:"updateUser"`
}

// TypeOfModel 用泛型实现类型的自动匹配
type TypeOfModel interface {
	Employee | Category | Dish | DishFlavor |
		DishDto | Setmeal | SetmealDish | User |
		AddressBook | ShoppingCart | Orders
}

// TypeOfListModel 用泛型实现的自动匹配
type TypeOfListModel interface {
	EmployeeList | CategoryList | DishList |
		DishFlavorList | SetmealList | SetmealDishList |
		AddressBookList | ShoppingCartList | OrdersList |
		OrderDetailList
}

// GetOne 通过queryName查询query值对应数据是否存在，只查询单个，类型自动匹配
func GetOne[T TypeOfModel](data *T, queryName []string, query ...any) {
	db.Where(querySplit(queryName), query...).First(&data)
}

// List 通过queryName查询query值对应数据是否存在，返回的是一个切片，类型自动匹配
func List[T TypeOfListModel](list *T, order string, query []string, args ...any) {
	db.Order(order).Where(querySplit(query), args...).Find(&list)
}

// Insert 插入数据，类型自动匹配
func Insert[T TypeOfModel](data T) {
	db.Create(&data)
}

func InsertBatch[T TypeOfListModel](list T) {
	db.Create(&list)
}

func InsertWithList[T TypeOfModel, TL TypeOfListModel](data T, list TL) {
	db.Create(&data)
	if len(list) > 0 {
		db.Create(&list)
	}
}

// Update 更新全部，不包括null对应的字段，类型自动匹配
func Update[T TypeOfModel](data T, updateMap ...map[string]any) {
	if len(updateMap) > 0 {
		db.Model(&data).Updates(updateMap[0])
	} else {
		db.Model(&data).Updates(&data) // 注意 &
	}
}

func UpdateByIds[T TypeOfModel](data T, ids []int, updateMap map[string]any) {
	db.Model(&data).Where("id in ?", ids).Updates(updateMap)
}

func UpdateWith[T TypeOfModel](data T, query string, queryData any, updateMap map[string]any) {
	db.Model(&data).Where(query+" = ?", queryData).Updates(updateMap)
}

// DeleteByIds 指定主键进行删除
func DeleteByIds[T TypeOfModel](data T, ids ...int) {
	db.Delete(&data, ids)
}

// DeleteBatch 批量删除，不指定主键，指定删除queryString对应的数据
// queryString = 字段名  |  query = 字段值
func DeleteBatch[T TypeOfModel](data T, queryString string, query any) {
	db.Delete(&data, queryString, query)
}

// Page 分页查询，只能模糊匹配一个字段
func Page[T TypeOfListModel](page int, pageSize int, searchMap map[string]string, list *T, total *int64, order string) {
	if page == 0 {
		page = -1
	}
	if pageSize == 0 {
		pageSize = -1
	}
	offset := (page - 1) * pageSize
	if page == -1 && pageSize == -1 {
		offset = -1
	}
	var query string
	var arg string
	for k, v := range searchMap {
		query = k + " LIKE ?"
		arg = "%" + v + "%"
	}
	if len(searchMap) > 0 {
		db.Order(order).Where(query, arg).Limit(pageSize).Offset(offset).Find(&list)
		db.Model(&list).Where(query, arg).Count(total)
	} else {
		db.Order(order).Limit(pageSize).Offset(offset).Find(&list)
		db.Model(&list).Count(total)
	}
}

func PageWhere[T TypeOfListModel](page int, pageSize int, list *T, total *int64, order string, query []string, args ...any) {
	if page == 0 {
		page = -1
	}
	if pageSize == 0 {
		pageSize = -1
	}
	offset := (page - 1) * pageSize
	if page == -1 && pageSize == -1 {
		offset = -1
	}
	db.Order(order).Where(querySplit(query), args).Limit(pageSize).Offset(offset).Find(&list)
	db.Model(&list).Where(querySplit(query), args).Count(total)
}

func querySplit(queryName []string) string {
	var sb strings.Builder
	length := len(queryName)
	for i := 0; i < length-1; i++ {
		sb.WriteString(queryName[i])
		sb.WriteString(" = ? AND ")
	}
	sb.WriteString(queryName[length-1])
	sb.WriteString(" = ?")
	return sb.String()
}
