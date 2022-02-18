package model

import (
	"go.mongodb.org/mongo-driver/bson"
)

//==========================================================
//preference Management
type Pref struct {
	Name    string      `json:"name,omitempty"`
	Content interface{} `json:"content,omitempty"`
}

func AddNewPref(name string, content interface{}) {
	colle := DB.Collection(DBCOLN_PREFS)
	pref := Pref{name, content}
	colle.InsertOne(CTX, pref)
}

func UpdatePref(name string, content interface{}) {
	colle := DB.Collection(DBCOLN_PREFS)
	filter := bson.M{"name": name}
	update := bson.D{
		{"$set", bson.M{
			"content": content,
		}},
	}
	colle.UpdateOne(CTX, filter, update)
}

func GetPref(name string) interface{} {
	colle := DB.Collection(DBCOLN_PREFS)
	filter := bson.M{"name": name}
	cursor := colle.FindOne(CTX, filter)
	result := Pref{}
	cursor.Decode(&result)
	return result
}

func GetAllPrefs() []Pref {
	prefs := []Pref{}
	colle := DB.Collection(DBCOLN_PREFS)
	filter := bson.D{}
	cursor, _ := colle.Find(CTX, filter)
	for cursor.Next(CTX) {
		var cnl Pref
		cursor.Decode(&cnl)
		prefs = append(prefs, cnl)
	}
	return prefs
}

func DeletePref(name string) {
	colle := DB.Collection(DBCOLN_PREFS)
	colle.DeleteOne(CTX, bson.M{"name": name})
}
