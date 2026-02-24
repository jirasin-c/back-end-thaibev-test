package service

type CreatePersonRequest struct {
	FirstName       string `json:"firstName" binding:"required"`
	LastName        string `json:"lastName"  binding:"required"`
	Email           string `json:"email"     binding:"required,email"`
	Phone           string `json:"phone"     binding:"required"` // validate
	BirthDay        string `json:"birthDay"  binding:"required"` // dd/MM/yyyy or yyyy-MM-dd
	Sex             string `json:"sex"       binding:"required,oneof=M F"`
	OccupationCode  string `json:"occupationCode" binding:"required"`
	ProfileFileName string `json:"profileFileName" binding:"required"`
	ProfileBase64   string `json:"profileBase64"   binding:"required"`
}

type CreatePersonResponse struct {
	ID string `json:"id"`
}

type OccupationDTO struct {
	Code string `db:"code"`
	Name string `db:"name"`
}
