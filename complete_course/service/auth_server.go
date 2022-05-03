package service


type AuthServer struct {
	userStore UserStore
	jwtManager *JWTManager
}

func NewAuthServer(userStore UserStore, jwtManager *JWTManager) *AuthServer {
	return &AuthServer{
		userStore: userStore,
		jwtManager: jwtManager,
	}
}
