package system

import (
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

type conf struct {
	Interface string
}

var (
	C           conf
	DeviceNames []string
)

func TestWifi() error {
	cmd := exec.Command("ping", "-c", "1", "google.com")

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func Vim(args []string) error {
	cmd := exec.Command("vim", args...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func IWCTL(args []string) error {
	cmd := exec.Command("iwctl", args...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func GetInterfaceFromConfig() string {
	confPath := viper.ConfigFileUsed()
	confFile, err := ioutil.ReadFile(confPath)
	cobra.CheckErr(err)
	err = yaml.Unmarshal(confFile, &C)
	cobra.CheckErr(err)

	return C.Interface
}

func ShowRouteInfo() error {
	cmd := exec.Command("ip", "route")

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func ShowInterfaceInfo() error {
	cmd := exec.Command("ip", "addr")

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func GetDriverInfo() error {
	cmd := exec.Command("bash", "-c", "lspci -k | grep -A 3 'Network'")

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func PrependArgument(arg string, args []string) []string {
	args = append(args, arg)
	args[0], args[len(args)-1] = args[len(args)-1], args[0]
	return args
}
