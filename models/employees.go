package models

/*
sql.NullString is not supported: Oracle database does not differentiate between an empty string ("") and a NULL.
*/
type Employee struct {
	EmployeeID string `json:"employee_id"`
	EmployeeEmail *string `json:"employee_email"`
	EmployeeLastName *string `json:"employee_last_name"`
	EmployeeFirstName *string `json:"employee_first_name"`
	EmployeeMiddleName *string `json:"employee_middle_name"`
	EmployeeSex *string `json:"employee_sex"`
	EmployeePositionLevel *string `json:"employee_position_level"`
	EmployeePositionName *string `json:"employee_position_name"`
	EmployeeBirthday *string `json:"employee_birthday"`
	EmployeeAge *int `json:"employee_age"`
	EmployeeGroup *string `json:"employee_group"`
	SupervisorEmail *string `json:"supervisor_email"`
	SupervisorLastName *string `json:"supervisor_last_name"`
	SupervisorFirstName *string `json:"supervisor_first_name"`
	SupervisorMiddleName *string `json:"supervisor_middle_name"`
	Organization *string `json:"organization"`
}
