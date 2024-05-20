package main

import (
	"fmt"
	"log"
	"os"

	"github.com/claudiolor/textfsmgo/pkg/textfsmgo"
	"github.com/claudiolor/textfsmgo/pkg/utils"
	"github.com/melbahja/goph"
	"github.com/spf13/viper"
)

type NetworkInterface struct {
	Interface string
	IPAddress string
	Status    string
	Protocol  string
}

func showError(err error, ecode int) {
	fmt.Println(err.Error())
	os.Exit(ecode)
}

func useAddressList(l []NetworkInterface) {
	fmt.Println(l)
}

func main() {
	// Load configuration from config file
	viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	// Get username, password, and device from config file
	username := viper.GetString("username")
	password := viper.GetString("password")
	devices := viper.GetStringSlice("devices")
	device := devices[0]

	// SSH client configuration
	auth := goph.Password(password)
	client, err := goph.NewUnknown(username, device, auth)
	if err != nil {
		log.Fatalf("Failed to create SSH client: %v", err)
	}
	defer client.Close()

	// Run command
	out, err := client.Run("show ip interface brief")
	if err != nil {
		log.Fatalf("Failed to run command: %v", err)
	}

	// Save command output to a file for textfsmgo to read
	outputFile := "show_ip_interface_brief.raw"
	err = os.WriteFile(outputFile, out, 0644)
	if err != nil {
		log.Fatalf("Failed to write command output to file: %v", err)
	}

	// Load TextFSM template
	tmplFile := "show_ip_interface_brief.textfsm"
	parser, err := textfsmgo.NewTextFSMParser(tmplFile)
	if err != nil {
		log.Fatalf("Failed to parse template: %v", err)
	}

	// Read command output from file
	inputStr, err := os.ReadFile(outputFile)
	if err != nil {
		log.Fatalf("Failed to read command output file: %v", err)
	}

	// Parse command output with TextFSM
	res, err := parser.ParseTextToDicts(string(inputStr))
	if err != nil {
		showError(err, 1)
	}

	// Result to struct
	interfaces := []NetworkInterface{}
	for _, entry := range res {
		interfaces = append(
			interfaces,
			NetworkInterface{
				Interface: entry["INTERFACE"].(string),
				IPAddress: entry["IP_ADDRESS"].(string),
				Status:    entry["STATUS"].(string),
				Protocol:  entry["PROTOCOL"].(string),
			},
		)
	}
	useAddressList(interfaces)

	// Convert to JSON
	if jsonRes, err := utils.ConvertResToJson(&res, true); err != nil {
		showError(err, 1)
	} else {
		fmt.Println(string(jsonRes))
	}
}
