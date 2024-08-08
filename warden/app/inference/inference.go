package inference

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var c = &http.Client{
	Timeout: time.Millisecond * 500,
}

type Input []float64

func (i Input) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	for _, val := range i {
		err := binary.Write(buf, binary.LittleEndian, val)
		if err != nil {
			return nil, err
		}
	}
	return buf.Bytes(), nil
}

func DeserializeInputData(data []byte) (Input, error) {
	buf := bytes.NewReader(data)
	var result []float64

	for {
		var val float64
		err := binary.Read(buf, binary.LittleEndian, &val)
		if err != nil {
			// Check for EOF
			if err.Error() == "EOF" {
				break
			}
			return nil, err
		}
		result = append(result, val)
	}

	return result, nil
}

type Request struct {
	Data Input `json:"data"`
}

type Response struct {
	Data    Input  `json:"value_output"`
	Receipt []byte `json:"receipt"`
}

func Solve(input Input) (Response, error) {
	req := Request{
		Data: input,
	}

	jsonBz, err := json.Marshal(req)
	if err != nil {
		return Response{}, err
	}

	res, err := c.Post("http://localhost:9001/job/solve", "application/json", bytes.NewReader(jsonBz))
	if err != nil {
		return Response{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return Response{}, fmt.Errorf("inference endpoint returned non-200 status code: %d", res.StatusCode)
	}

	var resp Response
	err = json.NewDecoder(res.Body).Decode(&resp)
	if err != nil {
		return Response{}, err
	}

	return resp, nil
}
