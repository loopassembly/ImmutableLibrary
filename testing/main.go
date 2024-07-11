package main

import (
    "bytes"
    "fmt"
    "os/exec"
    "strings"
)


func UploadFileAndGetCID(filePath string) (string, error) {
    // Command to run node with upload.js and pass the file path as an argument
    cmd := exec.Command("node", "/home/loopassembly/Documents/hack4bengal-backend/upload.js", filePath)

    var stdout, stderr bytes.Buffer
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr

    // Execute the command
    if err := cmd.Run(); err != nil {
        return "", fmt.Errorf("error running upload.js: %v\nStderr: %s", err, stderr.String())
    }

    // Capture the output from JavaScript
    output := strings.TrimSpace(stdout.String())

    // Check if output starts with "CID(" and ends with ")"
    if strings.HasPrefix(output, "CID(") && strings.HasSuffix(output, ")") {
        cid := output[4 : len(output)-1] // Extract CID excluding "CID(" and ")"
        return cid, nil
    }

    return "", fmt.Errorf("unexpected output from upload.js: %s", output)
}


// func main() {
//     filePath := "/home/loopassembly/Documents/hack4bengal-backend/testing/168.png"

//     // Call the function to upload file and get CID
//     cid, err := UploadFileAndGetCID(filePath)
//     if err != nil {
//         fmt.Println("Error:", err)
//         return
//     }


//     fmt.Println("CID returned from JavaScript:", cid)
// }
