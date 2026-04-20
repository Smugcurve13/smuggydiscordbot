package main

import (
	"fmt"
	"os"
	
	"golang.org/x/crypto/ssh"
)

func connectSSH() (*ssh.Client, error) {
	host := os.Getenv("SSH_HOST")
	user := os.Getenv("SSH_USER")
	pass := os.Getenv("SSH_PASS")
	config := ssh.ClientConfig{
		User:	user,
		Auth:	[]ssh.AuthMethod{ssh.Password(pass)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, err := ssh.Dial("tcp", host+":22", &config)
	if err != nil {
		fmt.Printf("Error Ocurred: %v", err)
	}
	
	fmt.Println("SSH CONNECTED !!")
	return client , nil
}

func runCommand(client *ssh.Client, command string) (string, error) {
	session, err := client.NewSession()
	if err != nil {
		fmt.Printf("Error in Creating New Session : %v", err)
	}
	defer session.Close()
	output, err3 := session.Output(command)
	if err3 != nil {
		fmt.Printf("Error in Running Command : %v", err3)
	}
	str_output := string(output)
	return str_output, nil
}

func getStats() (string, error) {
	client, err := connectSSH()
	if err != nil {
		fmt.Printf("Error in connecting to SSH: %v", err)
	}
	defer client.Close()

	ram_cmd := "free -m"
	cpu_cmd := `grep 'cpu ' /proc/stat | awk '{usage=($2+$4)*100/($2+$3+$4+$5)} END {print usage"%"}'`
	disk_cmd := "df -h /"
	temp_cmd := `sensors | grep "Package id 0" | awk '{print $4}'`
	
	ram , err1 := runCommand(client, ram_cmd)
	if err1 != nil {
		fmt.Println("Error while excecuting RAM stats command")
	}
	cpu , err2 := runCommand(client, cpu_cmd)
	if err2 != nil {
		fmt.Println("Error while excecuting CPU stats command")
	}
	disk , err3 := runCommand(client, disk_cmd)
	if err3 != nil {
		fmt.Println("Error while excecuting DISK stats command")
	}
	temp , err4 := runCommand(client, temp_cmd)
	if err4 != nil {
		fmt.Println("Error while excecuting TEMPERATURE stats command")
	}
	result := fmt.Sprintf("```\nSmuggyServer Stats\n\nRAM:\n%v\nCPU:\n%v\nDISK:\n%v\nTEMPERATURE:\n%v\n```", ram, cpu, disk, temp)
	fmt.Println(result)
	return result , nil
}

type ServerStats struct {
	RAMUsedPercent  float64
	CPUPercent 		float64
	DiskUsedPercent float64
	TempCelsius  	float64
}