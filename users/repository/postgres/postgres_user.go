package postgres

type UserPostgresRepository struct{}

func (upr *UserPostgresRepository) InsertNewUser() {}

func (upr *UserPostgresRepository) GetUserByEmail(email string) {}
