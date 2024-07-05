package config

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

type AllConfiguration struct {
	Port        string
	MySqlHost   string
	MySqlPort   string
	MySqlName   string
	MySqlPass   string
	MySqlDbName string
	MySqlDriver string
}

func (c *AllConfiguration) Required(key string, value string) error {
	if value == "" {
		return errors.New(fmt.Sprintf("config %s is required", key))
	}
	return nil
}

func InitialAllConfig() (*AllConfiguration, error) {
	c := &AllConfiguration{
		Port:        os.Getenv("PORT"),
		MySqlHost:   os.Getenv("MYSQL_HOST"),
		MySqlPort:   os.Getenv("MYSQL_PORT"),
		MySqlName:   os.Getenv("MYSQL_NAME"),
		MySqlPass:   os.Getenv("MYSQL_PASS"),
		MySqlDbName: os.Getenv("MYSQL_DB_NAME"),
		MySqlDriver: os.Getenv("MYSQL_DRIVER_NAME"),
	}

	// check required for fields
	err := c.Required("PORT", c.Port)
	err = c.Required("MYSQL_HOST", c.MySqlHost)
	err = c.Required("MYSQL_PORT", c.MySqlPort)
	err = c.Required("MYSQL_NAME", c.MySqlName)
	err = c.Required("MYSQL_PASS", c.MySqlPass)
	err = c.Required("MYSQL_DATABASE", c.MySqlDbName)

	if err != nil {
		return nil, err
	}
	return c, nil
}
