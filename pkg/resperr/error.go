package resperr

const (
	// error codes
	ERRCONFLICT = "conflict"
	ERRNOTFOUND = "not_found"
	ERRINVALID  = "invalid"
	ERRINTERNAL = "internal"
)

type Error struct {
	Code string

	UserMessage string

	Op  string
	Err error
}



func New(statusCode int, errMessage string) Error {
	return &Error{
		
	}

}

func WithCodeAndMesssage(err error, Code )


func (e *Error) Wrap() {

}
