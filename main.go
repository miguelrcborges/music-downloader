package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func clearScreen() {
	fmt.Print("\033[2J\033[H")
}

func main() {
	var input string
	stdin_reader := bufio.NewReader(os.Stdin)

	for {
		clearScreen()
		fmt.Print("Search query: ")
		input, _ = stdin_reader.ReadString('\n')
		// Remove newline
		input = input[:len(input)-1]

		results := search(input)

		if len(results) == 0 {
			fmt.Println("No results found.\nPress enter to continue.")
			stdin_reader.ReadString('\n')
			continue
		}

		for i, result := range results {
			fmt.Printf("%d: [%s] %s - %s\n", i+1, result.Type, result.Artist, result.Name)
		}

		var chosenOption int
	select_track:
		fmt.Println("What do you want to download. Type 0 if none.")
		fmt.Scan(&chosenOption)
		stdin_reader.ReadString('\n')
		if chosenOption <= 0 {
			continue
		} else if chosenOption > len(results) {
			fmt.Println("Invalid option")
			goto select_track
		}

		// fmt.Printf("%d: [%s] %s - %s\n", choosenTrack, results[choosenTrack-1].Type, results[choosenTrack-1].Artist, results[choosenTrack-1].Name)
		// stdin_reader.ReadString('\n')
		download_options := getDownloadOptions(results[chosenOption-1].Link)

		clearScreen()

	select_stream:
		for i, k := range download_options {
			fmt.Printf("%d: %s\n", i+1, k[0])
		}
		fmt.Println("Select an available streaming service. Type 0 to cancel.")
		fmt.Scan(&chosenOption)
		stdin_reader.ReadString('\n')
		if chosenOption <= 0 {
			continue
		} else if chosenOption > len(download_options) {
			fmt.Println("Invalid option")
			goto select_stream
		}

		request_id := requestDownload(download_options[chosenOption-1][1])
		if request_id == "" {
			fmt.Println("Failed to download.\nPress enter to continue.")
			stdin_reader.ReadString('\n')
			continue
		}

		for {
			clearScreen()
			status := getDownloadStatus(request_id)
			if status.Status == "Finished." {
				break
			}
			fmt.Println(status.Status)
			if status.Percent != 0 {
				fmt.Printf("\t%d%%\n", status.Percent)
			}
			time.Sleep(1 * time.Second)
		}

		fmt.Printf("Downloading track as %s.zip", request_id)
		if startDownload(request_id) != nil {
			clearScreen()
			fmt.Println("Failed to download the file.")
		} else {
			clearScreen()
			fmt.Println("Download completed.")
		}
		stdin_reader.ReadString('\n')
	}
}
