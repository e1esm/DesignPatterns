package main

type UserCollection struct {
	Users []*User
}

func (uc *UserCollection) NewIterator() Iterator {
	return &UserIterator{
		Users: uc.Users,
	}
}
