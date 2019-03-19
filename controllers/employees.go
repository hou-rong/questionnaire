package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"questionnaire/database"
	"questionnaire/models"
	"questionnaire/utils"
	"strconv"
)

var GetEmployees = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Execute the SQL query to take information about all employees.
	rows, err := database.OracleDB.Query(`SELECT 
		P_TABEL_ID,
		P_EMAIL,
		P_LASTNAME,
		P_FIRSTNAME,
		P_MIDDLENAME,
		P_SEX,
		P_PL,
		P_POSITION_NAME,
		P_DATE_OF_BIRTH,
		AGE,
		A_PEOPLE_GROUP,
		C_EMAIL,
		C_LASTNAME,
		C_FIRSTNAME,
		C_MIDDLENAME,
       REPLACE(NVL(REGEXP_REPLACE(TRIM('\' from TREE_ORGANIZATION_NAME), '\\+', ' > '), RANG1_ORGANIZATION_NAME), chr(9), ' > ')
	FROM
		DMP_ORG_STR_PEOPLE`)

	// Send response with detailed information about the error if the above process was unsuccessful.
	if err != nil {
		log.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	// Call "Close" function.
	defer rows.Close()

	// Initialize the variable called "employees" and assign an array to the variable.
	var employees []models.ExtendedEmployee

	// Parse the result set of the SQL query.
	for rows.Next() {
		// Initialize the variable called "employee" and assign an "Employee" struct to the variable.
		var employee models.ExtendedEmployee

		// Call "Scan()" function on the result set of the SQL query.
		if err := rows.Scan(&employee.EmployeeID,
			&employee.EmployeeEmail,
			&employee.EmployeeLastName,
			&employee.EmployeeFirstName,
			&employee.EmployeeMiddleName,
			&employee.EmployeeSex,
			&employee.EmployeePositionLevel,
			&employee.EmployeePositionName,
			&employee.EmployeeBirthday,
			&employee.EmployeeAge,
			&employee.EmployeeGroup,
			&employee.SupervisorEmail,
			&employee.SupervisorLastName,
			&employee.SupervisorFirstName,
			&employee.SupervisorMiddleName,
			&employee.Organization); err != nil {
			log.Println(err)
			utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
			return
		}

		// Set questions to the final array.
		employees = append(employees, employee)
	}

	// Send successful response with status code "200" and JSON.
	utils.Response(responseWriter, http.StatusOK, employees)
}

var GetEmployee = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Take query parameters with the help of "mux.Vars" function.
	vars := mux.Vars(request)

	// Take email from query based parameter.
	email := vars["email"]

	// Execute the SQL query to take information about all employees.
	rows, err := database.OracleDB.Query(`SELECT 
		P_TABEL_ID,
		P_EMAIL,
		P_LASTNAME,
		P_FIRSTNAME,
		P_MIDDLENAME,
		P_SEX,
		P_PL,
		P_POSITION_NAME,
		P_DATE_OF_BIRTH,
		AGE,
		A_PEOPLE_GROUP,
		C_EMAIL,
		C_LASTNAME,
		C_FIRSTNAME,
		C_MIDDLENAME,
       	REPLACE(NVL(REGEXP_REPLACE(TRIM('\' from TREE_ORGANIZATION_NAME), '\\+', ' > '), RANG1_ORGANIZATION_NAME), chr(9), ' > ')
	FROM
		DMP_ORG_STR_PEOPLE
	WHERE
		P_EMAIL = '` + email + `'`)

	// Send response with detailed information about the error if the above process was unsuccessful.
	if err != nil {
		log.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	// Call "Close" function.
	defer rows.Close()

	// Initialize the variable called "employees" and assign an array to the variable.
	var employees []models.ExtendedEmployee

	// Parse the result set of the SQL query.
	for rows.Next() {
		// Initialize the variable called "employee" and assign an "Employee" struct to the variable.
		var employee models.ExtendedEmployee

		// Call "Scan()" function on the result set of the SQL query.
		if err := rows.Scan(&employee.EmployeeID,
			&employee.EmployeeEmail,
			&employee.EmployeeLastName,
			&employee.EmployeeFirstName,
			&employee.EmployeeMiddleName,
			&employee.EmployeeSex,
			&employee.EmployeePositionLevel,
			&employee.EmployeePositionName,
			&employee.EmployeeBirthday,
			&employee.EmployeeAge,
			&employee.EmployeeGroup,
			&employee.SupervisorEmail,
			&employee.SupervisorLastName,
			&employee.SupervisorFirstName,
			&employee.SupervisorMiddleName,
			&employee.Organization); err != nil {
			log.Println(err)
			utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
			return
		}

		// Set questions to the final array.
		employees = append(employees, employee)
	}

	// Send successful response with status code "200" and JSON.
	utils.Response(responseWriter, http.StatusOK, employees)
}

