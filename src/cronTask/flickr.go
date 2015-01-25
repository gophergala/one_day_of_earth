package cronTask

import (
	"APIs"
	"APIs/flickr"
	"config"
	"lib"
	"mongodatabase"
)

func Flickr_Cron(lat, lng, date, distance string) (cerr *lib.CError) {
	index := lib.MD5strings(lat, lng, date, distance)
	loc_hash := lib.MD5strings(lat, lng)
	flickr_collection := mongodatabase.FlickrCollection{}
	m := mongodatabase.Mongo{}
	m.Connect()
	elem_condition := map[string]interface{}{
		"index": index,
	}
	found, err := m.FindOne(config.FLICKR_DB_COLLECTION, elem_condition, &flickr_collection)

	if err != nil {
		cerr = &lib.CError{}
		cerr.SetMessage(err.Error())
		return
	}

	images, rerr := flickr.SearchImages(lat, lng, date, distance)
	if rerr != nil {
		cerr = rerr
		return
	}
	if !found {
		flickr_collection.Index = index
		flickr_collection.DateStr = date
		flickr_collection.LocationHash = loc_hash
		flickr_collection.Images = append(flickr_collection.Images, images...)
		err = m.Insert(config.FLICKR_DB_COLLECTION, flickr_collection)
		if err != nil {
			cerr = &lib.CError{}
			cerr.SetMessage(err.Error())
			return
		}
	} else {
		var (
			temp_images []APIs.ApiImage
			contains    bool
		)
		for _, t := range images {
			contains = false
			for i, v := range flickr_collection.Images {
				if t.Id == v.Id {
					flickr_collection.Images[i] = t
					contains = true
					break
				}
			}
			if !contains {
				temp_images = append(temp_images, t)
			}
		}
		flickr_collection.Images = append(flickr_collection.Images, temp_images...)
		err = m.Update(config.FLICKR_DB_COLLECTION, elem_condition, flickr_collection)
		if err != nil {
			cerr = &lib.CError{}
			cerr.SetMessage(err.Error())
			return
		}
	}

	return
}
