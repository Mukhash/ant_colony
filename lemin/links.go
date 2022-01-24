package lemin

import (
	"errors"
	"regexp"
	"strings"
)

func getGraph(strArr []string, offset int, buffer *Colony) error {
	if offset == -1 {
		return errors.New("No links")
	}
	for i := offset; i < len(strArr); i++ {
		if strings.HasPrefix(strArr[i], "#") {
			continue
		}
		match, _ := regexp.MatchString("[ -~]+-[ -~]+", strArr[i])
		if !match {
			return errors.New("Invalid link: " + strArr[i])
		}
		if err := handleLink(strArr[i], buffer); err != nil {
			return err
		}
	}
	return nil
}

func handleLink(link string, buffer *Colony) error {
	rooms := strings.Split(link, "-")
	if rooms[0] == rooms[1] {
		return errors.New(rooms[0] + " room links to itself")
	}
	mapValue := buffer.uniqueRoomNames[rooms[0]]
	if mapValue == 0 {
		return errors.New("In link " + link + " no such room " + rooms[0])
	}
	mapValue = buffer.uniqueRoomNames[rooms[1]]
	if mapValue == 0 {
		return errors.New("In link " + link + " no such room " + rooms[1])
	}
	if !contains(rooms[1], buffer.graph[rooms[0]]) {
		buffer.graph[rooms[0]] = append(buffer.graph[rooms[0]], rooms[1])
	}
	if !contains(rooms[0], buffer.graph[rooms[1]]) {
		buffer.graph[rooms[1]] = append(buffer.graph[rooms[1]], rooms[0])
	}

	return nil
}
