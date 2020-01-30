package ksoftgo

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type KSession struct {
	Token          string
	Debug          bool
	Client         *http.Client
	UserAgent      string
	MaxRestRetries int
	RetryAfter     time.Duration
}

type Album struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Year   int    `json:"year"`
	Artist struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"artist"`
	Tracks []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"tracks"`
}

type Artist struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Albums []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Year int    `json:"year"`
	} `json:"albums"`
	Tracks []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"tracks"`
}

type APIErrorMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Image struct {
	URL       string `json:"url"`
	Snowflake string `json:"snowflake"`
	NSFW      bool   `json:"nsfw"`
	Tag       string `json:"tag"`
}
type Tags struct {
	Models []struct {
		Name string `json:"name"`
		Nsfw bool   `json:"nsfw"`
	} `json:"models"`
	Tags     []string `json:"tags"`
	NsfwTags []string `json:"nsfw_tags"`
}

type BanCheck struct {
	Banned bool `json:"is_banned"`
}

type BanInfo struct {
	ID            string      `json:"id"`
	Name          string      `json:"name"`
	Discriminator string      `json:"discriminator"`
	ModeratorID   string      `json:"moderator_id"`
	Reason        string      `json:"reason"`
	Proof         string      `json:"proof"`
	IsBanActive   bool        `json:"is_ban_active"`
	CanBeAppealed bool        `json:"can_be_appealed"`
	Timestamp     string      `json:"timestamp"`
	AppealReason  string      `json:"appeal_reason"`
	AppealDate    interface{} `json:"appeal_date"`
	RequestedBy   string      `json:"requested_by"`
	Exists        bool        `json:"exists"`
}

type BansList struct {
	BanCount     int         `json:"ban_count"`
	PageCount    int         `json:"page_count"`
	PerPage      int         `json:"per_page"`
	Page         int         `json:"page"`
	OnPage       int         `json:"on_page"`
	NextPage     int         `json:"next_page"`
	PreviousPage interface{} `json:"previous_page"`
	Data         []struct {
		ID            string      `json:"id"`
		Name          string      `json:"name"`
		Discriminator string      `json:"discriminator"`
		ModeratorID   string      `json:"moderator_id"`
		Reason        string      `json:"reason"`
		Proof         string      `json:"proof"`
		IsBanActive   bool        `json:"is_ban_active"`
		CanBeAppealed bool        `json:"can_be_appealed"`
		Timestamp     string      `json:"timestamp"`
		AppealReason  interface{} `json:"appeal_reason"`
		AppealDate    interface{} `json:"appeal_date"`
	} `json:"data"`
}

type Currency struct {
	Value  float64 `json:"value"`
	Pretty string  `json:"pretty"`
}

type GeoIP struct {
	Error bool `json:"error"`
	Code  int  `json:"code"`
	Data  struct {
		City          string      `json:"city"`
		ContinentCode string      `json:"continent_code"`
		ContinentName string      `json:"continent_name"`
		CountryCode   string      `json:"country_code"`
		CountryName   string      `json:"country_name"`
		DmaCode       interface{} `json:"dma_code"`
		Latitude      float64     `json:"latitude"`
		Longitude     float64     `json:"longitude"`
		PostalCode    string      `json:"postal_code"`
		Region        string      `json:"region"`
		TimeZone      string      `json:"time_zone"`
		Apis          struct {
			Weather       string `json:"weather"`
			Gis           string `json:"gis"`
			Openstreetmap string `json:"openstreetmap"`
			Googlemaps    string `json:"googlemaps"`
		} `json:"apis"`
	} `json:"data"`
}

type GIS struct {
	Error bool `json:"error"`
	Code  int  `json:"code"`
	Data  struct {
		Address     string   `json:"address"`
		Lat         float64  `json:"lat"`
		Lon         float64  `json:"lon"`
		BoundingBox []string `json:"bounding_box"`
		Type        []string `json:"type"`
		Map         string   `json:"map"`
	} `json:"data"`
}

type LyricsSearch struct {
	Total int `json:"total"`
	Took  int `json:"took"`
	Data  []struct {
		Artist      string  `json:"artist"`
		ArtistID    int     `json:"artist_id"`
		Album       string  `json:"album"`
		AlbumIds    string  `json:"album_ids"`
		AlbumYear   string  `json:"album_year"`
		Name        string  `json:"name"`
		Lyrics      string  `json:"lyrics"`
		SearchStr   string  `json:"search_str"`
		AlbumArt    string  `json:"album_art"`
		Popularity  int     `json:"popularity"`
		ID          string  `json:"id"`
		SearchScore float64 `json:"search_score"`
		URL         string  `json:"url"`
	} `json:"data"`
}

type Reddit struct {
	Title     string  `json:"title"`
	ImageURL  string  `json:"image_url"`
	Source    string  `json:"source"`
	Subreddit string  `json:"subreddit"`
	Upvotes   int     `json:"upvotes"`
	Downvotes int     `json:"downvotes"`
	Comments  int     `json:"comments"`
	CreatedAt float64 `json:"created_at"`
	NSFW      bool    `json:"nsfw"`
	Author    string  `json:"author"`
}

type Track struct {
	Name   string `json:"name"`
	Artist struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"artist"`
	Albums []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Year int    `json:"year"`
	} `json:"albums"`
	Lyrics string `json:"lyrics"`
}

