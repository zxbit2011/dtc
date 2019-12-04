package main

import (
	"fmt"
	"strings"
)

const (
	DatabaseOracle    = "Oracle"
	DatabaseMySQL     = "MySQL"
	DatabaseSQLServer = "SQLServer"
)

//表定义
type TableMetaInfo struct {
	DatabaseType string            //数据库类型
	TableSchema  string            //数据库
	TableName    string            //表名
	COMMENT      string            //表备注
	Columns      []TableColumnInfo //表字段
	Index        []TableIndex      //表索引
	CurrentLen   int64             //
}

//表字段属性
type TableColumnInfo struct {
	Name          string //名称
	Type          string //类型
	Length        int64  //长度
	Precision     int64  //大小
	Scale         int64  //比例
	IsPK          bool   //是否为主键
	IsNull        bool   //是否为空
	ColumnComment string //字段备注
}

// FieldType 字段类型长度精度的描述
type FieldType struct {
	HasLength         bool // 是否存在长度概念
	HasPrecisionScale bool //是否存在精度概念
}

// TableIndex 表索引
type TableIndex struct {
	KeyName    string   //索引键名称
	Type       string   //索引类型，MYSQL：（FULLTEXT、NORMAL、SPATIAL（未用）、UNIQUE），Oracle：（Normal、Bitmap、Unique）
	Func       string   //索引方法，BTREE、HASH、RTREE
	ColumnName []string //索引关联的字段，排序顺序
	Comment    string   //索引备注
}

// ToMySQL 转换MySQL表结构SQL脚本
func (t TableMetaInfo) ToMySQL() (sql string, err error) {
	switch t.DatabaseType {
	case DatabaseOracle:
		return t.oracleToMysql()
	case DatabaseSQLServer:
		return t.sqlServerToMysql()
	case DatabaseMySQL:
		return t.mysqlToMysql()
	default:
		err = fmt.Errorf("ToMySQL 不支持的数据类型：%s", t.DatabaseType)
		return
	}
	return
}

func (t TableMetaInfo) oracleToMysql() (scriptSQL string, err error) {
	//字段类型转换
	for _, v := range t.Columns {
		switch v.Type {
		case "BFILE":
			v.Type = "VARCHAR"
			v.Length = 255
			break
		case "BINARY_FLOAT":
			v.Type = "FLOAT"
			break
		case "BLOB", "LONG RAW":
			v.Type = "LONGBLOB"
			break
		case "CHAR", "CHARACTER":
			if v.Length > 255 {
				v.Type = "VARCHAR"
			}
			break
		case "DATE":
			v.Type = "DATETIME"
			break
		case "DECIMAL", "DEC":
			v.Type = "DECIMAL"
			break
		case "INTEGER", "INT":
			break
		case "INTERVAL YEAR TO MONTH", "INTERVAL DAY TO SECOND":
			v.Type = "VARCHAR"
			v.Length = 30
			break
		case "NCHAR":
			if v.Length <= 255 {
				v.Type = "CHAR"
			} else {
				v.Type = "VARCHAR"
			}
			break
		case "NCHAR VARYING":
			v.Type = "VARCHAR"
			v.Length = 4000
			break
		case "NUMBER":
			if v.Scale <= 0 {
				if v.Precision < 3 {
					v.Type = "TINYINT"
				} else if v.Precision >= 3 && v.Precision < 5 {
					v.Type = "SMALLINT"
				} else if v.Precision >= 5 && v.Precision < 9 {
					v.Type = "INT"
				} else if v.Precision >= 9 && v.Precision < 19 {
					v.Type = "BIGINT"
				} else if v.Precision >= 19 && v.Precision <= 38 {
					v.Type = "DECIMAL"
				} else {
					v.Type = "DOUBLE"
				}
			} else {
				v.Type = "DECIMAL"
			}
			break
		case "NUMERIC":
			//与MySQL一致
			break
		case "NVARCHAR2":
			v.Type = "VARCHAR"
			v.Length = 4000
			break
		case "RAW":
			if v.Length <= 255 {
				v.Type = "BINARY"
			} else {
				v.Type = "VARBINARY"
			}
			break
		case "DOUBLE PRECISION", "FLOAT", "BINARY_DOUBLE", "REAL":
			v.Type = "DOUBLE"
			break
		case "ROWID":
			v.Type = "CHAR"
			v.Length = 10
			break
		case "SMALLINT":
			v.Type = "DECIMAL"
			v.Length = 38
			break
		case "TIMESTAMP", "TIMESTAMP WITH TIME ZONE", "TIMESTAMP WITH LOCAL TIME ZONE":
			v.Type = "DATETIME"
			break
		case "UROWID", "VARCHAR", "VARCHAR2":
			v.Type = "VARCHAR"
			break
		case "CLOB", "LONG", "NCLOB", "XMLTYPE":
			v.Type = "LONGTEXT"
			break
		}
	}
	return t.mysqlToMysql()
}

