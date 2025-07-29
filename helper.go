package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please run 'go run helper.go kibana.yaml' to update the Kibana YAML file with the Elasticsearch pod IP.")
		return
	}

	yamlFile := os.Args[1]
	namespace := "aa-upwork"
	podName := "elasticsearch-0"
	placeholder := "<IP>"

	// Step 1: Get Pod IP
	cmd := exec.Command("kubectl", "get", "pod", podName, "-n", namespace, "-o", "jsonpath={.status.podIP}")
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error getting pod IP: %v\n", err)
		return
	}
	podIP := out.String()
	if podIP == "" {
		fmt.Println("Pod IP not found.")
		return
	}
	fmt.Printf("Pod IP: %s\n", podIP)

	// Step 2: Read file
	content, err := os.ReadFile(yamlFile)
	if err != nil {
		fmt.Printf("Error reading YAML file %s: %v\n", yamlFile, err)
		return
	}

	// Step 3: Replace placeholder
	updatedContent := strings.ReplaceAll(string(content), placeholder, podIP)

	// Step 4: Write updated file
	err = os.WriteFile(yamlFile, []byte(updatedContent), 0644)
	if err != nil {
		fmt.Printf("Error writing updated YAML file: %v\n", err)
		return
	}

	fmt.Printf("Successfully updated %s with pod IP.\n", yamlFile)
}

