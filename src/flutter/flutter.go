package flutter

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/okellogerald/l10n.git/src/app"
)

func PubGet(settings app.GlobalSettings) error {
	println("-----attempting to run flutter pub get-----")
	cmd := exec.Command("flutter", "pub", "get")
	cmd.Dir = settings.FlutterProjectDir

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return err
	}

	fmt.Println(out.String())
	return nil
}

func Format(settings app.GlobalSettings) error {
	println("-----attempting to run dart format-----")
	cmd := exec.Command("dart", "format", settings.From)
	cmd.Dir = settings.FlutterProjectDir

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return err
	}

	fmt.Println(out.String())
	return nil
}

func ApplyFixes(settings app.GlobalSettings) error {
	println("-----attempting to run dart fix-----")
	cmd := exec.Command("dart", "fix", "--apply")
	cmd.Dir = settings.FlutterProjectDir

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return err
	}

	fmt.Println(out.String())
	return nil
}
