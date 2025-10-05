package createtemplate

import (
	"fmt"
	"os"
)

func CreateCTemplate() {
	err := os.Mkdir("src", 0755)
	if err != nil {
		fmt.Println("Error during the creation of the src dir: ", err)
	}

	file, err := os.Create("src/main.c")
	if err != nil {
		fmt.Println("Error during the creation of main.c: ", err)
	}
	defer file.Close()

	_, err = file.WriteString(`#include <stdio.h>

int main(void) {
    printf("Hello World!\n");

    return 0;
}
`)
	if err != nil {
		fmt.Println("Error writing the file: ", err)
	}

	file, err = os.Create("tasks.yaml")
	if err != nil {
		fmt.Println("Error during the creation of tasks.yaml: ", err)
	}
	defer file.Close()

	yamlContent := `tasks:
  - name: build
    command: "mkdir -p build && gcc src/main.c -o build/main"
    depents_on: []

  - name: run
    command: ./build/main
    depents_on: [build]

  - name: clean
    command: rm -rf build
    depents_on: []
`
	_, err = file.WriteString(yamlContent)
	if err != nil {
		fmt.Println("Error writing the file: ", err)
	}

	fmt.Println("[*] => C template created successfully")
}
