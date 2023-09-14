package patient

type Patient struct {
	ID          int    `json:"id"`
	Name        string `json:"name" binding:"required"`
	LastName    string `json:"lastname" binding:"required"`
	Address     string `json:"address" binding:"required"`
	Dni         string `json:"dni" binding:"required"`
	DateOfBirth string `json:"dateofbirth" binding:"required"`
}