package teq_log

import "log"

func PrintLog(userID int64, userName string, err string) {
	log.Println("[", userID, ":", userName, "] ", err)
}
