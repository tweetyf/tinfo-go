package model

import (
	"tinfo-go/utils"

	"go.mongodb.org/mongo-driver/bson"
)

// ==========================================================
// User management

type User struct {
	Uid       string `json:"uid,omitempty"`
	UnameShow string `json:"unameshow,omitempty"`
	Uemail    string `json:"uemail,omitempty"`
	Upwd      string `json:"upwd,omitempty"`
	USalt     uint   `json:"usalt,omitempty"`
	Urole     string `json:"urole,omitempty"`
	Uavatar   string `json:"uavatar,omitempty"`
	Udesc     string `json:"udesc,omitempty"`
}

func GetAllUsers() []User {
	colle := DB.Collection(DBCOLN_USERS)
	cursor, err := colle.Find(CTX, bson.M{})
	if err != nil {
		utils.LogE("OOps! something wrong: %v", err)
	}
	users := []User{}
	for cursor.Next(CTX) {
		var user User
		cursor.Decode(&user)
		users = append(users, user)
	}
	return users
}

func GetUserByID(uid string) User {
	colle := DB.Collection(DBCOLN_USERS)
	cursor := colle.FindOne(CTX, bson.M{"uid": uid})
	user := User{}
	cursor.Decode(&user)
	return user
}

func AddNewUser(uname string, upwd string, uemail string, urole string, uavatar string, udesc string) User {
	salt := utils.GenNewSalt()
	pwd := utils.GenNewPasswd(upwd, salt)
	user := User{Uid: utils.GenUid(),
		UnameShow: uname,
		Uemail:    uemail,
		Upwd:      pwd,
		Urole:     urole,
		USalt:     salt,
		Uavatar:   uavatar,
		Udesc:     udesc,
	}
	colle := DB.Collection(DBCOLN_USERS)
	result, err := colle.InsertOne(CTX, user)
	if err != nil {
		utils.LogE("%v", err)
	} else {
		utils.LogD("add new user:%v", result)
	}
	return user
}

func UpdateUser(user User) {
	colle := DB.Collection(DBCOLN_USERS)
	filter := bson.M{"uid": user.Uid}
	update := bson.D{
		{"$set", bson.M{"unameshow": user.UnameShow,
			"urole":   user.Urole,
			"uavatar": user.Uavatar,
			"udesc":   user.Udesc}},
	}
	colle.UpdateOne(CTX, filter, update)
}

func UpdateUserPWD(uid string, newPwd string) {
	colle := DB.Collection(DBCOLN_USERS)
	filter := bson.M{"uid": uid}
	update := bson.D{
		{"$set", bson.M{"upwd": newPwd}},
	}
	colle.UpdateOne(CTX, filter, update)
}

func DeleteUser(uid string) {
	colle := DB.Collection(DBCOLN_USERS)
	colle.DeleteOne(CTX, bson.M{"uid": uid})
}
