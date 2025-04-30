package ini

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

const DefaultSection = ""

type Config map[string]Section
type Section map[string]string

func (config Config) Section(name string) Section {
	section, exists := config[name]
	if !exists {
		config[name] = make(Section)
		section = config[name]
	}
	return section
}

func (config Config) DefaultSection() Section {
	return config.Section(DefaultSection)
}

func (config Config) DeleteSection(name string) {
	delete(config, name)
}

func (config Config) HasSection(name string) bool {
	_, exists := config[name]
	return exists
}

func (config Config) SectionNames() []string {
	var sectionNames []string
	for sn := range config {
		sectionNames = append(sectionNames, sn)
	}
	sort.Strings(sectionNames)
	return sectionNames
}

func (config Config) String() string {
	var sb strings.Builder
	sectionNames := config.SectionNames()
	lastIndex := len(sectionNames) - 1
	for i, sn := range sectionNames {
		if sn != DefaultSection {
			sb.WriteString(fmt.Sprintf("[%s]\n", sn))
		}
		sb.WriteString(config[sn].String())
		if i != lastIndex {
			sb.WriteString("\n")
		}
	}
	return sb.String()
}

func (config Config) SaveToFile(path string) error {
	return os.WriteFile(path, []byte(config.String()), 0644)
}

func (s Section) Keys() []string {
	keys := []string{}
	for k := range s {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func (s Section) Has(key string) bool {
	_, exists := s[key]
	return exists
}

func (s Section) Get(key string) string {
	value, _ := s[key]
	return value
}

func (s Section) Delete(key string) {
	delete(s, key)
}

func (s Section) Set(key string, value string) {
	s[key] = value
}

func (s Section) String() string {
	var sb strings.Builder
	for _, k := range s.Keys() {
		sb.WriteString(fmt.Sprintf("%s = %s\n", k, s[k]))
	}
	return sb.String()
}

func New() Config {
	return make(Config)
}

func NewFromFile(path string) (Config, error) {
	config := New()

	f, err := os.Open(path)
	if err != nil {
		return config, err
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	sectionName := DefaultSection
	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		if line == "" || strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}
		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			sectionName = strings.TrimSpace(line[1 : len(line)-1])
			config.Section(sectionName) // this allows empty sections to be prensent
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			section := config.Section(sectionName)
			key := strings.TrimSpace(parts[0])
			if key != "" {
				val := strings.TrimSpace(parts[1])
				section.Set(key, val)
			}
		}
	}
	return config, s.Err()
}
