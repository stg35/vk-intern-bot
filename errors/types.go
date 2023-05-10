package errors

type botError string

const (
	RequestFailed         botError = "can't do request"
	SessionDataFailed     botError = "can't get server session data"
	LongPollRequestFailed botError = "can't make long poll request"
	GetUpdatesFailed      botError = "can't get updates"
)
