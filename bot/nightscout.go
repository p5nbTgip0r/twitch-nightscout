package bot

import (
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
	"twitch-nightscout/diabetes"
	"twitch-nightscout/nightscout"
)

type PresentableError struct {
	Msg string
}

func (e *PresentableError) Error() string {
	return e.Msg
}

type DiabetesSnapshot struct {
	Timestamp time.Time
	Glucose   diabetes.GlucoseMgdl
	Trend     nightscout.Trend
	Delta     diabetes.GlucoseMgdl
	Iob       string
	Cob       string
}

func doApi(req *nightscout.NSRequest) ([]byte, error) {
	client := &http.Client{}
	res, err := client.Do(req.AsRequest())
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("non-200 response code: %s", res.Status)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func getPebble(i *nightscout.Instance) (*DiabetesSnapshot, error) {
	var ss DiabetesSnapshot

	req := i.NewRequest("pebble")
	req.AddToken(i)
	res, err := doApi(req)
	if err != nil {
		return nil, err
	}
	// grab data for first bg
	raw := gjson.GetBytes(res, "bgs.0")
	if raw.Exists() {
		// ensure the timestamp is within 15 minutes
		pebbleTime := time.UnixMilli(raw.Get("datetime").Int())
		timeDiff := time.Since(pebbleTime)
		if timeDiff > (time.Minute * 15) {
			return nil, &PresentableError{
				Msg: fmt.Sprintf("BG data is too old: %s", timeDiff.Round(time.Second)),
			}
		}

		// assume units based on the presence of a decimal point in the glucose value
		var bgUnit diabetes.GlucoseUnit
		if strings.Contains(raw.Get("sgv").String(), ".") {
			bgUnit = diabetes.MMOL
		} else {
			bgUnit = diabetes.MGDL
		}

		var bgRaw = raw.Get("sgv").Float()
		var deltaRaw = raw.Get("bgdelta").Float()
		if bgUnit == diabetes.MGDL {
			ss.Glucose = diabetes.GlucoseMgdl(bgRaw)
			ss.Delta = diabetes.GlucoseMgdl(deltaRaw)
		} else {
			ss.Glucose = diabetes.GlucoseMmol(bgRaw).ToMgdl()
			ss.Delta = diabetes.GlucoseMmol(deltaRaw).ToMgdl()
		}

		ss.Trend = nightscout.Trends[raw.Get("trend").Int()]
		ss.Iob = raw.Get("iob").String()
		ss.Cob = raw.Get("cob").String()
		ss.Timestamp = time.UnixMilli(raw.Get("datetime").Int())
	}

	return &ss, nil
}
