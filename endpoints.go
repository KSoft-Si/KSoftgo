package ksoftgo

import (
	"fmt"
	"github.com/google/go-querystring/query"
	"net/url"
	"strconv"
)

var (
	EndpointRest = "https://api.ksoft.si/"

	// --------- IMAGES ENDPOINTS ----------------------------------------------

	EndpointMemeRandomMeme  = EndpointRest + "images/random-meme"
	EndpointMemeTags        = EndpointRest + "images/tags"
	EndpointMemeRandomAww   = EndpointRest + "images/random-aww"
	EndpointMemeImage       = func(snowflake string) string { return EndpointRest + "images/image/" + snowflake }
	EndpointMemeRandomImage = func(param ParamRandomImage) string {
		q, _ := query.Values(param)
		return EndpointRest + "images/random-image?" + q.Encode()
	}
	EndpointMemeWikihow = func(param ParamWikiHow) string {
		q, _ := query.Values(param)
		return EndpointRest + "images/random-wikihow?" + q.Encode()
	}
	EndpointMemeRandomReddit = func(param ParamRandomReddit) string {
		q, _ := query.Values(param.Options)
		return EndpointRest + "images/rand-reddit/" + param.SubReddit + "?" + q.Encode()
	}
	EndpointMemeRandomNSFW = func(param ParamRandomNSFW) string {
		q, _ := query.Values(param)
		return EndpointRest + "images/random-nsfw?" + q.Encode()
	}

	// --------- BANS ENDPOINTS ------------------------------------------------

	EndpointBansAdd  = EndpointRest + "bans/add"
	EndpointBansInfo = func(param ParamBans) string {
		q, _ := query.Values(param)
		return EndpointRest + "bans/info?" + q.Encode()
	}
	EndpointBansCheck = func(param ParamBans) string {
		q, _ := query.Values(param)
		return EndpointRest + "bans/check?" + q.Encode()
	}
	EndpointBansDelete = func(param ParamDeleteBan) string {
		q, _ := query.Values(param)
		return EndpointRest + "bans/delete?" + q.Encode()
	}
	EndpointBansList = func(param ParamListBans) string {
		q, _ := query.Values(param)
		return EndpointRest + "bans/list?" + q.Encode()
	}

	// --------- KUMO ENDPOINTS ------------------------------------------------

	EndpointKumoGis = func(param ParamGIS) string {
		return EndpointRest + "kumo/gis?q=" + url.QueryEscape(param.Location)
	}
	EndpointKumoWeather = func(param ParamWeather) string {
		return EndpointRest + "kumo/weather/" + param.ReportType + "?q=" + url.QueryEscape(param.Location)
	}
	EndpointKumoWeatherAdv = func(param ParamAdvWeather) string {
		q, _ := query.Values(param.Options)
		return EndpointRest + fmt.Sprintf("kumo/weather/%s,%s/%s?%s",
			strconv.FormatFloat(param.Latitude, 'f', -1, 64),
			strconv.FormatFloat(param.Longitude, 'f', -1, 64),
			param.ReportType, q.Encode())
	}
	EndpointKumoGeoIP = func(param ParamIP) string {
		q, _ := query.Values(param)
		return EndpointRest + "kumo/geoip?" + q.Encode()
	}
	EndpointKumoCurrency = func(param ParamCurrency) string {
		q, _ := query.Values(param)
		return EndpointRest + "kumo/currency?" + q.Encode()
	}

	// --------- MUSIC ENDPOINTS -----------------------------------------------

	EndpointLyricsSearch = func(param ParamSearchLyrics) string {
		q, _ := query.Values(param)
		return EndpointRest + "lyrics/search?" + q.Encode()
	}
	EndpointLyricsArtist         = func(id int64) string { return EndpointRest + "lyrics/artist/" + strconv.FormatInt(id, 10) }
	EndpointLyricsAlbum          = func(id int64) string { return EndpointRest + "lyrics/album/" + strconv.FormatInt(id, 10) }
	EndpointLyricsTrack          = func(id int64) string { return EndpointRest + "lyrics/track/" + strconv.FormatInt(id, 10) }
	EndpointMusicRecommendations = EndpointRest + "music/recommendations"
)
