package common

// const for error message
const (
	ERROR                      = "error"
	ERROR_BIND_JSON            = "failed to bind json request body"
	ERROR_USER_EXISTED         = "the email was used for another user"
	ERROR_QUERY_LIMIT          = "limit must be a number"
	ERROR_QUERY_LIMIT_ZERO     = "limit must not be 0"
	ERROR_QUERY_OFFSET         = "offset must be a number"
	ERROR_REQUIRE_LIMIT_OFFSET = "limit/offset must co-exist or not"
)

const USER_ID_HEADER = "user-id"
