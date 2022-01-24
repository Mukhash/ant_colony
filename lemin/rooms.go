package lemin

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

func isComment(str string) bool {
	if strings.HasPrefix(str, "#") {
		return true
	}
	return false
}

func findRoom(i int, arr []string, re *regexp.Regexp) (int, error) {
	var match bool
	for j := i + 1; j < len(arr); j++ {
		if arr[j] == "##end" || arr[j] == "##start" {
			return 0, errors.New("No room after " + arr[i])
		}
		if isComment(arr[j]) {
			continue
		}
		if match = re.MatchString(arr[j]); match {
			return j, nil
		}
	}
	return 0, errors.New("Invalid room input")
}

func populateRooms(strArr []string, buffer *Colony) (int, error) {
	linkReg, err := regexp.Compile("[ -~]+-[ -~]+")
	if err != nil {
		return -1, err
	}
	roomReg, err := regexp.Compile("[ -~]+ [0-9]+ [0-9]+")
	if err != nil {
		return -1, err
	}
	for i := 1; i < len(strArr); i++ {
		if match := linkReg.MatchString(strArr[i]); match {
			return i, nil
		}
		if strArr[i] == "##end" || strArr[i] == "##start" {
			ind, err := findRoom(i, strArr, roomReg)
			if err != nil {
				return -1, errors.New("No room after " + strArr[i])
			}
			if i+1 == len(strArr) {
				return -1, errors.New("No room after " + strArr[i])
			}
			if err := isStartEnd(strArr[i], strArr[ind], buffer, roomReg); err != nil {
				return -1, err
			}
			i = ind
			continue
		}
		if strings.HasPrefix(strArr[i], "#") {
			continue
		}
		room, err := getRoom(strArr[i], roomReg, buffer)
		if err != nil {
			return -1, err
		}
		buffer.rooms = append(buffer.rooms, room)
	}
	return -1, nil
}

func isStartEnd(startEnd string, roomStr string, buffer *Colony, re *regexp.Regexp) error {
	room, err := getRoom(roomStr, re, buffer)
	if err != nil {
		return err
	}
	buffer.rooms = append(buffer.rooms, room)
	if startEnd == "##end" {
		if buffer.endRoom.name != "" {
			return errors.New("More than one end room")
		}
		buffer.endRoom = room
	} else {
		if buffer.StartRoom.name != "" {
			return errors.New("More than one start room")
		}
		buffer.StartRoom = room
	}
	return err
}

func getRoom(str string, re *regexp.Regexp, buffer *Colony) (room, error) {
	res := room{}
	match := re.MatchString(str)
	if strings.HasPrefix(str, "L") || strings.HasPrefix(str, "#") {
		return res, errors.New("Invalid room: " + str)
	}
	if match {
		arr := strings.Split(str, " ")
		if len(arr) > 3 {
			return res, errors.New("Invalid room: " + str)
		}
		res.name = arr[0]
		res.X, _ = strconv.Atoi(arr[1])
		res.Y, _ = strconv.Atoi(arr[2])
	} else {
		return res, errors.New("Invalid room: " + str)
	}
	if isDuplicate(res, buffer) {
		return res, errors.New("Invalid room: " + str)
	}
	return res, nil
}

func isDuplicate(room room, buffer *Colony) bool {
	mapValue := buffer.uniqueRoomNames[room.name]
	if mapValue > 0 {
		return true
	}
	buffer.uniqueRoomNames[room.name] = mapValue + 1
	join := room.X*10 + room.Y
	mapValue = buffer.uniqueRoomDesc[join]
	if mapValue > 0 {
		return true
	}
	buffer.uniqueRoomDesc[join] = mapValue + 1
	return false
}
