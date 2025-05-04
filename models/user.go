package models

type User struct {
	ID       string `json:"id"`
	UserName string `json:"userName"`
	Pass     string `json:"pass"`
	IsAdmin  bool   `json:"isAdmin"`
}

//Mock Users table
var Users = map[string]User{
	"3f0f5f7a-2e4a-4a25-9c8c-2c1147f153ed": {ID: "3f0f5f7a-2e4a-4a25-9c8c-2c1147f153ed", UserName: "Abdullah", Pass: "1,@3A", IsAdmin: false},
	"gf0fdsfa-sg4a-52s5-9c8c-2asd47fdfb3a": {ID: "gf0fdsfa-sg4a-52s5-9c8c-2asd47fdfb3a", UserName: "Admin", Pass: "A2K,2@S", IsAdmin: true},
}

func GetUserByName(name string) (User, bool) {
	for _, user := range Users {
		if name == user.UserName {
			return user, true
		}
	}
	return User{}, false
}


func GetUserById(Id string)(User,bool){
	user,ok := Users[Id]
	return user,ok
}
