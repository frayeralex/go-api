package models

import (
	"fmt"
	"github.com/globalsign/mgo/bson"
	"github.com/go-bongo/bongo"
	"time"
)

type Activity struct {
	bongo.DocumentBase `bson:",inline"`
	Channel            bson.ObjectId `bson:"channel" json:"channel"`
	ActivityAt         time.Time     `bson:"activityAt" json:"activityAt"`
	Type               string        `bson:"type" json:"type"`
	Value              int           `bson:"value" json:"value"`
}

func (a Activity) Id() string {
	return fmt.Sprintf(`%x`, string(a.GetId()))
}