type Weather struct {
	Error  bool `json:"error"`
	Status int  `json:"status"`
	Data   struct {
		Time                string      `json:"time"`
		Summary             string      `json:"summary"`
		Icon                string      `json:"icon"`
		PrecipIntensity     float64     `json:"precipIntensity"`
		PrecipProbability   float64     `json:"precipProbability"`
		Temperature         float64     `json:"temperature"`
		ApparentTemperature float64     `json:"apparentTemperature"`
		DewPoint            float64     `json:"dewPoint"`
		Humidity            float64     `json:"humidity"`
		Pressure            float64     `json:"pressure"`
		WindSpeed           float64     `json:"windSpeed"`
		WindGust            float64     `json:"windGust"`
		WindBearing         int         `json:"windBearing"`
		CloudCover          float64     `json:"cloudCover"`
		UvIndex             int         `json:"uvIndex"`
		Visibility          float64     `json:"visibility"`
		Ozone               float64     `json:"ozone"`
		SunriseTime         interface{} `json:"sunriseTime"`
		SunsetTime          interface{} `json:"sunsetTime"`
		IconURL             string      `json:"icon_url"`
		Alerts              []struct {
			Title       string   `json:"title"`
			Regions     []string `json:"regions"`
			Severity    string   `json:"severity"`
			Time        int      `json:"time"`
			Expires     int      `json:"expires"`
			Description string   `json:"description"`
			URI         string   `json:"uri"`
		} `json:"alerts"`
		Units    string `json:"units"`
		Location struct {
			Lat     float64 `json:"lat"`
			Lon     float64 `json:"lon"`
			Address string  `json:"address,omitempty"`
		} `json:"location"`
	} `json:"data"`
}

type ParamAddBan struct {
	ID            int64  `json:"user,omitempty"`
	Reason        string `json:"reason,omitempty"`
	Proof         string `json:"proof,omitempty"`
	Name          string `json:"user_name,omitempty"`
	Discriminator int    `json:"user_discriminator,omitempty"`
	ModeratorID   int64  `json:"mod,omitempty"`
	CanBeAppealed bool   `json:"appeal_possible,omitempty"`
}

type ParamAdvWeather struct {
	Latitude   float64
	Longitude  float64
	ReportType string
	Units      string
	Lang       string
	Icons      string
}

type ParamCurrency struct {
	From  string
	To    string
	Value string
}

type ParamDeleteBan struct {
	User  int64
	Force bool
}

type ParamGIS struct {
	Location   string
	Fast       bool
	More       bool
	Zoom       int
	IncludeMap bool
}

type ParamSearchLyrics struct {
	Query    string
	TextOnly bool
	Limit    int
}

type ParamTag struct {
	Name string
	NSFW bool
}

type ParamWeather struct {
	Location   string
	ReportType string
	Units      string
	Lang       string
	Icons      string
}

const (
	ErrCodeMissingParameters = 123
	ErrCodeInvalidValue      = 124
	ErrCodeAlreadyExists     = 125
)

func dumpAsValues(m map[string]interface{}) (data url.Values) {
	data = url.Values{}
	for k, v := range m {
		switch v.(type) {
		case float64:
			in, _ := ParseFloat(fmt.Sprintf("%v", v))
			data.Set(k, strconv.FormatFloat(in, 'f', -1, 64))
			break
		default:
			data.Set(k, fmt.Sprintf("%v", v))
		}
	}
	return
}

/**
 * Thanks bwmarrin
 **/

type RESTError struct {
	Request      *http.Request
	Response     *http.Response
	ResponseBody []byte

	Message *APIErrorMessage
}

func newRestError(req *http.Request, resp *http.Response, body []byte) *RESTError {
	restErr := &RESTError{
		Request:      req,
		Response:     resp,
		ResponseBody: body,
	}

	var msg *APIErrorMessage
	err := json.Unmarshal(body, &msg)
	if err == nil {
		restErr.Message = msg
	}

	return restErr
}

func (r RESTError) Error() string {
	return "HTTP " + r.Response.Status + ", " + string(r.ResponseBody)
}

/*
 * Thanks yyscamper (https://gist.github.com/yyscamper/5657c360fadd6701580f3c0bcca9f63a)
 */
func ParseFloat(str string) (float64, error) {
	val, err := strconv.ParseFloat(str, 64)
	if err == nil {
		return val, nil
	}

	//Some number may be seperated by comma, for example, 23,120,123, so remove the comma firstly
	str = strings.Replace(str, ",", "", -1)

	//Some number is specifed in scientific notation
	pos := strings.IndexAny(str, "eE")
	if pos < 0 {
		return strconv.ParseFloat(str, 64)
	}

	var baseVal float64
	var expVal int64

	baseStr := str[0:pos]
	baseVal, err = strconv.ParseFloat(baseStr, 64)
	if err != nil {
		return 0, err
	}

	expStr := str[(pos + 1):]
	expVal, err = strconv.ParseInt(expStr, 10, 64)
	if err != nil {
		return 0, err
	}

	return baseVal * math.Pow10(int(expVal)), nil
}
