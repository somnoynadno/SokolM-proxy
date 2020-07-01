package main

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"sokol_proxy/api"
	"sokol_proxy/db"
	"sokol_proxy/models"
	"time"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func main() {
	log.Info("everything is fine")
	for true {
		log.Debug("sleeping for 1 hour")
		time.Sleep(time.Hour * 1)

		// get access cookie
		r, err := api.Login()
		if err != nil {
			log.Error(err)
			continue
		}

		if r.StatusCode != 200 {
			log.Error("login status code is " + r.Status)
			continue
		} else {
			log.Info("login succeed")
		}

		// fetch data for each device
		for _, d := range db.Devices {
			log.Info("fetching data for device " + d.ID)
			resp, err := api.GetDailyData(d.ID, r.Cookies())
			if err != nil {
				log.Error(err)
				continue
			}

			if resp.StatusCode != 200 {
				log.Error("fetching data status code is " + resp.Status)
				continue
			}

			var sokolData []models.SokolM
			err = json.NewDecoder(resp.Body).Decode(&sokolData)
			if err != nil {
				log.Error(err)
				continue
			}

			filteredData := filterSokolData(sokolData)

			if len(filteredData) != 0 {
				log.Debug("last measurement:, ", fmt.Sprintf("%+v\n", filteredData[len(filteredData)-1]))
				log.Debug("data size: ", len(filteredData))
			} else {
				log.Debug("no data received")
				continue
			}

			// insert fetched data
			err = db.InsertData(d.NumberInDB, filteredData)
			if err != nil {
				log.Error(err)
				continue
			}
		}
	}
}

func filterSokolData(ss []models.SokolM) (ret []models.SokolM) {
	for _, s := range ss {
		if s.TR != 0 {
			s.Time = s.Date.Unix()
			ret = append(ret, s)
		}
	}
	return ret
}
