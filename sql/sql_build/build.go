package sqlBuilder

import "lv-utils/sql"

func Select(table string) (*sql.SqlBuilder) {
	auto := new(sql.SqlBuilder)
	auto.Select(table)
	return auto
}
func Insert(table string) (*sql.SqlBuilder) {
	auto := new(sql.SqlBuilder)
	auto.Insert(table)
	return auto
}
func Update(table string) (*sql.SqlBuilder) {
	auto := new(sql.SqlBuilder)
	auto.Update(table)
	return auto
}
func Delete(table string) (*sql.SqlBuilder) {
	auto := new(sql.SqlBuilder)
	auto.Delete(table)
	return auto
}



