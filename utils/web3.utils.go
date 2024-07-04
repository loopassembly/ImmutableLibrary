package utils

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func UploadFileAndGetCID(filePath string) (string, error) {
	// kernal cms=10000233452
	fmt.Println("working")
	cmd := exec.Command("node", "utils/upload.js", filePath)
	fmt.Println("now workingh")

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	print("pkasa")

	// Execute the command
	if err := cmd.Run(); err != nil {
		print(stderr.String())
		return "", fmt.Errorf("error running upload.js: %v\nStderr: %s", err, stderr.String())
	}

	
	output := strings.TrimSpace(stdout.String())
	print(output)


	if strings.HasPrefix(output, "CID(") && strings.HasSuffix(output, ")") {
		cid := output[4 : len(output)-1] 
		return cid, nil
	}
	fmt.Println(output)

	return output,nil
}
