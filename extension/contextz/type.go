package contextz

type contextKey struct {
	name string
}

var (
	sessionStoreContextKey = &contextKey{name: "session_store"}
	userContextKey         = &contextKey{name: "user_entity"}
)
