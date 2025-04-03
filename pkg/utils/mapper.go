package utils

import (
	"database/sql"
	"reflect"
	"fmt"
)

func MapRowsToStruct(rows *sql.Rows, dest interface{}) error {
	// Obtener el tipo del destino (debe ser un slice de estructuras)
	sliceType := reflect.TypeOf(dest).Elem()
	if sliceType.Kind() != reflect.Slice {
		return fmt.Errorf("el destino debe ser un slice de estructuras")
	}

	// Obtener la referencia al slice destino
	sliceValue := reflect.ValueOf(dest).Elem()

	// Obtener las columnas de la consulta
	columns, err := rows.Columns()
	if err != nil {
		return err
	}

	// Crear un mapa de nombres de columna a índice
	columnIndex := make(map[string]int)
	for i, col := range columns {
		columnIndex[col] = i
	}

	// Iterar sobre las filas
	for rows.Next() {
		// Crear una nueva instancia del struct
		structPtr := reflect.New(sliceType.Elem()) // Crear un nuevo struct del tipo destino
		structValue := structPtr.Elem()

		// Crear un slice de interfaces para Scan
		fieldValues := make([]interface{}, len(columns))
		for i := range fieldValues {
			fieldValues[i] = new(interface{}) // Puntero a interface{} para recibir los valores
		}

		// Escanear los valores en el slice
		if err := rows.Scan(fieldValues...); err != nil {
			return err
		}

		// Asignar valores a los campos del struct
		for i, colName := range columns {
			// Buscar el campo por el nombre de la etiqueta JSON
			for j := 0; j < structValue.NumField(); j++ {
				fieldStruct := structValue.Type().Field(j)
				jsonTag := fieldStruct.Tag.Get("json")
				if jsonTag == colName {
					field := structValue.Field(j)

					// Si el campo es asignable, establecer el valor
					if field.CanSet() {
						val := reflect.ValueOf(*(fieldValues[i].(*interface{})))
						if val.IsValid() {
							field.Set(val.Convert(field.Type())) // Convertir y asignar el valor
						}
					}
					break
				}
			}
		}

		// Agregar el struct al slice destino
		sliceValue.Set(reflect.Append(sliceValue, structValue))
	}

	// Manejar caso de consulta vacía
	if sliceValue.Len() == 0 {
		return sql.ErrNoRows
	}

	return nil
}




func MapRowToStruct(rows *sql.Rows, dest interface{}) error {
	// Obtener el tipo del destino (debe ser un puntero a una estructura)
	destType := reflect.TypeOf(dest)
	if destType.Kind() != reflect.Ptr || destType.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("el destino debe ser un puntero a una estructura")
	}

	// Obtener la referencia al valor de destino
	destValue := reflect.ValueOf(dest).Elem()

	// Obtener las columnas de la consulta
	columns, err := rows.Columns()
	if err != nil {
		return err
	}

	// Crear un slice de interfaces para Scan
	fieldValues := make([]interface{}, len(columns))
	for i := range fieldValues {
		fieldValues[i] = new(interface{}) // Puntero a interface{} para recibir los valores
	}

	// Leer una sola fila
	if !rows.Next() {
		return sql.ErrNoRows
	}

	// Escanear los valores en el slice
	if err := rows.Scan(fieldValues...); err != nil {
		return err
	}

	// Asignar valores a los campos de la estructura
	for i, colName := range columns {
		// Buscar el campo por el nombre de la etiqueta JSON
		for j := 0; j < destValue.NumField(); j++ {
			fieldStruct := destValue.Type().Field(j)
			jsonTag := fieldStruct.Tag.Get("json")
			if jsonTag == colName {
				field := destValue.Field(j)

				// Si el campo es asignable, establecer el valor
				if field.CanSet() {
					val := reflect.ValueOf(*(fieldValues[i].(*interface{})))
					if val.IsValid() {
						field.Set(val.Convert(field.Type())) // Convertir y asignar el valor
					}
				}
				break
			}
		}
	}

	return nil
}




// OPCION 1

// func MapRowToStruct(row *sql.Row, dest interface{}) error {
// 	columns, err := getColumns(row)
// 	if err != nil {
// 		return err
// 	}

