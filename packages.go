package main

import (
	"bytes"
	"log"
	"os/exec"
	"strings"
	"time"
)

type Package struct {
	Name          string
	Version       string
	BuildDate     time.Time
	InstallDate   time.Time
	InstallReason string
	Value         string
}

const DATE_LAYOUT = "Mon 02 Jan 2006 15:04:05"

func (p Package) String() string      { return p.Name }
func (p Package) Title() string       { return p.Name }
func (p Package) Description() string { return p.Version }
func (p Package) FilterValue() string { return p.Name }

// Return all packages available from yay
func GetAllPackages() []byte {
	out, err := exec.Command("yay", "-Slq").Output()
	if err != nil {
		log.Fatal(err)
	}
	return out
}

// get all the installed packages on the system
func GetInstalled() [][]byte {
	out, err := exec.Command("yay", "-Qqi").Output()
	if err != nil {
		log.Fatal(err)
	}
	return bytes.Split(out, []byte("\n\n"))
}

// map the result of a yay query to a package struct
func NewPackage(p []byte) (Package, bool) {
	fields := bytes.Split(p, []byte("\n"))
	var name, version string
	var result []time.Time

	for _, v := range fields {
		field := string(v)
		if strings.Contains(field, "Install Reason") && !strings.Contains(field, "Explicitly installed") {
			return Package{}, false
		}
		if strings.Contains(field, "Optional Deps") {
			continue
		}
		splitStr := strings.SplitN(string(field), ":", 2)
		if len(splitStr) == 1 {
			continue
		}
		for i, v := range splitStr {
			splitStr[i] = strings.Trim(v, " ")
		}

		key := splitStr[0]
		value := splitStr[1]

		switch key {
		case "Name":
			name = value
		case "Version":
			version = value
		case "Install Date", "Build Date":
			date, err := time.Parse(DATE_LAYOUT, value)
			if err != nil {
				log.Fatal(err)
			}
			result = append(result, date)
		}
	}

	return Package{
		Name:        name,
		Version:     version,
		InstallDate: result[1],
		BuildDate:   result[0],
		Value:       string(p),
	}, true
}

func GetUserInstalledPackages() []Package {
	installed := GetInstalled()
	packages := []Package{}
	for _, v := range installed {
		if len(v) == 0 {
			continue
		}
		pack, ok := NewPackage(v)
		if ok {
			packages = append(packages, pack)
		}
	}

	return packages
}