func (t TableMetaInfo) sqlServerToMysql() (scriptSQL string, err error) {
	//字段类型转换
	for _, v := range t.Columns {
		switch v.Type {
		case "BIGINT", "BINARY", "DATE", "DECIMAL", "INT", "INTEGER", "NUMERIC", "REAL", "SMALLINT", "TIME", "TINYINT":
			break
		case "BIT":
			v.Type = "TINYINT"
			break
		case "CHAR":
			if v.Length > 255 {
				v.Type = "TEXT"
			}
			break
		case "DATETIME":
			v.Length = 3
			break
		case "DATETIME2", "DATETIMEOFFSET":
			v.Type = "DATETIME"
			break
		case "DOUBLE PRECISION", "FLOAT":
			v.Type = "DOUBLE"
			break
		case "IMAGE":
			v.Type = "LONGBLOB"
			break
		case "MONEY":
			v.Type = "DECIMAL"
			v.Precision = 15
			v.Scale = 4
			break
		case "NCHAR":
			if v.Length > 255 {
				v.Type = "TEXT"
			} else {
				v.Type = "CHAR"
			}
			break
		case "NTEXT":
			v.Type = "LONGTEXT"
			break
		case "NVARCHAR":
			if v.Length > 4000 {
				v.Type = "VARCHAR"
				v.Length = 4000
			} else {
				v.Type = "LONGTEXT"
			}
			break
		case "ROWVERSION", "TIMESTAMP":
			v.Type = "BINARY"
			v.Length = 8
			break
		case "SMALLDATETIME":
			v.Type = "DATETIME"
			break
		case "SMALLMONEY":
			v.Type = "DECIMAL"
			v.Precision = 6
			v.Scale = 4
			break
		case "TEXT":
			v.Type = "LONGTEXT"
			break
		case "UNIQUEIDENTIFIER":
			v.Type = "CHAR"
			v.Length = 16
			break
		case "VARBINARY":
			if v.Length > 8000 {
				v.Type = "LONGBLOB"
			}
			break
		case "VARCHAR":
			if v.Length > 8000 {
				v.Type = "LONGTEXT"
			}
			break
		case "XML":
			v.Type = "LONGTEXT"
			break
		}
	}
	return t.mysqlToMysql()
}

func (t TableMetaInfo) mysqlToMysql() (scriptSQL string, err error) {
	var cols string
	primaryKeys := make([]string, 0)
	//当字符长度达到指定长度时，将剩余字段的类型更改为text，防止大表创建时因存储长度报错
	t.CurrentLen = int64(MysqlFieldFixLen * len(t.Columns))
	for idx, col := range t.Columns {
		if idx != 0 {
			cols += ","
		}
		if tpe, ok := MySQLFieldType[strings.ToUpper(col.Type)]; !ok {
			err = fmt.Errorf("table %s,column %s,unsupported field type: %s", t.TableName, col.Name, col.Type)
			return
		} else {
			if strings.EqualFold(MysqlTypeChar, col.Type) || strings.EqualFold(MysqlTypeVarchar, col.Type) {
				if t.CurrentLen+col.Length > MysqlCharMaxLen {
					col.Type = MysqlTypeText
					tpe, _ = MySQLFieldType[MysqlTypeText]
				} else {
					t.CurrentLen += col.Length
				}
			}
			null := ""
			if !col.IsNull {
				null = "NOT NULL"
			}
			if tpe.HasPrecisionScale {
				cols += fmt.Sprintf("`%s` %s(%d,%d) %s COMMENT '%s'", col.Name, col.Type, col.Precision, col.Scale, null, col.ColumnComment)
			} else if tpe.HasLength {
				cols += fmt.Sprintf("`%s` %s(%d) %s COMMENT '%s'", col.Name, col.Type, col.Length, null, col.ColumnComment)
			} else {
				cols += fmt.Sprintf("`%s` %s %s COMMENT '%s'", col.Name, col.Type, null, col.ColumnComment)
			}
		}
		if col.IsPK {
			primaryKeys = append(primaryKeys, col.Name)
		}
	}

	//索引
	var indexs []string
	for _, value := range t.Index {
		indexs = append(indexs, assembleIndexMySQL(value))
	}
	var indexSQL string
	if len(indexs) > 0 {
		indexSQL = fmt.Sprintf(",%s", strings.Join(indexs, ","))
	}

	if len(primaryKeys) > 0 {
		scriptSQL = fmt.Sprintf("CREATE TABLE IF NOT EXISTS `%s`.`%s` (%s,PRIMARY KEY (`%s`)%s) COMMENT '%s';",
			t.TableSchema, t.TableName, cols, strings.Join(primaryKeys, "`,`"), indexSQL, t.COMMENT)
	} else {
		scriptSQL = fmt.Sprintf("CREATE TABLE IF NOT EXISTS `%s`.`%s` (%s%s) COMMENT '%s';",
			t.TableSchema, t.TableName, cols, indexSQL, t.COMMENT)
	}
	return
}

// assembleIndexMySQL 组装索引SQL
func assembleIndexMySQL(newIndex TableIndex) (sql string) {
	if newIndex.Type == "FULLTEXT" || newIndex.Type == "SPATIAL" {
		sql = fmt.Sprintf("%s", newIndex.Type)
	} else {
		sql = "KEY"
	}
	sql += fmt.Sprintf(" `%s`(`%s`)", newIndex.KeyName, strings.Join(newIndex.ColumnName, "`,`"))
	if newIndex.Type != "FULLTEXT" && newIndex.Type != "SPATIAL" && newIndex.Func != "" {
		sql += fmt.Sprintf(" USING %s", strings.ToUpper(newIndex.Func))
	}
	if newIndex.Comment != "" {
		sql += fmt.Sprintf(" COMMENT '%s'", newIndex.Comment)
	}
	return
}
