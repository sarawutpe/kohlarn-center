package model

type Response struct {
	Data  interface{} `bson:"data,omitempty" json:"data,omitempty"`   // data
	Error *Error      `bson:"error,omitempty" json:"error,omitempty"` // error
}

type Error struct {
	Status  string      `bson:"status,omitempty" json:"status,omitempty"`   // HTTP status
	Name    string      `bson:"name,omitempty" json:"name,omitempty"`       // error name ('ApplicationError' or 'ValidationError')
	Message string      `bson:"message,omitempty" json:"message,omitempty"` // A human readable error message
	Details interface{} `bson:"details,omitempty" json:"details,omitempty"` // error info specific to the error type
}
