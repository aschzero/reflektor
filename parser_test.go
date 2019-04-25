package reflektor

import (
	"testing"
)

var validJobYaml = []byte("jobs:\n  - name: recipes\n    source: /my/recipes\n    schedule: \"* * * * *\"\n\n  " +
	"- name: docs\n    source: /my/docs\n    schedule: \"* * * * *\"")

var expectedJobCollection = []*Job{
	{Name: "recipes", SourcePath: "/my/recipes", Schedule: "@monthly"},
	{Name: "documents", SourcePath: "/my/docs", Schedule: "@weekly"},
}

func TestNewJobCollection(t *testing.T) {
	jobs, err := NewJobCollection(validJobYaml)

	if err != nil {
		t.Error(err)
	}

	if len(jobs) != len(expectedJobCollection) {
		t.Errorf("Unexpected amount of parsed jobs, got %d", len(jobs))
	}

	var actual *Job
	var actual2 *Job

	expected := expectedJobCollection[0]
	expected2 := expectedJobCollection[1]

	for _, j := range jobs {
		if j.Name == expected.Name {
			actual = j
			return
		}

		if j.Name == expected2.Name {
			actual2 = j
			return
		}
	}

	testJobAttributes(actual, expected, t)
	testJobAttributes(actual2, expected2, t)
}

func TestValidate(t *testing.T) {
	jc := &JobConfig{
		SourcePath: "",
		Schedule:   "@weekly",
	}

	jc2 := &JobConfig{
		SourcePath: "/abc",
		Schedule:   "",
	}

	err := jc.Validate()
	if err == nil {
		t.Error("Expected error, got nil")
	}

	err = jc2.Validate()
	if err == nil {
		t.Error("Expected error, got nil")
	}
}

func testJobAttributes(actual *Job, expected *Job, t *testing.T) {
	if actual == nil {
		t.Error("Got nil job")
	}

	if actual.Name != expected.Name {
		t.Errorf("Unexpected job name, got %s", actual.Name)
	}

	if actual.SourcePath != expected.SourcePath {
		t.Errorf("Unexpected job source, got %s", actual.Name)
	}

	if actual.Schedule != expected.Schedule {
		t.Errorf("Unexpected job schedule, got %s", actual.Name)
	}
}
