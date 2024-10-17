package helpers

import (
	"os"
)

// Reads the entire contents of a file into a string.
//
// It returns a pointer to the string for potential performance
// optimization when dealing with large files content.
func ReadFileContentToString(path string) (*string, error) {
	content, errRead := os.ReadFile(path)
	if errRead != nil {
		return nil, errRead
	}
	result := string(content)
	return &result, nil
}

//func SaveFile(path string, content string) error {
//	errWrite := os.WriteFile(path, []byte(content), 0644)
//	if errWrite != nil {
//		return errWrite
//	}
//	return nil
//}
//
//func SaveFile(file multipart.File, handler *multipart.FileHeader) (string, error) {
//	//2. Retrieve file from form-data
//	//<Form-id> is the form key that we will read from. Client should use the same form key when uploading the file
//	defer func(file multipart.File) {
//		err := file.Close()
//		if err != nil {
//
//		}
//	}(file)
//
//	//3. Create a temporary file to our directory
//	tempFolderPath := fmt.Sprintf("%s%s", "./public/", "/tempFiles")
//	tempFileName := fmt.Sprintf("upload-%s-*.%s", fileNameWithoutExtension(handler.Filename), filepath.Ext(handler.Filename))
//
//	tempFile, err := os.CreateTemp(tempFolderPath, tempFileName)
//	if err != nil {
//		errStr := fmt.Sprintf("Error in creating the file %s\n", err)
//		fmt.Println(errStr)
//		return errStr, err
//	}
//
//	defer func(tempFile *os.File) {
//		err := tempFile.Close()
//		if err != nil {
//
//		}
//	}(tempFile)
//
//	//4. Write upload file bytes to your new file
//	filebytes, err := io.ReadAll(file)
//	if err != nil {
//		errStr := fmt.Sprintf("Error in reading the file buffer %s\n", err)
//		fmt.Println(errStr)
//		return errStr, err
//	}
//
//	write, err := tempFile.Write(filebytes)
//	if err != nil {
//		return "", err
//	}
//	return "Successfully uploaded\n", nil
//}
