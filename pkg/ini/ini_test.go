package ini_test

import (
	"os"
	"testing"

	"github.com/fabiant7t/hobot/pkg/ini"
)

func TestProgramatically(t *testing.T) {
	config := ini.New()
	section := config.DefaultSection()
	section.Set("global", "variable")
	section = config.Section("foo")
	section.Set("bar", "bar")
	section.Set("baz", "baz")
	got := config.String()
	want := `global = variable

[foo]
bar = bar
baz = baz
`
	if got != want {
		t.Errorf("Got:\n%s\nWant:\n%s\n", got, want)
	}
}

func TestSectionCheckCreateDelete(t *testing.T) {
	config := ini.New()
	if exists := config.HasSection("test"); exists != false {
		t.Error("Section test should not exist")
	}
	config.Section("test")
	if exists := config.HasSection("test"); exists != true {
		t.Error("Section test should exist")
	}
	config.DeleteSection("test")
	if exists := config.HasSection("test"); exists != false {
		t.Error("Section test should not exist")
	}
}

func TestKeyCheckCreateDelete(t *testing.T) {
	config := ini.New()
	section := config.Section("test")
	if exists := section.Has("foo"); exists != false {
		t.Error("Key should not exist")
	}
	if value := section.Get("foo"); value != "" {
		t.Error("Key should not be empty string (zero value)")
	}
	section.Set("foo", "bar")
	if exists := section.Has("foo"); exists != true {
		t.Error("Key should exist")
	}
	if value := section.Get("foo"); value != "bar" {
		t.Error("Key should be set to bar")
	}
	section.Delete("foo")
	if exists := section.Has("foo"); exists != false {
		t.Error("Key should not exist")
	}
}

func TestSaveToFile(t *testing.T) {
	f, err := os.CreateTemp("", "*.ini")
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(f.Name())

	config := ini.New()
	section := config.DefaultSection()
	section.Set("global", "variable")
	config.SaveToFile(f.Name())

	configFromFile, err := ini.NewFromFile(f.Name())
	if err != nil {
		t.Error(err)
	}
	if got, want := configFromFile.String(), config.String(); got != want {
		t.Errorf("Got:\n%s\nWant:\n%s\n", got, want)
	}
}

func TestIntegration(t *testing.T) {
	var source = `; a comment
global=one
# another comment in another syntax
[section1]
a=1
 b=2
c =3
d= 4
e=5 

[section2]
[sectionnotclosed
[section3]
favorite color = red
favorite color = blue
i am a flag =
=invalid
mass energy = e=m*c^2
semicolon="looks like this: ;"
semicolon2=looks like this: ;
someone with no syntactic clue added nonsense
the quick brown fox = jumps over the lazy dog
the lazy dog = sleeps on the equal key ==================`

	var want = `global = one

[section1]
a = 1
b = 2
c = 3
d = 4
e = 5

[section2]

[section3]
favorite color = blue
i am a flag = 
mass energy = e=m*c^2
semicolon = "looks like this: ;"
semicolon2 = looks like this: ;
the lazy dog = sleeps on the equal key ==================
the quick brown fox = jumps over the lazy dog
`
	f, err := os.CreateTemp("", "*.ini")
	if err != nil {
		t.Error(err)
	}
	f.Write([]byte(source))
	defer os.Remove(f.Name())

	config, err := ini.NewFromFile(f.Name())
	if err != nil {
		t.Error(err)
	}
	got := config.String()
	if got != want {
		t.Errorf("Got:\n%s\nWant:\n%s", got, want)
	}
}
