package main

import (
    "os"
    "fmt"
    "net/http"
    "io/ioutil"
    "strconv"
    _"strings"
    "regexp"
    "strings"
    "encoding/base64"
    _"encoding/binary"
    "reflect"
    _"bytes"
	_"github.com/go-audio/audio"
	"github.com/go-audio/wav"
    "io"
)

var URL, BODY, filename, base64string string
const offset int = 64

func main(){

    // parsing done in init
    base64uint8, _ := base64.StdEncoding.DecodeString( base64string )
    fmt.Println("base64 data/head:", base64uint8[ :offset], reflect.TypeOf(base64uint8))
    
    err := os.WriteFile("indian.wav", base64uint8, 0644)



	// Open the source WAV file
	sourceFile, err := os.Open("indian.wav")
	if err != nil {
		fmt.Println("Error opening source file:", err)
		return
	}
	defer sourceFile.Close()

	// Create a new reader to read the WAV file
	decoder := wav.NewDecoder(sourceFile)

	// Create the output file for the reversed WAV
	outputFile, err := os.Create("reversed.wav")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	// Create a new writer for the output WAV file
encoder := wav.NewEncoder(outputFile, decoder.SampleRate, decoder.BitDepth, decoder.NumChans, decoder.WavAudioFormat)

	// Buffer to hold the frames
	frameBuffer := &audio.IntBuffer{
		Format:         &audio.Format{SampleRate: decoder.SampleRate, NumChannels: decoder.NumChans},
		Data:           make([]int, decoder.NumChans),
		SourceBitDepth: decoder.BitDepth,
	}

	// Iterate over each frame, reverse it, and write to the output file
	for {
		// Read one frame (i.e., one sample for each channel)
		if err := decoder.PCMFrame(frameBuffer); err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error reading frame:", err)
			return
		}

		// Reverse the samples in the frame (for stereo, swap channels)
		for i, j := 0, len(frameBuffer.Data)-1; i < j; i, j = i+1, j-1 {
			frameBuffer.Data[i], frameBuffer.Data[j] = frameBuffer.Data[j], frameBuffer.Data[i]
		}

		// Write the reversed frame to the output file
		if err := encoder.Write(frameBuffer); err != nil {
			fmt.Println("Error writing frame:", err)
			return
		}
	}

	// Close the encoder to finalize the file
	if err := encoder.Close(); err != nil {
		fmt.Println("Error closing output file:", err)
	}




}








func init(){

    URL = "http://www.pythonchallenge.com/pc/hex/bin.html"
    ups, sep := "butterfly", 6
    conn := & http.Client{}
    req, err := http.NewRequest("GET", URL, nil)
    if err != nil {fmt.Println("err/", err)}

    req.SetBasicAuth(ups[: sep], ups[sep :])
    resp, _ := conn.Do(req)
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println(string(body), "\nbody ends/\n\n")
    BODY = string(body)

    // filename
    re := regexp.MustCompile(`(?s)name="(.*?)"`)
    matches := re.FindAllStringSubmatch(BODY, -1)
    filename = matches[0][1]

    // get the bound
    re = regexp.MustCompile(`(?s)boundary="(.*?)"`)
    matches = re.FindAllStringSubmatch(BODY, -1)
    bound := "--" + matches[0][1]
    N := len(matches[0])
    fmt.Println("len/", len(matches[0]), "matches/", matches)
    fmt.Println("match/", strconv.Quote( matches[0][1] ))
    fmt.Println("modf./", strconv.Quote( bound ))
    fmt.Println("check/", "--===============1295515792==")


    // get the bounded trunk which should look base64 encoded
    re = regexp.MustCompile(fmt.Sprintf(`(?s)%s(.*?)%s`, bound, bound))
    matches = re.FindAllStringSubmatch(BODY, -1)
    N = len(matches[0][1])
    end := N - offset
    fmt.Println("\nlen/", len(matches[0]))
    fmt.Println("aft/", matches[0][1][: offset], "bef/", matches[0][1][end :])

    trunk := strings.Split(matches[0][1], "\n\n")
    base64string = trunk[1]
    N = len( base64string )
    end = N - offset
    fmt.Println("len/", len(trunk))
    fmt.Println("aft/", trunk[1][: offset], "bef/", trunk[1][end :])
}