// 	values, err := getValues(dest, columns)
// 	if err != nil {
// 		return err
// 	}

// 	if err := row.Scan(values...); err != nil {
// 		if errors.Is(err, sql.ErrNoRows) {
// 			return nil // Retorna nil para indicar que no hay resultados
// 		}
// 		return err
// 	}
// 	return nil
// }

// // MapRowsToStructs mapea múltiples filas de SQL a una slice de estructuras en Go.
// func MapRowsToStructs(rows *sql.Rows, dest interface{}) error {
// 	sliceType := reflect.TypeOf(dest).Elem()
// 	if sliceType.Kind() != reflect.Slice {
// 		return errors.New("dest debe ser un puntero a un slice")
// 	}

// 	elemType := sliceType.Elem()
// 	sliceValue := reflect.ValueOf(dest).Elem()

// 	columns, err := rows.Columns()
// 	if err != nil {
// 		return err
// 	}

// 	for rows.Next() {
// 		elem := reflect.New(elemType).Interface()
// 		values, err := getValues(elem, columns)
// 		if err != nil {
// 			return err
// 		}

// 		if err := rows.Scan(values...); err != nil {
// 			return err
// 		}
// 		sliceValue.Set(reflect.Append(sliceValue, reflect.ValueOf(elem).Elem()))
// 	}

// 	return rows.Err()
// }

// // getColumns obtiene los nombres de las columnas de una consulta
// func getColumns(row *sql.Row) ([]string, error) {
// 	var cols []string
// 	err := row.Scan(make([]interface{}, len(cols))...) // Intentamos escanear para obtener las columnas
// 	if err != nil {
// 		return nil, err
// 	}
// 	return cols, nil
// }

// // getValues crea un slice de punteros a valores correspondientes a los campos de la estructura
// func getValues(dest interface{}, columns []string) ([]interface{}, error) {
// 	val := reflect.ValueOf(dest).Elem()
// 	typ := val.Type()

// 	values := make([]interface{}, len(columns))
// 	for i, col := range columns {
// 		field := val.FieldByNameFunc(func(s string) bool {
// 			field, _ := typ.FieldByName(s)
// 			return field.Tag.Get("json") == col
// 		})
// 		if field.IsValid() && field.CanSet() {
// 			values[i] = field.Addr().Interface()
// 		} else {
// 			var dummy interface{}
// 			values[i] = &dummy // Ignorar columnas desconocidas
// 		}
// 	}
// 	return values, nil
// }




// OPCION 2

// func MapRowToStruct(row *sql.Row, dest interface{}) error {
// 	// Obtener los campos de la estructura de destino
// 	destValue := reflect.ValueOf(dest).Elem()
// 	destType := destValue.Type()

// 	// Crear punteros a los valores de la estructura
// 	values := make([]interface{}, destType.NumField())
// 	for i := 0; i < destType.NumField(); i++ {
// 			values[i] = destValue.Field(i).Addr().Interface()
// 	}

// 	// Escanear los valores
// 	err := row.Scan(values...)
// 	if err != nil {
// 			return err
// 	}

// 	return nil
// }
// func MapRowToStruct(rows *sql.Rows, dest interface{}) error {
// 	// Obtener los campos de la estructura de destino
// 	destValue := reflect.ValueOf(dest).Elem()
// 	// destType := destValue.Type()

// 	// Obtener las columnas de la consulta
// 	columns, err := rows.Columns()
// 	if err != nil {
// 		return err
// 	}

// 	// Crear un slice para almacenar los valores
// 	values := make([]interface{}, len(columns))
// 	valuePtrs := make([]interface{}, len(columns))

// 	// Crear punteros para el escaneo
// 	for i := range values {
// 		valuePtrs[i] = &values[i]
// 	}

// 	// Verificar si hay una fila disponible
// 	if !rows.Next() {
// 		return sql.ErrNoRows
// 	}

// 	// Escanear la fila en valores temporales
// 	err = rows.Scan(valuePtrs...)
// 	if err != nil {
// 		return err
// 	}

