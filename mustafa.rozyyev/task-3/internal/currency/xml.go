package currency

import (
	"encoding/xml"
	"errors"
	"fmt"
	"os"

	"golang.org/x/net/html/charset"
)

type xmlDocument struct {
	Rates []XMLRate `xml:"Valute"`
}

func ReadXML(path string) (rates []XMLRate, err error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open XML file %q: %w", path, err)
	}
	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			err = errors.Join(err, fmt.Errorf("close XML file %q: %w", path, closeErr))
		}
	}()

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = charset.NewReaderLabel

	var document xmlDocument
	if err := decoder.Decode(&document); err != nil {
		return nil, fmt.Errorf("decode XML file %q: %w", path, err)
	}

	return document.Rates, nil
}
