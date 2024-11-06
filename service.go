package main

import (
	"log"
	"os"
	"time"
	"user/app/Models"
	"user/database"

	"gopkg.in/yaml.v3"
)

type ServiceConfig struct {
	Name     string         `yaml:"name"`
	Security SecurityConfig `yaml:"security"`
}

type SecurityConfig struct {
	Permissions []Models.Permission `yaml:"permissions"`
}

func InitService() {
	// Open the YAML file
	file, err := os.Open("config/service.yml")
	if err != nil {
		log.Fatalf("Error opening YAML file: %v", err)
	}
	defer file.Close()

	// Decode the YAML file
	var serviceConfigs []ServiceConfig
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&serviceConfigs); err != nil {
		log.Fatalf("Error decoding YAML file: %v", err)
	}

	// Save permissions to the database
	for _, serviceConfig := range serviceConfigs {
		for _, permission := range serviceConfig.Security.Permissions {
			permission.Service = "user-service"

			var existingPermission Models.Permission
			if err := database.DBConn.Where("identifier = ? AND service = ?", permission.Identifier, permission.Service).First(&existingPermission).Error; err == nil {
				isDirty := false

				if existingPermission.Name != permission.Name {
					existingPermission.Name = permission.Name
					isDirty = true
				}

				if existingPermission.Description != permission.Description {
					existingPermission.Description = permission.Description
					isDirty = true
				}

				if existingPermission.Olac != permission.Olac {
					existingPermission.Olac = permission.Olac
					isDirty = true
				}

				if isDirty {
					existingPermission.UpdatedAt = time.Now()
					if err := database.DBConn.Save(&existingPermission).Error; err != nil {
						log.Printf("Error updating permission %s: %v", permission.Name, err)
					} else {
						log.Printf("Permission %s updated successfully\n", permission.Name)
					}
				}
			} else {
				permission.CreatedAt = time.Now()
				if err := database.DBConn.Create(&permission).Error; err != nil {
					log.Printf("Error saving permission %s: %v", permission.Name, err)
				} else {
					log.Printf("Permission %s saved successfully\n", permission.Name)
				}
			}
		}
	}
}