// 	// Mapear los valores a la estructura destino
// 	for i, colName := range columns {
// 		field := destValue.FieldByName(colName)
// 		if field.IsValid() && field.CanSet() {
// 			// Convertir el tipo del valor antes de asignarlo
// 			val := reflect.ValueOf(values[i])
// 			if val.IsValid() && val.Type().ConvertibleTo(field.Type()) {
// 				field.Set(val.Convert(field.Type()))
// 			}
// 		}
// 	}

// 	return nil
// }

// func MapRowsToStructs(rows *sql.Rows, dest interface{}) error {
// 	// Obtener el tipo del slice destino
// 	destValue := reflect.ValueOf(dest).Elem()
// 	destType := destValue.Type().Elem()

// 	for rows.Next() {
// 			// Crear una nueva instancia de la estructura
// 			item := reflect.New(destType).Elem()

// 			// Crear punteros a los valores de la estructura
// 			values := make([]interface{}, destType.NumField())
// 			for i := 0; i < destType.NumField(); i++ {
// 					values[i] = item.Field(i).Addr().Interface()
// 			}

// 			// Escanear la fila en la estructura
// 			err := rows.Scan(values...)
// 			if err != nil {
// 					return err
// 			}

// 			// Agregar la estructura al slice destino
// 			destValue.Set(reflect.Append(destValue, item))
// 	}

// 	return nil
// }

// OPCION 3 

// func GetRow[T any](db *sql.DB, query string, args ...interface{}) (*T, error) {
// 	row := db.QueryRow(query, args...)

// 	// Crear una instancia del tipo T (User en este caso)
// 	var result T
// 	structValue := reflect.ValueOf(&result).Elem()

// 	// Obtener los nombres de las columnas
// 	columns, err := getColumnNames(db, query)
// 	if err != nil {
// 			return nil, err
// 	}

// 	// Crear un slice de interfaces para recibir los valores
// 	values := make([]interface{}, len(columns))
// 	for i := range values {
// 			var v interface{}
// 			values[i] = &v
// 	}

// 	// Hacer el escaneo de la fila
// 	if err := row.Scan(values...); err != nil {
// 			return nil, err
// 	}

// 	// Asignar los valores a la estructura
// 	for i, col := range columns {
// 			field := structValue.FieldByName(col)
// 			if field.IsValid() && field.CanSet() {
// 					field.Set(reflect.ValueOf(*(values[i].(*interface{}))))
// 			}
// 	}

// 	return &result, nil
// }

// func GetRows[T any](db *sql.DB, query string, args ...interface{}) ([]T, error) {
// 	rows, err := db.Query(query, args...)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	// Obtener los nombres de las columnas
// 	columns, err := rows.Columns()
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Lista de resultados
// 	var results []T

// 	// Iterar sobre cada fila
// 	for rows.Next() {
// 		// Crear una nueva instancia del tipo T
// 		var result T
// 		structValue := reflect.ValueOf(&result).Elem()

// 		// Crear un slice de interfaces para recibir los valores
// 		values := make([]interface{}, len(columns))
// 		for i := range values {
// 			var v interface{}
// 			values[i] = &v
// 		}

// 		// Escanear los valores de la fila
// 		if err := rows.Scan(values...); err != nil {
// 			return nil, err
// 		}

// 		// Asignar los valores a la estructura
// 		for i, col := range columns {
// 			field := structValue.FieldByName(col)
// 			if field.IsValid() && field.CanSet() {
// 				field.Set(reflect.ValueOf(*(values[i].(*interface{}))))
// 			}
// 		}

// 		// Agregar a la lista de resultados
// 		results = append(results, result)
// 	}

// 	// Verificar si hubo errores en el recorrido de las filas
// 	if err := rows.Err(); err != nil {
// 		return nil, err
// 	}

// 	return results, nil
// }

// // getColumnNames obtiene los nombres de las columnas de una consulta
// func getColumnNames(db *sql.DB, query string) ([]string, error) {
// 	rows, err := db.Query(query + " LIMIT 1")
// 	if err != nil {
// 			return nil, err
// 	}
// 	defer rows.Close()

// 	return rows.Columns()
// }
