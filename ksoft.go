package ksoftgo

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strings"
	"time"
)

// Following https://semver.org/
const VERSION string = "2.0.0"

// New creates a new KSoft instance.
func New(token string) (s *KSession, err error) {
	s = &KSession{
		Client:    &http.Client{Timeout: 30 * time.Second},
		UserAgent: "KSoftgo (https://github.com/KSoft-Si/KSoftgo, v" + VERSION + ")",
	}

	if token == "" {
		err = fmt.Errorf("no token provided")
		return
	} else {
		s.Token = token
	}
	return
}

// Get a random image
// Example:
//		image, err := ksession.RandomImage(kosftgo.ParamRandomImage{Tag: "doge"})
func (s *KSession) RandomImage(tag ParamRandomImage) (i Image, err error) {
	i = Image{}
	res, err := s.request("GET", EndpointMemeRandomImage(tag), nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(res, &i)
	return
}

// Get a random meme
// Example:
//		reddit, err := ksession.RandomMeme()
func (s *KSession) RandomMeme() (r Reddit, err error) {
	r = Reddit{}
	res, err := s.request("GET", EndpointMemeRandomMeme, nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(res, &r)
	return
}

// Get a random picture that makes you say awwwww
// Example:
//		reddit, err := ksession.RandomAww()
func (s *KSession) RandomAww() (reddit Reddit, err error) {
	reddit = Reddit{}
	res, err := s.request("GET", EndpointMemeRandomAww, nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(res, &reddit)
	return
}

// Get a random reddit post
// Example:
//		reddit, err := ksession.RandomReddit(ksoftgo.ParamRandomReddit{SubReddit: "memes", Options: ksoftgo.OptionalRandomReddit{Span: "month"}})
func (s *KSession) RandomReddit(param ParamRandomReddit) (reddit Reddit, err error) {
	reddit = Reddit{}
	res, err := s.request("GET", EndpointMemeRandomReddit(param), nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(res, &reddit)
	return
}

// Get a random NSFW post
// Example:
//		reddit, err := ksession.RandomNSFW()
func (s *KSession) RandomNSFW() (reddit Reddit, err error) {
	reddit = Reddit{}
	res, err := s.request("GET", EndpointMemeRandomNSFW(ParamRandomNSFW{}), nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(res, &reddit)
	return
}

// Get a random NSFW post with options
// Example:
//		reddit, err := ksession.RandomNSFW(ksoftgo.ParamRandomNSFW{GIFsOnly: true})
func (s *KSession) RandomNSFWOptions(options ParamRandomNSFW) (reddit Reddit, err error) {
	reddit = Reddit{}
	res, err := s.request("GET", EndpointMemeRandomNSFW(options), nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(res, &reddit)
	return
}

// Get a random WikiHow article
// Example:
//		image, err := ksession.RandomWikiHow()
func (s *KSession) RandomWikiHow() (i WikiHowImage, err error) {
	i = WikiHowImage{}
	res, err := s.request("GET", EndpointMemeWikihow(ParamWikiHow{}), nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(res, &i)
	return
}

// Get a random WikiHow article with options
// Example:
//		image, err := ksession.RandomWikiHow(ksoftgo.ParamWikiHow{NSFW: true})
func (s *KSession) RandomWikiHowOptions(options ParamWikiHow) (i WikiHowImage, err error) {
	i = WikiHowImage{}
	res, err := s.request("GET", EndpointMemeWikihow(options), nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(res, &i)
	return
}

// Get an image by it's snowflake
// Example:
//		image, err := ksession.ImageBySnowflake("i-ix63ra_m-12")
func (s *KSession) ImageBySnowflake(snowflake string) (i Image, err error) {
	i = Image{}
	res, err := s.request("GET", EndpointMemeImage(snowflake), nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(res, &i)
	return
}

// Get tags
// Example:
//		tags, err := ksession.GetTags()
func (s *KSession) GetTags() (tags Tags, err error) {
	tags = Tags{}
	res, err := s.request("GET", EndpointMemeTags, nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(res, &tags)
	return
}

// Add a ban to the ban list
// Example:
//		err := ksession.AddBan(ksoftgo.ParamAddBan{ID: 0000000000000000, Reason: "bad guy", Proof: "imgur.com"})
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

// Get ban info
// Example:
//		baninfo, err := ksession.GetBanInfo(ksoftgo.ParamBans{UserID: "123456789123456789"})
func (s *KSession) GetBanInfo(param ParamBans) (info BanInfo, err error) {
	info = BanInfo{}
	res, err := s.request("GET", EndpointBansInfo(param), nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(res, &info)
	return
}

// Check user
// Example:
//		isbanned, err := ksession.CheckBan(ksoftgo.ParamBans{UserID: "123456789123456789"})
func (s *KSession) CheckBan(param ParamBans) (c bool, err error) {
	res, err := s.request("GET", EndpointBansCheck(param), nil)
	if err != nil {
		return
	}

	bc := BanCheck{}
	err = json.Unmarshal(res, &bc)
	return bc.Banned, err
}

// Delete ban
// Example:
//		ksession.DeleteBan(ksoftgo.ParamDeleteBan{User: "123456789123456789", Force: false})
func (s *KSession) DeleteBan(delete ParamDeleteBan) {
	_, err := s.request("DELETE", EndpointBansDelete(delete), nil)
	if err != nil {
		return
	}
}

// List of bans
// Example:
//		banlist, err := ksession.GetBans(ksoftgo.ParamListBans{Page: 1})
func (s *KSession) GetBans(param ParamListBans) (banlist BansList, err error) {
	banlist = BansList{}
	res, err := s.request("GET", EndpointBansList(param), nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(res, &banlist)
	return
}

// Search for locations and get maps
// Example:
//		gis, err := ksession.GetGis(ksoftgo.ParamGIS{Location: "Montreal"})
func (s *KSession) GetGIS(params ParamGIS) (gis GIS, err error) {
	gis = GIS{}

	res, err := s.request("GET", EndpointKumoGis(params), nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(res, &gis)
	return
}

// Weather - easy
// Example:
//		weather, err := ksession.GetWeather(ksoftgo.ParamWeather{Location: "Montreal", ReportType: "currently"})
func (s *KSession) GetWeather(params ParamWeather) (weather Weather, err error) {
	weather = Weather{}

	res, err := s.request("GET", EndpointKumoWeather(params), nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(res, &weather)
	return
}

// Weather - advanced
// Example:
//		weather, err := ksession.GetAdvWeather(ksoftgo.ParamAdvWeather{Latitude: 0.0, Longitude: 0.0, ReportType: "currently"})
func (s *KSession) GetAdvWeather(params ParamAdvWeather) (weather Weather, err error) {
	weather = Weather{}

	res, err := s.request("GET", EndpointKumoWeatherAdv(params), nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(res, &weather)
	return
}

// GeoIP
// Example:
//		geoip, err := ksession.GeoIP(ksoftgo.ParamIP{IP: "8.8.8.8"})
func (s *KSession) GeoIP(param ParamIP) (geoip GeoIP, err error) {
	geoip = GeoIP{}
	res, err := s.request("GET", EndpointKumoGeoIP(param), nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(res, &geoip)
	return
}

// Currency conversion
// Example:
//		currency, err := ksession.CurrenyConversion(ksoftgo.ParamCurrency{FromCurrency: "USD", ToCurrency: "EUR", Value: 1.50})
func (s *KSession) CurrencyConversion(param ParamCurrency) (curr Currency, err error) {
	curr = Currency{}
	res, err := s.request("GET", EndpointKumoCurrency(param), nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(res, &curr)
	return
}

// Get lyrics
// Example:
//		lyricssearch, err := ksession.SearchLyrics(ksoftgo.ParamSearchLyrics{Query: "Rick never gonna give you up"})
func (s *KSession) SearchLyrics(param ParamSearchLyrics) (results LyricsSearch, err error) {
	results = LyricsSearch{}

	res, err := s.request("GET", EndpointLyricsSearch(param), nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(res, &results)
	return
}

// Get artist by ID
// Example:
//		artist, err := ksession.GetArtist(628942)
func (s *KSession) GetArtist(id int64) (results Artist, err error) {
	results = Artist{}
	res, err := s.request("GET", EndpointLyricsArtist(id), nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(res, &results)
	return
}

// Get album by ID
// Example:
//		album, err := ksession.GetAlbum(88287)
func (s *KSession) GetAlbum(id int64) (results Album, err error) {
	results = Album{}
	res, err := s.request("GET", EndpointLyricsAlbum(id), nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(res, &results)
	return
}

// Get track by ID
// Example:
//		track, err := ksession.GetTrack(680639)
func (s *KSession) GetTrack(id int64) (results Track, err error) {
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
