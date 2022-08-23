package internal

import (
	"github.com/aws/aws-sdk-go/aws/session"
	log "github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"runtime"
	"syscall"
)

const (
	DocumentShell      string = ""
	DocumentForwarding string = "AWS-StartPortForwardingSession"
	Reason             string = "connecting via ssmcli"
)

var (
	sess *session.Session
)

func RunSSMShell(instance string) {
	arch := runtime.GOOS
	switch arch {
	case "windows":
		path := "C:\\Program Files\\Amazon\\AWSCLIV2\\aws.exe"
		cmd := exec.Command(path, "ssm", "start-session", "--target", instance)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			log.WithError(err).Error(err)
		}
	case "darwin":
		err := syscall.Exec("/opt/homebrew/bin/aws", []string{"aws", "ssm", "start-session", "--target", instance}, os.Environ())
		if err != nil {
			log.WithError(err).Error(err)
		}
	case "linux":
		err := syscall.Exec("/usr/local/bin/aws", []string{"aws", "ssm", "start-session", "--target", instance}, os.Environ())
		if err != nil {
			log.WithError(err).Error(err)
		}
	}
}

func RunSSMResume(sessionId string) {
	arch := runtime.GOOS
	switch arch {
	case "windows":
		path := "C:\\Program Files\\Amazon\\AWSCLIV2\\aws.exe"
		cmd := exec.Command(path, "ssm", "resume-session", "--session-id", sessionId)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			log.WithError(err).Error(err)
		}
	case "darwin":
		err := syscall.Exec("/opt/homebrew/bin/aws", []string{"aws", "ssm", "resume-session", "--session-id", sessionId}, os.Environ())
		if err != nil {
			log.WithError(err).Error(err)
		}
	case "linux":
		err := syscall.Exec("/usr/local/bin/aws", []string{"aws", "ssm", "resume-session", "--session-id", sessionId}, os.Environ())
		if err != nil {
			log.WithError(err).Error(err)
		}
	}
}
