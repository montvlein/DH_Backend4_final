package dentist

type Dentist struct {
	ID          int     `json:"id"`
	LastName    string  `json:"lastname" binding:"required"`	
	Name        string  `json:"name" binding:"required"`
	License    	string  `json:"license" binding:"required"`
	}