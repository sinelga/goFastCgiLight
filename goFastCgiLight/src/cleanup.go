package main

import (
	"encoding/csv"
	"flag"
	"io"
	"log"
	"log/syslog"
	"orphance"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var hoursFlag *int = flag.Int("hours", 0, "Hours was not visited")

func main() {
	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")

	}

	flag.Parse()
	var hoursFlagint int

	if *hoursFlag > 0 {

		hoursFlagint = *hoursFlag

	} else {

		oldfile := false

		if _, err := os.Stat("cleanupspace.csv"); err != nil {
			if os.IsNotExist(err) {

				golog.Info("cleanupspace.csv dont exist create default")
				csvFile, err := os.Create("cleanupspace.csv")
				defer csvFile.Close()
				if err != nil {

					golog.Crit("cleanupspace: " + err.Error())
				}
				writer := csv.NewWriter(csvFile)

				csvstring := []string{"800"}

				err = writer.Write(csvstring)
				if err != nil {

					golog.Crit("cleanupspace: " + err.Error())
				}
				writer.Flush()
				csvFile.Close()

			}
		}
		finfo, err := os.Stat("cleanupspace.csv")
		if err != nil {
			golog.Crit("cleanupspace: " + err.Error())
		}

		lasmod := finfo.ModTime().Unix()

		if (time.Now().Unix() - lasmod) > 84400 {
			golog.Info("old don't change createdflagint in cleanupspace.csv")
			oldfile = true
		} else {

			golog.Info("resently updated (less 1 day)  cleanupspace.csv modify !!")
		}

		csvFile, err := os.Open("cleanupspace.csv")
		defer csvFile.Close()
		if err != nil {

			golog.Crit("cleanupspace: " + err.Error())
		}
		csvReader := csv.NewReader(csvFile)

		var hours int
		for {
			fields, err := csvReader.Read()
			if err == io.EOF {
				break
			} else if err != nil {
				golog.Crit(err.Error())
			}

			hours, err = strconv.Atoi(fields[0])
			if err != nil {

				golog.Crit("cleanupspace: " + err.Error())
			}
			golog.Info("cleanup: hours from file: " + fields[0])

		}

		if !oldfile && (hours > 9) {

			csvFile, err = os.Create("cleanupspace.csv")
			defer csvFile.Close()
			if err != nil {

				golog.Crit("cleanupspace: " + err.Error())
			}

			hours = hours - 5

			writer := csv.NewWriter(csvFile)

			csvstring := []string{strconv.Itoa(hours)}

			err = writer.Write(csvstring)
			if err != nil {

				golog.Crit("cleanupspace: " + err.Error())
			}
			writer.Flush()
			csvFile.Close()

		}
		hoursFlagint = hours

	}

	golog.Info("cleanup: so final hours " + strconv.Itoa(hoursFlagint))
	startCleanup(*golog, hoursFlagint)

}

func startCleanup(golog syslog.Writer, hours int) {

	hoursint64 := float64(hours)

	var numScanned = 0
	var i64size int64 = 10000
	var i64sizesmall int64 = 1000
	var i64sizebig int64=80000

	var scan = func(path string, fileInfo os.FileInfo, inpErr error) (err error) {
		numScanned++

		if !fileInfo.IsDir() {
			filessize := strconv.FormatInt(fileInfo.Size(), 10)

			if fileInfo.Size() < i64sizesmall {
				golog.Info("Small < 1000 filessize delete " + path + " " + filessize)
				orphance.LookUp(golog, path)

			} else {

				if hoursint64 < time.Since(fileInfo.ModTime()).Hours() {

					if fileInfo.Size() < i64size {

						orphance.LookUp(golog, path)
					} else if fileInfo.Size() > i64sizebig {
											
						golog.Info("Not visited but BIG size 80000 delete "+path+" "+filessize)
						orphance.LookUp(golog, path)
					
					
					}

				}
			}
		}
		return
	}

	err := filepath.Walk("www", scan)
	if err != nil {

		golog.Err(err.Error())
	}

	golog.Info("cleanup:Total files scanned " + strconv.Itoa(numScanned))

}
