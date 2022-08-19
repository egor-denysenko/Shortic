package shortener

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	apiUtility "shortic/pkg/api/utility"
	"shortic/pkg/dbservice"
	"strings"
	"time"
)

const ErrMarshallingJson = "Error provviding the shorted url."
const ErrWrongApiVersion = "Wrong Api Version Used."
const ErrUrlNotValid = "Provvided url is not valid."

// Struct for interpreting the excpected POST request body.
type postUrlBody struct {
	Url string `json:"url"`
}

// The response struct that will be sent in JSON format to the client.
type responsePostBody struct {
	ShortenUrl string `json:"shortenUrl"`
}

// Shortic  ... Portable url shortener
// @Summary
// @Description
// @Tags Shortener
// @Param Url body postUrlBody true "Url"
// @Success 200 {object} responsePostBody
// @Failure 501 {object} object
// @Failure 410 {object} object
// @Failure 400 {object} object
// @Router /data/shortener [post]
func FormUrlShortenerHandler(dbservice *dbservice.DatabaseService) http.HandlerFunc {
	// Wrapper in order to accept and use the redis connection provvided.
	return func(w http.ResponseWriter, r *http.Request) {

		// Check if only the supported method is called with the proper helper function
		if !apiUtility.ApiAllowedMethods(r.Method, apiUtility.AllowedMethods{"POST": {}}) {
			w.WriteHeader(http.StatusNotImplemented)
		}

		// Confront if the Accept-Version header is compatible with the enpoint version
		if !apiUtility.ApiVersionHeaderCheck(r, "0.0.1") {
			http.Error(w, ErrWrongApiVersion, http.StatusGone)
		}

		var urlBodyStruct postUrlBody

		// Decode body payload into golang struct
		errDecodingJson := json.NewDecoder(r.Body).Decode(&urlBodyStruct)
		if errDecodingJson != nil {
			http.Error(w, errDecodingJson.Error(), http.StatusBadRequest)
			return
		}

		if !isUrlValid(urlBodyStruct.Url) {
			http.Error(w, ErrUrlNotValid, http.StatusInternalServerError)
			return
		}

		md5UrlHash := urlToMd5(urlBodyStruct.Url)

		var shortenUrl string
		for {
			shortenUrl = reduceUrl(md5UrlHash)
			if verifyUrlUniqueness(dbservice, shortenUrl) {
				dbservice.SaveUrl(urlBodyStruct.Url, shortenUrl)
				break
			}
		}

		shortenUrl = os.Getenv("ShortenerUrl") + "/" + shortenUrl

		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")

		response := responsePostBody{ShortenUrl: shortenUrl}
		//log.Default().Fatalln("qui va dopo url check")

		jsonResp, err := json.Marshal(response)
		if err != nil {
			http.Error(w, ErrMarshallingJson, http.StatusInternalServerError)
			return
		}

		w.Write(jsonResp)
		return
	}
}

// Helper funciton that checks if the provvided url is a valid url
func isUrlValid(fullUserUrl string) bool {
	_, err := url.ParseRequestURI(fullUserUrl)

	if err != nil {
		return false
	}

	return true
}

func reduceUrl(fullUserUrl string) string {
	var b strings.Builder
	hashRune := []rune(urlToMd5(fullUserUrl))

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 7; i++ {
		randomIndex := rand.Intn(len(hashRune))
		b.WriteRune(hashRune[randomIndex])
	}

	return b.String()
}

func urlToMd5(url string) string {
	hash := md5.Sum([]byte(url))
	return hex.EncodeToString(hash[:])
}

func verifyUrlUniqueness(dbservice *dbservice.DatabaseService, computedUrl string) bool {
	if dbservice.CheckUrlCollision(computedUrl) {
		return true
	}
	return false
}
