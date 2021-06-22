package entity

//User represents users table in database
type User struct {
	ID       uint64  `gorm:"primary_key:auto_increment" json:"id"`
	Name     string  `gorm:"type:varchar(255)" json:"name"`
	Email    string  `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Password string  `gorm:"->;<-;not null" json:"-"`  //maksud json:"-" tuh berarti dia gaakan muncul di json
	Token    string  `gorm:"-" json:"token,omitempty"` //ini berarti Token gaakan masuk ke database tapi muncul di json nya
	Books    *[]Book `json:"books,omitempty"`          //Pointer ke entity book
}
