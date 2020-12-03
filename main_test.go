package main

import "testing"

func TestGetNameAndNumber(t *testing.T) {
	contact := Contact{"45565594216", "kamoliddin", "xojiyev", 9940008293, "asaka"}
	name, number := contact.getNameAndNumber()
	if name != "kamoliddin" || number != 9940008293 {
		t.Errorf("Error Name  or Number is incorrect. \n Name: %q Number: %d \n But must be Name:kamoliddin, Number:9940008293", name, number)
	}
}

func TestGetTitleAndBody(t *testing.T) {
	task := Task{"55485784485", "Shopping", "Need do shopping and buy apple,....", 1}
	id, title, priority := task.getTitleAndBodyTask()

	if id != "55485784485" || title != "Shopping" || priority != 1 {
		t.Errorf("Error ID, Title or Priority is incorrect. \n "+
			"ID: %q, Title: %q , Priority: %d \n But must be ID:55485784485, "+
			"Title:Shopping, Priority:1", id, title, priority)
	}
}
