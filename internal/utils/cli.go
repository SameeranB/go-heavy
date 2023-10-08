package utils

import (
	"bufio"
	"fmt"
	"github.com/SameeranB/go-heavy/pkg"
	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

func GetWorkflowConfigFromTest(pathToWorkflow string) (pkg.Workflow, error) {

	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current directory: %s", err)
	}

	// Create a new Yaegi interpreter instance with module support
	i := interp.New(
		interp.Options{
			GoPath:    currentDir, // Set the GoPath to the current directory
			BuildTags: []string{"modules"},
		},
	)

	err = i.Use(stdlib.Symbols)
	if err != nil {
		return pkg.Workflow{}, err
	}

	moduleName, err := getModuleName("go.mod")
	if err != nil {
		log.Fatalf("Failed to get module name: %s", err)
	}

	fullImportPath := moduleName + "/" + pathToWorkflow

	importStatement := fmt.Sprintf(`import "%s"`, fullImportPath)

	_, err = i.Eval(importStatement)
	if err != nil {
		log.Fatalf("Failed to import package: %s", err)
	}

	functionCall := fmt.Sprintf("%s.main()", pathToWorkflow)
	v, err := i.Eval(functionCall)
	if err != nil {
		log.Fatalf("Failed to execute function: %s", err)
	}

	workflow := v.Interface().(pkg.Workflow)

	return workflow, nil

}

func RunWorkflow(workflow pkg.Workflow, duration int, concurrency int) error {
	// Use a WaitGroup to wait for all goroutines (simulating users) to complete
	var wg sync.WaitGroup

	// Start the test
	fmt.Println("Starting the test...")

	fmt.Println("Running setup...")
	pctx, err := workflow.SetUp(pkg.PassedContext{})
	if err != nil {
		return err
	}

	// Create a channel to signal all goroutines to stop after the duration
	stopChan := make(chan struct{})

	// Launch goroutines simulating concurrent users
	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-stopChan:
					return
				default:
					pctx, err = runIteration(workflow, pctx)
					if err != nil {
						log.Println(err)
						close(stopChan)
					}
				}
			}
		}()
	}

	// Let the test run for the specified duration
	time.Sleep(time.Duration(duration) * time.Second)

	// Signal all goroutines to stop
	close(stopChan)

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("Test completed!")
	return nil
}

func runIteration(workflow pkg.Workflow, pctx pkg.PassedContext) (pkg.PassedContext, error) {
	pctx, err := workflow.Init(pctx)
	if err != nil {
		return pctx, err
	}

	for _, step := range workflow.Steps {
		pctx, err = step(pctx)
		if err != nil {
			return pctx, err
		}
	}

	pctx, err = workflow.Teardown(pctx)
	if err != nil {
		return pctx, err
	}

	return pctx, nil
}

func getModuleName(goModPath string) (string, error) {
	file, err := os.Open(goModPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "module ") {
			return strings.TrimSpace(line[len("module "):]), nil
		}
	}
	return "", nil
}
