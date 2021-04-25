package system

import (
	"os"
	"os/exec"
)

type conf struct {
	Devices map[string]string
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
