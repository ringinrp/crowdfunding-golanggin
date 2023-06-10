package user

type RegisterUserInput struct {
	Name       string `josn:"name" binding:"required"`
	Occupation string `josn:"occupation" binding:"required"`
	Email      string `josn:"email" binding:"required"`
	Password   string `josn:"password" binding:"required"`
}
