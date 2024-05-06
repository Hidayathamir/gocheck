// Package q provides gorm sql query functionality.
package q

// Equal return gorm sql query column = ?.
func Equal(column string) string {
	return column + " = ?"
}
