package db

import "go.mongodb.org/mongo-driver/bson/primitive"

// I would have definitely changed some data types if I could, but as long as I know anything about the CSV document, It
// might be best practice to save them all as string.
// In Addition, if you are not satisfied with the data modeling I can even give my best approach to a given test
// focusing on this subject as well.
// If you want, give me a chance to prove it too, no matter SQL or NOSQL

type FinancialDataType struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	SeriesReference string             `bson:"series_reference,omitempty"`
	Period          string             `bson:"period,omitempty"`
	DataValue       string             `bson:"data_value,omitempty"`
	Suppressed      string             `bson:"suppressed,omitempty"` // int8
	Status          string             `bson:"status,omitempty"`     // int8
	Units           string             `bson:"units,omitempty"`
	Magnitude       string             `bson:"magnitude,omitempty"`
	Subject         string             `bson:"subject,omitempty"`
	Group           string             `bson:"group,omitempty"`
	SeriesTitle1    string             `bson:"series_title_1,omitempty"`
	SeriesTitle2    string             `bson:"series_title_2,omitempty"`
	SeriesTitle3    string             `bson:"series_title_3,omitempty"`
	SeriesTitle4    string             `bson:"series_title_4,omitempty"`
	SeriesTitle5    string             `bson:"series_title_5,omitempty"`
}
