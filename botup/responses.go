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

// Bot sucess responses
var (
	OK_BOT_ADDED      = []byte("OK_BOT_ADDED")
	OK_PLUGIN_ADDED   = []byte("OK_PLUGIN_ADDED")
	OK_BOT_REMOVED    = []byte("OK_BOT_REMOVED")
	OK_PLUGIN_REMOVED = []byte("OK_PLUGIN_REMOVED")
)

// Bpt error responses
var (
	ERR_BOT_ALREADY_EXISTS    = []byte("ERR_BOT_ALREADY_EXISTS")
	ERR_PLUGIN_ALREADY_EXISTS = []byte("ERR_PLUGIN_ALREADY_EXISTS")
	ERR_BOT_DOESNT_EXISTS     = []byte("ERR_BOT_DOESNT_EXISTS")
	ERR_PLUGIN_DOESNT_EXISTS  = []byte("ERR_PLUGIN_DOESNT_EXISTS")
)
