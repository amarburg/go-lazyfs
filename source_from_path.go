package lazyfs

import (
	"fmt"
	"net/url"
	"regexp"
)

// Source from path will automatically build an appropriate FileSource
// depending on whether path is an URL os a local path.
func SourceFromPath(path string) (FileSource, error) {

	// TODO.  Additional validatation of path
	if len(path) == 0 {
		return nil, fmt.Errorf("Zero length filename: %s", path)
	}

	var file FileSource
	var err error

	match, _ := regexp.MatchString("^http", path)
	if match {
		uri, err := url.Parse(path)
		file, err = OpenHttpSource(*uri)
		if err != nil {
			return nil, fmt.Errorf("Error opening URL: %s", err.Error())
		}
	} else {
		file, err = OpenLocalFile(path)
		if err != nil {
			return nil, fmt.Errorf("Error opening file: %s", err.Error())
		}
	}

	if file == nil {
		return nil, fmt.Errorf("Error creating file")
	}

	return file, nil
}
