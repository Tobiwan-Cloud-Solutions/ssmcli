package internal

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"runtime"
	"ssmcli"
	"ssmcli/utils"
	"strings"
)

func selectProfiles() map[string]string {
	arch := runtime.GOOS
	profiles := map[string]string{}
	switch arch {
	case "windows":
		home := os.Getenv("USERPROFILE")
		profileFileLocation := fmt.Sprintf("%s\\.aws\\credentials", home)
		profiles := getProfiles(profileFileLocation)
		return profiles
	case "darwin":
		home := os.Getenv("HOME")
		profileFileLocation := getenv("AWS_CONFIG_FILE", fmt.Sprintf("%s/.aws/credentials", home))
		profiles := getProfiles(profileFileLocation)
		return profiles
	case "linux":
		home := os.Getenv("HOME")
		profileFileLocation := getenv("AWS_CONFIG_FILE", fmt.Sprintf("%s/.aws/credentials", home))
		profiles := getProfiles(profileFileLocation)
		return profiles
	}
	return profiles
}

func getProfiles(profileFileLocation string) map[string]string {
	profileMap := map[string]string{}

	file, err := os.Open(profileFileLocation)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	r, err := regexp.Compile(`\[.+?\]`)

	if err != nil {
		log.Fatal(err)
	}

	for scanner.Scan() {
		if r.MatchString(scanner.Text()) {
			s := scanner.Text()
			reg := regexp.MustCompile(`\[([^\[\]]*)\]`)
			//res := reg.ReplaceAllString(s, "")
			submatchall := reg.FindAllString(s, -1)
			for _, element := range submatchall {
				element = strings.Trim(element, "[")
				element = strings.Trim(element, "]")
				profileMap[s] = element
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return profileMap
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func ProfileSelector() string {
	profileMap := selectProfiles()
	profilePrompt := ssmcli.BuildPrompt(ssmcli.ProfileLabel, utils.MapKeysToSlice(profileMap))
	selectedProfileName, _ := ssmcli.SelectFromMap(profilePrompt, profileMap)
	os.Setenv("AWS_PROFILE", selectedProfileName)
	return selectedProfileName
}

func profileSelect() string {
	profile := ProfileSelector()
	return profile
}

func RegionSelector() string {
	regionMap := map[string]string{
		"us-east-1": "us-east-1",
		"us-east-2": "us-east-2",
		"us-west-1": "us-west-1",
		"us-west-2": "us-west-2",
	}
	regionPrompt := ssmcli.BuildPrompt(ssmcli.RegionLabel, utils.MapKeysToSlice(regionMap))
	selectedRegionName, _ := ssmcli.SelectFromMap(regionPrompt, regionMap)
	os.Setenv("AWS_REGION", selectedRegionName)
	return selectedRegionName
}

func regionSelect() string {
	region := RegionSelector()
	return region
}
