package cronTask

import (
	"config"
	"lib"
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
