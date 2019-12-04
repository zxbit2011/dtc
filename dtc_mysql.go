package main

const (
	MysqlCharMaxLen  = 16183
	MysqlFieldFixLen = 2
	MysqlTypeChar    = "char"
	MysqlTypeVarchar = "varchar"
	MysqlTypeText    = "text"
)

// OracleFieldType Oracle字段类型长度精度的描述
var MySQLFieldType = map[string]FieldType{
	"BIGINT":             {false, false},
	"BINARY":             {true, false},
	"BIT":                {true, false},
	"BLOB":               {false, false},
	"BOOL":               {false, false},
	"BOOLEAN":            {false, false},
	"CHAR":               {true, false},
	"DATE":               {false, false},
	"DATETIME":           {false, false},
	"DECIMAL":            {true, true},
	"DOUBLE":             {false, false},
	"ENUM":               {true, false},
	"FLOAT":              {false, false},
	"GEOMETRY":           {false, false},
	"GEOMETRYCOLLECTION": {false, false},
	"INT":                {false, false},
	"INTEGER":            {false, false},
	"JSON":               {false, false},
	"LINESTRING":         {false, false},
	"LONGBLOB":           {false, false},
	"LONGTEXT":           {false, false},
	"MEDIUMBLOB":         {false, false},
	"MEDIUMINT":          {false, false},
	"MEDIUMTEXT":         {false, false},
	"MULTILINESTRING":    {false, false},
	"MULTIPOINT":         {false, false},
	"MULTIPOLYGON":       {false, false},
	"NUMERIC":            {true, true},
	"POINT":              {false, false},
	"POLYGON":            {false, false},
	"REAL":               {true, true},
	"SET":                {true, false},
	"SMALLINT":           {false, false},
	"TEXT":               {false, false},
	"TIME":               {false, false},
	"TIMESTAMP":          {false, false},
	"TINYBLOB":           {false, false},
	"TINYINT":            {false, false},
	"TINYTEXT":           {false, false},
	"VARBINARY":          {true, false},
	"VARCHAR":            {true, false},
}
