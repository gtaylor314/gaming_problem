package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

const timeForm = "15:04"

func main() {
	// given two strings in the format of HH:MM, determine the number of 15 minute games that can be played from A to B
	// games require a full 15 minutes and start every quarter hour (no earlier) - in other words 12:02 - 12:17 would be
	// zero playable games as the games would run from 12:00 - 12:15 and 12:15 - 12:30 and so on
	// time is in 24-hr format, if A == B then no games are playable, if A > B, then time has rolled over to the next day

	timeA := "17:01"
	timeB := "19:02"

	games := numOfGames(timeA, timeB)
	fmt.Printf("From time %s to time %s, there are %d playable games", timeA, timeB, games)
}

func numOfGames(A string, B string) int {
	// if timeA == timeB, then no games are playable
	if strings.Compare(A, B) == 0 {
		fmt.Println("timeA equals timeB")
		return 0
	}

	// parse both A and B and log and exit on errors
	timeA, err := time.Parse(timeForm, A)
	if err != nil {
		log.Println("error parsing timeA: " + err.Error())
		return 0
	}
	timeB, err := time.Parse(timeForm, B)
	if err != nil {
		log.Println("error parsing timeB: " + err.Error())
		return 0
	}

	numOfGamesPlayed := 0

	// if timeA is before timeB, any game play occurs within the same day
	if timeA.Before(timeB) {

		timeA = adjustTimeA(timeA)
		timeB = adjustTimeB(timeB)
		// fmt.Println(timeA)
		// fmt.Println(timeB)

		// calculate the duration of time between timeB and timeA - in hms format
		durationOfGamePlay := timeB.Sub(timeA)
		// convert duration to total number of minutes
		durationOfGamePlayMinutes := durationOfGamePlay.Minutes()
		// divide the total number of minutes by 15 to determine how many 15 minute games can be played
		numOfGamesPlayed = int(durationOfGamePlayMinutes / 15.0)

		// in the event that timeA is within a few minutes of timeB:
		// timeA will be rounded up to the next hour/quarter-hour
		// and timeB will be rounded down to the nearest hour/quarter-hour
		// which will return a negative duration - in this case, return 0
		if numOfGamesPlayed < 0 {
			return 0
		}

		return numOfGamesPlayed

	}

	// if timeA does not equal timeB, and timeA is not before timeB, then the game play goes from one day into the next
	timeA = adjustTimeA(timeA)
	timeB = adjustTimeB(timeB)
	// fmt.Println(timeA)
	// fmt.Println(timeB)

	// since the date of timeA and timeB are not given, we set them as both January 1, 0000
	// this allows us to use two different forms of expressing midnight (24:00 and 00:00)
	// the 24 hour expression auto updates to January 2, 0000 which we can use to calculate timeA's duration
	// the 00 hour expression allows us to calculate timeB's duration
	timeMidNight24 := time.Date(0000, 01, 01, 24, 0, 0, 0, time.UTC)
	timeMidNight00 := time.Date(0000, 01, 01, 0, 0, 0, 0, time.UTC)

	durationTimeA := timeMidNight24.Sub(timeA)
	durationTimeB := timeB.Sub(timeMidNight00)
	// total the duration in hms format
	totalDuration := durationTimeA + durationTimeB
	// convert the duration to minutes
	durationOfGamePlayMinutes := totalDuration.Minutes()
	// divide the total number of minutes by 15 to determine how many 15 minute games can be played
	numOfGamesPlayed = int(durationOfGamePlayMinutes / 15.0)

	return numOfGamesPlayed
}

// adjustTimeA() rounds timeA up to the next hour/quarter-hour if needed
func adjustTimeA(timeA time.Time) time.Time {
	// grab the minutes from timeA
	timeAMinutes := timeA.Minute()
	// mod the minutes with 15 to determine if timeA needs to be adjusted to the next hour/quarter-hour
	timeAMod := timeAMinutes % 15

	// if the mod returns zero, timeA does not need to be adjusted up
	if timeAMod != 0 {
		// subtract the mod result from 15 to determine what amount the time needs to adjusted
		timeAChange := 15 - timeAMod
		// add the additional minutes to round the time up to the next hour/quarter-hour mark
		timeA = timeA.Add(time.Minute * time.Duration(timeAChange))
	}
	return timeA
}

// adjustTimeB() rounds timeB down to the nearest hour/quarter-hour if needed
func adjustTimeB(timeB time.Time) time.Time {
	// grab the minutes from timeB
	timeBMinutes := timeB.Minute()
	// mod the minutes with 15 to determine how far off the hour/quarter-hour the time is
	timeBMod := timeBMinutes % 15
	// subtract that amount via adding a negative value (timeBMod will be zero if already on the hour/quarter-hour)
	timeB = timeB.Add(-(time.Minute * time.Duration(timeBMod)))

	return timeB
}
