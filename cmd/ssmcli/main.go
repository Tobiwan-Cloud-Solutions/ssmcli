package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"ssmcli"
	"ssmcli/internal"
	"ssmcli/utils"
)

func main() {
	var err error
	defer func() {
		if err != nil {
			log.Fatal(err)
		}
	}()

	sessionMap := internal.GetSessions()
	if len(sessionMap) > 0 {
		sessionPrompt := ssmcli.BuildPrompt(ssmcli.SessionLabel, utils.MapKeysToSlice(sessionMap))
		selectedSessionId, err := ssmcli.SelectFromMap(sessionPrompt, sessionMap)
		if err != nil {
			return
		}
		if selectedSessionId == ssmcli.SkipPrompt {
			internal.RunSSMResume(selectedSessionId)
		}
	}

	instanceMap := internal.GetInstances()
	if len(instanceMap) == 0 {
		return
	}

	instancePrompt := ssmcli.BuildPrompt(ssmcli.InstanceLabel, utils.MapKeysToSlice(instanceMap))
	selectedInstanceId, err := ssmcli.SelectFromMap(instancePrompt, instanceMap)
	if err != nil {
		return
	}

	actionPrompt := ssmcli.BuildPrompt(ssmcli.ActionLabel, ssmcli.Actions)

	action, err := ssmcli.PromptRunner(actionPrompt)
	if err != nil {
		return
	}

	if action == ssmcli.ActionSSH {
		internal.RunSSMShell(selectedInstanceId)
	} else if action == ssmcli.SkipPrompt {
		// continue
	} else {
		err = fmt.Errorf("unsupported action")
	}
}
