package repo

// 快捷语法

func AND(column string) string {
	return "AND " + column + " = ?"
}

func AND_LIKE(column string) string {
	return "AND " + column + " LIKE ?"
}

func AND_IN(column string) string {
	return "AND " + column + " IN (?)"
}

func DESC(column string) string { // 降序
	return column + " DESC"
}