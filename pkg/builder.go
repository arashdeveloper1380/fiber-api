package pkg

type UserInfo struct {
	Name  string
	Phone string
	Email string
}

type UserBuilder struct {
	user UserInfo
}

func NewBuilder() *UserBuilder {
	return &UserBuilder{}
}

func (b *UserBuilder) SetName(name string) *UserBuilder {
	b.user.Name = name
	return b
}

func (b *UserBuilder) SetPhone(phone string) *UserBuilder {
	b.user.Phone = phone
	return b
}

func (b *UserBuilder) SetEmail(email string) *UserBuilder {
	b.user.Email = email
	return b
}

func (b *UserBuilder) build() UserInfo {
	return b.user
}

func exampleBuilder() (UserInfo, error) {
	user := NewBuilder().
		SetName("arash").
		SetEmail("arash@gmail.com").
		SetPhone("09030613817").build()

	return user, nil
}