/*
Function name:
"GenerateSQLQuery"

Function description:
The main task of the function is to create an sql query string depending on the query based parameters.
*/
func GenerateSQLQuery(organizationID string, organizationName string, employeePositionName string) string {
	// Initialize variable called "sqlQueryString".
	var sqlQueryString string

	// Create sql query string.
	if employeePositionName == "Стажеры" {
		sqlQueryString = `
		SELECT
			P_TABEL_ID,
			P_EMAIL
		FROM
			DMP_ORG_STR_PEOPLE
		WHERE
			P_POSITION_NAME = 'Стажер'`
		return sqlQueryString
	} else if employeePositionName == "Сотрудники" {
		sqlQueryString = `
		SELECT
			P_TABEL_ID,
			P_EMAIL
		FROM
			DMP_ORG_STR_PEOPLE
		WHERE
			ORGANIZATION_ID = ` + organizationID + `
		AND
			TREE_ORGANIZATION_NAME LIKE '%` + organizationName + `%'
		AND
			A_PEOPLE_GROUP NOT IN ('Сотрудницы, находящиеся в декрете')`
		return sqlQueryString
	} else if employeePositionName == "Руководители" {
		sqlQueryString = `
		SELECT	
			P_TABEL_ID,
			P_EMAIL
		FROM
			DMP_ORG_STR_PEOPLE
		WHERE
			P_EMAIL IN (
				SELECT 
					C_EMAIL
				FROM
					DMP_ORG_STR_PEOPLE
				WHERE
					ORGANIZATION_ID = ` + organizationID + `
				AND
					TREE_ORGANIZATION_NAME LIKE '%` + organizationName + `%'
				AND
					A_PEOPLE_GROUP NOT IN ('Сотрудницы, находящиеся в декрете')
				GROUP BY
					C_EMAIL	
			)
		`
	}

	// Return sql query string.
	return sqlQueryString
}

var GetOrganizationEmployees = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Take position of the employee from query based parameters.
	employeePosition := request.URL.Query().Get("employee_position")
	if len(employeePosition) == 0 {
		utils.ResponseWithError(responseWriter, http.StatusBadRequest, "Server could not understand the request due to invalid syntax.")
		return
	}

	// Create struct called "RequestBody".
	type RequestBody struct {
		OrganizationID int `json:"organization_id"`
		OrganizationName string `json:"organization_name"`
	}

	// Initialize the variable called "requestBody" and assign "RequestBody" struct to the variable.
	requestBody := RequestBody{}

	// NewDecoder returns a new decoder that reads from request body.
	decoder := json.NewDecoder(request.Body)

	// Decode reads the next JSON-encoded value from its input and stores it in the value pointed to by "&requestBody".
	_ = decoder.Decode(&requestBody)

	fmt.Println(GenerateSQLQuery(strconv.Itoa(requestBody.OrganizationID), requestBody.OrganizationName, employeePosition))

	// Execute the SQL query to take information about all employees.
	rows, err := database.OracleDB.Query(GenerateSQLQuery(strconv.Itoa(requestBody.OrganizationID), requestBody.OrganizationName, employeePosition))

	// Send response with detailed information about the error if the above process was unsuccessful.
	if err != nil {
		log.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	// Call "Close" function.
	defer rows.Close()

	// Initialize the variable called "employees" and assign an array to the variable.
	var employees []models.AbridgedEmployee

	// Parse the result set of the SQL query.
	for rows.Next() {
		// Initialize the variable called "employee" and assign an "Employee" struct to the variable.
		var employee models.AbridgedEmployee

		// Call "Scan()" function on the result set of the SQL query.
		if err := rows.Scan(&employee.EmployeeID, &employee.EmployeeEmail); err != nil {
			log.Println(err)
			utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
			return
		}

		// Set questions to the final array.
		employees = append(employees, employee)
	}

	// Send successful response with status code "200" and JSON.
	utils.Response(responseWriter, http.StatusOK, employees)
}
