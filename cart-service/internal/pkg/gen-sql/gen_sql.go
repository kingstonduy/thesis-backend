package gensql

import (
	"reflect"

	"github.com/doug-martin/goqu/v9"
)

func extractDBTags(data interface{}) map[string]interface{} {
	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	result := make(map[string]interface{})

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		dbTag := field.Tag.Get("db")
		if dbTag != "" {
			result[dbTag] = value.Interface()
		}
	}

	return result
}

// Extract the `db` tags and corresponding values from the struct
func structToDBMap(data interface{}) (map[string]interface{}, error) {
	result := make(map[string]interface{})

	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		dbTag := field.Tag.Get("db")
		if dbTag != "" {
			if !value.IsNil() { // Check if the value is not nil
				result[dbTag] = value.Elem().Interface()
			}
		}
	}

	return result, nil
}

func GenUpdateSql(tableName string, columns map[string]interface{}, conditions map[string]interface{}) (string, error) {
	// Generate the SQL query dynamically
	updateQuery := goqu.Update(tableName)

	// Extract columns to update from the struct
	updateRecord := goqu.Record{}
	for key, value := range columns {
		updateRecord[key] = value
	}
	updateQuery = updateQuery.Set(updateRecord)

	// Add conditions
	whereCondition := goqu.Ex{}
	for key, value := range conditions {
		whereCondition[key] = value
	}
	updateQuery = updateQuery.Where(whereCondition)

	// Generate the final SQL query
	sql, _, err := updateQuery.ToSQL()
	if err != nil {
		return "", err
	}

	return sql, nil
}

func GenInsertSql(tableName string, data interface{}) (s string, err error) {
	columns, err := structToDBMap(data)
	if err != nil {
		return "", err
	}
	// Generate the SQL query dynamically
	insertQuery := goqu.Insert(tableName)

	// Extract columns and their values from the struct
	insertRecord := goqu.Record{}
	colunnmNames := extractDBTags(data)
	for key, _ := range colunnmNames {
		val, ok := columns[key]
		if ok {
			insertRecord[key] = val
		} else {
			insertRecord[key] = nil
		}
	}
	insertQuery = insertQuery.Rows(insertRecord)

	// Generate the final SQL query
	sql, _, err := insertQuery.ToSQL()
	if err != nil {
		return "", err
	}

	return sql, nil
}

func SelectByParams(tableName string, conditions map[string]interface{}) (string, error) {
	// Generate the SQL query dynamically
	selectQuery := goqu.From(tableName)

	// Add conditions
	whereCondition := goqu.Ex{}
	for key, value := range conditions {
		whereCondition[key] = value
	}
	selectQuery = selectQuery.Where(whereCondition)

	// Generate the final SQL query
	sql, _, err := selectQuery.ToSQL()
	if err != nil {
		return "", err
	}

	return sql, nil
}
