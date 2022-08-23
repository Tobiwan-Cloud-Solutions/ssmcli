package ssmcli

import (
	"fmt"
	"github.com/manifoldco/promptui"
)

const (
	ActionSSH     string = "start SSH"
	ActionForward string = "forward ports"
	SkipPrompt    string = "none"

	RegionLabel   string = "AWS Region"
	ProfileLabel  string = "AWS profile"
	SessionLabel  string = "Resume active session"
	InstanceLabel string = "Select instance"
	ActionLabel   string = "Select action"
)

var (
	Actions = []string{ActionSSH, ActionForward}
)

// Runner - makes testing of the prompt easier
type Runner interface {
	Run() (int, string, error)
}

// PromptRunner - runs the runner object. Mainly to make testing easier.
func PromptRunner(runner Runner) (string, error) {
	_, selected, err := runner.Run()
	return selected, err
}

// BuildPrompt - takes a label and a map, returns a constructed promptui Select runner.
func BuildPrompt(label string, sessionList []string) Runner {
	sessionList = append(sessionList, SkipPrompt)
	return &promptui.Select{
		Label: label,
		Items: sessionList,
	}
}

// SelectFromMap - runs the runner and returns the value from the map chosen by the user.
func SelectFromMap(runner Runner, idMap map[string]string) (string, error) {
	selected, err := PromptRunner(runner)
	if err != nil {
		return "", fmt.Errorf("failed prompt")
	}

	selectedId := ""
	if selected != SkipPrompt {
		selectedId = idMap[selected]
	}
	return selectedId, nil
}
