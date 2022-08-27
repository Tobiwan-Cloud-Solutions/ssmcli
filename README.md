[![TWSLogo](https://github.com/Tobiwan-Cloud-Solutions/images/blob/main/TWSBannerGHLogo.png)](https://tobiwansolutions.net)

# SSM-CLI #

This is a simple golang CLI to enable users to easily start AWS Systems Manager 
Sessions from the terminal, without having to remember the command themselves 
or list instance IDs.

## Dependencies

At the moment, this tool requires both the [AWS CLI](https://aws.amazon.com/cli/) and the [Session Manager Plugin](https://docs.aws.amazon.com/systems-manager/latest/userguide/session-manager-working-with-install-plugin.html).

## Usage

With some enviroment variables for AWS set, just run `ssmcli`. The script will list existing sessions first so you can reconnect if one died, then it will list all instances.

Session Manager needs to be working for this script to work, it's just a wrapper! So if the agent isn't installed or can't connect to AWS, it won't work.

## Detailed Usage

If you've no profile or variables or token, you'll get this:

    $ ssmcli
    ERRO[0004] aws error                                     code=AuthFailure message="AWS was not able to validate the provided access credentials"

With variables set, you'll get a list of instances that Session Manager knows about:

    Use the arrow keys to navigate: ↓ ↑ → ←
    ? Select instance:
      ▸ ec2.jenkins
        ec2.centos.7

Upon selecting an instance, you will get the option to choose your region:

    Use the arrow keys to navigate: ↓ ↑ → ←
    ? AWS Region:
      ▸ us-west-1
        us-west-2
        us-east-1
        us-east-2

When you select one, you get the option to start a shell, or forward ports:

    Use the arrow keys to navigate: ↓ ↑ → ←
    ? Select action:
      ▸ start SSH
        forward ports

If you then select SSH, you get a session:
    
    Starting session with SessionId: 0a0              
    sh-4.2$

If you want to forward ports, then you get the following:

    ✔ forward ports
    DEBU[0002] selected Action                               action="forward ports"
    Use the arrow keys to navigate: ↓ ↑ → ←
    ? Select port to forward:
      ▸ 22
        80
        443
        8080

(TODO) And then selecting the port runs the script:

## Building

Needs a version of Go that supports modules, and then:

    go install 

Put it somewhere on your PATH or add GOBIN to your PATH.
