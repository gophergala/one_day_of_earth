package cronTask

import (
	"config"
	"fmt"
	"lib"
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
		var err *lib.CError
		err = RunCronTask(config.TWITTER_CRON, "37.7624499", "-122.4602593", lib.YesterdayTime().Format("2006-01-02"), "1000")

		if err != nil {
			fmt.Println(err.Message())
		} else {
			fmt.Println("Twitter Done")
		}

		err = RunCronTask(config.YOUTUBE_CRON, "37.7624499", "-122.4602593", lib.YesterdayTime().Format(time.RFC3339), "1000")

		if err != nil {
			fmt.Println(err.Message())
		} else {
			fmt.Println("Youtube Done")
		}

		err = RunCronTask(config.INSTAGRAM_CRON, "37.7624499", "-122.4602593", strconv.Itoa(lib.YesterdayTime().Second()), "5000")

		if err != nil {
			fmt.Println(err.Message())
		} else {
			fmt.Println("Instagram Done")
		}

		err = RunCronTask(config.FLICKR_CRON, "37.7624499", "-122.4602593", lib.YesterdayTime().Format("2006-01-02"), "20")

		if err != nil {
			fmt.Println(err.Message())
		} else {
			fmt.Println("Flickr Done")
		}

		time.Sleep(time.Duration(sleep_seconds) * time.Second)
	}
}
