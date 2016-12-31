package botup

// User sucess responses
var (
	OK_LOGGED_IN    = []byte("OK_LOGGED_IN")
	OK_LOGGED_OUT   = []byte("OK_LOGGED_OUT")
	OK_USER_CREATED = []byte("OK_USER_CREATED")
)

// User error responses
var (
	ERR_NOT_LOGGED_IN     = []byte("ERR_NOT_LOGGED_IN")
	ERR_INTERNAL          = []byte("ERR_INTERNAL")
	ERR_FIELDS_MISSING    = []byte("ERR_FIELDS_MISSING")
	ERR_WRONG_CREDENTIALS = []byte("ERR_WRONG_CREDENTIALS")
	ERR_USER_TAKEN        = []byte("ERR_USER_TAKEN")
)
