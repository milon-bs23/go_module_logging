package teq_log

import "log"

func PrintLog(userID int, userName string, err error) {
	log.Println("[", userID, ":", userName, "] ", err)
}
