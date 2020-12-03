package main

func main() {

}

//Contact ...
type Contact struct {
	ID         string
	name       string
	familyName string
	number     int
	address    string
}

// Task ...
type Task struct {
	ID       string
	title    string
	body     string
	priority int
}

func (c *Contact) getNameAndNumber() (string, int) {

	return c.name, c.number
}
func (t *Task) getTitleAndBodyTask() (string, string, int) {
	return t.ID, t.title, t.priority
}
