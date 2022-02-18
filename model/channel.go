package model

import (
	"time"
	"tinfo-go/utils"

	"go.mongodb.org/mongo-driver/bson"
)

//==========================================================
//Channel Management

type Channel struct {
	ChanId       string `json:"chanid,omitempty"`
	ChanName     string `json:"channame,omitempty"`
	ChanpicUrl   string `json:"chanpicurl,omitempty"`
	ChancoverUrl string `json:"chancoverurl,omitempty"`
	ChanDesc     string `json:"chandesc,omitempty"`
	Issuer       string `json:"issuer,omitempty"`
	IssueTime    int64  `json:"issuetime,omitempty"`
	EditTime     int64  `json:"edittime,omitempty"`
}

func GetAllChannels() []Channel {
	chans := []Channel{}
	colle := DB.Collection(DBCOLN_CHANS)
	cursor, err := colle.Find(CTX, bson.M{})
	if err != nil {
		utils.LogE("OOps! something wrong: %v", err)
	}
	for cursor.Next(CTX) {
		var cnl Channel
		cursor.Decode(&cnl)
		chans = append(chans, cnl)
	}
	return chans
}

func GetChannelByID(Cid string) Channel {
	result := Channel{}
	colle := DB.Collection(DBCOLN_CHANS)
	filter := bson.M{"chanid": Cid}
	cursor := colle.FindOne(CTX, filter)
	cursor.Decode(&result)
	return result
}

func AddNewChannel(channame string, picurl string, coverurl string, desc string, issuer string) Channel {
	cnl := Channel{ChanId: utils.GenPid(),
		ChanName:     channame,
		ChanpicUrl:   picurl,
		ChancoverUrl: coverurl,
		ChanDesc:     desc,
		Issuer:       issuer,
		IssueTime:    time.Now().Unix(),
		EditTime:     time.Now().Unix()}
	colle := DB.Collection(DBCOLN_CHANS)
	result, err := colle.InsertOne(CTX, cnl)
	if err != nil {
		utils.LogE("%v", err)
	} else {
		utils.LogD("add new channel:%v", result)
	}
	return cnl
}
func UpdateChannel(cnl Channel) {
	colle := DB.Collection(DBCOLN_CHANS)
	filter := bson.M{"chanid": cnl.ChanId}
	update := bson.D{
		{"$set", bson.M{"channame": cnl.ChanName,
			"chanpicurl":   cnl.ChanpicUrl,
			"chancoverurl": cnl.ChancoverUrl,
			"chandesc":     cnl.ChanDesc,
			"issuer":       cnl.Issuer,
			"edittime":     time.Now().Unix()}},
	}
	colle.UpdateOne(CTX, filter, update)
}
func DeleteChannel(chanId string) {
	colle := DB.Collection(DBCOLN_CHANS)
	colle.DeleteOne(CTX, bson.M{"chanid": chanId})
}
