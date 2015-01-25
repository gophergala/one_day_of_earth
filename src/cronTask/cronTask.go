package cronTask

import (
	"config"
	"fmt"
	"lib"
	"mongodatabase"
	"strconv"
	"time"
)

func RunCronTask(task_type int, lat, lng, date, distance string) (cerr *lib.CError) {
	cerr = nil
	switch task_type {
	case config.TWITTER_CRON:
		{
			cerr = Twitter_Cron(lat, lng, date, distance)
		}
	case config.YOUTUBE_CRON:
		{
			cerr = Youtube_Cron(lat, lng, date, distance)
		}
	case config.INSTAGRAM_CRON:
		{
			cerr = Instagram_Cron(lat, lng, date, distance)
		}
	case config.FLICKR_CRON:
		{
			cerr = Flickr_Cron(lat, lng, date, distance)
		}
	default:
		{
			cerr = &lib.CError{}
			cerr.SetMessage("Wrong Cron Task Command")
			return
		}
	}
	return
}

func StartCron(sleep_seconds int) {
	for {
		cities := make([]mongodatabase.CityCollection, 0)
		m := mongodatabase.Mongo{}
		m.Connect()
		found, cerr := m.FindAll(config.CITIES_DB_COLLECTION, map[string]interface{}{}, &cities)
		if cerr != nil {
			fmt.Println(cerr.Error())
			time.Sleep(time.Duration(sleep_seconds) * time.Second)
			continue
		}
		if !found {
			fmt.Println("No Cities Found")
			time.Sleep(time.Duration(sleep_seconds) * time.Second)
			continue
		}
		fmt.Println(len(cities))
		for _, city := range cities {
			fmt.Println(city.Name, city.Lat, city.Lng)
			var err *lib.CError

			go func() {
				err = RunCronTask(config.TWITTER_CRON, city.Lat, city.Lng, lib.YesterdayTime().Format("2006-01-02"), "1000")

				if err != nil {
					fmt.Println(err.Message())
					fmt.Println(city.Lat, city.Lng, "Tweet")
				}
			}()

			go func() {
				err = RunCronTask(config.YOUTUBE_CRON, city.Lat, city.Lng, lib.YesterdayTime().Format(time.RFC3339), "1000")

				if err != nil {
					fmt.Println(err.Message())
					fmt.Println(city.Lat, city.Lng, "TYoutue")
				}
			}()

			go func() {
				err = RunCronTask(config.INSTAGRAM_CRON, city.Lat, city.Lng, strconv.Itoa(lib.YesterdayTime().Second()), "5000")

				if err != nil {
					fmt.Println(err.Message())
					fmt.Println(city.Lat, city.Lng, "Insta")
				}
			}()

			go func() {
				err = RunCronTask(config.FLICKR_CRON, city.Lat, city.Lng, lib.YesterdayTime().Format("2006-01-02"), "20")

				if err != nil {
					fmt.Println(err.Message())
					fmt.Println(city.Lat, city.Lng, "Flick")
				}
			}()
			time.Sleep(10 * time.Second)
		}
		time.Sleep(time.Duration(sleep_seconds) * time.Second)
	}
}
