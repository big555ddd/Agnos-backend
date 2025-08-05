package patientdto

type GetPatientByIdRequest struct {
	ID string `uri:"id" binding:"required"`
}

type ListPatientRequest struct {
	Page        int    `form:"page"`
	Size        int    `form:"size"`
	SortBy      string `form:"sort_by"`
	OrderBy     string `form:"order_by"`
	NationalID  string `form:"national_id"`
	PassportID  string `form:"passport_id"`
	FirstName   string `form:"first_name"`
	MiddleName  string `form:"middle_name"`
	LastName    string `form:"last_name"`
	DateOfBirth string `form:"date_of_birth"`
	Email       string `form:"email"`
	PhoneNumber string `form:"phone_number"`
}
