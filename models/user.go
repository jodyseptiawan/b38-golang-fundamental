package models

import "time"

// User model struct
type User struct {
  ID          int 			`json:"id"`
  Name 		    string		`json:"name" gorm:"type: varchar(255)"`
  Email		    string 		`json:"email" gorm:"type: varchar(255)"`
  Password 	  string		`json:"password" gorm:"type: varchar(255)"`
  Role        bool      `json:"role"`
  Image       string    `json:"image"`
  CreatedAt 	time.Time	`json:"created_at"`
  UpdatedAt 	time.Time	`json:"updated_at"`
}

type Profile struct {
  ID          int 			`json:"id"`
  Address 		string		`json:"address" gorm:"type: varchar(255)"`
  PhoneNumber	string 		`json:"phoneNumber" gorm:"type: varchar(255)"`
  CreatedAt 	time.Time	`json:"created_at"`
  UpdatedAt 	time.Time	`json:"updated_at"`
}