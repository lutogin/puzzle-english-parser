package config

import (
	"fmt"
	"strings"
)

type Config struct {
	API APIConfig
	APP AppConfig
}

type APIConfig struct {
	BaseURL   string
	Endpoints EndpointsConfig
	Params    map[string]string
}

type AppConfig struct {
	Config map[string]string
}

type EndpointsConfig struct {
	Dictionary string
}

var cfg *Config

func GetConfig() (*Config, error) {
	if cfg != nil {
		return cfg, nil
	}

	cfg = &Config{
		API: APIConfig{
			BaseURL: "https://puzzle-english.com",
			Endpoints: EndpointsConfig{
				Dictionary: "/change-my-dictionary",
			},
			Params: map[string]string{
				"dictionaryChange": "for_dictionary_change=true",
				"ajaxAction":       "ajax_action=ajax_pe_get_next_page_dictionary",
				"page":             "page=",
			},
		},
		APP: AppConfig{
			Config: map[string]string{
				"fileName":     "words.csv",
				"csvSeparator": "\t", // Using tab as default separator
				"wordsPerPage": "100",
				"cookieFile":   "cookies.txt",
			},
		},
	}

	return cfg, nil
}

func (c *Config) GetDictionaryEndpoint() string {
	return c.API.BaseURL + c.API.Endpoints.Dictionary
}

func (c *Config) GetDictionaryQueryParams(page int) string {
	params := []string{
		c.API.Params["dictionaryChange"],
		c.API.Params["ajaxAction"],
		fmt.Sprintf("%s%d", c.API.Params["page"], page),
	}

	return strings.Join(params, "&")
}
