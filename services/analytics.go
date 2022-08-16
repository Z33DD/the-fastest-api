package services

import "github.com/Z33DD/Napoleon/db"

type AnalyticsFlied struct {
	linkId string
	value  string
}

func Track(linkId string) error {
	db.Client.Incr("analytics_" + linkId)
	return nil
}

func Retrieve(linkId string) ([]AnalyticsFlied, error) {
	var data []AnalyticsFlied
	val, err := db.Client.Get("analytics_" + linkId).Result()
	if err != nil {
		return nil, err
	}
	data = append(data, AnalyticsFlied{linkId: linkId, value: string(val)})

	return data, nil
}
