package flutter

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/okellogerald/l10n.git/src/app"
)

func PubGet() error {
	println("-----attempting to run flutter pub get-----")
	cmd := exec.Command("flutter", "pub", "get")
	cmd.Dir = app.FlutterProjectDir

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return err
	}

	fmt.Println(out.String())
	return nil
}

func Format() error {
	println("-----attempting to run dart format-----")
	cmd := exec.Command("dart", "format", app.LocalizationsDir)
	cmd.Dir = app.FlutterProjectDir

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return err
	}

	fmt.Println(out.String())
	return nil
}

func ApplyFixes() error {
	println("-----attempting to run dart fix-----")
	cmd := exec.Command("dart", "fix", "--apply")
	cmd.Dir = app.FlutterProjectDir

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return err
	}

	fmt.Println(out.String())
	return nil
}
