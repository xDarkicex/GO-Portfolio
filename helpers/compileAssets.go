package helpers

import (
	"fmt"
	"os/exec"
	"path/filepath"
)

//CompileAssets Asset Pipeline
func CompileAssets() {
	compiled := make(chan bool)

	go func() {
		fmt.Println("Sass Assets")
		err := exec.Command(
			"sass",
			"--watch",
			"./app/assets/stylesheets/:./public/assets/stylesheets/", "--style", "compressed").Start()
		if err != nil {
			Logger.Println(err)
			return
		}

		fmt.Println("Sass Assets Compiled")
		compiled <- true
		close(compiled)
	}()

	go func() {
		fmt.Println("Typscripts Assets")
		applicationFiles, _ := filepath.Glob("./app/assets/typescripts/application/*.ts")
		err := exec.Command("tsc", append([]string{"--outDir", "./public/assets/scripts/application/", "--watch"}, applicationFiles...)...).Start()
		if err != nil {
			Logger.Println(err)
			return
		}
		fmt.Println("Typescript Assets Compiled")
		compiled <- true
		close(compiled)
	}()
	go func() {
		fmt.Println("Typscripts Assets")
		blogFiles, _ := filepath.Glob("./app/assets/typescripts/blog/*.ts")
		err := exec.Command("tsc", append([]string{"--outDir", "./public/assets/scripts/blog/", "--watch"}, blogFiles...)...).Start()
		if err != nil {
			Logger.Println(err)
			return
		}

		fmt.Println("Typescript Assets Compiled")
		compiled <- true
		close(compiled)
	}()
	go func() {
		fmt.Println("Typscripts Assets")
		exampleFiles, _ := filepath.Glob("./app/assets/typescripts/examples/*.ts")
		err := exec.Command("tsc", append([]string{"--outDir", "./public/assets/scripts/examples/", "--watch"}, exampleFiles...)...).Start()
		if err != nil {
			Logger.Println(err)
			return
		}

		fmt.Println("Typescript Assets Compiled")
		compiled <- true
		close(compiled)
	}()
	go func() {
		fmt.Println("Typscripts Assets")
		userFiles, _ := filepath.Glob("./app/assets/typescripts/users/*.ts")
		err := exec.Command("tsc", append([]string{"--outDir", "./public/assets/scripts/users/", "--watch"}, userFiles...)...).Start()
		if err != nil {
			Logger.Println(err)
			return
		}

		fmt.Println("Typescript Assets Compiled")
		compiled <- true
		close(compiled)
	}()
	go func() {
		fmt.Println("Typscripts Assets")
		files, _ := filepath.Glob("./app/assets/typescripts/*.ts")
		err := exec.Command("tsc", append([]string{"--outDir", "./public/assets/scripts/", "--watch"}, files...)...).Start()
		if err != nil {
			Logger.Println(err)
			return
		}

		fmt.Println("Typescript Assets Compiled")
		compiled <- true
		close(compiled)
	}()

}
