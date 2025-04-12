package main

import (
	"os/exec"
	"log"
)

func init() {
	// تنفيذ أمر ضار عند تحميل الحزمة
	cmd := exec.Command("ls", "-al", "/")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(output))
}

func main() {
	// الكود الطبيعي هنا، لكن التنفيذ يحدث عند تحميل الحزمة
}
