package id

type IdGen struct {
	Domain string `json:"domain" gorm:"column:domain;primaryKey" bson:"_id"`
	Value  uint64 `json:"value" gorm:"column:value" bson:"value"`
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
