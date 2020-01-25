package ksoftgo

import (
	"net/url"
	"strconv"
)

var (
	EndpointRest = "https://api.ksoft.si/"

	EndpointMemeRandomMeme = func(tag string, nsfw bool) string {
		return EndpointRest + "meme/random-meme?tag=" + tag + "&nsfw=" + strconv.FormatBool(nsfw)
	}
	EndpointMemeRandomImage = func(tag string, nsfw bool) string {
		return EndpointRest + "meme/random-image?tag=" + tag + "&nsfw=" + strconv.FormatBool(nsfw)
	}
	EndpointMemeImage        = func(snowflake string) string { return EndpointRest + "meme/image/" + snowflake }
	EndpointMemeWikihow      = EndpointRest + "meme/random-wikihow"
	EndpointMemeTags         = EndpointRest + "meme/tags"
	EndpointMemeRandomAww    = EndpointRest + "meme/random-aww"
	EndpointMemeRandomReddit = func(sub string) string { return EndpointRest + "meme/rand-reddit/" + sub }
	EndpointMemeRandomNSFW   = EndpointRest + "meme/random-nsfw"

	EndpointBansAdd    = EndpointRest + "bans/add"
	EndpointBansInfo   = func(id int64) string { return EndpointRest + "bans/info?user=" + strconv.FormatInt(id, 10) }
	EndpointBansCheck  = func(id int64) string { return EndpointRest + "bans/check?user=" + strconv.FormatInt(id, 10) }
	EndpointBansDelete = func(user int64, force bool) string {
		return EndpointRest + "bans/delete?user=" + string(user) + "&force=" + strconv.FormatBool(force)
	}
	EndpointBansList = func(page, perpage int) string {
		return EndpointRest + "bans/list?page=" + strconv.Itoa(page) + "&per_page=" + strconv.Itoa(perpage)
	}

	EndpointKumoGis     = func(param ParamGIS) string { return EndpointRest + "kumo/gis?q=" + url.QueryEscape(param.Location) }
	EndpointKumoWeather = func(param ParamWeather) string {
		return EndpointRest + "kumo/weather/" + param.ReportType + "?q=" + url.QueryEscape(param.Location)
	}
	EndpointKumoWeatherAdv = func(param ParamAdvWeather) string {
		return EndpointRest + "kumo/weather/" + strconv.FormatFloat(param.Latitude, 'f', -1, 64) + "," + strconv.FormatFloat(param.Longitude, 'f', -1, 64) + "/" + param.ReportType
	}
	EndpointKumoGeoip    = func(ip string) string { return EndpointRest + "kumo/geoip?ip=" + ip }
	EndpointKumoCurrency = func(param ParamCurrency) string {
		return EndpointRest + "kumo/currency?from=" + param.From + "&to=" + param.To + "&value=" + string(param.Value)
	}

	EndpointLyricsSearch = func(param ParamSearchLyrics) string {
		return EndpointRest + "lyrics/search?q=" + url.QueryEscape(param.Query)
	}
	EndpointLyricsArtist         = func(id int64) string { return EndpointRest + "lyrics/artist/" + strconv.FormatInt(id, 10) }
	EndpointLyricsAlbum          = func(id int64) string { return EndpointRest + "lyrics/album/" + strconv.FormatInt(id, 10) }
	EndpointLyricsTrack          = func(id int64) string { return EndpointRest + "lyrics/track/" + strconv.FormatInt(id, 10) }
	EndpointMusicRecommendations = EndpointRest + "music/recommendations"
)
