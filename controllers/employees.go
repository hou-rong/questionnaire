package controllers

import (
	"log"
	"net/http"
	"questionnaire/database"
	"questionnaire/models"
	"questionnaire/utils"
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
	var employees []models.Employee

	// Parse the result set of the SQL query.
	for rows.Next() {
		// Initialize the variable called "employee" and assign an "Employee" struct to the variable.
		var employee models.Employee

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