package ksoftgo

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const VERSION string = "1.0"

func New(token string) (s *KSession, err error) {
	s = &KSession{
		Client:    &http.Client{Timeout: 20 * time.Second},
		UserAgent: "KSoftgo",
		//UserAgent: "KSoftgo (https://github.com/Noctember/KSoftgo, v" + VERSION + ")",
		Debug: true,
	}

	if token == "" {
		err = fmt.Errorf("no token provided")
		return
	} else {
		s.Token = token
	}
	return
}

func (s *KSession) RandomImage(tag ParamTag) (i Image, err error) {
	i = Image{}
	res, err := s.request("GET", EndpointMemeRandomImage(tag.Name, tag.NSFW), nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(res, &i)
	return
}

func (s *KSession) RandomMeme(tag ParamTag) (i Image, err error) {
	i = Image{}
	res, err := s.request("GET", EndpointMemeRandomMeme(tag.Name, tag.NSFW), nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(res, &i)
	return
}

func (s *KSession) RandomAww() (reddit Reddit, err error) {
	reddit = Reddit{}
	res, err := s.request("GET", EndpointMemeRandomAww, nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(res, &reddit)
	return
}

func (s *KSession) RandomReddit(sub string) (reddit Reddit, err error) {
	reddit = Reddit{}
	res, err := s.request("GET", EndpointMemeRandomReddit(sub), nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(res, &reddit)
	return
}

func (s *KSession) RandomNSFW() (reddit Reddit, err error) {
	reddit = Reddit{}
	res, err := s.request("GET", EndpointMemeRandomNSFW, nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(res, &reddit)
	return
}

func (s *KSession) RandomWikiHow() (i Image, err error) {
	i = Image{}
	res, err := s.request("GET", EndpointMemeWikihow, nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(res, &i)
	return
}

func (s *KSession) ImageBySnowflake(snowflake string) (i Image, err error) {
	i = Image{}
	res, err := s.request("GET", EndpointMemeImage(snowflake), nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(res, &i)
	return
}

func (s *KSession) GetTags() (tags Tags, err error) {
	tags = Tags{}
	res, err := s.request("GET", EndpointMemeTags, nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(res, &tags)
	return
}

func (s *KSession) AddBan(info ParamAddBan) (err error) {
	jsonMap := make(map[string]interface{})
	data, err := json.Marshal(info)

	err = json.Unmarshal(data, &jsonMap)
	if err != nil {
		return
	}

	err = s.PostForm(EndpointBansAdd, dumpAsValues(jsonMap))
	return
}

func (s *KSession) GetBanInfo(user int64) (info BanInfo, err error) {
	info = BanInfo{}
	res, err := s.request("GET", EndpointBansInfo(user), nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(res, &info)
	return
}

func (s *KSession) CheckBan(user int64) (c bool, err error) {
	res, err := s.request("GET", EndpointBansCheck(user), nil)
	if err != nil {
		return
	}

	bc := BanCheck{}
	err = json.Unmarshal(res, &bc)
	return bc.Banned, err
}

func (s *KSession) DeleteBan(delete ParamDeleteBan) {
	_, err := s.request("DELETE", EndpointBansDelete(delete.User, delete.Force), nil)
	if err != nil {
		return
	}
}

func (s *KSession) GetBans(page, perpage int) (banlist BansList, err error) {
	banlist = BansList{}
	res, err := s.request("GET", EndpointBansList(page, perpage), nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(res, &banlist)
	return
}

func (s *KSession) GetGIS(params ParamGIS) (gis GIS, err error) {
	gis = GIS{}
	url := EndpointKumoGis(params)
	if params.Fast {
		url += "&fast=" + strconv.FormatBool(params.Fast)
	} else if params.IncludeMap {
		url += "&include_map=" + strconv.FormatBool(params.IncludeMap)
	} else if params.More {
		url += "&more=" + strconv.FormatBool(params.More)
	} else if params.Zoom != 0 {
		url += "&map_zoom=" + string(params.Zoom)
	}

	res, err := s.request("GET", url, nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(res, &gis)
	return
}

func (s *KSession) GetWeather(params ParamWeather) (weather Weather, err error) {
	weather = Weather{}
	url := EndpointKumoWeather(params)
	if params.Units != "" {
		url += "&units=" + params.Units
	} else if params.Lang != "" {
		url += "&lang=" + params.Lang
	} else if params.Icons != "" {
		url += "&icons=" + params.Icons
	}

	res, err := s.request("GET", url, nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(res, &weather)
	return
}

func (s *KSession) GetAdvWeather(params ParamAdvWeather) (weather Weather, err error) {
	weather = Weather{}
	url := EndpointKumoWeatherAdv(params)
	if params.Units != "" {
		url += "&units=" + params.Units
	} else if params.Lang != "" {
		url += "&lang=" + params.Lang
	} else if params.Icons != "" {
		url += "&icons=" + params.Icons
	}

	res, err := s.request("GET", url, nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(res, &weather)
	return
}

func (s *KSession) GeoIP(ip string) (geoip GeoIP, err error) {
	geoip = GeoIP{}
	res, err := s.request("GET", EndpointKumoGeoip(ip), nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(res, &geoip)
	return
}

func (s *KSession) CurrencyConversion(param ParamCurrency) (curr Currency, err error) {
	curr = Currency{}
	res, err := s.request("GET", EndpointKumoCurrency(param), nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(res, &curr)
	return
}

func (s *KSession) SearchLyrics(param ParamSearchLyrics) (results LyricsSearch, err error) {
	results = LyricsSearch{}
	url := EndpointLyricsSearch(param)
	if param.Limit != 0 {
		url += "&limit=" + string(param.Limit)
	} else if param.TextOnly {
		url += "&text_only=" + strconv.FormatBool(param.TextOnly)
	}

	res, err := s.request("GET", url, nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(res, &results)
	return
}

func (s *KSession) GetArtist(id int) (results Artist, err error) {
	results = Artist{}
	res, err := s.request("GET", EndpointLyricsArtist(id), nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(res, &results)
	return
}

func (s *KSession) GetAlbum(id int) (results Album, err error) {
	results = Album{}
	res, err := s.request("GET", EndpointLyricsAlbum(id), nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(res, &results)
	return
}

func (s *KSession) GetTrack(id int) (results Track, err error) {
	results = Track{}
	res, err := s.request("GET", EndpointLyricsTrack(id), nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(res, &results)
	return
}

func (s *KSession) log(caller int, format string, a ...interface{}) {
	pc, file, line, ok := runtime.Caller(caller)
	fmt.Println(ok)
	files := strings.Split(file, "/")
	file = files[len(files)-1]

	name := runtime.FuncForPC(pc).Name()
	fns := strings.Split(name, ".")
	name = fns[len(fns)-1]

	msg := fmt.Sprintf(format, a...)

	log.Printf("[KSoft] %s:%d:%s() %s\n", file, line, name, msg)
}
