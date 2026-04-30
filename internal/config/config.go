package config

type Config struct {
	Calendars []Calendar `yaml:"calendars"`
}

type Calendar struct {
	Name  string `yaml:"name"`
	Feeds []Feed `yaml:"feeds"`
}

type Feed struct {
	URL string `yaml:"url"`
}
