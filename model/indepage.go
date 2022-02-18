package model

import (
	"time"
	"tinfo-go/utils"

	"go.mongodb.org/mongo-driver/bson"
)

//==========================================================
//Web independant page Management
type Indepage struct {
	Pname       string `json:"pname,omitempty"`
	Ptitle      string `json:"ptitle,omitempty"`
	Pcontent    string `json:"pcontent,omitempty"`
	Pcontentraw string `json:"pcontentraw,omitempty"`
	Issuer      string `json:"issuer,omitempty"`
	IssueTime   int64  `json:"issuetime,omitempty"`
	EditTime    int64  `json:"edittime,omitempty"`
}

func GetIndepages() []Indepage {
	chans := []Indepage{}
	colle := DB.Collection(DBCOLN_INDEPAGES)
	cursor, err := colle.Find(CTX, bson.M{})
	if err != nil {
		utils.LogE("OOps! something wrong: %v", err)
	}
	for cursor.Next(CTX) {
		var cnl Indepage
		cursor.Decode(&cnl)
		chans = append(chans, cnl)
	}
	return chans
}

func GetIndepageByName(name string) Channel {
	result := Channel{}
	colle := DB.Collection(DBCOLN_INDEPAGES)
	filter := bson.M{"pname": name}
	cursor := colle.FindOne(CTX, filter)
	cursor.Decode(&result)
	return result
}

func AddNewIndepage(pname string, ptitle string, pcntraw string, issuer string) Indepage {
	cnl := Indepage{
		Pname:       pname,
		Ptitle:      ptitle,
		Pcontent:    utils.MarkDown2HTML(pcntraw),
		Pcontentraw: pcntraw,
		Issuer:      issuer,
		IssueTime:   time.Now().Unix(),
		EditTime:    time.Now().Unix()}
	colle := DB.Collection(DBCOLN_INDEPAGES)
	result, err := colle.InsertOne(CTX, cnl)
	if err != nil {
		utils.LogE("%v", err)
	} else {
		utils.LogD("add new indePage:%v", result)
	}
	return cnl
}
func UpdateIndepage(cnl Indepage) {
	colle := DB.Collection(DBCOLN_INDEPAGES)
	filter := bson.M{"pname": cnl.Pname}
	update := bson.D{
		{"$set", bson.M{"pname": cnl.Pname,
			"ptitle":      cnl.Ptitle,
			"pcontent":    utils.MarkDown2HTML(cnl.Pcontentraw),
			"pcontentraw": cnl.Pcontentraw,
			"issuer":      cnl.Issuer,
			"edittime":    time.Now().Unix()}},
	}
	colle.UpdateOne(CTX, filter, update)
}
func DeleteIndepage(pname string) {
	colle := DB.Collection(DBCOLN_INDEPAGES)
	colle.DeleteOne(CTX, bson.M{"pname": pname})
}
